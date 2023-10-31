package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlbAscripts = `{
  "block": {
    "attributes": {
      "ascript_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ascripts": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ascript_id": "string",
              "ascript_name": "string",
              "enabled": "bool",
              "ext_attribute_enabled": "bool",
              "ext_attributes": [
                "list",
                [
                  "object",
                  {
                    "attribute_key": "string",
                    "attribute_value": "string"
                  }
                ]
              ],
              "id": "string",
              "listener_id": "string",
              "load_balancer_id": "string",
              "position": "string",
              "script_content": "string",
              "status": "string"
            }
          ]
        ]
      },
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
      "listener_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudAlbAscriptsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlbAscripts), &result)
	return &result
}
