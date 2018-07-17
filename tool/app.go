package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/apollon/workshops/api/models"
	"github.com/gobuffalo/uuid"

	"github.com/codegangsta/cli"
	"github.com/pkg/errors"
)

func main() {
	app := cli.NewApp()
	app.Name = "tool"

	app.Commands = []cli.Command{
		{
			Name: "import",
			Subcommands: []cli.Command{
				{
					Name:      "category",
					ArgsUsage: "[fileName]",
					Action: func(c *cli.Context) {
						fmt.Println("args? ", c.Args().Present())
						if !c.Args().Present() {
							return
						}

						fileName := c.Args().First()
						fmt.Println("fileName: ", fileName)
						/*
							resp, err := http.Get("http://127.0.0.1:3000/api/v1/categories/list")
							if err != nil {
								fmt.Println("request error", errors.WithStack(err))
								return
							}
							defer resp.Body.Close()

							type CategoryItem struct {
								ID    uuid.UUID `json:"id"`
								Title string    `json:"title"`
							}

							categories := []CategoryItem{}

							err = json.NewDecoder(resp.Body).Decode(&categories)
							if err != nil {
								fmt.Println("decode error", errors.WithStack(err))
							}

							fout, err := os.Create(fileName)
							defer fout.Close()
							for _, category := range categories {
								fout.WriteString(category.ID.String() + "," + category.Title + "\n")
							}
						*/
						fin, err := os.Open(fileName)
						if err != nil {
							fmt.Println("open file error ", errors.WithStack(err))
							return
						}
						defer fin.Close()
						scaner := bufio.NewScanner(fin)
						for scaner.Scan() {
							line := scaner.Text()
							csvReader := csv.NewReader(strings.NewReader(line))
							csvLine, err := csvReader.Read()
							if err != nil {
								fmt.Println("error parsing data", errors.WithStack(err))
								return
							}
							id, err := uuid.NewV4()
							parentID, err := uuid.FromString(csvLine[4])
							category := &models.Category{
								ID:          id,
								Alias:       csvLine[0],
								Title:       csvLine[1],
								Description: csvLine[2],
								Logo:        csvLine[3],
								ParentID:    parentID,
							}
							err = models.DB.Create(category)
							if err != nil {
								fmt.Println("error creating record", errors.WithStack(err))
								return
							}
							fmt.Println("created record", id)
						}
					},
				},
				{
					Name:      "items",
					ArgsUsage: "[fileName]",
					Action: func(c *cli.Context) {
						fmt.Println("args? ", c.Args().Present())
						if !c.Args().Present() {
							return
						}

						fileName := c.Args().First()
						fmt.Println("fileName: ", fileName)
						fin, err := os.Open(fileName)
						if err != nil {
							fmt.Println("open file error ", errors.WithStack(err))
							return
						}
						defer fin.Close()
						scaner := bufio.NewScanner(fin)
						for scaner.Scan() {
							line := scaner.Text()
							csvReader := csv.NewReader(strings.NewReader(line))
							csvLine, err := csvReader.Read()
							if err != nil {
								fmt.Println("error parsing data", errors.WithStack(err))
								return
							}
							id, err := uuid.NewV4()
							price, err := strconv.Atoi(csvLine[4])
							count, err := strconv.Atoi(csvLine[5])
							item := &models.Item{
								ID:          id,
								Alias:       csvLine[0],
								Title:       csvLine[1],
								Description: csvLine[2],
								Pictures:    csvLine[3],
								Price:       price,
								Count:       count,
							}
							err = models.DB.Create(item)
							if err != nil {
								fmt.Println("error creating record", errors.WithStack(err))
								return
							}
							fmt.Println("created record", id)
						}
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
