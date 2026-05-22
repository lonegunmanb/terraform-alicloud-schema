package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPolardbBatchTask = `{
  "block": {
    "attributes": {
      "batch_id": {
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
      "instance_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "task_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "task_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "task_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "task_params": {
        "block": {
          "attributes": {
            "skill_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
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

func AlicloudPolardbBatchTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPolardbBatchTask), &result)
	return &result
}
