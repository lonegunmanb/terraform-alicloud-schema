package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudNlbListeners = `{
  "block": {
    "attributes": {
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
      "listener_protocol": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "listeners": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "alpn_enabled": "bool",
              "alpn_policy": "string",
              "ca_certificate_ids": [
                "list",
                "string"
              ],
              "ca_enabled": "bool",
              "certificate_ids": [
                "list",
                "string"
              ],
              "cps": "number",
              "end_port": "string",
              "id": "string",
              "idle_timeout": "number",
              "listener_description": "string",
              "listener_id": "string",
              "listener_port": "number",
              "listener_protocol": "string",
              "load_balancer_id": "string",
              "mss": "number",
              "proxy_protocol_enabled": "bool",
              "sec_sensor_enabled": "bool",
              "security_policy_id": "string",
              "server_group_id": "string",
              "start_port": "string",
              "status": "string"
            }
          ]
        ]
      },
      "load_balancer_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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

func AlicloudNlbListenersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudNlbListeners), &result)
	return &result
}
