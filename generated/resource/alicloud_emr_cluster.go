package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEmrCluster = `{
  "block": {
    "attributes": {
      "charge_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "deposit_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "eas_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "emr_ver": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "high_availability_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "is_open_public_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "key_pair_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_pwd": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "meta_store_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "option_software_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "related_cluster_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssh_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "use_local_metadb": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "user_defined_emr_ecs_role": {
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "bootstrap_action": {
        "block": {
          "attributes": {
            "arg": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "execution_fail_strategy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "execution_moment": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "execution_target": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "configs": {
        "block": {
          "attributes": {
            "config_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "config_value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "file_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "service_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "host_group": {
        "block": {
          "attributes": {
            "auto_renew": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "charge_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "decommission_timeout": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "disk_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enable_graceful_decommission": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "gpu_driver": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host_group_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host_group_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_list": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "node_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "period": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "sys_disk_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sys_disk_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "meta_store_conf": {
        "block": {
          "attributes": {
            "db_password": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "db_url": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "db_user_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
      },
      "modify_cluster_service_config": {
        "block": {
          "attributes": {
            "comment": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "config_params": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "config_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "custom_config_params": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "gateway_cluster_id_list": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "group_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host_instance_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "refresh_host_config": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "service_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AlicloudEmrClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEmrCluster), &result)
	return &result
}
