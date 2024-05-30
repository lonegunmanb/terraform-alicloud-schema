package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOssAccountPublicAccessBlock = `{
  "block": {
    "attributes": {
      "block_public_access": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
            },
            "update": {
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

func AlicloudOssAccountPublicAccessBlockSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOssAccountPublicAccessBlock), &result)
	return &result
}
