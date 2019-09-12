package proto_quick_start

import (
	"log"
	"testing"
)

func TestPerson(t *testing.T) {

	p := &Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*Person_PhoneNumber{
			{Number: "555-4321", Type: Person_HOME},
		},
	}

	log.Println(p.String())
}
