//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlSource) DeepCopyInto(out *GraphqlSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlSource.
func (in *GraphqlSource) DeepCopy() *GraphqlSource {
	if in == nil {
		return nil
	}
	out := new(GraphqlSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GraphqlSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlSourceList) DeepCopyInto(out *GraphqlSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GraphqlSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlSourceList.
func (in *GraphqlSourceList) DeepCopy() *GraphqlSourceList {
	if in == nil {
		return nil
	}
	out := new(GraphqlSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GraphqlSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlSourceSpec) DeepCopyInto(out *GraphqlSourceSpec) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(HTTPSource)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlSourceSpec.
func (in *GraphqlSourceSpec) DeepCopy() *GraphqlSourceSpec {
	if in == nil {
		return nil
	}
	out := new(GraphqlSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlSourceStatus) DeepCopyInto(out *GraphqlSourceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlSourceStatus.
func (in *GraphqlSourceStatus) DeepCopy() *GraphqlSourceStatus {
	if in == nil {
		return nil
	}
	out := new(GraphqlSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSource) DeepCopyInto(out *HTTPSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSource.
func (in *HTTPSource) DeepCopy() *HTTPSource {
	if in == nil {
		return nil
	}
	out := new(HTTPSource)
	in.DeepCopyInto(out)
	return out
}
