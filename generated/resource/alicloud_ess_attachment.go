package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEssAttachment = `{
  "block": {
    "attributes": {
      "entrusted": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "force": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "lifecycle_hook": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "load_balancer_weights": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "number"
        ]
      },
      "scaling_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEssAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEssAttachment), &result)
	return &result
}
