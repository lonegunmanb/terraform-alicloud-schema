package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpcIpamIpamScopes = `{
  "block": {
    "attributes": {
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
      "ipam_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipam_scope_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipam_scope_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipam_scope_type": {
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
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scopes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "id": "string",
              "ipam_id": "string",
              "ipam_scope_description": "string",
              "ipam_scope_id": "string",
              "ipam_scope_name": "string",
              "ipam_scope_type": "string",
              "region_id": "string",
              "resource_group_id": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ]
            }
          ]
        ]
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

func AlicloudVpcIpamIpamScopesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpcIpamIpamScopes), &result)
	return &result
}
