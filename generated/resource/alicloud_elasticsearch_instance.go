package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudElasticsearchInstance = `{
  "block": {
    "attributes": {
      "auto_renew_duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "client_node_amount": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "client_node_spec": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_node_amount": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "data_node_disk_encrypted": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "data_node_disk_performance_level": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_node_disk_size": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "data_node_disk_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "data_node_spec": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_kibana_private_network": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_kibana_public_network": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_public": {
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
      "instance_charge_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kibana_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kibana_node_spec": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kibana_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "kibana_private_security_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kibana_private_whitelist": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "kibana_whitelist": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
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
      "master_node_disk_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_node_spec": {
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
      "period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "private_whitelist": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "protocol": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "public_whitelist": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "renew_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "renewal_duration_unit": {
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
      "setting_config": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "status": {
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
      "version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "warm_node_amount": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "warm_node_disk_encrypted": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "warm_node_disk_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "warm_node_disk_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "warm_node_spec": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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

func AlicloudElasticsearchInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudElasticsearchInstance), &result)
	return &result
}
