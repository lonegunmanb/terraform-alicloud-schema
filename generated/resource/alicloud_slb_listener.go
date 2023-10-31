package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlbListener = `{
  "block": {
    "attributes": {
      "acl_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "acl_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "acl_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backend_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "bandwidth": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ca_certificate_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
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
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_http2": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "established_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "forward_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "frontend_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "gzip": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
        "computed": true,
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
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "instance_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "lb_port": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "lb_protocol": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "listener_forward": {
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
      "master_slave_server_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "persistence_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "proxy_protocol_v2_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "server_certificate_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_certificate_id": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
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
      "tls_cipher_policy": {
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
      "x_forwarded_for": {
        "block": {
          "attributes": {
            "retrive_client_ip": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "retrive_slb_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "retrive_slb_ip": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "retrive_slb_proto": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSlbListenerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlbListener), &result)
	return &result
}
