package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAdbDbClusterLakeVersion = `{
  "block": {
    "attributes": {
      "backup_set_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "commodity_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "compute_resource": {
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
      "db_cluster_description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_cluster_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "disk_encryption": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_default_resource_group": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_ssl": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "engine": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expire_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expired": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lock_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "lock_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
        "type": "string"
      },
      "product_form": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "product_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reserved_node_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "reserved_node_size": {
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
      "restore_to_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "restore_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secondary_vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secondary_zone_id": {
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
      "source_db_cluster_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_resource": {
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

func AlicloudAdbDbClusterLakeVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAdbDbClusterLakeVersion), &result)
	return &result
}
