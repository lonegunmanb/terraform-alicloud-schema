package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudMonitorServiceMetricAlarmRules = `{
  "block": {
    "attributes": {
      "dimensions": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "metric_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "namespace": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_name": {
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
              "composite_expression": [
                "list",
                [
                  "object",
                  {
                    "expression_list": [
                      "list",
                      [
                        "object",
                        {
                          "comparison_operator": "string",
                          "metric_name": "string",
                          "period": "number",
                          "statistics": "string",
                          "threshold": "string"
                        }
                      ]
                    ],
                    "expression_list_join": "string",
                    "expression_raw": "string",
                    "level": "string",
                    "times": "number"
                  }
                ]
              ],
              "contact_groups": "string",
              "dimensions": "string",
              "effective_interval": "string",
              "email_subject": "string",
              "escalations": [
                "list",
                [
                  "object",
                  {
                    "critical": [
                      "list",
                      [
                        "object",
                        {
                          "comparison_operator": "string",
                          "pre_condition": "string",
                          "statistics": "string",
                          "threshold": "string",
                          "times": "number"
                        }
                      ]
                    ],
                    "info": [
                      "list",
                      [
                        "object",
                        {
                          "comparison_operator": "string",
                          "pre_condition": "string",
                          "statistics": "string",
                          "threshold": "string",
                          "times": "number"
                        }
                      ]
                    ],
                    "warn": [
                      "list",
                      [
                        "object",
                        {
                          "comparison_operator": "string",
                          "pre_condition": "string",
                          "statistics": "string",
                          "threshold": "string",
                          "times": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "id": "string",
              "labels": [
                "list",
                [
                  "object",
                  {
                    "key": "string",
                    "value": "string"
                  }
                ]
              ],
              "metric_name": "string",
              "namespace": "string",
              "no_data_policy": "string",
              "no_effective_interval": "string",
              "period": "string",
              "prometheus": [
                "list",
                [
                  "object",
                  {
                    "annotations": [
                      "list",
                      [
                        "object",
                        {
                          "key": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "level": "string",
                    "prom_ql": "string",
                    "times": "number"
                  }
                ]
              ],
              "resources": "string",
              "rule_name": "string",
              "silence_time": "string",
              "source_type": "string",
              "status": "bool",
              "webhook": "string"
            }
          ]
        ]
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudMonitorServiceMetricAlarmRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudMonitorServiceMetricAlarmRules), &result)
	return &result
}
