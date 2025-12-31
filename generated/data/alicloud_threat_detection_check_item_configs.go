package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionCheckItemConfigs = `{
  "block": {
    "attributes": {
      "configs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "check_id": "number",
              "check_show_name": "string",
              "check_type": "string",
              "custom_configs": [
                "list",
                [
                  "object",
                  {
                    "default_value": "string",
                    "name": "string",
                    "show_name": "string",
                    "type_define": "string",
                    "value": "string"
                  }
                ]
              ],
              "description": [
                "list",
                [
                  "object",
                  {
                    "type": "string",
                    "value": "string"
                  }
                ]
              ],
              "estimated_count": "number",
              "instance_sub_type": "string",
              "instance_type": "string",
              "risk_level": "string",
              "section_ids": [
                "list",
                "number"
              ],
              "vendor": "string"
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
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "task_sources": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudThreatDetectionCheckItemConfigsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionCheckItemConfigs), &result)
	return &result
}
