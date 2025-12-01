package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudNlbServerGroup = `{
  "block": {
    "attributes": {
      "address_ip_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "any_port_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "connection_drain": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "connection_drain_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "connection_drain_timeout": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preserve_client_ip_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "protocol": {
        "computed": true,
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
      "scheduler": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "server_group_type": {
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
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "health_check": {
        "block": {
          "attributes": {
            "health_check_connect_port": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "health_check_connect_timeout": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "health_check_domain": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "health_check_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "health_check_exp": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "health_check_http_code": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "health_check_interval": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "health_check_req": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "health_check_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "health_check_url": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "healthy_threshold": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "http_check_method": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "unhealthy_threshold": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AlicloudNlbServerGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudNlbServerGroup), &result)
	return &result
}
