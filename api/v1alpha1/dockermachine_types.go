/*
Copyright 2023.

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
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

const (
	// MachineFinalizer allows cleaning up resources associated with
	// DockerMachine before removing it from the API Server.
	MachineFinalizer = "dockermachine.infrastructure.cluster.x-k8s.io"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DockerMachineSpec defines the desired state of DockerMachine
type DockerMachineSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// ProviderID is the identifier for the DockerMachine instance
	ProviderID *string `json:"providerID,omitempty"`

	// CustomImage allows customizing the container image that is used for
	// running the machine
	// +optional
	CustomImage string `json:"customImage,omitempty"`

	// Bootstrapped is true when the kubeadm bootstrapping has been run
	// against this machine
	// +optional
	Bootstrapped bool `json:"bootstrapped,omitempty"`
}

// DockerMachineStatus defines the observed state of DockerMachine
type DockerMachineStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Ready indicates the docker infrastructure has been provisioned and is ready
	// +optional
	Ready bool `json:"ready"`

	// Conditions defines current service state of the DockerMachine.
	// +optional
	Conditions clusterv1.Conditions `json:"conditions,omitempty"`

	// LoadBalancerConfigured denotes that the machine has been
	// added to the load balancer
	// +optional
	LoadBalancerConfigured bool `json:"loadBalancerConfigured"`

	// Addresses contains the associated addresses for the docker machine.
	// +optional
	Addresses []clusterv1.MachineAddress `json:"addresses,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DockerMachine is the Schema for the dockermachines API
type DockerMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DockerMachineSpec   `json:"spec,omitempty"`
	Status DockerMachineStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DockerMachineList contains a list of DockerMachine
type DockerMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DockerMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DockerMachine{}, &DockerMachineList{})
}

// GetConditions returns the conditions of ByoMachine status
func (dockerMachine *DockerMachine) GetConditions() clusterv1.Conditions {
	return dockerMachine.Status.Conditions
}

// SetConditions sets the conditions of ByoMachine status
func (dockerMachine *DockerMachine) SetConditions(conditions clusterv1.Conditions) {
	dockerMachine.Status.Conditions = conditions
}
