package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAdbResourceGroup = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_num": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "update_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "users": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
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

func AlicloudAdbResourceGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAdbResourceGroup), &result)
	return &result
}
