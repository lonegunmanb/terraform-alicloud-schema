package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCsEdgeKubernetes = `{
  "block": {
    "attributes": {
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
      "cluster_spec": {
        "computed": true,
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
      "deletion_protection": {
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
      "kube_config": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancer_spec": {
        "computed": true,
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
      "name_prefix": {
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
      "password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "pod_cidr": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "service_cidr": {
        "description_kind": "plain",
        "optional": true,
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "worker_instance_types": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "worker_nodes": {
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
      "worker_number": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "worker_ram_role_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "worker_vswitch_ids": {
        "description_kind": "plain",
        "required": true,
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
          "deprecated": true,
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCsEdgeKubernetesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCsEdgeKubernetes), &result)
	return &result
}
