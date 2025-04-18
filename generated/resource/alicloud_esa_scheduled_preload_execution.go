package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaScheduledPreloadExecution = `{
  "block": {
    "attributes": {
      "end_time": {
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
      "interval": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "scheduled_preload_execution_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "scheduled_preload_job_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "slice_len": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "start_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
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

func AlicloudEsaScheduledPreloadExecutionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaScheduledPreloadExecution), &result)
	return &result
}
