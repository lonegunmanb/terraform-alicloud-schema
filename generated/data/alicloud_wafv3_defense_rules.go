package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudWafv3DefenseRules = `{
  "block": {
    "attributes": {
      "defense_type": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "query": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_type": {
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
              "config": [
                "list",
                [
                  "object",
                  {
                    "abroad_regions": "string",
                    "account_identifiers": [
                      "set",
                      [
                        "object",
                        {
                          "decode_type": "string",
                          "key": "string",
                          "position": "string",
                          "priority": "number",
                          "sub_key": "string"
                        }
                      ]
                    ],
                    "auto_update": "bool",
                    "bypass_regular_rules": [
                      "set",
                      "string"
                    ],
                    "bypass_regular_types": [
                      "set",
                      "string"
                    ],
                    "bypass_tags": [
                      "set",
                      "string"
                    ],
                    "cc_effect": "string",
                    "cc_status": "number",
                    "cn_regions": "string",
                    "codec_list": [
                      "set",
                      "string"
                    ],
                    "conditions": [
                      "set",
                      [
                        "object",
                        {
                          "key": "string",
                          "op_value": "string",
                          "sub_key": "string",
                          "values": "string"
                        }
                      ]
                    ],
                    "gray_config": [
                      "list",
                      [
                        "object",
                        {
                          "gray_rate": "number",
                          "gray_sub_key": "string",
                          "gray_target": "string"
                        }
                      ]
                    ],
                    "gray_status": "number",
                    "mode": "number",
                    "protocol": "string",
                    "rate_limit": [
                      "list",
                      [
                        "object",
                        {
                          "interval": "number",
                          "status": [
                            "list",
                            [
                              "object",
                              {
                                "code": "number",
                                "count": "number",
                                "ratio": "number"
                              }
                            ]
                          ],
                          "sub_key": "string",
                          "target": "string",
                          "threshold": "number",
                          "ttl": "number"
                        }
                      ]
                    ],
                    "remote_addr": [
                      "set",
                      "string"
                    ],
                    "rule_action": "string",
                    "throttle_threhold": "number",
                    "throttle_type": "string",
                    "time_config": [
                      "list",
                      [
                        "object",
                        {
                          "time_periods": [
                            "set",
                            [
                              "object",
                              {
                                "end": "number",
                                "start": "number"
                              }
                            ]
                          ],
                          "time_scope": "string",
                          "time_zone": "number",
                          "week_time_periods": [
                            "set",
                            [
                              "object",
                              {
                                "day": "string",
                                "day_periods": [
                                  "set",
                                  [
                                    "object",
                                    {
                                      "end": "number",
                                      "start": "number"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "ua": "string",
                    "url": "string",
                    "waf_base_config": [
                      "set",
                      [
                        "object",
                        {
                          "rule_batch_operation_config": "string",
                          "rule_detail": [
                            "set",
                            [
                              "object",
                              {
                                "rule_action": "string",
                                "rule_id": "string",
                                "rule_status": "number"
                              }
                            ]
                          ],
                          "rule_type": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "defense_origin": "string",
              "defense_scene": "string",
              "defense_type": "string",
              "gmt_modified": "string",
              "id": "string",
              "resource": "string",
              "rule_id": "number",
              "rule_name": "string",
              "rule_status": "number",
              "template_id": "number"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudWafv3DefenseRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudWafv3DefenseRules), &result)
	return &result
}
