package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
	pb "test/example"
)

func main() {
	// Создаем экземпляр сообщения User
	user := &pb.User{
		Username: "alice",
		Email:    "alice@example.com",
		Address: &pb.Address{
			Street: "123 Main St",
			City:   "Anytown",
		},
	}

	// Сериализуем сообщение в бинарный формат
	data, err := proto.Marshal(user)
	if err != nil {
		log.Fatalf("Failed to serialize user: %v", err)
	}

	// Десериализуем бинарные данные обратно в экземпляр User
	var newUser pb.User
	err = proto.Unmarshal(data, &newUser)
	if err != nil {
		log.Fatalf("Failed to deserialize user: %v", err)
	}

	// Выводим данные из нового экземпляра User
	fmt.Println("Deserialized Username:", newUser.GetUsername())
	fmt.Println("Deserialized Email:", newUser.GetEmail())
	fmt.Println("Deserialized Address:")
	fmt.Println("- Street:", newUser.GetAddress().GetStreet())
	fmt.Println("- City:", newUser.GetAddress().GetCity())
}
