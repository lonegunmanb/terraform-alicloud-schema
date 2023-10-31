package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudBssOpenApiPricingModules = `{
  "block": {
    "attributes": {
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
      "modules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "code": "string",
              "id": "string",
              "pricing_module_name": "string",
              "product_code": "string",
              "product_type": "string",
              "subscription_type": "string",
              "unit": "string",
              "values": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "remark": "string",
                    "type": "string",
                    "value": "string"
                  }
                ]
              ]
            }
          ]
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
      "product_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscription_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudBssOpenApiPricingModulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudBssOpenApiPricingModules), &result)
	return &result
}
