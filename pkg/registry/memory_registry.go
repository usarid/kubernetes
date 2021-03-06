/*
Copyright 2014 Google Inc. All rights reserved.

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
package registry

import (
	. "github.com/GoogleCloudPlatform/kubernetes/pkg/api"
)

// An implementation of PodRegistry and ControllerRegistry that is backed by memory
// Mainly used for testing.
type MemoryRegistry struct {
	podData       map[string]Pod
	controllerData map[string]ReplicationController
	serviceData    map[string]Service
}

func MakeMemoryRegistry() *MemoryRegistry {
	return &MemoryRegistry{
		podData:       map[string]Pod{},
		controllerData: map[string]ReplicationController{},
		serviceData:    map[string]Service{},
	}
}

func (registry *MemoryRegistry) ListPods(labelQuery *map[string]string) ([]Pod, error) {
	result := []Pod{}
	for _, value := range registry.podData {
		if LabelsMatch(value, labelQuery) {
			result = append(result, value)
		}
	}
	return result, nil
}

func (registry *MemoryRegistry) GetPod(podID string) (*Pod, error) {
	pod, found := registry.podData[podID]
	if found {
		return &pod, nil
	} else {
		return nil, nil
	}
}

func (registry *MemoryRegistry) CreatePod(machine string, pod Pod) error {
	registry.podData[pod.ID] = pod
	return nil
}

func (registry *MemoryRegistry) DeletePod(podID string) error {
	delete(registry.podData, podID)
	return nil
}

func (registry *MemoryRegistry) UpdatePod(pod Pod) error {
	registry.podData[pod.ID] = pod
	return nil
}

func (registry *MemoryRegistry) ListControllers() ([]ReplicationController, error) {
	result := []ReplicationController{}
	for _, value := range registry.controllerData {
		result = append(result, value)
	}
	return result, nil
}

func (registry *MemoryRegistry) GetController(controllerID string) (*ReplicationController, error) {
	controller, found := registry.controllerData[controllerID]
	if found {
		return &controller, nil
	} else {
		return nil, nil
	}
}

func (registry *MemoryRegistry) CreateController(controller ReplicationController) error {
	registry.controllerData[controller.ID] = controller
	return nil
}

func (registry *MemoryRegistry) DeleteController(controllerId string) error {
	delete(registry.controllerData, controllerId)
	return nil
}

func (registry *MemoryRegistry) UpdateController(controller ReplicationController) error {
	registry.controllerData[controller.ID] = controller
	return nil
}

func (registry *MemoryRegistry) ListServices() (ServiceList, error) {
	var list []Service
	for _, value := range registry.serviceData {
		list = append(list, value)
	}
	return ServiceList{Items: list}, nil
}

func (registry *MemoryRegistry) CreateService(svc Service) error {
	registry.serviceData[svc.ID] = svc
	return nil
}

func (registry *MemoryRegistry) GetService(name string) (*Service, error) {
	svc, found := registry.serviceData[name]
	if found {
		return &svc, nil
	} else {
		return nil, nil
	}
}

func (registry *MemoryRegistry) DeleteService(name string) error {
	delete(registry.serviceData, name)
	return nil
}

func (registry *MemoryRegistry) UpdateService(svc Service) error {
	return registry.CreateService(svc)
}

func (registry *MemoryRegistry) UpdateEndpoints(e Endpoints) error {
	return nil
}
