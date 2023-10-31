package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDmsEnterpriseLogicDatabase = `{
  "block": {
    "attributes": {
      "alias": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "database_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "db_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "env_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logic": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "logic_database_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owner_id_list": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "owner_name_list": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "schema_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "search_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDmsEnterpriseLogicDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDmsEnterpriseLogicDatabase), &result)
	return &result
}
