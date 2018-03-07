// Copyright © 2017 The Kubicorn Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cluster

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "k8s.io/kube-deploy/cluster-api/api/cluster/v1alpha1"
)

const (
	CloudAmazon       = "amazon"
	CloudAzure        = "azure"
	CloudGoogle       = "google"
	CloudBaremetal    = "baremetal"
	CloudDigitalOcean = "digitalocean"
	CloudOVH          = "ovh"
	CloudPacket       = "packet"
)

type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Name              string         `json:"name,omitempty"`
	Project           *Project       `json:"project,omitempty"`
	CloudId           string         `json:"cloudId,omitempty"`
	ServerPools       []*ServerPool  `json:"serverPools,omitempty"`
	Cloud             string         `json:"cloud,omitempty"`
	Location          string         `json:"location,omitempty"`
	SSH               *SSH           `json:"SSH,omitempty"`
	Network           *Network       `json:"network,omitempty"`
	Values            *Values        `json:"values,omitempty"`
	KubernetesAPI     *KubernetesAPI `json:"kubernetesAPI,omitempty"`
	GroupIdentifier   string         `json:"groupIdentifier,omitempty"`
	Components        *Components    `json:"components,omitempty"`

	// Add the official Kubernetes Cluster API for us to begin transitioning (lol) to the API
	ClusterAPI *clusterv1.Cluster
}

func NewCluster(name string) *Cluster {
	return &Cluster{
		Name: name,
	}
}
