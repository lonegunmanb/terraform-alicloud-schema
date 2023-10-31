package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOtsSecondaryIndex = `{
  "block": {
    "attributes": {
      "defined_columns": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "include_base_data": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "index_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "index_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "primary_keys": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "table_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudOtsSecondaryIndexSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOtsSecondaryIndex), &result)
	return &result
}
