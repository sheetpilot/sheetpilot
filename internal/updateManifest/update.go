package updateManifest

import (
	"errors"
	"io"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/client-go/kubernetes/scheme"
)

var (
	ErrorNotSupport = errors.New("not support resource type")
	ErrMismatch     = errors.New("mismatch deployment name")
)

// CheckDeployment check deployment name and return k8s Object if correct
func CheckDeployment(name string, yamlData []byte) (runtime.Object, error) {
	//var deployment *appsv1.Deployment
	decoder := scheme.Codecs.UniversalDeserializer()
	obj, _, err := decoder.Decode(yamlData, nil, nil)
	if err != nil {
		log.Println(err)
	}

	switch deployment := obj.(type) {
	case *appsv1.Deployment:
		if deployment.Name != name {
			return nil, ErrMismatch
		}
	default:
		return nil, ErrorNotSupport
	}

	return obj, nil
}

func int32Ptr(i int32) *int32 {
	return &i
}

// UpdateResourceValues update Deployment resource values
func UpdateResourceValues(obj runtime.Object, cpuRequest, memRequest, cpuLimit, memLimit string, replicaCount int32) runtime.Object {
	deployment := obj.(*appsv1.Deployment)
	cpuCondition := -1
	memCondition := -1
	//deployment.Spec.Template.Spec.Containers[0].Resources.Requests = make(map[corev1.ResourceName]resource.Quantity)
	if cpuLimit != "" {
		cpuLMT := resource.MustParse(cpuLimit)
		cpuCondition = cpuLMT.Cmp(resource.MustParse(cpuRequest))
	} else {
		cpuCondition = deployment.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu().Cmp(resource.MustParse(cpuRequest))
	}

	if memLimit != "" {
		memLMT := resource.MustParse(memLimit)
		memCondition = memLMT.Cmp(resource.MustParse(memRequest))
	} else {
		memCondition = deployment.Spec.Template.Spec.Containers[0].Resources.Limits.Memory().Cmp(resource.MustParse(memRequest))
	}

	if cpuRequest != "" && resource.MustParse(cpuRequest).Format == "DecimalSI" && cpuCondition != -1 {
		deployment.Spec.Template.Spec.Containers[0].Resources.Requests[corev1.ResourceCPU] = resource.MustParse(cpuRequest)
	}
	if memRequest != "" && resource.MustParse(memRequest).Format == "BinarySI" && memCondition != -1 {
		deployment.Spec.Template.Spec.Containers[0].Resources.Requests[corev1.ResourceMemory] = resource.MustParse(memRequest)
	}
	if cpuLimit != "" && resource.MustParse(cpuLimit).Format == "DecimalSI" && cpuCondition != -1 {
		deployment.Spec.Template.Spec.Containers[0].Resources.Limits[corev1.ResourceCPU] = resource.MustParse(cpuLimit)
	}

	if memLimit != "" && resource.MustParse(memLimit).Format == "BinarySI" && memCondition != -1 {
		deployment.Spec.Template.Spec.Containers[0].Resources.Limits[corev1.ResourceMemory] = resource.MustParse(memLimit)
	}

	deployment.Spec.Replicas = int32Ptr(replicaCount)

	return obj
}

// PrintDeployment Print correct deployment yaml to file
func PrintDeployment(out io.Writer, obj runtime.Object) error {
	printr := printers.NewTypeSetter(scheme.Scheme).ToPrinter(&printers.YAMLPrinter{})

	deployment := obj.(*appsv1.Deployment)

	return printr.PrintObj(deployment, out)
}
