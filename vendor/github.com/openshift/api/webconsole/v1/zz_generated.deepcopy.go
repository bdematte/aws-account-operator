// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertInfo) DeepCopyInto(out *CertInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertInfo.
func (in *CertInfo) DeepCopy() *CertInfo {
	if in == nil {
		return nil
	}
	out := new(CertInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfo) DeepCopyInto(out *ClusterInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfo.
func (in *ClusterInfo) DeepCopy() *ClusterInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionsConfiguration) DeepCopyInto(out *ExtensionsConfiguration) {
	*out = *in
	if in.ScriptURLs != nil {
		in, out := &in.ScriptURLs, &out.ScriptURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StylesheetURLs != nil {
		in, out := &in.StylesheetURLs, &out.StylesheetURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionsConfiguration.
func (in *ExtensionsConfiguration) DeepCopy() *ExtensionsConfiguration {
	if in == nil {
		return nil
	}
	out := new(ExtensionsConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeaturesConfiguration) DeepCopyInto(out *FeaturesConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeaturesConfiguration.
func (in *FeaturesConfiguration) DeepCopy() *FeaturesConfiguration {
	if in == nil {
		return nil
	}
	out := new(FeaturesConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPServingInfo) DeepCopyInto(out *HTTPServingInfo) {
	*out = *in
	in.ServingInfo.DeepCopyInto(&out.ServingInfo)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPServingInfo.
func (in *HTTPServingInfo) DeepCopy() *HTTPServingInfo {
	if in == nil {
		return nil
	}
	out := new(HTTPServingInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedCertificate) DeepCopyInto(out *NamedCertificate) {
	*out = *in
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.CertInfo = in.CertInfo
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedCertificate.
func (in *NamedCertificate) DeepCopy() *NamedCertificate {
	if in == nil {
		return nil
	}
	out := new(NamedCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingInfo) DeepCopyInto(out *ServingInfo) {
	*out = *in
	out.CertInfo = in.CertInfo
	if in.NamedCertificates != nil {
		in, out := &in.NamedCertificates, &out.NamedCertificates
		*out = make([]NamedCertificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CipherSuites != nil {
		in, out := &in.CipherSuites, &out.CipherSuites
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingInfo.
func (in *ServingInfo) DeepCopy() *ServingInfo {
	if in == nil {
		return nil
	}
	out := new(ServingInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebConsoleConfiguration) DeepCopyInto(out *WebConsoleConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ServingInfo.DeepCopyInto(&out.ServingInfo)
	out.ClusterInfo = in.ClusterInfo
	out.Features = in.Features
	in.Extensions.DeepCopyInto(&out.Extensions)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebConsoleConfiguration.
func (in *WebConsoleConfiguration) DeepCopy() *WebConsoleConfiguration {
	if in == nil {
		return nil
	}
	out := new(WebConsoleConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebConsoleConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}
