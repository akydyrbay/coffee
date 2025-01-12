package service

import "fmt"

type Status struct {
	ErrorMessage error
	Code         int
}

var (
	IdAlreadyExists Status = Status{ErrorMessage: fmt.Errorf("inventory item ID already exists"), Code: 409}
	NoIdGiven       Status = Status{ErrorMessage: fmt.Errorf("no ID given"), Code: 404}
	NotFound        Status = Status{ErrorMessage: fmt.Errorf("item not found"), Code: 404}
	// TODO
	// ....
	Success Status = Status{ErrorMessage: nil, Code: 200}
)
