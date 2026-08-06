package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSslCertificatesServiceCertificateValidation = `{
  "block": {
    "attributes": {
      "cert_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "certificate_status": {
        "computed": true,
        "description_kind": "plain",
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
      "validation_record_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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

func AlicloudSslCertificatesServiceCertificateValidationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSslCertificatesServiceCertificateValidation), &result)
	return &result
}
