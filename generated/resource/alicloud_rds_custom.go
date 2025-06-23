package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRdsCustom = `{
  "block": {
    "attributes": {
      "amount": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "auto_pay": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_renew": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "create_extra_param": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deployment_set_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "direction": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dry_run": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "force": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "force_stop": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "host_name": {
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
      "image_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_charge_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_type": {
        "description_kind": "plain",
        "required": true,
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
      "io_optimized": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_pair_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "password": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "period_unit": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_enhancement_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "spot_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "support_case": {
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
      "vswitch_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "data_disk": {
        "block": {
          "attributes": {
            "category": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "performance_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "system_disk": {
        "block": {
          "attributes": {
            "category": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "size": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AlicloudRdsCustomSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRdsCustom), &result)
	return &result
}
