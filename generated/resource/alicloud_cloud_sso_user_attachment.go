package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudSsoUserAttachment = `{
  "block": {
    "attributes": {
      "directory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_id": {
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
      "user_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudSsoUserAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudSsoUserAttachment), &result)
	return &result
}
