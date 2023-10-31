package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDcdnWafRules = `{
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "query_args": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "waf_rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "action": "string",
              "cc_status": "string",
              "cn_region_list": "string",
              "conditions": [
                "list",
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
              "defense_scene": "string",
              "effect": "string",
              "gmt_modified": "string",
              "id": "string",
              "other_region_list": "string",
              "policy_id": "string",
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
                          "code": "string",
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
              "regular_rules": [
                "list",
                "string"
              ],
              "regular_types": [
                "list",
                "string"
              ],
              "remote_addr": [
                "list",
                "string"
              ],
              "rule_name": "string",
              "scenes": [
                "list",
                "string"
              ],
              "status": "string",
              "waf_group_ids": "string",
              "waf_rule_id": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDcdnWafRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDcdnWafRules), &result)
	return &result
}
