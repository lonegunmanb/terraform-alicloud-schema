package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCsKubernetesNodePool = `{
  "block": {
    "attributes": {
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
      "cis_enabled": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "compensate_with_on_demand": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cpu_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deployment_set_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "desired_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "force_delete": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "format_disk": {
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
      "image_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "install_cloud_monitor": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "instance_charge_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_types": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "instances": {
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
      "internet_max_bandwidth_out": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "keep_instance_name": {
        "computed": true,
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
        "sensitive": true,
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
      "login_as_non_root": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "multi_az_policy": {
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
      "node_count": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "node_name_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "node_pool_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "on_demand_base_capacity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "on_demand_percentage_above_base_capacity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "platform": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rds_instances": {
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
      "runtime_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scaling_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "scaling_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_id": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "security_hardening_os": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "soc_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "spot_instance_pools": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "spot_instance_remedy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "spot_strategy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_disk_bursting_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "system_disk_categories": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "system_disk_category": {
        "computed": true,
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
      "system_disk_kms_key": {
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
      "system_disk_snapshot_policy_id": {
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
      "unschedulable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "update_nodes": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "user_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vswitch_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "data_disks": {
        "block": {
          "attributes": {
            "auto_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "auto_snapshot_policy_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bursting_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "category": {
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
              "type": "string"
            },
            "file_system": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mount_target": {
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
      "kubelet_configuration": {
        "block": {
          "attributes": {
            "allowed_unsafe_sysctls": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "container_log_max_files": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "container_log_max_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cpu_manager_policy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "event_burst": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "event_record_qps": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "eviction_hard": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "eviction_soft": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "eviction_soft_grace_period": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "feature_gates": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "bool"
              ]
            },
            "kube_api_burst": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kube_api_qps": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kube_reserved": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "max_pods": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "read_only_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "registry_burst": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "registry_pull_qps": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "serialize_image_pulls": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "system_reserved": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "labels": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "management": {
        "block": {
          "attributes": {
            "auto_repair": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "auto_upgrade": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "auto_vul_fix": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "max_unavailable": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "surge": {
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "surge_percentage": {
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "auto_repair_policy": {
              "block": {
                "attributes": {
                  "restart_node": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "auto_upgrade_policy": {
              "block": {
                "attributes": {
                  "auto_upgrade_kubelet": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "auto_vul_fix_policy": {
              "block": {
                "attributes": {
                  "restart_node": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "vul_level": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "private_pool_options": {
        "block": {
          "attributes": {
            "private_pool_options_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "private_pool_options_match_criteria": {
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
      "rolling_policy": {
        "block": {
          "attributes": {
            "max_parallelism": {
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
      "rollout_policy": {
        "block": {
          "attributes": {
            "max_unavailable": {
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
      "scaling_config": {
        "block": {
          "attributes": {
            "eip_bandwidth": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "eip_internet_charge_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enable": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "is_bond_eip": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "max_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "type": {
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
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "taints": {
        "block": {
          "attributes": {
            "effect": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "tee_config": {
        "block": {
          "attributes": {
            "tee_enable": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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

func AlicloudCsKubernetesNodePoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCsKubernetesNodePool), &result)
	return &result
}
