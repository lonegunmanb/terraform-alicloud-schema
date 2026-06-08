package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsAggTaskGroup = `{
  "block": {
    "attributes": {
      "agg_task_group_config": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "agg_task_group_config_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "agg_task_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "agg_task_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cron_expr": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delay": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
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
      "max_retries": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_run_time_in_seconds": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "override_if_exists": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "precheck_string": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "schedule_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schedule_time_expr": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_prometheus_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_prometheus_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "to_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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

func AlicloudCmsAggTaskGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsAggTaskGroup), &result)
	return &result
}
