package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDbfsInstance = `{
  "block": {
    "attributes": {
      "advanced_features": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "category": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "delete_snapshot": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_raid": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "encryption": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "fs_name": {
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
      "instance_name": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "performance_level": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "raid_stripe_unit_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "size": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "snapshot_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "used_scene": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "ecs_list": {
        "block": {
          "attributes": {
            "ecs_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
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

func AlicloudDbfsInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDbfsInstance), &result)
	return &result
}
