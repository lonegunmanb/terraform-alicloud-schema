package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCommonBandwidthPackages = `{
  "block": {
    "attributes": {
      "bandwidth_package_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dry_run": {
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
      "include_reservation_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
              "bandwidth": "string",
              "bandwidth_package_id": "string",
              "bandwidth_package_name": "string",
              "business_status": "string",
              "deletion_protection": "bool",
              "description": "string",
              "expired_time": "string",
              "has_reservation_data": "bool",
              "id": "string",
              "internet_charge_type": "string",
              "isp": "string",
              "name": "string",
              "payment_type": "string",
              "public_ip_addresses": [
                "list",
                [
                  "object",
                  {
                    "allocation_id": "string",
                    "bandwidth_package_ip_relation_status": "string",
                    "ip_address": "string"
                  }
                ]
              ],
              "ratio": "number",
              "reservation_active_time": "string",
              "reservation_bandwidth": "string",
              "reservation_internet_charge_type": "string",
              "reservation_order_type": "string",
              "resource_group_id": "string",
              "service_managed": "number",
              "status": "string"
            }
          ]
        ]
      },
      "resource_group_id": {
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

func AlicloudCommonBandwidthPackagesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCommonBandwidthPackages), &result)
	return &result
}
