package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEnsInstance = `{
  "block": {
    "attributes": {
      "amount": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "auto_release_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_renew": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_use_coupon": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "billing_cycle": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "carrier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ens_region_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "force_stop": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "host_name": {
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
      "image_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "include_data_disks": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "instance_charge_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_name": {
        "computed": true,
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
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ip_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_pair_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "net_district_code": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "net_work_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "password_inherit": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "payment_type": {
        "description_kind": "plain",
        "required": true,
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
      "private_ip_address": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_ip_identification": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "schedule_area_level": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scheduling_price_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scheduling_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "unique_suffix": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "user_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_id": {
        "computed": true,
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
            "disk_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "encrypt_key_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encrypted": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "system_disk": {
        "block": {
          "attributes": {
            "category": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "size": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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

func AlicloudEnsInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEnsInstance), &result)
	return &result
}
