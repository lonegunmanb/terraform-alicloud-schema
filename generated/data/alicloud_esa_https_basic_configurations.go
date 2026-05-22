package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaHttpsBasicConfigurations = `{
  "block": {
    "attributes": {
      "config_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "config_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "configurations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ciphersuite": "string",
              "ciphersuite_group": "string",
              "config_id": "string",
              "config_type": "string",
              "http2": "string",
              "http3": "string",
              "https": "string",
              "id": "string",
              "ocsp_stapling": "string",
              "rule": "string",
              "rule_enable": "string",
              "rule_name": "string",
              "sequence": "number",
              "tls10": "string",
              "tls11": "string",
              "tls12": "string",
              "tls13": "string"
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
      "rule_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "site_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEsaHttpsBasicConfigurationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaHttpsBasicConfigurations), &result)
	return &result
}
