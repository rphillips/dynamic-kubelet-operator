// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicKubeletOperator) DeepCopyInto(out *DynamicKubeletOperator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicKubeletOperator.
func (in *DynamicKubeletOperator) DeepCopy() *DynamicKubeletOperator {
	if in == nil {
		return nil
	}
	out := new(DynamicKubeletOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamicKubeletOperator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicKubeletOperatorList) DeepCopyInto(out *DynamicKubeletOperatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DynamicKubeletOperator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicKubeletOperatorList.
func (in *DynamicKubeletOperatorList) DeepCopy() *DynamicKubeletOperatorList {
	if in == nil {
		return nil
	}
	out := new(DynamicKubeletOperatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamicKubeletOperatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicKubeletOperatorSpec) DeepCopyInto(out *DynamicKubeletOperatorSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicKubeletOperatorSpec.
func (in *DynamicKubeletOperatorSpec) DeepCopy() *DynamicKubeletOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(DynamicKubeletOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicKubeletOperatorStatus) DeepCopyInto(out *DynamicKubeletOperatorStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicKubeletOperatorStatus.
func (in *DynamicKubeletOperatorStatus) DeepCopy() *DynamicKubeletOperatorStatus {
	if in == nil {
		return nil
	}
	out := new(DynamicKubeletOperatorStatus)
	in.DeepCopyInto(out)
	return out
}
