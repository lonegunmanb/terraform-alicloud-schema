package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSslCertificatesServicePcaCertificate = `{
  "block": {
    "attributes": {
      "algorithm": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "common_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "country_code": {
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
      "locality": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "organization": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "organization_unit": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "years": {
        "description_kind": "plain",
        "required": true,
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

func AlicloudSslCertificatesServicePcaCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSslCertificatesServicePcaCertificate), &result)
	return &result
}
