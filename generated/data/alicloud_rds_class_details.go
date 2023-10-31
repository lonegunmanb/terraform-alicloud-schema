package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRdsClassDetails = `{
  "block": {
    "attributes": {
      "category": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "class_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "class_group": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "commodity_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cpu": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_storage_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "engine_version": {
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
      "instruction_set_arch": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "max_connections": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "max_iombps": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "max_iops": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "memory_class": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "reference_price": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudRdsClassDetailsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRdsClassDetails), &result)
	return &result
}
