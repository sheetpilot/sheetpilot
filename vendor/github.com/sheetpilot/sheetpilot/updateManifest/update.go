package updateManifest

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"log"
)

// CheckDeployment is deployment name correct or not?
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

func UpdateReplicas(obj runtime.Object, n int32) runtime.Object {
	deployment := obj.(*appsv1.Deployment)
	deployment.Spec.Replicas = int32Ptr(n)
	return obj
}
