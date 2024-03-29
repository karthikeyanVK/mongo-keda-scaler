package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "app",
		Usage: "Simple app for a keda related talk",
		Commands: []cli.Command{
			{
				Name:  "app",
				Usage: "Runs simple http server on 3232",
				Action: func(c *cli.Context) error {
					StartWebserver()
					return nil
				},
			},
			{
				Name:  "api",
				Usage: "Runs simple metric API server on 3232",
				Action: func(c *cli.Context) error {
					StartAPI()
					return nil
				},
			},
			//add subcommand for mongo
			{
				Name:  "mongo",
				Usage: "mongo methods",
				Subcommands: []cli.Command{
					{
						Name:  "insert",
						Usage: "Insert some values to db",
						Action: func(c *cli.Context) error {
							err := InsertMongoData()
							if err != nil {
								fmt.Println("Failed to insert values")
								log.Fatal(err)
							} else {
								fmt.Println("Records inserted")
							}
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "Delete all records from db",
						Action: func(c *cli.Context) error {
							err := DeleteMongoData()
							if err != nil {
								fmt.Println("Failed to delete records")
								log.Fatal(err)
							} else {
								fmt.Println("Records deleted")
							}
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
