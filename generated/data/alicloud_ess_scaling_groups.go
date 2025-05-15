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
              "allocation_strategy": "string",
              "az_balance": "bool",
              "capacity_options_compensate_with_on_demand": "bool",
              "capacity_options_on_demand_base_capacity": "number",
              "capacity_options_on_demand_percentage_above_base_capacity": "number",
              "capacity_options_spot_auto_replace_on_demand": "bool",
              "compensate_with_on_demand": "bool",
              "cooldown_time": "number",
              "creation_time": "string",
              "db_instance_ids": [
                "list",
                "string"
              ],
              "desired_capacity": "number",
              "enable_desired_capacity": "bool",
              "group_deletion_protection": "bool",
              "group_type": "string",
              "health_check_type": "string",
              "id": "string",
              "init_capacity": "number",
              "launch_template_id": "string",
              "launch_template_override": [
                "list",
                [
                  "object",
                  {
                    "instance_type": "string",
                    "spot_price_limit": "number",
                    "weighted_capacity": "number"
                  }
                ]
              ],
              "launch_template_version": "string",
              "lifecycle_state": "string",
              "load_balancer_ids": [
                "list",
                "string"
              ],
              "max_instance_lifetime": "number",
              "max_size": "number",
              "min_size": "number",
              "modification_time": "string",
              "monitor_group_id": "string",
              "multi_az_policy": "string",
              "name": "string",
              "on_demand_base_capacity": "number",
              "on_demand_percentage_above_base_capacity": "number",
              "pending_capacity": "number",
              "pending_wait_capacity": "number",
              "protected_capacity": "number",
              "region_id": "string",
              "removal_policies": [
                "list",
                "string"
              ],
              "removing_capacity": "number",
              "removing_wait_capacity": "number",
              "resource_group_id": "string",
              "scaling_policy": "string",
              "spot_allocation_strategy": "string",
              "spot_capacity": "number",
              "spot_instance_pools": "number",
              "spot_instance_remedy": "bool",
              "standby_capacity": "number",
              "stop_instance_timeout": "number",
              "stopped_capacity": "number",
              "suspended_processes": [
                "list",
                "string"
              ],
              "system_suspended": "bool",
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
