package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenBandwidthPackages = `{
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
      "include_reservation_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "packages": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bandwidth": "number",
              "bandwidth_package_charge_type": "string",
              "business_status": "string",
              "cen_bandwidth_package_id": "string",
              "cen_bandwidth_package_name": "string",
              "cen_ids": [
                "list",
                "string"
              ],
              "description": "string",
              "expired_time": "string",
              "geographic_region_a_id": "string",
              "geographic_region_b_id": "string",
              "geographic_span_id": "string",
              "has_reservation_data": "string",
              "id": "string",
              "instance_id": "string",
              "is_cross_border": "bool",
              "name": "string",
              "payment_type": "string",
              "reservation_active_time": "string",
              "reservation_bandwidth": "string",
              "reservation_internet_charge_type": "string",
              "reservation_order_type": "string",
              "status": "string"
            }
          ]
        ]
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

func AlicloudCenBandwidthPackagesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenBandwidthPackages), &result)
	return &result
}
