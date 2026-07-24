package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudApigDomains = `{
  "block": {
    "attributes": {
      "domains": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ca_cert_identifier": "string",
              "cert_identifier": "string",
              "client_ca_cert": "string",
              "domain_id": "string",
              "domain_name": "string",
              "domain_scope": "string",
              "force_https": "bool",
              "http2_option": "string",
              "id": "string",
              "m_tls_enabled": "bool",
              "protocol": "string",
              "resource_group_id": "string",
              "tls_cipher_suites_config": [
                "list",
                [
                  "object",
                  {
                    "config_type": "string",
                    "tls_cipher_suite": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "support_versions": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "tls_max": "string",
              "tls_min": "string"
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
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
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
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudApigDomainsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudApigDomains), &result)
	return &result
}
