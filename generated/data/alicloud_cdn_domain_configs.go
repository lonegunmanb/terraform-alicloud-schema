package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCdnDomainConfigs = `{
  "block": {
    "attributes": {
      "config_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "configs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "config_id": "string",
              "function_args": [
                "list",
                [
                  "object",
                  {
                    "arg_name": "string",
                    "arg_value": "string"
                  }
                ]
              ],
              "function_name": "string",
              "id": "string",
              "parent_id": "string",
              "status": "string"
            }
          ]
        ]
      },
      "domain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "function_name": {
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
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCdnDomainConfigsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCdnDomainConfigs), &result)
	return &result
}
