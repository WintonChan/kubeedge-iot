// +build !ignore_autogenerated

/*
Copyright 2021.

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

package v1alpha3

import (
	"k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessParameter) DeepCopyInto(out *AccessParameter) {
	*out = *in
	if in.Parameter != nil {
		in, out := &in.Parameter, &out.Parameter
		*out = make(map[string]apiextensionsv1.JSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessParameter.
func (in *AccessParameter) DeepCopy() *AccessParameter {
	if in == nil {
		return nil
	}
	out := new(AccessParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Conditions) DeepCopyInto(out *Conditions) {
	{
		in := &in
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in Conditions) DeepCopy() Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Device) DeepCopyInto(out *Device) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Device.
func (in *Device) DeepCopy() *Device {
	if in == nil {
		return nil
	}
	out := new(Device)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Device) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceAccess) DeepCopyInto(out *DeviceAccess) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceAccess.
func (in *DeviceAccess) DeepCopy() *DeviceAccess {
	if in == nil {
		return nil
	}
	out := new(DeviceAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceAccess) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceAccessSpec) DeepCopyInto(out *DeviceAccessSpec) {
	*out = *in
	if in.AccessParameters != nil {
		in, out := &in.AccessParameters, &out.AccessParameters
		*out = make([]AccessParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceAccessSpec.
func (in *DeviceAccessSpec) DeepCopy() *DeviceAccessSpec {
	if in == nil {
		return nil
	}
	out := new(DeviceAccessSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceCommand) DeepCopyInto(out *DeviceCommand) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceCommand.
func (in *DeviceCommand) DeepCopy() *DeviceCommand {
	if in == nil {
		return nil
	}
	out := new(DeviceCommand)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceCommandStatus) DeepCopyInto(out *DeviceCommandStatus) {
	*out = *in
	out.DeviceCommand = in.DeviceCommand
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceCommandStatus.
func (in *DeviceCommandStatus) DeepCopy() *DeviceCommandStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceCommandStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceList) DeepCopyInto(out *DeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Device, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceList.
func (in *DeviceList) DeepCopy() *DeviceList {
	if in == nil {
		return nil
	}
	out := new(DeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceModel) DeepCopyInto(out *DeviceModel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceModel.
func (in *DeviceModel) DeepCopy() *DeviceModel {
	if in == nil {
		return nil
	}
	out := new(DeviceModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceModel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceModelList) DeepCopyInto(out *DeviceModelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeviceModel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceModelList.
func (in *DeviceModelList) DeepCopy() *DeviceModelList {
	if in == nil {
		return nil
	}
	out := new(DeviceModelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceModelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceModelSpec) DeepCopyInto(out *DeviceModelSpec) {
	*out = *in
	if in.DeviceProperties != nil {
		in, out := &in.DeviceProperties, &out.DeviceProperties
		*out = make([]DeviceProperties, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceModelSpec.
func (in *DeviceModelSpec) DeepCopy() *DeviceModelSpec {
	if in == nil {
		return nil
	}
	out := new(DeviceModelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceModelStatus) DeepCopyInto(out *DeviceModelStatus) {
	*out = *in
	if in.DeviceRefs != nil {
		in, out := &in.DeviceRefs, &out.DeviceRefs
		*out = make([]NamespaceName, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceModelStatus.
func (in *DeviceModelStatus) DeepCopy() *DeviceModelStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceModelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceProperties) DeepCopyInto(out *DeviceProperties) {
	*out = *in
	out.ProfileProperty = in.ProfileProperty
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceProperties.
func (in *DeviceProperties) DeepCopy() *DeviceProperties {
	if in == nil {
		return nil
	}
	out := new(DeviceProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceProtocol) DeepCopyInto(out *DeviceProtocol) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceProtocol.
func (in *DeviceProtocol) DeepCopy() *DeviceProtocol {
	if in == nil {
		return nil
	}
	out := new(DeviceProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSpec) DeepCopyInto(out *DeviceSpec) {
	*out = *in
	if in.DeviceAccessRef != nil {
		in, out := &in.DeviceAccessRef, &out.DeviceAccessRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.DeviceService != nil {
		in, out := &in.DeviceService, &out.DeviceService
		*out = make(map[string]apiextensionsv1.JSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.Protocol.DeepCopyInto(&out.Protocol)
	if in.DeviceCommands != nil {
		in, out := &in.DeviceCommands, &out.DeviceCommands
		*out = make([]DeviceCommand, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSpec.
func (in *DeviceSpec) DeepCopy() *DeviceSpec {
	if in == nil {
		return nil
	}
	out := new(DeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceStatus) DeepCopyInto(out *DeviceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeviceCommands != nil {
		in, out := &in.DeviceCommands, &out.DeviceCommands
		*out = make([]DeviceCommandStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceStatus.
func (in *DeviceStatus) DeepCopy() *DeviceStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceAccessList) DeepCopyInto(out *DeviceAccessList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeviceAccess, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceAccessList.
func (in *DeviceAccessList) DeepCopy() *DeviceAccessList {
	if in == nil {
		return nil
	}
	out := new(DeviceAccessList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceAccessList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceName) DeepCopyInto(out *NamespaceName) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceName.
func (in *NamespaceName) DeepCopy() *NamespaceName {
	if in == nil {
		return nil
	}
	out := new(NamespaceName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileProperty) DeepCopyInto(out *ProfileProperty) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileProperty.
func (in *ProfileProperty) DeepCopy() *ProfileProperty {
	if in == nil {
		return nil
	}
	out := new(ProfileProperty)
	in.DeepCopyInto(out)
	return out
}
