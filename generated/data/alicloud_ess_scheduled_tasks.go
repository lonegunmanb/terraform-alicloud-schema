package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEssScheduledTasks = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scheduled_action": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scheduled_task_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tasks": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "id": "string",
              "launch_expiration_time": "number",
              "launch_time": "string",
              "max_value": "number",
              "min_value": "number",
              "name": "string",
              "recurrence_end_time": "string",
              "recurrence_type": "string",
              "recurrence_value": "string",
              "scheduled_action": "string",
              "task_enabled": "bool"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEssScheduledTasksSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEssScheduledTasks), &result)
	return &result
}
