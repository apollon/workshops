package actions

import (
	"github.com/apollon/workshops/api/models"
	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

// ItemsList default implementation.
func ItemsList(c buffalo.Context) error {
	items := []models.Item{}
	err := models.DB.All(&items)
	if err != nil {
		c.Logger().Error("DB error", errors.WithStack(err))
	}
	return c.Render(200, r.JSON(items))
}

// ItemsIndex default implementation.
func ItemsIndex(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "Items index"}))
}
