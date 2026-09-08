package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionAttackPathWhitelists = `{
  "block": {
    "attributes": {
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "path_name_desc": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "path_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "whitelist_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "whitelists": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "attack_path_asset_list": [
                "set",
                [
                  "object",
                  {
                    "asset_sub_type": "number",
                    "asset_type": "number",
                    "instance_id": "string",
                    "node_type": "string",
                    "region_id": "string",
                    "vendor": "number"
                  }
                ]
              ],
              "attack_path_whitelist_id": "string",
              "id": "string",
              "path_name": "string",
              "path_type": "string",
              "remark": "string",
              "whitelist_name": "string",
              "whitelist_type": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudThreatDetectionAttackPathWhitelistsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionAttackPathWhitelists), &result)
	return &result
}
