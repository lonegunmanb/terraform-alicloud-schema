package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallAddressBooks = `{
  "block": {
    "attributes": {
      "books": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address_list": [
                "list",
                "string"
              ],
              "auto_add_tag_ecs": "number",
              "description": "string",
              "ecs_tags": [
                "set",
                [
                  "object",
                  {
                    "tag_key": "string",
                    "tag_value": "string"
                  }
                ]
              ],
              "group_name": "string",
              "group_type": "string",
              "group_uuid": "string",
              "id": "string",
              "tag_relation": "string"
            }
          ]
        ]
      },
      "group_type": {
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

func AlicloudCloudFirewallAddressBooksSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallAddressBooks), &result)
	return &result
}
