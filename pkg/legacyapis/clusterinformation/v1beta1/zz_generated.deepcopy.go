//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	crdv1beta1 "antrea.io/antrea/pkg/apis/crd/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaAgentInfo) DeepCopyInto(out *AntreaAgentInfo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.PodRef = in.PodRef
	out.NodeRef = in.NodeRef
	if in.NodeSubnets != nil {
		in, out := &in.NodeSubnets, &out.NodeSubnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.OVSInfo.DeepCopyInto(&out.OVSInfo)
	out.NetworkPolicyControllerInfo = in.NetworkPolicyControllerInfo
	if in.AgentConditions != nil {
		in, out := &in.AgentConditions, &out.AgentConditions
		*out = make([]crdv1beta1.AgentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaAgentInfo.
func (in *AntreaAgentInfo) DeepCopy() *AntreaAgentInfo {
	if in == nil {
		return nil
	}
	out := new(AntreaAgentInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AntreaAgentInfo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaAgentInfoList) DeepCopyInto(out *AntreaAgentInfoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AntreaAgentInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaAgentInfoList.
func (in *AntreaAgentInfoList) DeepCopy() *AntreaAgentInfoList {
	if in == nil {
		return nil
	}
	out := new(AntreaAgentInfoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AntreaAgentInfoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaControllerInfo) DeepCopyInto(out *AntreaControllerInfo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.PodRef = in.PodRef
	out.NodeRef = in.NodeRef
	out.ServiceRef = in.ServiceRef
	out.NetworkPolicyControllerInfo = in.NetworkPolicyControllerInfo
	if in.ControllerConditions != nil {
		in, out := &in.ControllerConditions, &out.ControllerConditions
		*out = make([]crdv1beta1.ControllerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaControllerInfo.
func (in *AntreaControllerInfo) DeepCopy() *AntreaControllerInfo {
	if in == nil {
		return nil
	}
	out := new(AntreaControllerInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AntreaControllerInfo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaControllerInfoList) DeepCopyInto(out *AntreaControllerInfoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AntreaControllerInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaControllerInfoList.
func (in *AntreaControllerInfoList) DeepCopy() *AntreaControllerInfoList {
	if in == nil {
		return nil
	}
	out := new(AntreaControllerInfoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AntreaControllerInfoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
