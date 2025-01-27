package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		if err != nil {
			fmt.Printf("record %d: %+v validation errors: %v\n", count, emp, err)
			continue
		}
	}
}

const data = ` 
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

type EmptyFieldError struct {
	Field string
}

func (e EmptyFieldError) Error() string {
	return fmt.Sprintf("missing %s", e.Field)
}

var (
	validID = regexp.MustCompile(`\w{4}-\d{3}`)
)

var ErrInvalidID = errors.New("invalid ID")

func ValidateEmployee(e Employee) error {
	var errs []error
	if len(e.ID) == 0 {
		errs = append(errs, errors.New("Field ID is missing"))
	}
	if !validID.MatchString(e.ID) {
		return ErrInvalidID
	}
	if len(e.FirstName) == 0 {
		errs = append(errs, errors.New("Field FirstName is missing"))
	}
	if len(e.LastName) == 0 {
		errs = append(errs, errors.New("Field LastName is missing"))
	}
	if len(e.Title) == 0 {
		errs = append(errs, errors.New("Field Title is missing"))
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil

}
