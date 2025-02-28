package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaHttpsBasicConfiguration = `{
  "block": {
    "attributes": {
      "ciphersuite": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ciphersuite_group": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "config_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "http2": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "http3": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "https": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ocsp_stapling": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "site_id": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "tls10": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tls11": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tls12": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tls13": {
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

func AlicloudEsaHttpsBasicConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaHttpsBasicConfiguration), &result)
	return &result
}
