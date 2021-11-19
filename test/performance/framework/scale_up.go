// Copyright 2021 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package framework

import (
	"context"
	"fmt"
	"strings"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"

	"antrea.io/antrea/pkg/antctl/runtime"
	"antrea.io/antrea/test/performance/config"
	"antrea.io/antrea/test/performance/utils"
)

const (
	// ScaleTestNamespacePrefix is the Namespace of the scale test resources.
	ScaleTestNamespacePrefix = "antrea-scale-test-ns"

	AppLabelKey   = "app"
	AppLabelValue = "antrea-scale-test-workload"

	SimulatorNodeLabelKey   = "antrea/instance"
	SimulatorNodeLabelValue = "simulator"

	SimulatorTaintKey   = "simulator"
	SimulatorTaintValue = "true"

	ScaleClientContainerName   = "antrea-scale-test-client"
	ScaleClientPodTemplateName = "antrea-scale-test-client"
	ScaleTestClientDaemonSet   = "antrea-scale-test-client-daemonset"

	ShortTimeWait = time.Minute
)

var (
	// RealNodeAffinity is used to make a Pod not to be scheduled to a simulate node.
	RealNodeAffinity = corev1.Affinity{
		NodeAffinity: &corev1.NodeAffinity{
			RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
				NodeSelectorTerms: []corev1.NodeSelectorTerm{
					{
						MatchExpressions: []corev1.NodeSelectorRequirement{
							{
								Key:      SimulatorNodeLabelKey,
								Operator: corev1.NodeSelectorOpNotIn,
								Values:   []string{SimulatorNodeLabelValue},
							},
						},
					},
				},
			},
		},
	}

	// SimulateAffinity is used to make a Pod to be scheduled to a simulate node.
	SimulateAffinity = corev1.Affinity{
		NodeAffinity: &corev1.NodeAffinity{
			RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
				NodeSelectorTerms: []corev1.NodeSelectorTerm{
					{
						MatchExpressions: []corev1.NodeSelectorRequirement{
							{
								Key:      SimulatorNodeLabelKey,
								Operator: corev1.NodeSelectorOpIn,
								Values:   []string{SimulatorNodeLabelValue},
							},
						},
					},
				},
			},
		},
	}

	// SimulateToleration marks a Pod able to run on a simulate node.
	SimulateToleration = corev1.Toleration{
		Key:      SimulatorTaintKey,
		Operator: corev1.TolerationOpEqual,
		Value:    SimulatorTaintValue,
		Effect:   corev1.TaintEffectNoExecute,
	}

	// MasterToleration marks a Pod able to run on the master node.
	MasterToleration = corev1.Toleration{
		Key:      "node-role.kubernetes.io/master",
		Operator: corev1.TolerationOpExists,
		Effect:   corev1.TaintEffectNoSchedule,
	}

	// clientPodTemplate is the PodTemplateSpec of a scale test client Pod.
	clientPodTemplate = corev1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{ScaleClientPodTemplateName: ""}},
		Spec: corev1.PodSpec{
			Affinity:    &RealNodeAffinity,
			Tolerations: []corev1.Toleration{MasterToleration},
			Containers: []corev1.Container{
				{
					Name:            ScaleClientContainerName,
					Image:           "busybox",
					Command:         []string{"nc", "-lk", "-p", "10080"},
					ImagePullPolicy: corev1.PullIfNotPresent,
				},
			},
		},
	}
)

// ScaleData implemented the TestData interface and it provides clients for helping running
// scale test cases.
type ScaleData struct {
	kubernetesClientSet kubernetes.Interface
	kubeconfig          *rest.Config
	clientPods          []corev1.Pod
	Specification       *config.ScaleList
	nodesNum            int
	simulateNodesNum    int
	podsNum             int
}

func createTestPodClients(ctx context.Context, kClient kubernetes.Interface) error {
	if err := utils.DefaultRetry(func() error {
		_, err := kClient.AppsV1().DaemonSets(ScaleTestNamespacePrefix).Create(ctx, &appsv1.DaemonSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      ScaleTestClientDaemonSet,
				Namespace: ScaleTestNamespacePrefix,
				Labels:    map[string]string{ScaleClientPodTemplateName: ""},
			},
			Spec: appsv1.DaemonSetSpec{
				Selector: &metav1.LabelSelector{MatchLabels: map[string]string{ScaleClientPodTemplateName: ""}},
				Template: clientPodTemplate,
			},
		}, metav1.CreateOptions{})
		return err
	}); err != nil {
		return err
	}
	if err := wait.PollImmediateUntil(config.WaitInterval, func() (bool, error) {
		ds, err := kClient.AppsV1().DaemonSets(ScaleTestNamespacePrefix).
			Get(ctx, ScaleTestClientDaemonSet, metav1.GetOptions{})
		if err != nil {
			return false, nil
		}
		return ds.Status.DesiredNumberScheduled == ds.Status.NumberReady, nil
	}, ctx.Done()); err != nil {
		return fmt.Errorf("error when waiting scale test clients to be ready: %w", err)
	}
	if err := wait.PollImmediateUntil(config.WaitInterval, func() (bool, error) {
		podList, err := kClient.CoreV1().Pods(ScaleTestNamespacePrefix).
			List(ctx, metav1.ListOptions{LabelSelector: ScaleClientPodTemplateName})
		if err != nil {
			return false, nil
		}
		for _, pod := range podList.Items {
			if pod.Status.PodIP == "" {
				return false, nil
			}
		}
		return true, nil
	}, ctx.Done()); err != nil {
		return fmt.Errorf("error when waiting scale test clients to get IP: %w", err)
	}
	return nil
}

