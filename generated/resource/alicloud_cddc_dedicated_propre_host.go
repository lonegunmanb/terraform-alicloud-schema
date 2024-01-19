package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCddcDedicatedPropreHost = `{
  "block": {
    "attributes": {
      "auto_pay": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_renew": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dedicated_host_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ecs_deployment_set_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ecs_host_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ecs_instance_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ecs_instance_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ecs_unique_suffix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ecs_zone_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "engine": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internet_charge_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internet_max_bandwidth_out": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "key_pair_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "os_password": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "password_inherit": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payment_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "period_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_id": {
        "description_kind": "plain",
        "required": true,
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
      "user_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_data_encoded": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "ecs_class_list": {
        "block": {
          "attributes": {
            "data_disk_performance_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sys_disk_capacity": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "sys_disk_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "system_disk_performance_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCddcDedicatedPropreHostSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCddcDedicatedPropreHost), &result)
	return &result
}
