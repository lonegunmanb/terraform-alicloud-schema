package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlbListeners = `{
  "block": {
    "attributes": {
      "description_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "frontend_port": {
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
      "load_balancer_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "slb_listeners": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backend_port": "number",
              "bandwidth": "number",
              "ca_certificate_id": "string",
              "cookie": "string",
              "cookie_timeout": "number",
              "description": "string",
              "enable_http2": "string",
              "established_timeout": "number",
              "frontend_port": "number",
              "gzip": "string",
              "health_check": "string",
              "health_check_connect_port": "number",
              "health_check_connect_timeout": "number",
              "health_check_domain": "string",
              "health_check_http_code": "string",
              "health_check_interval": "number",
              "health_check_timeout": "number",
              "health_check_type": "string",
              "health_check_uri": "string",
              "healthy_threshold": "number",
              "idle_timeout": "number",
              "master_slave_server_group_id": "string",
              "persistence_timeout": "number",
              "protocol": "string",
              "proxy_protocol_v2_enabled": "bool",
              "request_timeout": "number",
              "scheduler": "string",
              "security_status": "string",
              "server_certificate_id": "string",
              "server_group_id": "string",
              "ssl_certificate_id": "string",
              "status": "string",
              "sticky_session": "string",
              "sticky_session_type": "string",
              "tls_cipher_policy": "string",
              "unhealthy_threshold": "number",
              "x_forwarded_for": "string",
              "x_forwarded_for_slb_id": "string",
              "x_forwarded_for_slb_ip": "string",
              "x_forwarded_for_slb_proto": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSlbListenersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlbListeners), &result)
	return &result
}
