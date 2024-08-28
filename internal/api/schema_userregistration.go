package api

import (
	"fmt"
)

type UserRegistration struct {
	DisplayName	string	`json:"displayName"`
	Email	string	`json:"email"`
}

// Checks if all of the required fields for UserRegistration are set
// and validates all of the constraints for the object.
func (obj *UserRegistration) Validate() error {
	if obj == nil {
		return nil
	}
	fields := map[string]interface{}{
		"email": obj.Email,
		"displayName": obj.DisplayName,
	}

	for field, value := range fields {
		if isEmpty := IsValEmpty(value); isEmpty{
			return fmt.Errorf("required field '%s' for object 'UserRegistration' is empty or unset", field)
		}
	}

	return nil
}

