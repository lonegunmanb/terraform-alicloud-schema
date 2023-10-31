package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGraphDatabaseDbInstance = `{
  "block": {
    "attributes": {
      "connection_string": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_category": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_instance_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_network_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_instance_storage_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_node_class": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_node_storage": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "db_version": {
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
      "payment_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
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
      "db_instance_ip_array": {
        "block": {
          "attributes": {
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
            "security_ips": {
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

func AlicloudGraphDatabaseDbInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGraphDatabaseDbInstance), &result)
	return &result
}
