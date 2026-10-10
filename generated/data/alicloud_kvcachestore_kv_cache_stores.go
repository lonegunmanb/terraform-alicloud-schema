package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudKvcachestoreKvCacheStores = `{
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
      "kvcs_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      },
      "stores": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "capacity": "number",
              "create_time": "string",
              "description": "string",
              "hpn_zone": "string",
              "id": "string",
              "kvcs_id": "string",
              "name": "string",
              "payment_type": "string",
              "region_id": "string",
              "status": "string",
              "zone_id": "string"
            }
          ]
        ]
      },
      "zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudKvcachestoreKvCacheStoresSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudKvcachestoreKvCacheStores), &result)
	return &result
}
