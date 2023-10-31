package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDdoscooDomainResources = `{
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
      "instance_ids": {
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
      },
      "query_domain_pattern": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resources": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "black_list": [
                "list",
                "string"
              ],
              "cc_enabled": "bool",
              "cc_rule_enabled": "bool",
              "cc_template": "string",
              "cert_name": "string",
              "domain": "string",
              "http2_enable": "bool",
              "https_ext": "string",
              "id": "string",
              "instance_ids": [
                "list",
                "string"
              ],
              "policy_mode": "string",
              "proxy_enabled": "bool",
              "proxy_types": [
                "list",
                [
                  "object",
                  {
                    "proxy_ports": [
                      "list",
                      "number"
                    ],
                    "proxy_type": "string"
                  }
                ]
              ],
              "real_servers": [
                "list",
                "string"
              ],
              "rs_type": "number",
              "ssl_ciphers": "string",
              "ssl_protocols": "string",
              "white_list": [
                "list",
                "string"
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDdoscooDomainResourcesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDdoscooDomainResources), &result)
	return &result
}
