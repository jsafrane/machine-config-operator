package v1

import (
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	kubeletconfigv1beta1 "k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1beta1"
)

// CustomResourceDefinition for MCOConfig
// apiVersion: apiextensions.k8s.io/v1beta1
// kind: CustomResourceDefinition
// metadata:
//   # name must match the spec fields below, and be in the form: <plural>.<group>
//   name: mcoconfigs.machineconfiguration.openshift.io
// spec:
//   # group name to use for REST API: /apis/<group>/<version>
//   group: machineconfiguration.openshift.io
//   # list of versions supported by this CustomResourceDefinition
//   versions:
//     - name: v1
//       # Each version can be enabled/disabled by Served flag.
//       served: true
//       # One and only one version must be marked as the storage version.
//       storage: true
//   # either Namespaced or Cluster
//   scope: Namespaced
//   names:
//     # plural name to be used in the URL: /apis/<group>/<version>/<plural>
//     plural: mcoconfigs
//     # singular name to be used as an alias on the CLI and for display
//     singular: mcoconfig
//     # kind is normally the CamelCased singular type. Your resource manifests use this.
//     kind: MCOConfig
//     # shortNames allow shorter string to match your resource on the CLI
//     shortNames:
//

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MCOConfig describes configuration for MachineConfigOperator.
type MCOConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MCOConfigSpec `json:"spec"`
}

// MCOConfigSpec is the spec for MCOConfig resource.
type MCOConfigSpec struct {
	ClusterDNSIP        string `json:"clusterDNSIP"`
	CloudProviderConfig string `json:"cloudProviderConfig"`
	ClusterName         string `json:"clusterName"`

	// The openshift platform, e.g. "libvirt", "openstack", "aws", or "none"
	Platform string `json:"platform"`

	BaseDomain string `json:"baseDomain"`

	SSHKey string `json:"sshKey"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MCOConfigList is a list of MCOConfig resources
type MCOConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MCOConfig `json:"items"`
}

// CustomResourceDefinition for ControllerConfig
// apiVersion: apiextensions.k8s.io/v1beta1
// kind: CustomResourceDefinition
// metadata:
//   # name must match the spec fields below, and be in the form: <plural>.<group>
//   name: controllerconfigs.machineconfiguration.openshift.io
// spec:
//   # group name to use for REST API: /apis/<group>/<version>
//   group: machineconfiguration.openshift.io
//   # list of versions supported by this CustomResourceDefinition
//   versions:
//     - name: v1
//       # Each version can be enabled/disabled by Served flag.
//       served: true
//       # One and only one version must be marked as the storage version.
//       storage: true
//   # either Namespaced or Cluster
//   scope: Namespaced
//   names:
//     # plural name to be used in the URL: /apis/<group>/<version>/<plural>
//     plural: controllerconfigs
//     # singular name to be used as an alias on the CLI and for display
//     singular: controllerconfig
//     # kind is normally the CamelCased singular type. Your resource manifests use this.
//     kind: ControllerConfig
//     # shortNames allow shorter string to match your resource on the CLI
//     shortNames:
//

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerConfig describes configuration for MachineConfigController.
type ControllerConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ControllerConfigSpec `json:"spec"`
}

// ControllerConfigSpec is the spec for ControllerConfig resource.
type ControllerConfigSpec struct {
	ClusterDNSIP        string `json:"clusterDNSIP"`
	CloudProviderConfig string `json:"cloudProviderConfig"`
	ClusterName         string `json:"clusterName"`

	// The openshift platform, e.g. "libvirt", "openstack", "aws", or "none"
	Platform string `json:"platform"`

	BaseDomain string `json:"baseDomain"`

	// CAs
	EtcdCAData []byte `json:"etcdCAData"`
	RootCAData []byte `json:"rootCAData"`

	// PullSecret is the default pull secret that needs to be installed
	// on all machines.
	PullSecret *corev1.ObjectReference `json:"pullSecret,omitempty"`

	// Public SSH
	SSHKey string `json:"sshKey"`

	// Images is map of images that are used by the controller.
	Images map[string]string `json:"images"`

	// Sourced from configmap/machine-config-osimageurl
	OSImageURL string `json:"osImageURL`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerConfigList is a list of ControllerConfig resources
type ControllerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ControllerConfig `json:"items"`
}

