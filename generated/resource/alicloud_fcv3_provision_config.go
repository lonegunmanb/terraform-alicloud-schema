package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudFcv3ProvisionConfig = `{
  "block": {
    "attributes": {
      "always_allocate_cpu": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "always_allocate_gpu": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "function_name": {
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
      "qualifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "scheduled_actions": {
        "block": {
          "attributes": {
            "end_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schedule_expression": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "time_zone": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "target_tracking_policies": {
        "block": {
          "attributes": {
            "end_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "metric_target": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "metric_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "time_zone": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
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

func AlicloudFcv3ProvisionConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudFcv3ProvisionConfig), &result)
	return &result
}
