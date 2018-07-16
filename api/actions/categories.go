package actions

import (
	"github.com/apollon/workshops/api/models"
	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

// CategoriesList default implementation.
func CategoriesList(c buffalo.Context) error {
	categories := []models.Category{}
	err := models.DB.All(&categories)
	if err != nil {
		c.Logger().Error("DB error", errors.WithStack(err))
	}
	return c.Render(200, r.JSON(categories))
}

// CategoriesIndex default implementation.
func CategoriesIndex(c buffalo.Context) error {
	query := models.DB.Q()

	categoryID := c.Param("id")
	if categoryID != "" {
		query = query.Where(`id = ?`, categoryID)
	}

	parentID := c.Param("parentID")
	if parentID != "" {
		query = query.Where(`parent_id = ?`, parentID)
	}

	categories := []models.Category{}
	err := query.All(&categories)
	if err != nil {
		c.Logger().Error("DB error", errors.WithStack(err))
	}
	return c.Render(200, r.JSON(categories))
}
