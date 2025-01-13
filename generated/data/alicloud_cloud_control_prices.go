package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudControlPrices = `{
  "block": {
    "attributes": {
      "desire_attributes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "prices": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "currency": "string",
              "discount_price": "number",
              "module_details": [
                "list",
                [
                  "object",
                  {
                    "cost_after_discount": "number",
                    "invoice_discount": "number",
                    "module_code": "string",
                    "module_name": "string",
                    "original_cost": "number",
                    "price_type": "string"
                  }
                ]
              ],
              "original_price": "number",
              "promotion_details": [
                "list",
                [
                  "object",
                  {
                    "promotion_desc": "string",
                    "promotion_id": "number",
                    "promotion_name": "string"
                  }
                ]
              ],
              "trade_price": "number"
            }
          ]
        ]
      },
      "product": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudControlPricesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudControlPrices), &result)
	return &result
}
