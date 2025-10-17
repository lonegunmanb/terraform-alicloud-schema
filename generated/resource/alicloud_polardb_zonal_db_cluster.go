package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPolardbZonalDbCluster = `{
  "block": {
    "attributes": {
      "auto_renew_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "cluster_latest_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "db_cluster_nodes_attributes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "db_cluster_nodes_configs": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "map",
          "string"
        ]
      },
      "db_cluster_nodes_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "db_minor_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_node_class": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ens_region_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pay_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "renewal_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_pay_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_space": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "storage_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_minor_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "used_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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

func AlicloudPolardbZonalDbClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPolardbZonalDbCluster), &result)
	return &result
}
