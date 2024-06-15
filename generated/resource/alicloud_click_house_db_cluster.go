package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudClickHouseDbCluster = `{
  "block": {
    "attributes": {
      "category": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connection_string": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_class": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_cluster_description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_cluster_network_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_cluster_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_node_group_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "db_node_storage": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "encryption_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_type": {
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
      "maintain_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
        "type": "string"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "renewal_status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "used_time": {
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
      "db_cluster_access_white_list": {
        "block": {
          "attributes": {
            "db_cluster_ip_array_attribute": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "db_cluster_ip_array_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_ip_list": {
              "description_kind": "plain",
              "optional": true,
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

func AlicloudClickHouseDbClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudClickHouseDbCluster), &result)
	return &result
}
