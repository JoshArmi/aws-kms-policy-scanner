package main

import (
	"aws-kms-policy-scanner/pkg/drivers"
	"aws-kms-policy-scanner/pkg/usecases"
	"fmt"
)

func main() {
	keys := drivers.GetKeys()
	for id, key := range keys {
		err := usecases.ValidatePolicy(key.Policy)
		if err != nil {
			fmt.Printf("Key %s has a too permissive policy\n", id)
		}
	}
}
