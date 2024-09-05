package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRdsDbProxy = `{
  "block": {
    "attributes": {
      "db_proxy_connect_string_port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "db_proxy_connection_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_proxy_connection_string": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_proxy_endpoint_aliases": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_proxy_endpoint_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_proxy_endpoint_read_write_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_proxy_features": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_proxy_instance_num": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "db_proxy_instance_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_proxy_ssl_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "effective_specific_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "effective_time": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_network_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "net_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "read_only_instance_distribution_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "read_only_instance_max_delay_time": {
        "computed": true,
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
      "ssl_expired_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "switch_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "upgrade_time": {
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
      }
    },
    "block_types": {
      "read_only_instance_weight": {
        "block": {
          "attributes": {
            "instance_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "weight": {
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

func AlicloudRdsDbProxySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRdsDbProxy), &result)
	return &result
}
