package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOosStateConfigurations = `{
  "block": {
    "attributes": {
      "configurations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "configure_mode": "string",
              "create_time": "string",
              "description": "string",
              "id": "string",
              "parameters": "string",
              "resource_group_id": "string",
              "schedule_expression": "string",
              "schedule_type": "string",
              "state_configuration_id": "string",
              "tags": [
                "map",
                "string"
              ],
              "targets": "string",
              "template_id": "string",
              "template_name": "string",
              "template_version": "string",
              "update_time": "string"
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudOosStateConfigurationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOosStateConfigurations), &result)
	return &result
}
