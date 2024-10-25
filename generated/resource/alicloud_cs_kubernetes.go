package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCsKubernetes = `{
  "block": {
    "attributes": {
      "api_audiences": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_authority": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "client_cert": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "client_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_ca_cert": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_domain": {
        "description": "cluster local domain",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_network_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connections": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "cpu_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_san": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deletion_protection": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_ssh": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "exclude_autoscaler_nodes": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "force_update": {
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
      "install_cloud_monitor": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "is_enterprise_security_group": {
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
      "kube_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancer_spec": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_auto_renew": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "master_auto_renew_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "master_disk_category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_disk_performance_level": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_disk_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "master_disk_snapshot_policy_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_instance_charge_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_instance_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_instance_types": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "master_nodes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "name": "string",
              "private_ip": "string"
            }
          ]
        ]
      },
      "master_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "master_period_unit": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_vswitch_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nat_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "new_nat_gateway": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "node_cidr_mask": {
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
      "node_port_range": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nodes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "os_type": {
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
      "platform": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pod_cidr": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pod_vswitch_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "proxy_mode": {
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
      "retain_resources": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "runtime": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "security_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_account_issuer": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_cidr": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "slb_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "slb_internet": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "slb_internet_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "slb_intranet": {
        "computed": true,
        "description_kind": "plain",
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
      "timezone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_ca": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "worker_auto_renew": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "worker_auto_renew_period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "worker_data_disk_category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "worker_data_disk_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "worker_disk_category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "worker_disk_performance_level": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "worker_disk_size": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "worker_disk_snapshot_policy_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "worker_instance_charge_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "worker_instance_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "worker_instance_types": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "worker_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "worker_numbers": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "number"
        ]
      },
      "worker_period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "worker_period_unit": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "worker_ram_role_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "worker_vswitch_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "addons": {
        "block": {
          "attributes": {
            "config": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "delete_options": {
        "block": {
          "attributes": {
            "delete_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "log_config": {
        "block": {
          "attributes": {
            "project": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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
              "optional": true,
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
      },
      "worker_data_disks": {
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
            "size": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
      "worker_nodes": {
        "block": {
          "attributes": {
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "private_ip": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCsKubernetesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCsKubernetes), &result)
	return &result
}
