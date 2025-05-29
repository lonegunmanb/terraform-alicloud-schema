package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlsAlerts = `{
  "block": {
    "attributes": {
      "alerts": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "alert_name": "string",
              "configuration": [
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
                    "auto_annotation": "bool",
                    "condition_configuration": [
                      "list",
                      [
                        "object",
                        {
                          "condition": "string",
                          "count_condition": "string"
                        }
                      ]
                    ],
                    "dashboard": "string",
                    "group_configuration": [
                      "list",
                      [
                        "object",
                        {
                          "fields": [
                            "list",
                            "string"
                          ],
                          "type": "string"
                        }
                      ]
                    ],
                    "join_configurations": [
                      "list",
                      [
                        "object",
                        {
                          "condition": "string",
                          "type": "string"
                        }
                      ]
                    ],
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
                    "mute_until": "number",
                    "no_data_fire": "bool",
                    "no_data_severity": "number",
                    "policy_configuration": [
                      "list",
                      [
                        "object",
                        {
                          "action_policy_id": "string",
                          "alert_policy_id": "string",
                          "repeat_interval": "string"
                        }
                      ]
                    ],
                    "query_list": [
                      "list",
                      [
                        "object",
                        {
                          "chart_title": "string",
                          "dashboard_id": "string",
                          "end": "string",
                          "power_sql_mode": "string",
                          "project": "string",
                          "query": "string",
                          "region": "string",
                          "role_arn": "string",
                          "start": "string",
                          "store": "string",
                          "store_type": "string",
                          "time_span_type": "string",
                          "ui": "string"
                        }
                      ]
                    ],
                    "send_resolved": "bool",
                    "severity_configurations": [
                      "list",
                      [
                        "object",
                        {
                          "eval_condition": [
                            "list",
                            [
                              "object",
                              {
                                "condition": "string",
                                "count_condition": "string"
                              }
                            ]
                          ],
                          "severity": "number"
                        }
                      ]
                    ],
                    "sink_alerthub": [
                      "list",
                      [
                        "object",
                        {
                          "enabled": "bool"
                        }
                      ]
                    ],
                    "sink_cms": [
                      "list",
                      [
                        "object",
                        {
                          "enabled": "bool"
                        }
                      ]
                    ],
                    "sink_event_store": [
                      "list",
                      [
                        "object",
                        {
                          "enabled": "bool",
                          "endpoint": "string",
                          "event_store": "string",
                          "project": "string",
                          "role_arn": "string"
                        }
                      ]
                    ],
                    "tags": [
                      "list",
                      "string"
                    ],
                    "template_configuration": [
                      "list",
                      [
                        "object",
                        {
                          "annotations": [
                            "map",
                            "string"
                          ],
                          "lang": "string",
                          "template_id": "string",
                          "tokens": [
                            "map",
                            "string"
                          ],
                          "type": "string",
                          "version": "string"
                        }
                      ]
                    ],
                    "threshold": "number",
                    "type": "string",
                    "version": "string"
                  }
                ]
              ],
              "description": "string",
              "display_name": "string",
              "id": "string",
              "schedule": [
                "list",
                [
                  "object",
                  {
                    "cron_expression": "string",
                    "delay": "number",
                    "interval": "string",
                    "run_immdiately": "bool",
                    "time_zone": "string",
                    "type": "string"
                  }
                ]
              ]
            }
          ]
        ]
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
      "project_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSlsAlertsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlsAlerts), &result)
	return &result
}
