package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudSsoAccessAssignments = `{
  "block": {
    "attributes": {
      "access_configuration_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "assignments": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_configuration_id": "string",
              "access_configuration_name": "string",
              "directory_id": "string",
              "id": "string",
              "principal_id": "string",
              "principal_name": "string",
              "principal_type": "string",
              "target_id": "string",
              "target_name": "string",
              "target_path_name": "string",
              "target_type": "string"
            }
          ]
        ]
      },
      "directory_id": {
        "description_kind": "plain",
        "required": true,
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
      "principal_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudSsoAccessAssignmentsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudSsoAccessAssignments), &result)
	return &result
}
