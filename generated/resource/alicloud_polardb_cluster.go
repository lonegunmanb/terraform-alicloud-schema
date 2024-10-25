package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPolardbCluster = `{
  "block": {
    "attributes": {
      "allow_shut_down": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_renew_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "backup_retention_policy_on_cluster_deletion": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "clone_data_point": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "collector_status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "compress_storage": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connection_string": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_category": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_option": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_node_class": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_node_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "db_node_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_node_num": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "db_revision_version_list": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "release_note": "string",
              "release_type": "string",
              "revision_version_code": "string",
              "revision_version_name": "string"
            }
          ]
        ]
      },
      "db_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_time_zone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deletion_lock": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encrypt_new_tables": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "from_time_service": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gdn_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hot_replica_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hot_standby_cluster": {
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
      "imci_switch": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "loose_polar_log_bin": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "loose_xengine": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "loose_xengine_use_memory_pct": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "lower_case_table_names": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "maintain_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "modify_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameter_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pay_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "planned_end_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "planned_start_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provisioned_iops": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "proxy_class": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "proxy_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "renewal_status": {
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
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scale_ap_ro_num_max": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "scale_ap_ro_num_min": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "scale_max": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "scale_min": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "scale_ro_num_max": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "scale_ro_num_min": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "seconds_until_auto_pause": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "security_ips": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "serverless_steady_switch": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "serverless_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_resource_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_pay_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_space": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "storage_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sub_category": {
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
      "target_db_revision_version_code": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tde_region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tde_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "upgrade_type": {
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
      "db_cluster_ip_array": {
        "block": {
          "attributes": {
            "db_cluster_ip_array_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "modify_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_ips": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "parameters": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
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

func AlicloudPolardbClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPolardbCluster), &result)
	return &result
}
