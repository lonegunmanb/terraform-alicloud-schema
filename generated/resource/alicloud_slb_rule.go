package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlbRule = `{
  "block": {
    "attributes": {
      "cookie": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cookie_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "delete_protection_validation": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "domain": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "frontend_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "health_check": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_connect_port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_domain": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_http_code": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "healthy_threshold": {
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
      "listener_sync": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancer_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scheduler": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sticky_session": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sticky_session_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "unhealthy_threshold": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "url": {
        "description_kind": "plain",
        "optional": true,
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

func AlicloudSlbRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlbRule), &result)
	return &result
}
