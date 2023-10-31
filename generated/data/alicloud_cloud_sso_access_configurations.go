package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudSsoAccessConfigurations = `{
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
              "access_configuration_id": "string",
              "access_configuration_name": "string",
              "create_time": "string",
              "description": "string",
              "directory_id": "string",
              "id": "string",
              "permission_policies": [
                "list",
                [
                  "object",
                  {
                    "add_time": "string",
                    "permission_policy_document": "string",
                    "permission_policy_name": "string",
                    "permission_policy_type": "string"
                  }
                ]
              ],
              "relay_state": "string",
              "session_duration": "number",
              "status_notifications": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "directory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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

func AlicloudCloudSsoAccessConfigurationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudSsoAccessConfigurations), &result)
	return &result
}
