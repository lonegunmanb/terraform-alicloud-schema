package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSagDnatEntry = `{
  "block": {
    "attributes": {
      "external_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "external_port": {
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
      "internal_ip": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "internal_port": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ip_protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sag_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSagDnatEntrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSagDnatEntry), &result)
	return &result
}
