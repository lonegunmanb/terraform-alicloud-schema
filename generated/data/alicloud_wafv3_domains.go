package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudWafv3Domains = `{
  "block": {
    "attributes": {
      "backend": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domains": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "domain": "string",
              "id": "string",
              "listen": [
                "list",
                [
                  "object",
                  {
                    "cert_id": "string",
                    "cipher_suite": "number",
                    "custom_ciphers": [
                      "list",
                      "string"
                    ],
                    "enable_tlsv3": "bool",
                    "exclusive_ip": "bool",
                    "focus_https": "bool",
                    "http2_enabled": "bool",
                    "http_ports": [
                      "list",
                      "number"
                    ],
                    "https_ports": [
                      "list",
                      "number"
                    ],
                    "ipv6_enabled": "bool",
                    "protection_resource": "string",
                    "tls_version": "string",
                    "xff_header_mode": "number",
                    "xff_headers": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "redirect": [
                "list",
                [
                  "object",
                  {
                    "backends": [
                      "list",
                      "string"
                    ],
                    "connect_timeout": "number",
                    "focus_http_backend": "bool",
                    "keepalive": "bool",
                    "keepalive_requests": "number",
                    "keepalive_timeout": "number",
                    "loadbalance": "string",
                    "read_timeout": "number",
                    "request_headers": [
                      "list",
                      [
                        "object",
                        {
                          "key": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "retry": "bool",
                    "sni_enabled": "bool",
                    "sni_host": "string",
                    "write_timeout": "number"
                  }
                ]
              ],
              "resource_manager_resource_group_id": "string",
              "status": "string"
            }
          ]
        ]
      },
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudWafv3DomainsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudWafv3Domains), &result)
	return &result
}
