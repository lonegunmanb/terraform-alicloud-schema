package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOosTemplates = `{
  "block": {
    "attributes": {
      "category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_by": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_date": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_date_after": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "has_trigger": {
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
      "share_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sort_field": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sort_order": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "template_format": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "templates": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "category": "string",
              "created_by": "string",
              "created_date": "string",
              "description": "string",
              "has_trigger": "bool",
              "id": "string",
              "share_type": "string",
              "tags": [
                "map",
                "string"
              ],
              "template_format": "string",
              "template_id": "string",
              "template_name": "string",
              "template_type": "string",
              "template_version": "string",
              "updated_by": "string",
              "updated_date": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudOosTemplatesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOosTemplates), &result)
	return &result
}
