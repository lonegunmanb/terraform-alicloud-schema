package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEssScalingGroup = `{
  "block": {
    "attributes": {
      "allocation_strategy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "az_balance": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "container_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "default_cooldown": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "desired_capacity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "group_deletion_protection": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "group_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_types": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "launch_template_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "launch_template_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "loadbalancer_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "max_instance_lifetime": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_size": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "min_size": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "multi_az_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "on_demand_base_capacity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "on_demand_percentage_above_base_capacity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "protected_instances": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "removal_policies": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scaling_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scaling_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spot_allocation_strategy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spot_instance_pools": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "spot_instance_remedy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "vswitch_id": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "alb_server_group": {
        "block": {
          "attributes": {
            "alb_server_group_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "weight": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "launch_template_override": {
        "block": {
          "attributes": {
            "instance_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "spot_price_limit": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "weighted_capacity": {
              "description_kind": "plain",
              "optional": true,
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

func AlicloudEssScalingGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEssScalingGroup), &result)
	return &result
}
