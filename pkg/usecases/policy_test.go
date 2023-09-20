package usecases

import (
	"testing"

	"github.com/josharmi/aws-kms-policy-scanner/pkg/entities/policy"
)

func TestTooOpenPolicy(t *testing.T) {
	want := policy.Polcy{}

	got := true

	if want != got {
		t.Fatalf("want %d, but got %d", want, got)
	}
}
