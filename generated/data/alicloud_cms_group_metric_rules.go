package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsGroupMetricRules = `{
  "block": {
    "attributes": {
      "dimensions": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_state": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "group_metric_rule_name": {
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
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "contact_groups": "string",
              "dimensions": "string",
              "effective_interval": "string",
              "email_subject": "string",
              "enable_state": "bool",
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
                          "statistics": "string",
                          "threshold": "string",
                          "times": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "group_id": "string",
              "group_metric_rule_name": "string",
              "id": "string",
              "metric_name": "string",
              "namespace": "string",
              "no_effective_interval": "string",
              "period": "number",
              "resources": "string",
              "rule_id": "string",
              "silence_time": "number",
              "source_type": "string",
              "status": "string",
              "webhook": "string"
            }
          ]
        ]
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCmsGroupMetricRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsGroupMetricRules), &result)
	return &result
}
