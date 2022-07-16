package main

import "errors"

type validator struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// * Validate a password in response to a valid POST request at /validator/{id}
func (v *validator) validate() error {
	return errors.New("not implemented")
}

// * Create a new validator in response to a valid POST request at /validator
func (v *validator) createValidator() error {
	return errors.New("not implemented")
}

func (v *validator) getValidator() error {
	return errors.New("not implemented")
}

// * Delete a validator in response to a valid DELETE request at /validator/{id}
func (v *validator) deleteValidator() error {
	return errors.New("not implemented")
}

// * Updata a validator in response to a valid PUT request at /validator/{id}
func (v *validator) updateValidator() error {
	return errors.New("not implemented")
}

// * Fetcha  list of all validators in response to a valid GET request at /validators
func fetchValidators(start, count int) ([]validator, error) {
	return nil, errors.New("not implemented")
}
