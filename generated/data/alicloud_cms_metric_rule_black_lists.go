package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsMetricRuleBlackLists = `{
  "block": {
    "attributes": {
      "category": {
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
      "lists": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "category": "string",
              "create_time": "string",
              "effective_time": "string",
              "enable_end_time": "string",
              "enable_start_time": "string",
              "id": "string",
              "instances": [
                "list",
                "string"
              ],
              "is_enable": "bool",
              "metric_rule_black_list_id": "string",
              "metric_rule_black_list_name": "string",
              "metrics": [
                "list",
                [
                  "object",
                  {
                    "metric_name": "string",
                    "resource": "string"
                  }
                ]
              ],
              "namespace": "string",
              "scope_type": "string",
              "scope_value": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "metric_rule_black_list_id": {
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
      "order": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCmsMetricRuleBlackListsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsMetricRuleBlackLists), &result)
	return &result
}
