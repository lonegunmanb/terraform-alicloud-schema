package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGpdbDbExtensions = `{
  "block": {
    "attributes": {
      "database_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "extensions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "current_version": "string",
              "description": "string",
              "extension_id": "string",
              "extension_name": "string",
              "id": "string",
              "is_install_need_restart": "bool",
              "is_latest_version": "bool",
              "latest_version": "string",
              "status": "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudGpdbDbExtensionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGpdbDbExtensions), &result)
	return &result
}
