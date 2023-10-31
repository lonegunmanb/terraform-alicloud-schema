package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRdsUpgradeDbInstance = `{
  "block": {
    "attributes": {
      "acl": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_upgrade_minor_version": {
        "computed": true,
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
      "certificate": {
        "description_kind": "plain",
        "optional": true,
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
      "collect_stat_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connection_string": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_string_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_class": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_instance_description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_storage": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "db_instance_storage_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dedicated_host_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deletion_protection": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "direction": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "force_restart": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ha_mode": {
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
      "instance_network_type": {
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
      "password": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payment_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      "private_key": {
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_arn": {
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
      "source_biz": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_db_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ssl_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "switch_over": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "switch_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "switch_time_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sync_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_major_version": {
        "description_kind": "plain",
        "required": true,
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
      "zone_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone_id_slave_1": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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

func AlicloudRdsUpgradeDbInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRdsUpgradeDbInstance), &result)
	return &result
}
