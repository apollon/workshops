package grifts

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/apollon/workshops/api/models"
	"github.com/gobuffalo/uuid"
	"github.com/icrowley/fake"
	"github.com/markbates/grift/grift"
	"github.com/pkg/errors"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		models.DB.TruncateAll()

		for category1Index := 0; category1Index < 15; category1Index++ {
			category1ID, _ := uuid.NewV4()
			category1 := &models.Category{
				ID:          category1ID,
				Alias:       fake.Model(),
				Title:       fake.Title(),
				Description: fake.Words(),
				Logo:        fake.Word(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			err := models.DB.Create(category1)
			if err != nil {
				return errors.WithStack(err)
			}

			fmt.Println("Category ", category1ID, " created")

			for category2Index := 0; category2Index < 20; category2Index++ {
				category2ID, _ := uuid.NewV4()
				category2 := &models.Category{
					ID:          category2ID,
					ParentID:    category1ID,
					Alias:       fake.Model(),
					Title:       fake.Title(),
					Description: fake.Words(),
					Logo:        fake.Word(),
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}

				err := models.DB.Create(category2)
				if err != nil {
					return errors.WithStack(err)
				}

				fmt.Println("Category ", category2ID, " created")

				for itemIndex := 0; itemIndex < 20; itemIndex++ {
					itemID, _ := uuid.NewV4()
					item := &models.Item{
						ID:          itemID,
						Alias:       fake.Word(),
						Title:       fake.ProductName(),
						Description: fake.Words(),
						Pictures:    fake.Word(),
						Price:       rand.Intn(1000),
						Count:       rand.Intn(10),
						CategoryID:  category2ID,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					}

					err := models.DB.Create(item)
					if err != nil {
						return errors.WithStack(err)
					}

					fmt.Println("Item ", itemID, " created")
				}
			}
		}

		return nil
	})
})
