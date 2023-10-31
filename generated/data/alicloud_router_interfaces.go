package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRouterInterfaces = `{
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
      "interfaces": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_point_id": "string",
              "creation_time": "string",
              "description": "string",
              "health_check_source_ip": "string",
              "health_check_target_ip": "string",
              "id": "string",
              "name": "string",
              "opposite_interface_id": "string",
              "opposite_interface_owner_id": "string",
              "opposite_region_id": "string",
              "opposite_router_id": "string",
              "opposite_router_type": "string",
              "role": "string",
              "router_id": "string",
              "router_type": "string",
              "specification": "string",
              "status": "string",
              "vpc_id": "string"
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
      "opposite_interface_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "opposite_interface_owner_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "router_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "router_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "specification": {
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

func AlicloudRouterInterfacesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRouterInterfaces), &result)
	return &result
}
