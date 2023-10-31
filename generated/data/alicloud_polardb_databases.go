package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPolardbDatabases = `{
  "block": {
    "attributes": {
      "databases": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "accounts": [
                "list",
                [
                  "object",
                  {
                    "account_name": "string",
                    "account_status": "string",
                    "privilege_status": "string"
                  }
                ]
              ],
              "character_set_name": "string",
              "db_description": "string",
              "db_name": "string",
              "db_status": "string",
              "engine": "string"
            }
          ]
        ]
      },
      "db_cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudPolardbDatabasesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPolardbDatabases), &result)
	return &result
}
