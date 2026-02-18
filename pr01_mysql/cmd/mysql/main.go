package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/TolkinSL/go-practice/pr01_mysql/internal/sqlite"
)

const (
	storagePatch = "./storage/storage.db"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	s, err := sqlite.New(storagePatch)
	if err != nil {
		log.Fatal("can't create/connect db: ", err)
	}

	if err = s.Init(ctx); err != nil {
		log.Fatal("can't init database: ", err)
	}

	fmt.Println("Ok!")

}