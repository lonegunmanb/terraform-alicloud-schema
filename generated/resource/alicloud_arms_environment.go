package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudArmsEnvironment = `{
  "block": {
    "attributes": {
      "aliyun_lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bind_resource_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "environment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "environment_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "environment_sub_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "environment_type": {
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
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
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

func AlicloudArmsEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudArmsEnvironment), &result)
	return &result
}
