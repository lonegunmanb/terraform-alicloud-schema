package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionAttackPathWhitelist = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "path_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "path_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "remark": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "whitelist_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "whitelist_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "attack_path_asset_list": {
        "block": {
          "attributes": {
            "asset_sub_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "asset_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "instance_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "node_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vendor": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
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

func AlicloudThreatDetectionAttackPathWhitelistSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionAttackPathWhitelist), &result)
	return &result
}
