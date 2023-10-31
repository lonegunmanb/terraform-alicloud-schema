package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEssScalingGroups = `{
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
              "active_capacity": "number",
              "active_scaling_configuration": "string",
              "cooldown_time": "number",
              "creation_time": "string",
              "db_instance_ids": [
                "list",
                "string"
              ],
              "group_deletion_protection": "bool",
              "health_check_type": "string",
              "id": "string",
              "launch_template_id": "string",
              "launch_template_version": "string",
              "lifecycle_state": "string",
              "load_balancer_ids": [
                "list",
                "string"
              ],
              "max_size": "number",
              "min_size": "number",
              "modification_time": "string",
              "name": "string",
              "pending_capacity": "number",
              "region_id": "string",
              "removal_policies": [
                "list",
                "string"
              ],
              "removing_capacity": "number",
              "suspended_processes": [
                "list",
                "string"
              ],
              "tags": [
                "map",
                "string"
              ],
              "total_capacity": "number",
              "total_instance_count": "number",
              "vpc_id": "string",
              "vswitch_id": "string",
              "vswitch_ids": [
                "list",
                "string"
              ]
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEssScalingGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEssScalingGroups), &result)
	return &result
}
