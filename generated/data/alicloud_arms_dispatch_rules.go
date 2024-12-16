package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudArmsDispatchRules = `{
  "block": {
    "attributes": {
      "dispatch_rule_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
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
              "dispatch_rule_id": "string",
              "dispatch_rule_name": "string",
              "group_rules": [
                "list",
                [
                  "object",
                  {
                    "group_interval": "number",
                    "group_wait_time": "number",
                    "grouping_fields": [
                      "list",
                      "string"
                    ],
                    "repeat_interval": "number"
                  }
                ]
              ],
              "id": "string",
              "label_match_expression_grid": [
                "list",
                [
                  "object",
                  {
                    "label_match_expression_groups": [
                      "list",
                      [
                        "object",
                        {
                          "label_match_expressions": [
                            "list",
                            [
                              "object",
                              {
                                "key": "string",
                                "operator": "string",
                                "value": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "notify_rules": [
                "list",
                [
                  "object",
                  {
                    "notify_channels": [
                      "list",
                      "string"
                    ],
                    "notify_end_time": "string",
                    "notify_objects": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "notify_object_id": "string",
                          "notify_type": "string"
                        }
                      ]
                    ],
                    "notify_start_time": "string"
                  }
                ]
              ],
              "notify_template": [
                "list",
                [
                  "object",
                  {
                    "email_content": "string",
                    "email_recover_content": "string",
                    "email_recover_title": "string",
                    "email_title": "string",
                    "robot_content": "string",
                    "sms_content": "string",
                    "sms_recover_content": "string",
                    "tts_content": "string",
                    "tts_recover_content": "string"
                  }
                ]
              ],
              "status": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudArmsDispatchRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudArmsDispatchRules), &result)
	return &result
}
