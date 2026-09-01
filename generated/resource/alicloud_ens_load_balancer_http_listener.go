package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEnsLoadBalancerHttpListener = `{
  "block": {
    "attributes": {
      "backend_server_port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "forward_port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "health_check_method": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "idle_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "listener_forward": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "listener_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "load_balancer_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "request_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "scheduler": {
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
      "unhealthy_threshold": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "x_forwarded_for": {
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

func AlicloudEnsLoadBalancerHttpListenerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEnsLoadBalancerHttpListener), &result)
	return &result
}
