package main

import (
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
	"test/example"
)

const msg = "Text to marshal and unmarshal"

func Json() {
	fmt.Println("--- JSON: ")
	type Greeting struct {
		Message string `json:"message"`
	}

	// Создаем экземпляр сообщения Greeting
	greeting := &Greeting{
		Message: msg,
	}

	// Сериализация сообщения в JSON
	data, err := json.Marshal(greeting)
	if err != nil {
		log.Fatalf("Failed to serialize: %v", err)
	}
	fmt.Println(data)
	fmt.Println(string(data))

	// Десериализация сообщения из JSON
	var newGreeting Greeting
	err = json.Unmarshal(data, &newGreeting)
	if err != nil {
		log.Fatalf("Failed to deserialize: %v", err)
	}
	fmt.Println(newGreeting.Message)
}

func Proto() {
	fmt.Println("--- PROTO: ")
	// Создаем экземпляр сообщения Greeting
	greeting := &example.Greeting{
		Message: msg,
	}

	// Сериализация сообщения в бинарный формат
	data, err := proto.Marshal(greeting)
	if err != nil {
		log.Fatalf("Failed to serialize: %v", err)
	}

	fmt.Println(data)
	fmt.Println(string(data))

	// Десериализация сообщения из бинарного формата
	var newGreeting example.Greeting
	err = proto.Unmarshal(data, &newGreeting)
	if err != nil {
		log.Fatalf("Failed to deserialize: %v", err)
	}

	// Выводим сообщение
	fmt.Println(newGreeting.Message)
}

func main() {
	Json()
	fmt.Println()
	Proto()
}
