package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRealtimeComputeJob = `{
  "block": {
    "attributes": {
      "deployment_id": {
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
      "job_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "namespace": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_queue_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stop_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "local_variables": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "restore_strategy": {
        "block": {
          "attributes": {
            "allow_non_restored_state": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "job_start_time_in_ms": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "kind": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "savepoint_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "status": {
        "block": {
          "attributes": {
            "current_job_status": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "failure": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "failed_at": "number",
                    "message": "string",
                    "reason": "string"
                  }
                ]
              ]
            },
            "health_score": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "risk_level": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "running": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "observed_flink_job_restarts": "number",
                    "observed_flink_job_status": "string"
                  }
                ]
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AlicloudRealtimeComputeJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRealtimeComputeJob), &result)
	return &result
}
