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
      "metric_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
