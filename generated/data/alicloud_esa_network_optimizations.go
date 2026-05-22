package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaNetworkOptimizations = `{
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
      "optimizations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "config_id": "string",
              "config_type": "string",
              "grpc": "string",
              "http2_origin": "string",
              "id": "string",
              "rule": "string",
              "rule_enable": "string",
              "rule_name": "string",
              "sequence": "number",
              "site_version": "number",
              "smart_routing": "string",
              "upload_max_filesize": "string",
              "websocket": "string"
            }
          ]
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
      },
      "site_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEsaNetworkOptimizationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaNetworkOptimizations), &result)
	return &result
}
