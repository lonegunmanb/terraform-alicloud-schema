package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudKmsKeys = `{
  "block": {
    "attributes": {
      "description_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "filters": {
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
      "keys": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "automatic_rotation": "string",
              "creation_date": "string",
              "creator": "string",
              "delete_date": "string",
              "description": "string",
              "id": "string",
              "key_id": "string",
              "key_spec": "string",
              "key_usage": "string",
              "last_rotation_date": "string",
              "material_expire_time": "string",
              "next_rotation_date": "string",
              "origin": "string",
              "primary_key_version": "string",
              "protection_level": "string",
              "rotation_interval": "string",
              "status": "string"
            }
          ]
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudKmsKeysSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudKmsKeys), &result)
	return &result
}
