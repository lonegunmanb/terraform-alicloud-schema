package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionAssets = `{
  "block": {
    "attributes": {
      "assets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "id": "string",
              "uuid": "string"
            }
          ]
        ]
      },
      "criteria": {
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
      "importance": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "logical_exp": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "machine_types": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "no_group_trace": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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

func AlicloudThreatDetectionAssetsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionAssets), &result)
	return &result
}
