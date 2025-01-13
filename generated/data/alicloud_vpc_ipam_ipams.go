package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpcIpamIpams = `{
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
      "ipam_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipams": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "default_resource_discovery_association_id": "string",
              "default_resource_discovery_id": "string",
              "id": "string",
              "ipam_description": "string",
              "ipam_id": "string",
              "ipam_name": "string",
              "private_default_scope_id": "string",
              "public_default_scope_id": "string",
              "region_id": "string",
              "resource_discovery_association_count": "number",
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

func AlicloudVpcIpamIpamsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpcIpamIpams), &result)
	return &result
}
