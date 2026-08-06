package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSslCertificatesServiceInstanceCertificate = `{
  "block": {
    "attributes": {
      "algorithm": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cert_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_id": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "certificate_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_source": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "common_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "exist_private_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "finger_print": {
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "issuer": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "not_after": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "not_before": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "serial": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subject_alternative_names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "using_product_list": {
        "computed": true,
        "description_kind": "plain",
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

func AlicloudSslCertificatesServiceInstanceCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSslCertificatesServiceInstanceCertificate), &result)
	return &result
}
