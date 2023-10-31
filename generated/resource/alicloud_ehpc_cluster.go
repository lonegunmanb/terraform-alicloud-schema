package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEhpcCluster = `{
  "block": {
    "attributes": {
      "account_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_renew": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_renew_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "client_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "compute_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "compute_enable_ht": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "compute_instance_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "compute_spot_price_limit": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "compute_spot_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deploy_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ecs_charge_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ehpc_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ha_enable": {
        "computed": true,
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
      "image_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_owner_alias": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "input_file_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "is_compute_ess": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "job_queue": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_pair_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "login_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "login_instance_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "manager_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "manager_instance_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "os_tag": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
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
      "plugin": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ram_node_types": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "ram_role_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "release_instance": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "remote_directory": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remote_vis_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scc_cluster_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scheduler_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "system_disk_level": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_disk_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "system_disk_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "volume_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "volume_mount_option": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "volume_mountpoint": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "volume_protocol": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "volume_type": {
        "computed": true,
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
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "without_agent": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "without_elastic_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "zone_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "additional_volumes": {
        "block": {
          "attributes": {
            "job_queue": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "local_directory": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "location": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "remote_directory": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "volume_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "volume_mount_option": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "volume_mountpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "volume_protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "volume_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "roles": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 8,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 10,
        "nesting_mode": "set"
      },
      "application": {
        "block": {
          "attributes": {
            "tag": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 100,
        "nesting_mode": "set"
      },
      "post_install_script": {
        "block": {
          "attributes": {
            "args": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 16,
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

func AlicloudEhpcClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEhpcCluster), &result)
	return &result
}
