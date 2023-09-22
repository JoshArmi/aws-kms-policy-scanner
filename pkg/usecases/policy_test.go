package usecases

import (
	"aws-kms-policy-scanner/pkg/entities"
	"testing"
)

func TestTooOpenPolicy(t *testing.T) {
	policy := entities.Policy{
		Version: "2012-10-17",
		Id:      "key-default-1",
		Statement: []entities.Statement{{
			Sid:      "",
			Effect:   "Allow",
			Action:   "kms:*",
			Resource: "*",
		}},
	}

	got := ValidatePolicy(policy)

	if _, ok := got.(*entities.TooPermissiveError); ok {
	} else {
		t.Fatalf("wanted TooPermissiveError, but got %s", got)
	}
}

func TestAllAccountPolicy(t *testing.T) {
	policy := entities.Policy{
		Version: "2012-10-17",
		Id:      "key-default-1",
		Statement: []entities.Statement{{
			Sid:    "",
			Effect: "Allow",
			Principal: map[string]string{
				"AWS": "arn:aws:iam::099267815798:root",
			},
			Action:   "kms:Decrypt",
			Resource: "*",
		}},
	}

	got := ValidatePolicy(policy)

	if _, ok := got.(*entities.TooPermissiveError); ok {
	} else {
		t.Fatalf("wanted TooPermissiveError, but got %s", got)
	}
}

func TestGoodPolicy(t *testing.T) {
	policy := entities.Policy{
		Version: "2012-10-17",
		Id:      "key-default-1",
		Statement: []entities.Statement{{
			Sid:    "",
			Effect: "Allow",
			Principal: map[string]string{
				"AWS": "arn:aws:iam::099267815798:kms-user",
			},
			Action:   "kms:Decrypt",
			Resource: "*",
		}},
	}

	got := ValidatePolicy(policy)

	if got != nil {
		t.Fatalf("wanted nothing, but got %s", got)
	}
}
