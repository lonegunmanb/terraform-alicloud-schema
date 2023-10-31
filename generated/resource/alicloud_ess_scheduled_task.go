package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEssScheduledTask = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "desired_capacity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "launch_expiration_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "launch_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_value": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_value": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "recurrence_end_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "recurrence_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "recurrence_value": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scaling_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scheduled_action": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scheduled_task_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "task_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEssScheduledTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEssScheduledTask), &result)
	return &result
}
