package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsEventRules = `{
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
      "name_prefix": {
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
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "event_pattern": [
                "list",
                [
                  "object",
                  {
                    "event_type_list": [
                      "list",
                      "string"
                    ],
                    "keyword_filter": [
                      "set",
                      [
                        "object",
                        {
                          "key_words": [
                            "list",
                            "string"
                          ],
                          "relation": "string"
                        }
                      ]
                    ],
                    "level_list": [
                      "list",
                      "string"
                    ],
                    "name_list": [
                      "list",
                      "string"
                    ],
                    "product": "string",
                    "sql_filter": "string"
                  }
                ]
              ],
              "event_rule_name": "string",
              "event_type": "string",
              "group_id": "string",
              "id": "string",
              "silence_time": "number",
              "status": "string"
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

func AlicloudCmsEventRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsEventRules), &result)
	return &result
}
