package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudApigPlugins = `{
  "block": {
    "attributes": {
      "gateway_id": {
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "plugin_class_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "plugin_class_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "plugins": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "gateway_id": "string",
              "gateway_name": "string",
              "id": "string",
              "plugin_class_id": "string",
              "plugin_class_name": "string",
              "plugin_id": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudApigPluginsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudApigPlugins), &result)
	return &result
}
