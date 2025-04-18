package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSelectdbDbInstance = `{
  "block": {
    "attributes": {
      "admin_pass": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "cache_size": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "cache_size_postpaid": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "cache_size_prepaid": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "cluster_count_postpaid": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "cluster_count_prepaid": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "cpu_postpaid": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "cpu_prepaid": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "db_instance_class": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_instance_description": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enable_public_network": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "engine": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine_minor_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gmt_created": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gmt_expired": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gmt_modified": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_net_infos": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connection_string": "string",
              "db_ip": "string",
              "net_type": "string",
              "port_list": [
                "list",
                [
                  "object",
                  {
                    "port": "string",
                    "protocol": "string"
                  }
                ]
              ],
              "vpc_instance_id": "string",
              "vswitch_id": "string"
            }
          ]
        ]
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
      "memory_postpaid": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "memory_prepaid": {
        "computed": true,
        "description_kind": "plain",
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
        "type": "string"
      },
      "period_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_ip_lists": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "group_name": "string",
              "group_tag": "string",
              "list_net_type": "string",
              "security_ip_list": "string",
              "security_ip_type": "string"
            }
          ]
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sub_domain": {
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
      "upgraded_engine_minor_version": {
        "computed": true,
        "deprecated": true,
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
      "desired_security_ip_lists": {
        "block": {
          "attributes": {
            "group_name": {
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

func AlicloudSelectdbDbInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSelectdbDbInstance), &result)
	return &result
}
