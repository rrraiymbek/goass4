package main

import (
	"cleanArch/pkg/store/postgres"
	"cleanArch/services/entity/internal/delivery/network"
	"cleanArch/services/entity/internal/repository"
	"cleanArch/services/entity/internal/useCase/contact"
	"cleanArch/services/entity/internal/useCase/group"
	"log"
	"net/http"
)

func main() {
	settings := postgres.Settings{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "1234",
		DBName:   "cleanArch",
	}

	conn, err := postgres.New(settings)

	if err != nil {
		panic(err)
	}
	contactRepo := repository.NewContactRepository(conn)

	contactUseCase := contact.NewContactUseCase(contactRepo)

	contactDelivery := network.NewContactDelivery(contactUseCase)

	groupRepo := repository.NewGroupRepository(conn)

	groupUseCase := group.NewGroupUseCase(groupRepo)

	groupDelivery := network.NewGroupDelivery(groupUseCase)
	defer conn.Pool.Close()

	http.HandleFunc("/contacts", contactDelivery.HandleRequests)
	http.HandleFunc("/groups", groupDelivery.HandleRequests)
	log.Println(conn.Pool.Stat())

	log.Println("Hello World!")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
