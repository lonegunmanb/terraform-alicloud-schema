package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMongodbShardingInstance = `{
  "block": {
    "attributes": {
      "account_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "auto_renew": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "backup_interval": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backup_period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "backup_retention_period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "backup_retention_policy_on_cluster_deletion": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "backup_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_release_protection": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_backup_log": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "engine_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "global_security_group_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "hidden_zone_id": {
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
      "instance_charge_type": {
        "computed": true,
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
      "log_backup_retention_period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "maintain_end_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "maintain_start_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "order_type": {
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
      "protocol_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "provisioned_iops": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retention_period": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "secondary_zone_id": {
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
      "security_ip_list": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "snapshot_backup_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_action": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_engine": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_type": {
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
      "tde_status": {
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
      "zone_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "config_server_list": {
        "block": {
          "attributes": {
            "connect_string": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "max_connections": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "max_iops": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "node_class": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "node_description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "node_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "node_storage": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "mongo_list": {
        "block": {
          "attributes": {
            "connect_string": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "node_class": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "node_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 32,
        "min_items": 2,
        "nesting_mode": "list"
      },
      "shard_list": {
        "block": {
          "attributes": {
            "node_class": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "node_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "node_storage": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "readonly_replicas": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 32,
        "min_items": 2,
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

func AlicloudMongodbShardingInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMongodbShardingInstance), &result)
	return &result
}
