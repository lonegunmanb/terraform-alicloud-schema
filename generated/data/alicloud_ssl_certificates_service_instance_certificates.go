package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSslCertificatesServiceInstanceCertificates = `{
  "block": {
    "attributes": {
      "certificate_source": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificates": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "algorithm": "string",
              "cert_identifier": "string",
              "certificate_id": "number",
              "certificate_name": "string",
              "certificate_source": "string",
              "certificate_status": "string",
              "common_name": "string",
              "domain": "string",
              "exist_private_key": "bool",
              "finger_print": "string",
              "id": "number",
              "instance_id": "string",
              "issuer": "string",
              "key_size": "number",
              "not_after": "number",
              "not_before": "number"
            }
          ]
        ]
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
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSslCertificatesServiceInstanceCertificatesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSslCertificatesServiceInstanceCertificates), &result)
	return &result
}
