package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGaForwardingRules = `{
  "block": {
    "attributes": {
      "accelerator_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "forwarding_rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "forwarding_rule_id": "string",
              "forwarding_rule_name": "string",
              "forwarding_rule_status": "string",
              "id": "string",
              "listener_id": "string",
              "priority": "number",
              "rule_actions": [
                "list",
                [
                  "object",
                  {
                    "forward_group_config": [
                      "list",
                      [
                        "object",
                        {
                          "server_group_tuples": [
                            "list",
                            [
                              "object",
                              {
                                "endpoint_group_id": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "order": "number",
                    "rule_action_type": "string"
                  }
                ]
              ],
              "rule_conditions": [
                "list",
                [
                  "object",
                  {
                    "host_config": [
                      "list",
                      [
                        "object",
                        {
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "path_config": [
                      "list",
                      [
                        "object",
                        {
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "rule_condition_type": "string"
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
      "listener_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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

func AlicloudGaForwardingRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGaForwardingRules), &result)
	return &result
}
