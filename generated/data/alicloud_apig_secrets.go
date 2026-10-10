package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudApigSecrets = `{
  "block": {
    "attributes": {
      "gateway_type": {
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
      "name_like": {
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
      },
      "secrets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_timestamp": "string",
              "description": "string",
              "gateway_type": "string",
              "id": "string",
              "kms_config": [
                "list",
                [
                  "object",
                  {
                    "kms_instance_id": "string",
                    "kms_key_id": "string",
                    "kms_secret_arn": "string",
                    "version_id": "string"
                  }
                ]
              ],
              "name": "string",
              "reference_count": "number",
              "secret_id": "string",
              "secret_source": "string",
              "status": "string",
              "update_timestamp": "string"
            }
          ]
        ]
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

func AlicloudApigSecretsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudApigSecrets), &result)
	return &result
}
