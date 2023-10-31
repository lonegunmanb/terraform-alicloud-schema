package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMarketProducts = `{
  "block": {
    "attributes": {
      "category_id": {
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "product_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "products": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "category_id": "number",
              "code": "string",
              "delivery_date": "string",
              "delivery_way": "string",
              "image_url": "string",
              "name": "string",
              "operation_system": "string",
              "score": "string",
              "short_description": "string",
              "suggested_price": "string",
              "supplier_id": "number",
              "supplier_name": "string",
              "tags": "string",
              "target_url": "string",
              "warranty_date": "string"
            }
          ]
        ]
      },
      "search_term": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sort": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "suggested_price": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "supplier_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "supplier_name_keyword": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMarketProductsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMarketProducts), &result)
	return &result
}
