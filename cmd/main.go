package main

import (
	"fmt"
	"log"

	"github.com/axelrhd/toaster"
	_ "modernc.org/sqlite"
)

var store toaster.Store

func main() {
	store, err := toaster.ConnectDB("sqlite", "file:db.sqlite?_journal=WAL&_fk=1")
	if err != nil {
		log.Fatal(err)
	}

	err = store.CreateSchema()
	if err != nil {
		log.Fatal(err)
	}

	id, err := store.Add(toaster.New("Hello you!"))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(id)
	}

	t, err := store.Get("25042101355691ncVDYN0001")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%+v\n", t)
	}

	err = store.Delete("25042101355691ncVDYN0001")
	if err != nil {
		fmt.Println(err)
	}

	t1, err := store.Get(id)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(t1.Render())
	}

	// withMapStore()
}

func withMapStore() {
	store = &toaster.MapStore{}
	id, err := store.Add(toaster.New("Hello World"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("id:", id)

	t, err := store.Get("abc123")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%+v\n", t)
	}

	t1, err := store.Get(id)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%+v\n", t1)
	}
}
