package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcsLaunchTemplate = `{
  "block": {
    "attributes": {
      "auto_release_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_renew": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_renew_period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "default_version_number": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "enable_vm_os_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "host_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "http_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "http_put_response_hop_limit": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "http_tokens": {
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
      "image_owner_alias": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_charge_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internet_charge_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internet_max_bandwidth_in": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "latest_version_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "launch_template_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "password_inherit": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "period_unit": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_ip_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ram_role_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_enhancement_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_id": {
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
      "spot_duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spot_price_limit": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "spot_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_disk_category": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_disk_description": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_disk_name": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_disk_size": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "template_resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "update_default_version_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "user_data": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "userdata": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "data_disks": {
        "block": {
          "attributes": {
            "category": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete_with_instance": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encrypted": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
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
            },
            "snapshot_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "network_interfaces": {
        "block": {
          "attributes": {
            "delete_on_release": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "primary_ip": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_group_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vswitch_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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
            "delete_with_instance": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encrypted": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "iops": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "performance_level": {
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
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEcsLaunchTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcsLaunchTemplate), &result)
	return &result
}
