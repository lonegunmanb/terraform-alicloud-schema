package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionHoneypotPresets = `{
  "block": {
    "attributes": {
      "current_page": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "honeypot_image_name": {
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
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_name": {
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
      "preset_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "presets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "honeypot_image_name": "string",
              "honeypot_preset_id": "string",
              "id": "string",
              "meta": [
                "list",
                [
                  "object",
                  {
                    "burp": "string",
                    "portrait_option": "bool",
                    "trojan_git": "string"
                  }
                ]
              ],
              "node_id": "string",
              "preset_name": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudThreatDetectionHoneypotPresetsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionHoneypotPresets), &result)
	return &result
}
