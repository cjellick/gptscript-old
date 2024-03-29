//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package engine

import (
	"github.com/acorn-io/gptscript/pkg/types"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *State) DeepCopyInto(out *State) {
	*out = *in
	in.Completion.DeepCopyInto(&out.Completion)
	if in.Pending != nil {
		in, out := &in.Pending, &out.Pending
		*out = make(map[string]types.CompletionToolCall, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = make(map[string]CallResult, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new State.
func (in *State) DeepCopy() *State {
	if in == nil {
		return nil
	}
	out := new(State)
	in.DeepCopyInto(out)
	return out
}
