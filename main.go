package main

import (
	"groupie-tracker/repository"
)

func main() {
	store := repository.New_Store()
	store.LoadData()
}
