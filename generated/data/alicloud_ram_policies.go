package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRamPolicies = `{
  "block": {
    "attributes": {
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "group_name": {
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
      "policies": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "attachment_count": "number",
              "create_date": "string",
              "default_version": "string",
              "description": "string",
              "document": "string",
              "id": "string",
              "name": "string",
              "policy_document": "string",
              "policy_name": "string",
              "tags": [
                "map",
                "string"
              ],
              "type": "string",
              "update_date": "string",
              "user_name": "string",
              "version_id": "string"
            }
          ]
        ]
      },
      "role_name": {
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
      },
      "type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudRamPoliciesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRamPolicies), &result)
	return &result
}
