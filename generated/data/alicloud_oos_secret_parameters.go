package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOosSecretParameters = `{
  "block": {
    "attributes": {
      "enable_details": {
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
      "parameters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "constraints": "string",
              "create_time": "string",
              "created_by": "string",
              "description": "string",
              "id": "string",
              "key_id": "string",
              "parameter_version": "number",
              "resource_group_id": "string",
              "secret_parameter_id": "string",
              "secret_parameter_name": "string",
              "share_type": "string",
              "tags": [
                "map",
                "string"
              ],
              "type": "string",
              "updated_by": "string",
              "updated_date": "string"
            }
          ]
        ]
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secret_parameter_name": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudOosSecretParametersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOosSecretParameters), &result)
	return &result
}
