package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEssScalingGroup = `{
  "block": {
    "attributes": {
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
      "id": {
        "computed": true,
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
      "scaling_group_name": {
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
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEssScalingGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEssScalingGroup), &result)
	return &result
}
