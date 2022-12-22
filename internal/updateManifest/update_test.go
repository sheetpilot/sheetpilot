package updateManifest

import (
	"errors"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"os"
	"testing"
)

func TestCheckDeployment(t *testing.T) {
	var testCases = []struct {
		name           string
		app            string
		deploymentPath string
		expErr         error
	}{
		{"SUCCESS", "test1", "./testdata/test1_deployment.yaml", nil},
		{"NOTSUPPORT", "test2", "./testdata/test2_deployment.yaml", ErrorNotSupport},
		{"MisMatch", "test4", "./testdata/test3_deployment.yaml", ErrMismatch},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			yamlData, err := os.ReadFile(tc.deploymentPath)
			if err != nil {
				t.Fatal(err)
			}
			_, err = CheckDeployment(tc.app, yamlData)
			if tc.expErr != nil {
				if err == nil {
					t.Errorf("Expected error: %q. Got 'nil' instead.", tc.expErr)
					return
				}
				if !errors.Is(err, tc.expErr) {
					t.Errorf("Expected error: %q. Got %q.", tc.expErr, err)
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error: %q", err)
			}
		})
	}
}

func TestUpdateResourceValues(t *testing.T) {
	var testCases = []struct {
		name           string
		app            string
		deploymentPath string
		updateReqCPU   string
		updateReqMem   string
		updateLmtCPU   string
		updateLmtMem   string
		updateRplCount int32
		expReqCPU      string
		expReqMem      string
		expLmtCPU      string
		expLmtMem      string
		expRplCount    int32
	}{
		{
			"SuccessTest1",
			"test1",
			"./testdata/test1_deployment.yaml",
			"100m",
			"500Mi",
			"300m",
			"1Gi",
			int32(1),
			"100m",
			"500Mi",
			"300m",
			"1Gi",
			int32(1),
		},
		{
			"FailTest1",
			"test1",
			"./testdata/test1_deployment.yaml",
			"300m",
			"1Gi",
			"100m",
			"500Mi",
			int32(1),
			"500m",
			"500Mi",
			"1",
			"1Gi",
			int32(1),
		},
		{
			"SuccessEmptyLimitTest2",
			"test3",
			"./testdata/test3_deployment.yaml",
			"1100m",
			"1500Mi",
			"",
			"",
			int32(2),
			"1100m",
			"1500Mi",
			"2000m",
			"2000Mi",
			int32(2),
		},
		{
			"FailEmptyLimitTest2",
			"test3",
			"./testdata/test3_deployment.yaml",
			"2100m",
			"2100Mi",
			"",
			"",
			int32(2),
			"1",
			"1Gi",
			"2000m",
			"2000Mi",
			int32(2),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			yamlData, err := os.ReadFile(tc.deploymentPath)
			if err != nil {
				t.Fatal(err)
			}
			obj, err := CheckDeployment(tc.app, yamlData)
			if err != nil {
				t.Fatal(err)
			}
			updatedObj := UpdateResourceValues(obj, tc.updateReqCPU, tc.updateReqMem, tc.updateLmtCPU, tc.updateLmtMem, tc.updateRplCount)

			deployment := updatedObj.(*appsv1.Deployment)

			if !deployment.Spec.Template.Spec.Containers[0].Resources.Requests.Cpu().Equal(resource.MustParse(tc.expReqCPU)) {
				t.Errorf("expected request CPU: %v, got %v instead.", deployment.Spec.Template.Spec.Containers[0].Resources.Requests.Cpu(), tc.expReqCPU)
			}

			if !deployment.Spec.Template.Spec.Containers[0].Resources.Requests.Memory().Equal(resource.MustParse(tc.expReqMem)) {
				t.Errorf("expected request Memory: %v, go %v instead.", deployment.Spec.Template.Spec.Containers[0].Resources.Requests.Memory(), tc.expReqMem)
			}

			if !deployment.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu().Equal(resource.MustParse(tc.expLmtCPU)) {
				t.Errorf("expected limit CPU: %v, got %v instead.", deployment.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu(), tc.expLmtCPU)
			}
			if !deployment.Spec.Template.Spec.Containers[0].Resources.Limits.Memory().Equal(resource.MustParse(tc.expLmtMem)) {
				t.Errorf("expected limit Memory: %v, got %v instead.", deployment.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu(), tc.expLmtMem)
			}

			if *deployment.Spec.Replicas != tc.expRplCount {
				t.Errorf("expected replica count: %d, got %d instead.", *deployment.Spec.Replicas, tc.expRplCount)
			}
		})
	}
}
