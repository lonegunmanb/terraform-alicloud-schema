package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAutoProvisioningGroup = `{
  "block": {
    "attributes": {
      "auto_provisioning_group_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_provisioning_group_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_target_capacity_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "excess_capacity_termination_policy": {
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
        "required": true,
        "type": "string"
      },
      "launch_template_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_spot_price": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "pay_as_you_go_allocation_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pay_as_you_go_target_capacity": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spot_allocation_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spot_instance_interruption_behavior": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spot_instance_pools_to_use_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "spot_target_capacity": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "terminate_instances": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "terminate_instances_with_expiration": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "total_target_capacity": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "valid_from": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "valid_until": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "launch_template_config": {
        "block": {
          "attributes": {
            "instance_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_price": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "priority": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vswitch_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "weighted_capacity": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudAutoProvisioningGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAutoProvisioningGroup), &result)
	return &result
}
