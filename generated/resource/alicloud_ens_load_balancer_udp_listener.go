package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEnsLoadBalancerUdpListener = `{
  "block": {
    "attributes": {
      "backend_server_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "eip_transmit": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "established_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_connect_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_connect_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_exp": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_req": {
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

func AlicloudEnsLoadBalancerUdpListenerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEnsLoadBalancerUdpListener), &result)
	return &result
}
