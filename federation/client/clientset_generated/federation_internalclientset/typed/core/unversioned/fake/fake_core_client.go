/*
Copyright 2016 The Kubernetes Authors.

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

package fake

import (
	unversioned "k8s.io/kubernetes/federation/client/clientset_generated/federation_internalclientset/typed/core/unversioned"
	restclient "k8s.io/kubernetes/pkg/client/restclient"
	core "k8s.io/kubernetes/pkg/client/testing/core"
)

type FakeCore struct {
	*core.Fake
}

func (c *FakeCore) ConfigMaps(namespace string) unversioned.ConfigMapInterface {
	return &FakeConfigMaps{c, namespace}
}

func (c *FakeCore) Events(namespace string) unversioned.EventInterface {
	return &FakeEvents{c, namespace}
}

func (c *FakeCore) Namespaces() unversioned.NamespaceInterface {
	return &FakeNamespaces{c}
}

func (c *FakeCore) Secrets(namespace string) unversioned.SecretInterface {
	return &FakeSecrets{c, namespace}
}

func (c *FakeCore) Services(namespace string) unversioned.ServiceInterface {
	return &FakeServices{c, namespace}
}

// GetRESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCore) GetRESTClient() *restclient.RESTClient {
	return nil
}
