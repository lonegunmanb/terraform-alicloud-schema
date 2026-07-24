package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSslCertificatesServiceCompany = `{
  "block": {
    "attributes": {
      "city": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "company_address": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "company_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "company_email": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "company_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "company_phone": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "company_type": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "country_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "department": {
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
      "lang": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "post_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "province": {
        "description_kind": "plain",
        "required": true,
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

func AlicloudSslCertificatesServiceCompanySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSslCertificatesServiceCompany), &result)
	return &result
}
