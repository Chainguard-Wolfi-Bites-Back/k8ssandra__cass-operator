// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1beta1

import (
	json "encoding/json"

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalVolumes) DeepCopyInto(out *AdditionalVolumes) {
	*out = *in
	in.PVCSpec.DeepCopyInto(&out.PVCSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalVolumes.
func (in *AdditionalVolumes) DeepCopy() *AdditionalVolumes {
	if in == nil {
		return nil
	}
	out := new(AdditionalVolumes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AdditionalVolumesSlice) DeepCopyInto(out *AdditionalVolumesSlice) {
	{
		in := &in
		*out = make(AdditionalVolumesSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalVolumesSlice.
func (in AdditionalVolumesSlice) DeepCopy() AdditionalVolumesSlice {
	if in == nil {
		return nil
	}
	out := new(AdditionalVolumesSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraDatacenter) DeepCopyInto(out *CassandraDatacenter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraDatacenter.
func (in *CassandraDatacenter) DeepCopy() *CassandraDatacenter {
	if in == nil {
		return nil
	}
	out := new(CassandraDatacenter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraDatacenter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraDatacenterList) DeepCopyInto(out *CassandraDatacenterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CassandraDatacenter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraDatacenterList.
func (in *CassandraDatacenterList) DeepCopy() *CassandraDatacenterList {
	if in == nil {
		return nil
	}
	out := new(CassandraDatacenterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraDatacenterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraDatacenterSpec) DeepCopyInto(out *CassandraDatacenterSpec) {
	*out = *in
	if in.DockerImageRunsAsCassandra != nil {
		in, out := &in.DockerImageRunsAsCassandra, &out.DockerImageRunsAsCassandra
		*out = new(bool)
		**out = **in
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	in.ManagementApiAuth.DeepCopyInto(&out.ManagementApiAuth)
	if in.NodeAffinityLabels != nil {
		in, out := &in.NodeAffinityLabels, &out.NodeAffinityLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	in.SystemLoggerResources.DeepCopyInto(&out.SystemLoggerResources)
	in.ConfigBuilderResources.DeepCopyInto(&out.ConfigBuilderResources)
	if in.Racks != nil {
		in, out := &in.Racks, &out.Racks
		*out = make([]Rack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.StorageConfig.DeepCopyInto(&out.StorageConfig)
	if in.ReplaceNodes != nil {
		in, out := &in.ReplaceNodes, &out.ReplaceNodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ForceUpgradeRacks != nil {
		in, out := &in.ForceUpgradeRacks, &out.ForceUpgradeRacks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DseWorkloads != nil {
		in, out := &in.DseWorkloads, &out.DseWorkloads
		*out = new(DseWorkloads)
		**out = **in
	}
	if in.PodTemplateSpec != nil {
		in, out := &in.PodTemplateSpec, &out.PodTemplateSpec
		*out = new(v1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]CassandraUser, len(*in))
		copy(*out, *in)
	}
	if in.Networking != nil {
		in, out := &in.Networking, &out.Networking
		*out = new(NetworkingConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalSeeds != nil {
		in, out := &in.AdditionalSeeds, &out.AdditionalSeeds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Reaper != nil {
		in, out := &in.Reaper, &out.Reaper
		*out = new(ReaperConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraDatacenterSpec.
func (in *CassandraDatacenterSpec) DeepCopy() *CassandraDatacenterSpec {
	if in == nil {
		return nil
	}
	out := new(CassandraDatacenterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraDatacenterStatus) DeepCopyInto(out *CassandraDatacenterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]DatacenterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.SuperUserUpserted.DeepCopyInto(&out.SuperUserUpserted)
	in.UsersUpserted.DeepCopyInto(&out.UsersUpserted)
	in.LastServerNodeStarted.DeepCopyInto(&out.LastServerNodeStarted)
	in.LastRollingRestart.DeepCopyInto(&out.LastRollingRestart)
	if in.NodeStatuses != nil {
		in, out := &in.NodeStatuses, &out.NodeStatuses
		*out = make(CassandraStatusMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeReplacements != nil {
		in, out := &in.NodeReplacements, &out.NodeReplacements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.QuietPeriod.DeepCopyInto(&out.QuietPeriod)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraDatacenterStatus.
func (in *CassandraDatacenterStatus) DeepCopy() *CassandraDatacenterStatus {
	if in == nil {
		return nil
	}
	out := new(CassandraDatacenterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraNodeStatus) DeepCopyInto(out *CassandraNodeStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraNodeStatus.
func (in *CassandraNodeStatus) DeepCopy() *CassandraNodeStatus {
	if in == nil {
		return nil
	}
	out := new(CassandraNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in CassandraStatusMap) DeepCopyInto(out *CassandraStatusMap) {
	{
		in := &in
		*out = make(CassandraStatusMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraStatusMap.
func (in CassandraStatusMap) DeepCopy() CassandraStatusMap {
	if in == nil {
		return nil
	}
	out := new(CassandraStatusMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraUser) DeepCopyInto(out *CassandraUser) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraUser.
func (in *CassandraUser) DeepCopy() *CassandraUser {
	if in == nil {
		return nil
	}
	out := new(CassandraUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatacenterCondition) DeepCopyInto(out *DatacenterCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatacenterCondition.
func (in *DatacenterCondition) DeepCopy() *DatacenterCondition {
	if in == nil {
		return nil
	}
	out := new(DatacenterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DseWorkloads) DeepCopyInto(out *DseWorkloads) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DseWorkloads.
func (in *DseWorkloads) DeepCopy() *DseWorkloads {
	if in == nil {
		return nil
	}
	out := new(DseWorkloads)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementApiAuthConfig) DeepCopyInto(out *ManagementApiAuthConfig) {
	*out = *in
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		*out = new(ManagementApiAuthInsecureConfig)
		**out = **in
	}
	if in.Manual != nil {
		in, out := &in.Manual, &out.Manual
		*out = new(ManagementApiAuthManualConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementApiAuthConfig.
func (in *ManagementApiAuthConfig) DeepCopy() *ManagementApiAuthConfig {
	if in == nil {
		return nil
	}
	out := new(ManagementApiAuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementApiAuthInsecureConfig) DeepCopyInto(out *ManagementApiAuthInsecureConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementApiAuthInsecureConfig.
func (in *ManagementApiAuthInsecureConfig) DeepCopy() *ManagementApiAuthInsecureConfig {
	if in == nil {
		return nil
	}
	out := new(ManagementApiAuthInsecureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementApiAuthManualConfig) DeepCopyInto(out *ManagementApiAuthManualConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementApiAuthManualConfig.
func (in *ManagementApiAuthManualConfig) DeepCopy() *ManagementApiAuthManualConfig {
	if in == nil {
		return nil
	}
	out := new(ManagementApiAuthManualConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkingConfig) DeepCopyInto(out *NetworkingConfig) {
	*out = *in
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(NodePortConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkingConfig.
func (in *NetworkingConfig) DeepCopy() *NetworkingConfig {
	if in == nil {
		return nil
	}
	out := new(NetworkingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePortConfig) DeepCopyInto(out *NodePortConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePortConfig.
func (in *NodePortConfig) DeepCopy() *NodePortConfig {
	if in == nil {
		return nil
	}
	out := new(NodePortConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rack) DeepCopyInto(out *Rack) {
	*out = *in
	if in.NodeAffinityLabels != nil {
		in, out := &in.NodeAffinityLabels, &out.NodeAffinityLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rack.
func (in *Rack) DeepCopy() *Rack {
	if in == nil {
		return nil
	}
	out := new(Rack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReaperConfig) DeepCopyInto(out *ReaperConfig) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReaperConfig.
func (in *ReaperConfig) DeepCopy() *ReaperConfig {
	if in == nil {
		return nil
	}
	out := new(ReaperConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageConfig) DeepCopyInto(out *StorageConfig) {
	*out = *in
	if in.CassandraDataVolumeClaimSpec != nil {
		in, out := &in.CassandraDataVolumeClaimSpec, &out.CassandraDataVolumeClaimSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalVolumes != nil {
		in, out := &in.AdditionalVolumes, &out.AdditionalVolumes
		*out = make(AdditionalVolumesSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageConfig.
func (in *StorageConfig) DeepCopy() *StorageConfig {
	if in == nil {
		return nil
	}
	out := new(StorageConfig)
	in.DeepCopyInto(out)
	return out
}
