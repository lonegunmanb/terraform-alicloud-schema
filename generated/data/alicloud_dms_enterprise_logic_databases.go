package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDmsEnterpriseLogicDatabases = `{
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
              "alias": "string",
              "database_ids": [
                "list",
                "string"
              ],
              "db_type": "string",
              "env_type": "string",
              "id": "string",
              "logic": "bool",
              "logic_database_id": "string",
              "owner_id_list": [
                "list",
                "string"
              ],
              "owner_name_list": [
                "list",
                "string"
              ],
              "schema_name": "string",
              "search_name": "string"
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

func AlicloudDmsEnterpriseLogicDatabasesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDmsEnterpriseLogicDatabases), &result)
	return &result
}
