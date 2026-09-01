package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEnsLoadBalancerUdpListeners = `{
  "block": {
    "attributes": {
      "enable_details": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "listeners": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backend_server_port": "number",
              "description": "string",
              "eip_transmit": "string",
              "established_timeout": "number",
              "health_check_connect_port": "number",
              "health_check_connect_timeout": "number",
              "health_check_exp": "string",
              "health_check_interval": "number",
              "health_check_req": "string",
              "healthy_threshold": "number",
              "id": "string",
              "listener_port": "number",
              "load_balancer_id": "string",
              "protocol": "string",
              "scheduler": "string",
              "status": "string",
              "unhealthy_threshold": "number"
            }
          ]
        ]
      },
      "load_balancer_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEnsLoadBalancerUdpListenersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEnsLoadBalancerUdpListeners), &result)
	return &result
}
