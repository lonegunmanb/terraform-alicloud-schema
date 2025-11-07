package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaNetworkOptimization = `{
  "block": {
    "attributes": {
      "config_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "grpc": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "http2_origin": {
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
      "sequence": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      },
      "smart_routing": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "upload_max_filesize": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "websocket": {
        "description_kind": "plain",
        "optional": true,
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

func AlicloudEsaNetworkOptimizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaNetworkOptimization), &result)
	return &result
}
