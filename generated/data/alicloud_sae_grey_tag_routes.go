package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSaeGreyTagRoutes = `{
  "block": {
    "attributes": {
      "app_id": {
        "description_kind": "plain",
        "required": true,
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
      "routes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "dubbo_rules": [
                "list",
                [
                  "object",
                  {
                    "condition": "string",
                    "group": "string",
                    "items": [
                      "list",
                      [
                        "object",
                        {
                          "cond": "string",
                          "expr": "string",
                          "index": "number",
                          "operator": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "method_name": "string",
                    "service_name": "string",
                    "version": "string"
                  }
                ]
              ],
              "grey_tag_route_name": "string",
              "id": "string",
              "sc_rules": [
                "list",
                [
                  "object",
                  {
                    "condition": "string",
                    "items": [
                      "list",
                      [
                        "object",
                        {
                          "cond": "string",
                          "name": "string",
                          "operator": "string",
                          "type": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "path": "string"
                  }
                ]
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSaeGreyTagRoutesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSaeGreyTagRoutes), &result)
	return &result
}
