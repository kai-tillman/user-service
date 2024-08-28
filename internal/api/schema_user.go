package api

import (
	"fmt"
)

type User struct {
	DisplayName	string	`json:"displayName"`
	Email	string	`json:"email"`
	Id	string	`json:"id"`
}

// Checks if all of the required fields for User are set
// and validates all of the constraints for the object.
func (obj *User) Validate() error {
	if obj == nil {
		return nil
	}
	fields := map[string]interface{}{
		"id": obj.Id,
		"email": obj.Email,
		"displayName": obj.DisplayName,
	}

	for field, value := range fields {
		if isEmpty := IsValEmpty(value); isEmpty{
			return fmt.Errorf("required field '%s' for object 'User' is empty or unset", field)
		}
	}

	return nil
}

