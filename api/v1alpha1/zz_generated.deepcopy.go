//go:build !ignore_autogenerated

// Copyright: This file is part of korrel8r, released under https://github.com/korrel8r/korrel8r/blob/main/LICENSE"

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/korrel8r/korrel8r/pkg/config"
	korrel8r "github.com/korrel8r/korrel8r/pkg/korrel8r"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]config.Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]config.Group, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Stores != nil {
		in, out := &in.Stores, &out.Stores
		*out = make([]korrel8r.StoreConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(korrel8r.StoreConfig, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DebugSpec) DeepCopyInto(out *DebugSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DebugSpec.
func (in *DebugSpec) DeepCopy() *DebugSpec {
	if in == nil {
		return nil
	}
	out := new(DebugSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Korrel8r) DeepCopyInto(out *Korrel8r) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Korrel8r.
func (in *Korrel8r) DeepCopy() *Korrel8r {
	if in == nil {
		return nil
	}
	out := new(Korrel8r)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Korrel8r) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Korrel8rList) DeepCopyInto(out *Korrel8rList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Korrel8r, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Korrel8rList.
func (in *Korrel8rList) DeepCopy() *Korrel8rList {
	if in == nil {
		return nil
	}
	out := new(Korrel8rList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Korrel8rList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Korrel8rSpec) DeepCopyInto(out *Korrel8rSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(Config)
		(*in).DeepCopyInto(*out)
	}
	if in.Debug != nil {
		in, out := &in.Debug, &out.Debug
		*out = new(DebugSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Korrel8rSpec.
func (in *Korrel8rSpec) DeepCopy() *Korrel8rSpec {
	if in == nil {
		return nil
	}
	out := new(Korrel8rSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Korrel8rStatus) DeepCopyInto(out *Korrel8rStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Korrel8rStatus.
func (in *Korrel8rStatus) DeepCopy() *Korrel8rStatus {
	if in == nil {
		return nil
	}
	out := new(Korrel8rStatus)
	in.DeepCopyInto(out)
	return out
}
