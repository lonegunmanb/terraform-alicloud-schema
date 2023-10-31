package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPolardbNodeClasses = `{
  "block": {
    "attributes": {
      "category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "classes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "supported_engines": [
                "list",
                [
                  "object",
                  {
                    "available_resources": [
                      "list",
                      [
                        "object",
                        {
                          "db_node_class": "string"
                        }
                      ]
                    ],
                    "engine": "string"
                  }
                ]
              ],
              "zone_id": "string"
            }
          ]
        ]
      },
      "db_node_class": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_version": {
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pay_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudPolardbNodeClassesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPolardbNodeClasses), &result)
	return &result
}
