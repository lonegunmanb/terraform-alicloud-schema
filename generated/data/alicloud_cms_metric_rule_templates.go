package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsMetricRuleTemplates = `{
  "block": {
    "attributes": {
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "keyword": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metric_rule_template_name": {
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "templates": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "alert_templates": [
                "list",
                [
                  "object",
                  {
                    "category": "string",
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
                                "times": "string"
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
                                "times": "string"
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
                                "times": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "metric_name": "string",
                    "namespace": "string",
                    "rule_name": "string",
                    "selector": "string",
                    "webhook": "string"
                  }
                ]
              ],
              "description": "string",
              "group_id": "string",
              "id": "string",
              "metric_rule_template_name": "string",
              "rest_version": "string",
              "template_id": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCmsMetricRuleTemplatesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsMetricRuleTemplates), &result)
	return &result
}
