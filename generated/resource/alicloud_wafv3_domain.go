package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudWafv3Domain = `{
  "block": {
    "attributes": {
      "access_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain": {
        "description_kind": "plain",
        "required": true,
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
      "resource_manager_resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "listen": {
        "block": {
          "attributes": {
            "cert_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cipher_suite": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "custom_ciphers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "enable_tlsv3": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "exclusive_ip": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "focus_https": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "http2_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "http_ports": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            },
            "https_ports": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            },
            "ipv6_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "protection_resource": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tls_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "xff_header_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "xff_headers": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "redirect": {
        "block": {
          "attributes": {
            "backends": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "connect_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "focus_http_backend": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "keepalive": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "keepalive_requests": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "keepalive_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "loadbalance": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "read_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "retry": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "sni_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "sni_host": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "write_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "request_headers": {
              "block": {
                "attributes": {
                  "key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func AlicloudWafv3DomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudWafv3Domain), &result)
	return &result
}
