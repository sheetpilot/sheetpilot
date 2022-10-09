package updateManifest

import (
	"errors"
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
