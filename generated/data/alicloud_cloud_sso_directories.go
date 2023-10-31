package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudSsoDirectories = `{
  "block": {
    "attributes": {
      "directories": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "directory_id": "string",
              "directory_name": "string",
              "id": "string",
              "mfa_authentication_status": "string",
              "region": "string",
              "saml_identity_provider_configuration": [
                "list",
                [
                  "object",
                  {
                    "create_time": "string",
                    "encoded_metadata_document": "string",
                    "entity_id": "string",
                    "login_url": "string",
                    "sso_status": "string"
                  }
                ]
              ],
              "scim_synchronization_status": "string",
              "tasks": [
                "list",
                [
                  "object",
                  {
                    "access_configuration_id": "string",
                    "access_configuration_name": "string",
                    "end_time": "string",
                    "failure_reason": "string",
                    "principal_id": "string",
                    "principal_name": "string",
                    "principal_type": "string",
                    "start_time": "string",
                    "status": "string",
                    "target_id": "string",
                    "target_name": "string",
                    "target_path": "string",
                    "target_type": "string",
                    "task_id": "string",
                    "task_type": "string"
                  }
                ]
              ]
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

func AlicloudCloudSsoDirectoriesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudSsoDirectories), &result)
	return &result
}
