package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDmsEnterpriseDatabases = `{
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
              "catalog_name": "string",
              "database_id": "string",
              "db_type": "string",
              "dba_id": "string",
              "dba_name": "string",
              "encoding": "string",
              "env_type": "string",
              "host": "string",
              "id": "string",
              "instance_id": "string",
              "owner_id_list": [
                "list",
                "string"
              ],
              "owner_name_list": [
                "list",
                "string"
              ],
              "port": "number",
              "schema_name": "string",
              "search_name": "string",
              "sid": "string",
              "state": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDmsEnterpriseDatabasesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDmsEnterpriseDatabases), &result)
	return &result
}
