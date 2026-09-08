package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpcRouteTargetGroups = `{
  "block": {
    "attributes": {
      "groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "config_mode": "string",
              "create_time": "string",
              "id": "string",
              "region_id": "string",
              "resource_group_id": "string",
              "route_target_group_description": "string",
              "route_target_group_id": "string",
              "route_target_group_name": "string",
              "route_target_member_list": [
                "list",
                [
                  "object",
                  {
                    "enable_status": "string",
                    "health_check_status": "string",
                    "member_id": "string",
                    "member_type": "string",
                    "weight": "number"
                  }
                ]
              ],
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "vpc_id": "string"
            }
          ]
        ]
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
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_target_group_id": {
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
      "vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "route_target_member_list": {
        "block": {
          "attributes": {
            "enable_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "health_check_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "member_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "member_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "weight": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudVpcRouteTargetGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpcRouteTargetGroups), &result)
	return &result
}
