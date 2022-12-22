package gitapiclient

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	clientArgs := new(gitApiClientFactory)
	clientArgs.clientName = "github"
	clientArgs.token = "xxxxx"

	client, _ := GitApiClientFactory(clientArgs)
	client("/sample-deployment/contents/deployment/admin/deployment.yaml")

	if 1 == 1 {
		t.Fatal("Failed Catastrphically")
		// t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}

	// name := "Gladys"
	// want := regexp.MustCompile(`\b` + name + `\b`)

	// if !want.MatchString(msg) || err != nil {
	// 	t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	// }
}
