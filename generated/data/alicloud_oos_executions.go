package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOosExecutions = `{
  "block": {
    "attributes": {
      "category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "end_date": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "end_date_after": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "executed_by": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "executions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "category": "string",
              "counters": "string",
              "create_date": "string",
              "end_date": "string",
              "executed_by": "string",
              "execution_id": "string",
              "id": "string",
              "is_parent": "bool",
              "mode": "string",
              "outputs": "string",
              "parameters": "string",
              "parent_execution_id": "string",
              "ram_role": "string",
              "start_date": "string",
              "status": "string",
              "status_message": "string",
              "status_reason": "string",
              "template_id": "string",
              "template_name": "string",
              "template_version": "string",
              "update_date": "string"
            }
          ]
        ]
      },
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
      "include_child_execution": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parent_execution_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ram_role": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sort_field": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sort_order": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_date_after": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_date_before": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
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
      },
      "template_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudOosExecutionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOosExecutions), &result)
	return &result
}
