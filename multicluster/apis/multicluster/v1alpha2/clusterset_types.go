/*
Copyright 2023 Antrea Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha2

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LeaderClusterInfo specifies information of a leader cluster.
type LeaderClusterInfo struct {
	// Identify a leader cluster in the ClusterSet.
	ClusterID string `json:"clusterID,omitempty"`
	// API server endpoint of the leader cluster.
	// E.g. "https://172.18.0.1:6443", "https://example.com:6443".
	Server string `json:"server,omitempty"`
	// Name of the Secret resource in the member cluster, which stores
	// the token to access the leader cluster's API server.
	Secret string `json:"secret,omitempty"`
	// ServiceAccount in the leader cluster, from which the member cluster's token
	// is generated. This is an optional field which helps admin to check
	// which ServiceAccount is used by a member cluster to access the leader cluster.
	//
	// DEPRECATED
	// This field is planned to be removed in the future releases.
	ServiceAccount string `json:"serviceAccount,omitempty"`
}

// ClusterSetSpec defines the desired state of ClusterSet.
type ClusterSetSpec struct {
	// ClusterID identifies the local cluster.
	// +kubebuilder:validation:Required
	ClusterID string `json:"clusterID"`
	// Leaders include leader clusters known to the member clusters.
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:Required
	Leaders []LeaderClusterInfo `json:"leaders"`
	// The leader cluster Namespace in which the ClusterSet is defined.
	// Used in a member cluster.
	Namespace string `json:"namespace,omitempty"`
}

type ClusterSetConditionType string

const (
	// ClusterSetReady indicates whether ClusterSet is ready.
	ClusterSetReady ClusterSetConditionType = "Ready"
)

// ClusterSetCondition indicates the readiness condition of the clusterSet.
type ClusterSetCondition struct {
	Type ClusterSetConditionType `json:"type,omitempty"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status,omitempty"`
	// +optional
	// Last time the condition transited from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// +optional
	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty"`
	// +optional
	// Unique, one-word, CamelCase reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
}

type ClusterConditionType string

const (
	// ClusterReady indicates whether cluster is ready and connected.
	ClusterReady ClusterConditionType = "Ready"
	// ClusterIsLeader indicates whether the cluster is the leader of the local cluster.
	// Used in member clusters only.
	ClusterIsLeader ClusterConditionType = "IsLeader"
	// ClusterConnected indicates whether the member cluster has connected to the
	// local cluster as the leader.
	// Used in leader clusters only.
	ClusterConnected ClusterConditionType = "ClusterConnected"
)

// ClusterCondition indicates the readiness condition of a cluster.
type ClusterCondition struct {
	Type ClusterConditionType `json:"type,omitempty"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status,omitempty"`

	// +optional
	// Last time the condition transited from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// +optional
	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty"`
	// +optional
	// Unique, one-word, CamelCase reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
}

type ClusterStatus struct {
	// ClusterID is the unique identifier of this cluster.
	ClusterID  string             `json:"clusterID,omitempty"`
	Conditions []ClusterCondition `json:"conditions,omitempty"`
}

// ClusterSetStatus defines the observed state of ClusterSet.
type ClusterSetStatus struct {
	// Total number of member clusters configured in the ClusterSet.
	TotalClusters int32 `json:"totalClusters,omitempty"`
	// Total number of clusters ready and connected.
	ReadyClusters int32 `json:"readyClusters,omitempty"`
	// The overall condition of the cluster set.
	Conditions []ClusterSetCondition `json:"conditions,omitempty"`
	// The status of individual member clusters.
	ClusterStatuses []ClusterStatus `json:"clusterStatuses,omitempty"`
	// The generation observed by the controller.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// +genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion

// +kubebuilder:printcolumn:name="Leader Cluster Namespace",type=string,JSONPath=`.spec.namespace`,description="The leader cluster Namespace for the ClusterSet"
// +kubebuilder:printcolumn:name="Total Clusters",type=string,JSONPath=`.status.totalClusters`,description="Total number of clusters in the ClusterSet"
// +kubebuilder:printcolumn:name="Ready Clusters",type=string,JSONPath=`.status.readyClusters`,description="Number of ready clusters in the ClusterSet"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=`.metadata.creationTimestamp`
// ClusterSet represents a ClusterSet.
type ClusterSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSetSpec   `json:"spec,omitempty"`
	Status ClusterSetStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterSetList contains a list of ClusterSet.
type ClusterSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterSet{}, &ClusterSetList{})
}
