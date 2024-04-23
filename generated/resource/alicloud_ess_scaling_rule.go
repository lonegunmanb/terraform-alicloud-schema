package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEssScalingRule = `{
  "block": {
    "attributes": {
      "adjustment_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "adjustment_value": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ari": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cooldown": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "disable_scale_in": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "estimated_instance_warmup": {
        "computed": true,
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
      "initial_max_size": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "metric_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "min_adjustment_magnitude": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "predictive_scaling_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "predictive_task_buffer_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "predictive_value_behavior": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "predictive_value_buffer": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "scale_in_evaluation_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "scale_out_evaluation_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "scaling_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scaling_rule_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scaling_rule_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_value": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "alarm_dimension": {
        "block": {
          "attributes": {
            "dimension_key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dimension_value": {
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
      "step_adjustment": {
        "block": {
          "attributes": {
            "metric_interval_lower_bound": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metric_interval_upper_bound": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scaling_adjustment": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEssScalingRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEssScalingRule), &result)
	return &result
}
