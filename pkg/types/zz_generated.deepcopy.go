//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package types

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompletionFunctionCall) DeepCopyInto(out *CompletionFunctionCall) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompletionFunctionCall.
func (in *CompletionFunctionCall) DeepCopy() *CompletionFunctionCall {
	if in == nil {
		return nil
	}
	out := new(CompletionFunctionCall)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompletionFunctionDefinition) DeepCopyInto(out *CompletionFunctionDefinition) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(JSONSchema)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompletionFunctionDefinition.
func (in *CompletionFunctionDefinition) DeepCopy() *CompletionFunctionDefinition {
	if in == nil {
		return nil
	}
	out := new(CompletionFunctionDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompletionMessage) DeepCopyInto(out *CompletionMessage) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = make([]ContentPart, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToolCall != nil {
		in, out := &in.ToolCall, &out.ToolCall
		*out = new(CompletionToolCall)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompletionMessage.
func (in *CompletionMessage) DeepCopy() *CompletionMessage {
	if in == nil {
		return nil
	}
	out := new(CompletionMessage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompletionRequest) DeepCopyInto(out *CompletionRequest) {
	*out = *in
	if in.Tools != nil {
		in, out := &in.Tools, &out.Tools
		*out = make([]CompletionTool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Messages != nil {
		in, out := &in.Messages, &out.Messages
		*out = make([]CompletionMessage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompletionRequest.
func (in *CompletionRequest) DeepCopy() *CompletionRequest {
	if in == nil {
		return nil
	}
	out := new(CompletionRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompletionTool) DeepCopyInto(out *CompletionTool) {
	*out = *in
	in.Function.DeepCopyInto(&out.Function)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompletionTool.
func (in *CompletionTool) DeepCopy() *CompletionTool {
	if in == nil {
		return nil
	}
	out := new(CompletionTool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompletionToolCall) DeepCopyInto(out *CompletionToolCall) {
	*out = *in
	if in.Index != nil {
		in, out := &in.Index, &out.Index
		*out = new(int)
		**out = **in
	}
	out.Function = in.Function
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompletionToolCall.
func (in *CompletionToolCall) DeepCopy() *CompletionToolCall {
	if in == nil {
		return nil
	}
	out := new(CompletionToolCall)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentPart) DeepCopyInto(out *ContentPart) {
	*out = *in
	if in.ToolCall != nil {
		in, out := &in.ToolCall, &out.ToolCall
		*out = new(CompletionToolCall)
		(*in).DeepCopyInto(*out)
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageURL)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentPart.
func (in *ContentPart) DeepCopy() *ContentPart {
	if in == nil {
		return nil
	}
	out := new(ContentPart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageURL) DeepCopyInto(out *ImageURL) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageURL.
func (in *ImageURL) DeepCopy() *ImageURL {
	if in == nil {
		return nil
	}
	out := new(ImageURL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONSchema) DeepCopyInto(out *JSONSchema) {
	*out = *in
	in.Property.DeepCopyInto(&out.Property)
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]Property, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Defs != nil {
		in, out := &in.Defs, &out.Defs
		*out = make(map[string]JSONSchema, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONSchema.
func (in *JSONSchema) DeepCopy() *JSONSchema {
	if in == nil {
		return nil
	}
	out := new(JSONSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Property) DeepCopyInto(out *Property) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JSONSchema, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Property.
func (in *Property) DeepCopy() *Property {
	if in == nil {
		return nil
	}
	out := new(Property)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tool) DeepCopyInto(out *Tool) {
	*out = *in
	if in.Arguments != nil {
		in, out := &in.Arguments, &out.Arguments
		*out = new(JSONSchema)
		(*in).DeepCopyInto(*out)
	}
	if in.Tools != nil {
		in, out := &in.Tools, &out.Tools
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tool.
func (in *Tool) DeepCopy() *Tool {
	if in == nil {
		return nil
	}
	out := new(Tool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ToolSet) DeepCopyInto(out *ToolSet) {
	{
		in := &in
		*out = make(ToolSet, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToolSet.
func (in ToolSet) DeepCopy() ToolSet {
	if in == nil {
		return nil
	}
	out := new(ToolSet)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Type) DeepCopyInto(out *Type) {
	{
		in := &in
		*out = make(Type, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Type.
func (in Type) DeepCopy() Type {
	if in == nil {
		return nil
	}
	out := new(Type)
	in.DeepCopyInto(out)
	return *out
}
