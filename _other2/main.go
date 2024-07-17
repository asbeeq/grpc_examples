package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
	pb "test/example"
)

func main() {
	// Create a Person instance
	person := &pb.Person{
		Name:         "John Doe",
		Id:           1234,
		Email:        "johndoe@example.com",
		PhoneNumbers: []string{"555-1234", "555-5678"},
		Attributes:   map[string]string{"key1": "value1", "key2": "value2"},
		Address:      &pb.Address{Street: "123 Elm St", City: "Springfield", State: "IL", Zip: "12345"},
		Status:       pb.Status_ACTIVE,
	}

	// Serialize the Person to binary format
	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatalf("Failed to serialize person: %v", err)
	}

	// Deserialize the binary data back into a new Person instance
	var newPerson pb.Person
	err = proto.Unmarshal(data, &newPerson)
	if err != nil {
		log.Fatalf("Failed to deserialize person: %v", err)
	}

	// Print the deserialized Person message
	fmt.Println("Name:", newPerson.GetName())
	fmt.Println("ID:", newPerson.GetId())
	fmt.Println("Email:", newPerson.GetEmail())
	fmt.Println("Phone Numbers:", newPerson.GetPhoneNumbers())
	fmt.Println("Attributes:", newPerson.GetAttributes())
	fmt.Println("Address:", newPerson.GetAddress())
	fmt.Println("Status:", newPerson.GetStatus())
}
