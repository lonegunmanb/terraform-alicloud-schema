package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSaeApplicationScalingRules = `{
  "block": {
    "attributes": {
      "app_id": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "app_id": "string",
              "create_time": "string",
              "id": "string",
              "scaling_rule_enable": "bool",
              "scaling_rule_metric": [
                "list",
                [
                  "object",
                  {
                    "max_replicas": "number",
                    "metrics": [
                      "list",
                      [
                        "object",
                        {
                          "metric_target_average_utilization": "number",
                          "metric_type": "string"
                        }
                      ]
                    ],
                    "metrics_status": [
                      "list",
                      [
                        "object",
                        {
                          "current_metrics": [
                            "list",
                            [
                              "object",
                              {
                                "current_value": "number",
                                "name": "string",
                                "type": "string"
                              }
                            ]
                          ],
                          "current_replicas": "number",
                          "desired_replicas": "number",
                          "last_scale_time": "string",
                          "max_replicas": "number",
                          "min_replicas": "number",
                          "next_scale_metrics": [
                            "list",
                            [
                              "object",
                              {
                                "name": "string",
                                "next_scale_in_average_utilization": "number",
                                "next_scale_out_average_utilization": "number"
                              }
                            ]
                          ],
                          "next_scale_time_period": "number"
                        }
                      ]
                    ],
                    "min_replicas": "number",
                    "scale_down_rules": [
                      "list",
                      [
                        "object",
                        {
                          "disabled": "bool",
                          "stabilization_window_seconds": "number",
                          "step": "number"
                        }
                      ]
                    ],
                    "scale_up_rules": [
                      "list",
                      [
                        "object",
                        {
                          "disabled": "bool",
                          "stabilization_window_seconds": "number",
                          "step": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "scaling_rule_name": "string",
              "scaling_rule_timer": [
                "list",
                [
                  "object",
                  {
                    "begin_date": "string",
                    "end_date": "string",
                    "period": "string",
                    "schedules": [
                      "list",
                      [
                        "object",
                        {
                          "at_time": "string",
                          "max_replicas": "number",
                          "min_replicas": "number",
                          "target_replicas": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "scaling_rule_type": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSaeApplicationScalingRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSaeApplicationScalingRules), &result)
	return &result
}
