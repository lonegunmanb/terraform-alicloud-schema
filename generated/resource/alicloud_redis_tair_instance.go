package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRedisTairInstance = `{
  "block": {
    "attributes": {
      "architecture_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_renew": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_renew_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backup_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_backup_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connection_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_string_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "effective_time": {
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
      "force_upgrade": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "global_instance_id": {
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
      "instance_class": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "intranet_bandwidth": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_connections": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "modify_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "node_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "param_no_loose_sentinel_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "param_repl_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "param_semisync_repl_timeout": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "param_sentinel_compat_enable": {
        "computed": true,
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
      "payment_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
        "type": "number"
      },
      "read_only_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "recover_config_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secondary_zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_ip_group_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_ips": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "shard_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "slave_read_only_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "src_db_instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_performance_level": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_size_gb": {
        "computed": true,
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
      "tair_instance_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tair_instance_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_auth_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zone_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
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

func AlicloudRedisTairInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRedisTairInstance), &result)
	return &result
}
