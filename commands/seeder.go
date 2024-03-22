package main

import (
	"fitness-tracker-api/testbackend/database"
	"flag"
	"fmt"
	"log"
)

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("command takes exactly one argument")
		return
	}

	db, err := database.ConnectToDatabase("/Users/isaacdugan/code/fitness-tracker-api/database/db.sqlite")
	if err != nil {
		fmt.Print("Error opening database")
		fmt.Print(err)
		return
	}
	defer db.Close()

	seeder := database.NewSeeder(db)

	switch firstArg := flag.Arg(0); firstArg {
	case "seed":
		err := seeder.Seed()
		if err != nil {
			log.Print(err.Error())
			return
		}
		fmt.Print("Done seeding")
	case "clear":
		err := seeder.Clear()
		if err != nil {
			log.Print(err.Error())
			return
		}
		fmt.Print("Done clearing")
	default:
		fmt.Printf("command not recognized %s", firstArg)
	}

}
