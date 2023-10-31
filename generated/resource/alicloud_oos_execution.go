package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOosExecution = `{
  "block": {
    "attributes": {
      "counters": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "create_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "end_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "executed_by": {
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
      "is_parent": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "loop_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "outputs": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "parameters": {
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "safety_check": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "template_content": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "template_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_date": {
        "computed": true,
        "description_kind": "plain",
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

func AlicloudOosExecutionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOosExecution), &result)
	return &result
}
