// +build !ignore_autogenerated

/*


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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentsDaemonSetSpec) DeepCopyInto(out *AgentsDaemonSetSpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentsDaemonSetSpec.
func (in *AgentsDaemonSetSpec) DeepCopy() *AgentsDaemonSetSpec {
	if in == nil {
		return nil
	}
	out := new(AgentsDaemonSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentsSpec) DeepCopyInto(out *AgentsSpec) {
	*out = *in
	in.DaemonSet.DeepCopyInto(&out.DaemonSet)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentsSpec.
func (in *AgentsSpec) DeepCopy() *AgentsSpec {
	if in == nil {
		return nil
	}
	out := new(AgentsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerDeploymentSpec) DeepCopyInto(out *AnalyzerDeploymentSpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerDeploymentSpec.
func (in *AnalyzerDeploymentSpec) DeepCopy() *AnalyzerDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(AnalyzerDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerSpec) DeepCopyInto(out *AnalyzerSpec) {
	*out = *in
	in.Deployment.DeepCopyInto(&out.Deployment)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerSpec.
func (in *AnalyzerSpec) DeepCopy() *AnalyzerSpec {
	if in == nil {
		return nil
	}
	out := new(AnalyzerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnableSpec) DeepCopyInto(out *EnableSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnableSpec.
func (in *EnableSpec) DeepCopy() *EnableSpec {
	if in == nil {
		return nil
	}
	out := new(EnableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowExporterDeploymentSpec) DeepCopyInto(out *FlowExporterDeploymentSpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowExporterDeploymentSpec.
func (in *FlowExporterDeploymentSpec) DeepCopy() *FlowExporterDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(FlowExporterDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Skydive) DeepCopyInto(out *Skydive) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Skydive.
func (in *Skydive) DeepCopy() *Skydive {
	if in == nil {
		return nil
	}
	out := new(Skydive)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Skydive) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkydiveFlowExporter) DeepCopyInto(out *SkydiveFlowExporter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkydiveFlowExporter.
func (in *SkydiveFlowExporter) DeepCopy() *SkydiveFlowExporter {
	if in == nil {
		return nil
	}
	out := new(SkydiveFlowExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SkydiveFlowExporter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkydiveFlowExporterList) DeepCopyInto(out *SkydiveFlowExporterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SkydiveFlowExporter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkydiveFlowExporterList.
func (in *SkydiveFlowExporterList) DeepCopy() *SkydiveFlowExporterList {
	if in == nil {
		return nil
	}
	out := new(SkydiveFlowExporterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SkydiveFlowExporterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkydiveFlowExporterSpec) DeepCopyInto(out *SkydiveFlowExporterSpec) {
	*out = *in
	in.Deployment.DeepCopyInto(&out.Deployment)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkydiveFlowExporterSpec.
func (in *SkydiveFlowExporterSpec) DeepCopy() *SkydiveFlowExporterSpec {
	if in == nil {
		return nil
	}
	out := new(SkydiveFlowExporterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkydiveFlowExporterStatus) DeepCopyInto(out *SkydiveFlowExporterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkydiveFlowExporterStatus.
func (in *SkydiveFlowExporterStatus) DeepCopy() *SkydiveFlowExporterStatus {
	if in == nil {
		return nil
	}
	out := new(SkydiveFlowExporterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkydiveList) DeepCopyInto(out *SkydiveList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Skydive, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkydiveList.
func (in *SkydiveList) DeepCopy() *SkydiveList {
	if in == nil {
		return nil
	}
	out := new(SkydiveList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SkydiveList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkydiveSpec) DeepCopyInto(out *SkydiveSpec) {
	*out = *in
	out.Enable = in.Enable
	in.Agents.DeepCopyInto(&out.Agents)
	in.Analyzer.DeepCopyInto(&out.Analyzer)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkydiveSpec.
func (in *SkydiveSpec) DeepCopy() *SkydiveSpec {
	if in == nil {
		return nil
	}
	out := new(SkydiveSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkydiveStatus) DeepCopyInto(out *SkydiveStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkydiveStatus.
func (in *SkydiveStatus) DeepCopy() *SkydiveStatus {
	if in == nil {
		return nil
	}
	out := new(SkydiveStatus)
	in.DeepCopyInto(out)
	return out
}
