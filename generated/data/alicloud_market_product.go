package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMarketProduct = `{
  "block": {
    "attributes": {
      "available_region": {
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
      "product": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "code": "string",
              "description": "string",
              "name": "string",
              "skus": [
                "list",
                [
                  "object",
                  {
                    "images": [
                      "list",
                      [
                        "object",
                        {
                          "image_id": "string",
                          "image_name": "string",
                          "region_id": "string"
                        }
                      ]
                    ],
                    "package_versions": [
                      "list",
                      [
                        "object",
                        {
                          "package_name": "string",
                          "package_version": "string"
                        }
                      ]
                    ],
                    "sku_code": "string",
                    "sku_name": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "product_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMarketProductSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMarketProduct), &result)
	return &result
}
