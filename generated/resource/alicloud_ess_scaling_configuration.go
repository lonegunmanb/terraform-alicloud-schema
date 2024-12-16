package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEssScalingConfiguration = `{
  "block": {
    "attributes": {
      "active": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "credit_specification": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deletion_protection": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "force_delete": {
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
      "image_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_options_login_as_non_root": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "instance_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_ids": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
      "instance_types": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "internet_charge_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internet_max_bandwidth_in": {
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
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "is_outdated": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "key_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_encrypted_password": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_encryption_context": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "override": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "password": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "password_inherit": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scaling_configuration_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scaling_group_id": {
        "description_kind": "plain",
        "required": true,
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
        "type": "number"
      },
      "spot_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "substitute": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_disk_auto_snapshot_policy_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_disk_category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_disk_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_disk_encrypt_algorithm": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_disk_encrypted": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "system_disk_kms_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_disk_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_disk_performance_level": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_disk_provisioned_iops": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "system_disk_size": {
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
      "user_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "custom_priorities": {
        "block": {
          "attributes": {
            "instance_type": {
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
        "nesting_mode": "set"
      },
      "data_disk": {
        "block": {
          "attributes": {
            "auto_snapshot_policy_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
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
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encrypted": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "kms_key_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
            "provisioned_iops": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
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
        "nesting_mode": "list"
      },
      "instance_pattern_info": {
        "block": {
          "attributes": {
            "architectures": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "burstable_performance": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cores": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "excluded_instance_types": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "instance_family_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_price": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "memory": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "instance_type_override": {
        "block": {
          "attributes": {
            "instance_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "weighted_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 10,
        "nesting_mode": "set"
      },
      "network_interfaces": {
        "block": {
          "attributes": {
            "instance_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ipv6_address_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "network_interface_traffic_mode": {
              "computed": true,
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "spot_price_limit": {
        "block": {
          "attributes": {
            "instance_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "price_limit": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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

func AlicloudEssScalingConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEssScalingConfiguration), &result)
	return &result
}
