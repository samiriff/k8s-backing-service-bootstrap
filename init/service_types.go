/*
Copyright 2018 samiriff.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// <BACKING_SERVICE>FeatureSetSpec to define sidecars for additional functionality
type <BACKING_SERVICE>FeatureSetSpec struct {
	Watcher         bool `json:"watcher,omitempty"`
	Logger          bool `json:"logger,omitempty"`
	Monitor         bool `json:"monitor,omitempty"`
	DiskFullChecker bool `json:"disk-full-checker,omitempty"`
	PITR            bool `json:"pitr,omitempty"`
}

// <BACKING_SERVICE>VersionUpgradeStatus to define sidecars for additional functionality
type <BACKING_SERVICE>VersionUpgradeStatus struct {
	CurrentVersion string `json:"current-version,omitempty"`

	// +kubebuilder:validation:Enum=InProgress,Succeeded,Failed
	Status string `json:"status,omitempty"`
}

// <BACKING_SERVICE>Spec defines the desired state of <BACKING_SERVICE>
type <BACKING_SERVICE>Spec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Version string `json:"version"`

	// +optional
	HighAvailability bool `json:"high-availability"`

	// +optional
	MultiAZ bool `json:"multi-az"`

	FeatureSet <BACKING_SERVICE>FeatureSetSpec `json:"feature-set,omitempty"`
}

// <BACKING_SERVICE>Status defines the observed state of <BACKING_SERVICE>
type <BACKING_SERVICE>Status struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Enum=InProgress,Succeeded,Failed
	Provisioning string `json:"provisioning,omitempty"`

	// +kubebuilder:validation:Enum=InProgress,Succeeded,Failed
	Deprovisioning string `json:"deprovisioning,omitempty"`

	VersionUpgrade <BACKING_SERVICE>VersionUpgradeStatus `json:"version-upgrade,omitempty"`

	// +kubebuilder:validation:Enum=InProgress,Succeeded,Failed
	MultiAZ string `json:"multi-az,omitempty"`

	// +kubebuilder:validation:Enum=BelowThreshold,AboveThreshold
	DiskFull string `json:"disk-full,omitempty"`

	LastOperation string `json:"last-operation"`

	LastMessage string `json:"last-error"`

	LastLogAt string `json:"last-log-at"`

	LastMonitorAt string `json:"last-monitor-at"`

	Volumes string `json:"volumes"`

	ScratchPad string `json:"scratch-pad"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// <BACKING_SERVICE> is the Schema for the <BACKING_SERVICE>s API
// +k8s:openapi-gen=true
type <BACKING_SERVICE> struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   <BACKING_SERVICE>Spec   `json:"spec,omitempty"`
	Status <BACKING_SERVICE>Status `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// <BACKING_SERVICE>List contains a list of <BACKING_SERVICE>
type <BACKING_SERVICE>List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []<BACKING_SERVICE> `json:"items"`
}

func init() {
	SchemeBuilder.Register(&<BACKING_SERVICE>{}, &<BACKING_SERVICE>List{})
}
