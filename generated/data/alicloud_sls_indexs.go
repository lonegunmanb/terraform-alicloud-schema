package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlsIndexs = `{
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
      "indexs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "keys": "string",
              "line": [
                "list",
                [
                  "object",
                  {
                    "case_sensitive": "bool",
                    "chn": "bool",
                    "exclude_keys": [
                      "list",
                      "string"
                    ],
                    "include_keys": [
                      "list",
                      "string"
                    ],
                    "token": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "log_reduce_black_list": [
                "list",
                "string"
              ],
              "log_reduce_white_list": [
                "list",
                "string"
              ],
              "max_text_len": "number",
              "ttl": "number"
            }
          ]
        ]
      },
      "logstore_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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

func AlicloudSlsIndexsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlsIndexs), &result)
	return &result
}
