package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEssScalingRules = `{
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
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "adjustment_type": "string",
              "adjustment_value": "number",
              "cooldown": "number",
              "disable_scale_in": "bool",
              "estimated_instance_warmup": "number",
              "hybrid_metrics": [
                "list",
                [
                  "object",
                  {
                    "dimensions": [
                      "set",
                      [
                        "object",
                        {
                          "dimension_key": "string",
                          "dimension_value": "string"
                        }
                      ]
                    ],
                    "expression": "string",
                    "id": "string",
                    "metric_name": "string",
                    "statistic": "string"
                  }
                ]
              ],
              "hybrid_monitor_namespace": "string",
              "id": "string",
              "initial_max_size": "number",
              "metric_name": "string",
              "metric_type": "string",
              "min_adjustment_magnitude": "number",
              "name": "string",
              "predictive_scaling_mode": "string",
              "predictive_task_buffer_time": "number",
              "predictive_value_behavior": "string",
              "predictive_value_buffer": "number",
              "scale_in_evaluation_count": "number",
              "scale_out_evaluation_count": "number",
              "scaling_group_id": "string",
              "scaling_rule_ari": "string",
              "step_adjustment": [
                "list",
                [
                  "object",
                  {
                    "metric_interval_lower_bound": "string",
                    "metric_interval_upper_bound": "string",
                    "scaling_adjustment": "number"
                  }
                ]
              ],
              "target_value": "number",
              "type": "string"
            }
          ]
        ]
      },
      "scaling_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEssScalingRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEssScalingRules), &result)
	return &result
}