func validScaleSpecification(c *config.ScaleConfiguration) error {
	if c.NpNumPerNode*2 > c.PodsNumPerNode {
		return fmt.Errorf("networkPolicy quantity is too larger than 1/2 workload Pods quantity, scale may fail")
	}
	if c.SvcNumPerNode*4 > c.PodsNumPerNode {
		return fmt.Errorf("service quantity is too larger than 1/4 workload Pods quantity, scale may fail")
	}
	return nil
}

func ScaleUp(ctx context.Context, kubeConfigPath, scaleConfigPath string) (*ScaleData, error) {
	var td ScaleData
	scaleConfig, err := config.ParseConfigs(scaleConfigPath)
	if err != nil {
		return nil, err
	}
	klog.InfoS("Scale config", "scaleConfig", scaleConfig)
	td.Specification = scaleConfig

	if err := validScaleSpecification(&scaleConfig.ScaleConfiguration); err != nil {
		return nil, err
	}

	kubeConfig, err := runtime.ResolveKubeconfig(kubeConfigPath)
	if err != nil {
		return nil, fmt.Errorf("error when retrieving incluster kubeconfig: %w", err)
	}
	td.kubeconfig = kubeConfig
	kClient, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("error when creating kubernetes client: %w", err)
	}
	masterNodes, err := kClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{LabelSelector: `node-role.kubernetes.io/master`})
	if err != nil {
		return nil, fmt.Errorf("error when getting Nodes in the cluster: %w", err)
	}
	if len(masterNodes.Items) == 0 {
		return nil, fmt.Errorf("can not find a master Node in the cluster")
	}

	td.kubernetesClientSet = kClient

	nodes, err := kClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("error when getting all Nodes: %w", err)
	}

	// Count simulate nodes.
	simulateNodesNum := 0
	for _, node := range nodes.Items {
		if v, ok := node.Labels[SimulatorNodeLabelKey]; ok && v == SimulatorNodeLabelValue {
			simulateNodesNum += 1
		}
	}
	td.nodesNum = len(nodes.Items)
	td.podsNum = td.nodesNum * scaleConfig.PodsNumPerNode
	td.simulateNodesNum = simulateNodesNum

	klog.Infof("Preflight checks and clean up")
	nsList, err := kClient.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	if scaleConfig.PreWorkload {
		klog.Infof("Deleting scale test namespaces")
		for _, ns := range nsList.Items {
			if !strings.HasPrefix(ns.Name, "antrea-scale-test") {
				continue
			}
			if err := kClient.CoreV1().Namespaces().Delete(ctx, ns.Name, metav1.DeleteOptions{}); err != nil {
				return nil, err
			}
		}
		klog.Infof("Checking clean up exist scale test namespaces")
		err = wait.PollImmediateUntil(config.WaitInterval, func() (bool, error) {
			list, err := kClient.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
			if err != nil {
				return false, err
			}
			ok := true
			for _, ns := range list.Items {
				if strings.HasPrefix(ns.Name, "antrea-scale-test") {
					ok = false
					if err := kClient.CoreV1().Namespaces().Delete(ctx, ns.Name, metav1.DeleteOptions{}); err != nil {
						return false, nil
					}
				}
			}
			klog.InfoS("Waiting for namespace clean up", "namespacesNum", len(list.Items))
			return ok, nil
		}, ctx.Done())
		if err != nil {
			return nil, err
		}

		klog.Infof("Creating scale test namespaces")
		if _, err := kClient.CoreV1().Namespaces().Create(ctx, &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ScaleTestNamespacePrefix}}, metav1.CreateOptions{}); err != nil {
			return nil, err
		}

		klog.Infof("Creating the scale test client DaemonSet")
		if err := createTestPodClients(ctx, kClient); err != nil {
			return nil, err
		}

		klog.Infof("Checking scale test client DaemonSet")
		err = wait.PollImmediateUntil(config.WaitInterval, func() (bool, error) {
			podList, err := kClient.CoreV1().Pods(ScaleTestNamespacePrefix).List(ctx, metav1.ListOptions{LabelSelector: ScaleClientPodTemplateName})
			if err != nil {
				return false, fmt.Errorf("error when getting scale test client pods: %w", err)
			}
			if len(podList.Items) == td.nodesNum-td.simulateNodesNum {
				td.clientPods = podList.Items
				return true, nil
			}
			klog.InfoS("Waiting test client DaemonSet Pods ready", "podsNum", len(podList.Items))
			return false, nil
		}, ctx.Done())
		if err != nil {
			return nil, err
		}
	}

	return &td, nil
}