// CustomResourceDefinition for MachineConfig
// apiVersion: apiextensions.k8s.io/v1beta1
// kind: CustomResourceDefinition
// metadata:
//   # name must match the spec fields below, and be in the form: <plural>.<group>
//   name: machineconfigs.machineconfiguration.openshift.io
// spec:
//   # group name to use for REST API: /apis/<group>/<version>
//   group: machineconfiguration.openshift.io
//   # list of versions supported by this CustomResourceDefinition
//   versions:
//     - name: v1
//       # Each version can be enabled/disabled by Served flag.
//       served: true
//       # One and only one version must be marked as the storage version.
//       storage: true
//   # either Namespaced or Cluster
//   scope: Cluster
//   names:
//     # plural name to be used in the URL: /apis/<group>/<version>/<plural>
//     plural: machineconfigs
//     # singular name to be used as an alias on the CLI and for display
//     singular: machineconfig
//     # kind is normally the CamelCased singular type. Your resource manifests use this.
//     kind: MachineConfig
//     # shortNames allow shorter string to match your resource on the CLI
//     shortNames:
//     - mc

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen=false

// MachineConfig defines the configuration for a machine
type MachineConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MachineConfigSpec `json:"spec"`
}

// MachineConfigSpec is the for MachineConfig
type MachineConfigSpec struct {
	// OSImageURL specifies the remote location that will be used to
	// fetch the OS.
	OSImageURL string `json:"osImageURL"`
	// Config is a Ignition Config object.
	Config ignv2_2types.Config `json:"config"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineConfigList is a list of MachineConfig resources
type MachineConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MachineConfig `json:"items"`
}

// CustomResourceDefinition for MachineConfigPool
// apiVersion: apiextensions.k8s.io/v1beta1
// kind: CustomResourceDefinition
// metadata:
//   # name must match the spec fields below, and be in the form: <plural>.<group>
//   name: machineconfigpools.machineconfiguration.openshift.io
// spec:
//   # group name to use for REST API: /apis/<group>/<version>
//   group: machineconfiguration.openshift.io
//   # list of versions supported by this CustomResourceDefinition
//   versions:
//     - name: v1
//       # Each version can be enabled/disabled by Served flag.
//       served: true
//       # One and only one version must be marked as the storage version.
//       storage: true
//   # either Namespaced or Cluster
//   scope: Cluster
//   names:
//     # plural name to be used in the URL: /apis/<group>/<version>/<plural>
//     plural: machineconfigpools
//     # singular name to be used as an alias on the CLI and for display
//     singular: machineconfigpool
//     # kind is normally the CamelCased singular type. Your resource manifests use this.
//     kind: MachineConfigPool
//     # shortNames allow shorter string to match your resource on the CLI
//     shortNames:
//

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineConfigPool describes a pool of MachineConfigs.
type MachineConfigPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MachineConfigPoolSpec   `json:"spec"`
	Status MachineConfigPoolStatus `json:"status"`
}

// MachineConfigPoolSpec is the spec for MachineConfigPool resource.
type MachineConfigPoolSpec struct {
	// Label selector for MachineConfigs.
	// Refer https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ on how label and selectors work.
	MachineConfigSelector *metav1.LabelSelector `json:"machineConfigSelector,omitempty"`

	// Label selector for Machines.
	MachineSelector *metav1.LabelSelector `json:"machineSelector,omitempty"`

	// If true, changes to this machine pool should be stopped.
	// This includes generating new desiredMachineConfig and update of machines.
	Paused bool `json:"paused"`

	// MaxUnavailable specifies the percentage or constant number of machines that can be updating at any given time.
	// default is 1.
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable"`
}

// MachineConfigPoolStatus is the status for MachineConfigPool resource.
type MachineConfigPoolStatus struct {
	// The generation observed by the controller.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// The current MachineConfig object for the machine pool.
	Configuration MachineConfigPoolStatusConfiguration `json:"configuration"`

	// Total number of machines in the machine pool.
	MachineCount int32 `json:"machineCount"`

	// Total number of machines targeted by the pool that have the CurrentMachineConfig as their config.
	UpdatedMachineCount int32 `json:"updatedMachineCount"`

	// Total number of ready machines targeted by the pool.
	ReadyMachineCount int32 `json:"readyMachineCount"`

	// Total number of unavailable (non-ready) machines targeted by the pool.
	// A node is marked unavailable if it is in updating state or NodeReady condition is false.
	UnavailableMachineCount int32 `json:"unavailableMachineCount"`

	// Represents the latest available observations of current state.
	Conditions []MachineConfigPoolCondition `json:"conditions"`
}

