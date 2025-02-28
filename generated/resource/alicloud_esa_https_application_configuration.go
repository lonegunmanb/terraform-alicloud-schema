package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaHttpsApplicationConfiguration = `{
  "block": {
    "attributes": {
      "alt_svc": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "alt_svc_clear": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "alt_svc_ma": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "alt_svc_persist": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "config_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "hsts": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hsts_include_subdomains": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hsts_max_age": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hsts_preload": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "https_force": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "https_force_code": {
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
      "site_version": {
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

func AlicloudEsaHttpsApplicationConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaHttpsApplicationConfiguration), &result)
	return &result
}
