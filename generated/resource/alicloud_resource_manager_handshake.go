package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudResourceManagerHandshake = `{
  "block": {
    "attributes": {
      "expire_time": {
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
      "master_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "master_account_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "modify_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "note": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_directory_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_entity": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudResourceManagerHandshakeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudResourceManagerHandshake), &result)
	return &result
}
