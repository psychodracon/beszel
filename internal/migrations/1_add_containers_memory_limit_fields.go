package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("containers")
		if err != nil {
			return err
		}
		if collection.Fields.GetByName("memory_limit") == nil {
			collection.Fields.Add(&core.NumberField{
				Name: "memory_limit",
			})
		}
		return app.Save(collection)
	}, nil)
}
