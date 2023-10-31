package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlbServerCertificates = `{
  "block": {
    "attributes": {
      "certificates": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "alicloud_certificate_id": "string",
              "alicloud_certificate_name": "string",
              "common_name": "string",
              "created_time": "string",
              "created_timestamp": "number",
              "expired_time": "string",
              "expired_timestamp": "number",
              "fingerprint": "string",
              "id": "string",
              "is_alicloud_certificate": "bool",
              "name": "string",
              "resource_group_id": "string",
              "subject_alternative_names": [
                "list",
                "string"
              ],
              "tags": [
                "map",
                "string"
              ]
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
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSlbServerCertificatesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlbServerCertificates), &result)
	return &result
}
