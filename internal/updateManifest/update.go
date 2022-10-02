package updateManifest

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"log"
)

// CheckDeployment check deployment name and return k8s Object if correct
func CheckDeployment(name string, yamlData []byte) (runtime.Object, error) {
	var deployment *appsv1.Deployment
	decoder := scheme.Codecs.UniversalDeserializer()
	obj, groupVersionKind, err := decoder.Decode(yamlData, nil, nil)
	if err != nil {
		log.Println(err)
	}
	if groupVersionKind.Group == "apps" && groupVersionKind.Version == "v1" && groupVersionKind.Kind == "Deployment" {
		deployment = obj.(*appsv1.Deployment)
	}
	if deployment.Name != name {
		return nil, fmt.Errorf("deployment is not correct")
	}

	return obj, nil
}

func int32Ptr(i int32) *int32 {
	return &i
}

// UpdateResourceValues update Deployment resource values
func UpdateResourceValues(obj runtime.Object, cpuRequest, menRequest, cpuLimit, memLimit int64, replicaCount int32) runtime.Object {
	deployment := obj.(*appsv1.Deployment)
	deployment.Spec.Template.Spec.Containers[0].Resources.Requests = make(map[corev1.ResourceName]resource.Quantity)

	deployment.Spec.Template.Spec.Containers[0].Resources.Requests[corev1.ResourceCPU] = *resource.NewMilliQuantity(cpuRequest, resource.DecimalSI)
	deployment.Spec.Template.Spec.Containers[0].Resources.Requests[corev1.ResourceMemory] = *resource.NewQuantity(menRequest*1024*1024*1024, resource.BinarySI)
	deployment.Spec.Template.Spec.Containers[0].Resources.Limits[corev1.ResourceCPU] = *resource.NewMilliQuantity(cpuLimit, resource.DecimalSI)
	deployment.Spec.Template.Spec.Containers[0].Resources.Limits[corev1.ResourceMemory] = *resource.NewQuantity(memLimit*1024*1024*1024, resource.BinarySI)
	deployment.Spec.Replicas = int32Ptr(replicaCount)

	return obj
}