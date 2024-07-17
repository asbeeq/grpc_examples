package main

import (
	"encoding/json"
	"log"
	"testing"

	"google.golang.org/protobuf/proto"
	"test/example"
)

// go test -bench=.

type GreetingJSON struct {
	Message string `json:"message"`
}

func BenchmarkJsonMarshal(b *testing.B) {
	greeting := &GreetingJSON{Message: msg}
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(greeting)
		if err != nil {
			log.Fatalf("Failed to serialize: %v", err)
		}
	}
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	greeting := &GreetingJSON{Message: msg}
	data, err := json.Marshal(greeting)
	if err != nil {
		log.Fatalf("Failed to serialize: %v", err)
	}
	for i := 0; i < b.N; i++ {
		var newGreeting GreetingJSON
		err := json.Unmarshal(data, &newGreeting)
		if err != nil {
			log.Fatalf("Failed to deserialize: %v", err)
		}
	}
}

func BenchmarkProtoMarshal(b *testing.B) {
	greeting := &example.Greeting{Message: msg}
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(greeting)
		if err != nil {
			log.Fatalf("Failed to serialize: %v", err)
		}
	}
}

func BenchmarkProtoUnmarshal(b *testing.B) {
	greeting := &example.Greeting{Message: msg}
	data, err := proto.Marshal(greeting)
	if err != nil {
		log.Fatalf("Failed to serialize: %v", err)
	}
	for i := 0; i < b.N; i++ {
		var newGreeting example.Greeting
		err := proto.Unmarshal(data, &newGreeting)
		if err != nil {
			log.Fatalf("Failed to deserialize: %v", err)
		}
	}
}
