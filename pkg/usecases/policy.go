package usecases

import (
	"aws-kms-policy-scanner/pkg/entities"
	"fmt"
	"regexp"
)

func ValidatePolicy(policy entities.Policy) error {
	r, _ := regexp.Compile(`.*arn:aws:iam::\d{12}:root.*`)
	for _, statement := range policy.Statement {
		if statement.Action == "kms:*" {
			return &entities.TooPermissiveError{}
		}
		match := r.FindString(fmt.Sprint(statement.Principal))
		if match != "" {
			return &entities.TooPermissiveError{}
		}
	}
	return nil
}
