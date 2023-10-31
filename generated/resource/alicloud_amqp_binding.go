package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAmqpBinding = `{
  "block": {
    "attributes": {
      "argument": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "binding_key": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "binding_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_name": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_exchange": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "virtual_host_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudAmqpBindingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAmqpBinding), &result)
	return &result
}
