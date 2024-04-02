package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCsKubernetesPermissions = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "permissions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "is_owner": "bool",
              "is_ram_role": "bool",
              "resource_id": "string",
              "resource_type": "string",
              "role_name": "string",
              "role_type": "string"
            }
          ]
        ]
      },
      "uid": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCsKubernetesPermissionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCsKubernetesPermissions), &result)
	return &result
}
