package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionHoneyPots = `{
  "block": {
    "attributes": {
      "honeypot_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "honeypot_name": {
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
      "name_regex": {
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
      "pots": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "honeypot_id": "string",
              "honeypot_image_id": "string",
              "honeypot_image_name": "string",
              "honeypot_name": "string",
              "id": "string",
              "node_id": "string",
              "preset_id": "string",
              "state": [
                "list",
                "string"
              ],
              "status": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudThreatDetectionHoneyPotsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionHoneyPots), &result)
	return &result
}
