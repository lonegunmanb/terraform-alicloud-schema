package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMarketOrder = `{
  "block": {
    "attributes": {
      "components": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "coupon_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "package_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pay_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pricing_cycle": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "quantity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMarketOrderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMarketOrder), &result)
	return &result
}
