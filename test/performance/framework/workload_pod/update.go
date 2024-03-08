// Copyright 2024 Antrea Authors.
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

package workload_pod

import (
	"context"
	"fmt"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog/v2"
)

const (
	ScaleTestPodProbeContainerName = "antrea-scale-test-pod-probe"
)

func CreateClientPod(ctx context.Context, kClient kubernetes.Interface, namespace, podName string, probes []string, containerName string) (*corev1.Pod, error) {
	var err error
	expectContainerNum := 0
	var newPod *corev1.Pod
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		pod, err := kClient.CoreV1().Pods(namespace).Get(ctx, podName, metav1.GetOptions{})
		if err != nil {
			return err
		}
		var containers []corev1.Container
		for _, probe := range probes {
			l := strings.Split(probe, ":")
			server, port := l[0], l[1]
			if server == "" {
				server = "$NODE_IP"
			}

			containers = append(containers, corev1.Container{
				Name:  containerName,
				Image: "busybox",
				// read up rest </proc/uptime; t1="${up%.*}${up#*.}"
				Command:         []string{"/bin/sh", "-c", fmt.Sprintf("server=%s; output_file=\"ping_log.txt\"; if [ ! -e \"$output_file\" ]; then touch \"$output_file\"; fi; last_status=\"unknown\"; last_change_time=$(adjtimex | awk '/(time.tv_sec|time.tv_usec):/ { printf(\"%%06d\", $2) }' && printf \"\\n\"); while true; do status=$(nc -vz -w 1 \"$server\" %s > /dev/null && echo \"up\" || echo \"down\"); current_time=$(adjtimex | awk '/(time.tv_sec|time.tv_usec):/ { printf(\"%%06d\", $2) }' && printf \"\\n\"); time_diff=$((current_time - last_change_time)); if [ \"$status\" != \"$last_status\" ]; then echo \"$(adjtimex | awk '/(time.tv_sec|time.tv_usec):/ { printf(\"%%06d\", $2) }' && printf \"\\n\") Status changed from $last_status to $status after ${time_diff} nanoseconds\"; echo \"$(adjtimex | awk '/(time.tv_sec|time.tv_usec):/ { printf(\"%%06d\", $2) }' && printf \"\\n\") Status changed from $last_status to $status after ${time_diff} nanoseconds\" >> \"$output_file\"; last_change_time=$current_time; last_status=$status; fi; sleep 0.1; done\n", server, port)},
				ImagePullPolicy: corev1.PullIfNotPresent,
				Env: []corev1.EnvVar{
					{
						Name: "NODE_IP",
						ValueFrom: &corev1.EnvVarSource{
							FieldRef: &corev1.ObjectFieldSelector{
								FieldPath: "status.hostIP",
							},
						},
					},
				},
			})
		}

		pod.Spec.Containers = append(pod.Spec.Containers, containers...)
		expectContainerNum = len(pod.Spec.Containers)

		newPod = &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      strings.Replace(pod.Name, "server", "client", 1),
				Namespace: pod.Namespace,
				Labels:    pod.Labels,
			},
			Spec: pod.Spec,
		}

		_, err = kClient.CoreV1().Pods(namespace).Create(ctx, newPod, metav1.CreateOptions{})
		return err
	})
	if err != nil {
		return nil, err
	}

	err = wait.Poll(time.Second, 30, func() (bool, error) {
		pod, err := kClient.CoreV1().Pods(namespace).Get(ctx, newPod.Name, metav1.GetOptions{})
		if err != nil {
			return false, err
		}

		if expectContainerNum == len(pod.Spec.Containers) {
			return true, nil
		}

		return false, nil
	})
	if err != nil {
		return nil, err
	}

	klog.InfoS("Create Client Pod successfully!")
	return newPod, nil
}