// MachineConfigPoolStatusConfiguration stores the current configuration for the pool, and
// optionally also stores the list of machineconfig objects used to generate the configuration.
type MachineConfigPoolStatusConfiguration struct {
	corev1.ObjectReference

	// source is the list of machineconfigs that were used to generate the single machineconfig object specified in `content`.
	// +optional
	Source []corev1.ObjectReference `json:"source,omitempty"`
}

// MachineConfigPoolCondition contains condition information for an MachineConfigPool.
type MachineConfigPoolCondition struct {
	// Type of the condition, currently ('Done', 'Updating', 'Failed').
	Type MachineConfigPoolConditionType `json:"type"`

	// Status of the condition, one of ('True', 'False', 'Unknown').
	Status corev1.ConditionStatus `json:"status"`

	// LastTransitionTime is the timestamp corresponding to the last status
	// change of this condition.
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`

	// Reason is a brief machine readable explanation for the condition's last
	// transition.
	Reason string `json:"reason"`

	// Message is a human readable description of the details of the last
	// transition, complementing reason.
	Message string `json:"message"`
}

// MachineConfigPoolConditionType valid conditions of a machineconfigpool
type MachineConfigPoolConditionType string

const (
	// MachineConfigPoolUpdated means machineconfigpool is updated completely.
	// When the all the machines in the pool are updated to the correct machine config.
	MachineConfigPoolUpdated MachineConfigPoolConditionType = "Updated"
	// MachineConfigPoolUpdating means machineconfigpool is updating.
	// When at least one of machine is not either not updated or is in the process of updating
	// to the desired machine config.
	MachineConfigPoolUpdating MachineConfigPoolConditionType = "Updating"
	// MachineConfigPoolDegraded means the update for one of the machine is not progressing
	// due to an error
	MachineConfigPoolDegraded MachineConfigPoolConditionType = "Degraded"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineConfigPoolList is a list of MachineConfigPool resources
type MachineConfigPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MachineConfigPool `json:"items"`
}

// CustomResourceDefinition for KubeletConfig
// apiVersion: apiextensions.k8s.io/v1beta1
// kind: CustomResourceDefinition
// metadata:
//   # name must match the spec fields below, and be in the form: <plural>.<group>
//   name: kubeletconfigs.machineconfiguration.openshift.io
// spec:
//   # group name to use for REST API: /apis/<group>/<version>
//   group: machineconfiguration.openshift.io
//   # list of versions supported by this CustomResourceDefinition
//   versions:
//     - name: v1
//       # Each version can be enabled/disabled by Served flag.
//       served: true
//       # One and only one version must be marked as the storage version.
//       storage: true
//   # either Namespaced or Cluster
//   scope: Cluster
//   names:
//     # plural name to be used in the URL: /apis/<group>/<version>/<plural>
//     plural: kubeletconfigs
//     # singular name to be used as an alias on the CLI and for display
//     singular: kubeletconfig
//     # kind is normally the CamelCased singular type. Your resource manifests use this.
//     kind: KubeletConfig
//     # shortNames allow shorter string to match your resource on the CLI
//     shortNames:
//

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeletConfig describes a customized Kubelet configuration.
type KubeletConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeletConfigSpec   `json:"spec,omitempty"`
	Status KubeletConfigStatus `json:"status,omitempty"`
}

// KubeletConfigSpec defines the desired state of KubeletConfig
type KubeletConfigSpec struct {
	MachineConfigPoolSelector *metav1.LabelSelector                      `json:"machineConfigPoolSelector,omitempty"`
	KubeletConfig             *kubeletconfigv1beta1.KubeletConfiguration `json:"kubeletConfig,omitempty"`
}

// KubeletConfigStatus defines the observed state of a KubeletConfig
type KubeletConfigStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeletConfigList is a list of KubeletConfig resources
type KubeletConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []KubeletConfig `json:"items"`
}
