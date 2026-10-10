package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRealtimeComputeVariables = `{
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
        "required": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "variables": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "id": "string",
              "kind": "string",
              "name": "string",
              "namespace": "string",
              "value": "string",
              "workspace": "string"
            }
          ]
        ]
      },
      "workspace": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudRealtimeComputeVariablesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRealtimeComputeVariables), &result)
	return &result
}
