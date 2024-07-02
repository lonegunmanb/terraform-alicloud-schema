package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGaBandwidthPackage = `{
  "block": {
    "attributes": {
      "auto_pay": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_renew_duration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "auto_use_coupon": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "bandwidth": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "bandwidth_package_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bandwidth_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "billing_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cbn_geographic_region_ida": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cbn_geographic_region_idb": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "duration": {
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
      "payment_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "promotion_option_no": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ratio": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "renewal_status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
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

func AlicloudGaBandwidthPackageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGaBandwidthPackage), &result)
	return &result
}
