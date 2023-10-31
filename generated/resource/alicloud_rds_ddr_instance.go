package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRdsDdrInstance = `{
  "block": {
    "attributes": {
      "acl": {
        "computed": true,
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
      "auto_upgrade_minor_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backup_set_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ca_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "category": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "client_ca_cert": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "client_ca_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "client_cert_revocation_list": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "client_crl_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "connection_string": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_string_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_ip_array_attribute": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_ip_array_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_storage_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "effective_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "engine_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "force_restart": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "fresh_white_list_readins": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ha_config": {
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
      "instance_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_storage": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "instance_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "maintain_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "manual_ha_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "modify_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "monitoring_period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "port": {
        "computed": true,
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
      "released_keep_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replication_acl": {
        "computed": true,
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
      "restore_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "restore_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      "security_ip_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_ip_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "server_cert": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_key": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_db_instance_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sql_collector_config_value": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "sql_collector_status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_action": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_auto_scale": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_threshold": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "storage_upper_bound": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "switch_time": {
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
      "target_minor_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tcp_connection_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tde_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "upgrade_db_instance_kernel_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "upgrade_time": {
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
      "whitelist_network_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone_id_slave_a": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "zone_id_slave_b": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
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
      "pg_hba_conf": {
        "block": {
          "attributes": {
            "address": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "database": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "mask": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "method": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "option": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "priority_id": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "user": {
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

func AlicloudRdsDdrInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRdsDdrInstance), &result)
	return &result
}
