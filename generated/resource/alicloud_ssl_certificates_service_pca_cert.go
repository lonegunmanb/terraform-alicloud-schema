package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSslCertificatesServicePcaCert = `{
  "block": {
    "attributes": {
      "after_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "algorithm": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "alias_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "before_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "common_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "country_code": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "days": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "enable_crl": {
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
      "immediately": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "locality": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "months": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "organization": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "organization_unit": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parent_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "san_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "san_value": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "upload_flag": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "years": {
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

func AlicloudSslCertificatesServicePcaCertSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSslCertificatesServicePcaCert), &result)
	return &result
}
