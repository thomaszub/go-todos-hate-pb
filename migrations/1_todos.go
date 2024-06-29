package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "2u23enfabtxwxpe",
				"created": "2024-06-27 18:47:01.812Z",
				"updated": "2024-06-27 18:47:01.812Z",
				"name": "todos",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "7vsmegan",
						"name": "content",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "mezrknnb",
						"name": "done",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, false, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
