package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEipAddresses = `{
  "block": {
    "attributes": {
      "address_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address_name": "string",
              "allocation_id": "string",
              "available_regions": [
                "list",
                "string"
              ],
              "bandwidth": "string",
              "bandwidth_package_bandwidth": "string",
              "bandwidth_package_id": "string",
              "bandwidth_package_type": "string",
              "create_time": "string",
              "deletion_protection": "bool",
              "description": "string",
              "expired_time": "string",
              "has_reservation_data": "string",
              "hd_monitor_status": "string",
              "id": "string",
              "instance_id": "string",
              "instance_region_id": "string",
              "instance_type": "string",
              "internet_charge_type": "string",
              "ip_address": "string",
              "isp": "string",
              "operation_locks": [
                "list",
                "string"
              ],
              "payment_type": "string",
              "reservation_active_time": "string",
              "reservation_bandwidth": "string",
              "reservation_internet_charge_type": "string",
              "reservation_order_type": "string",
              "resource_group_id": "string",
              "second_limited": "bool",
              "segment_instance_id": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ]
            }
          ]
        ]
      },
      "associated_instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "associated_instance_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dry_run": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "eips": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bandwidth": "string",
              "creation_time": "string",
              "deletion_protection": "bool",
              "id": "string",
              "instance_id": "string",
              "instance_type": "string",
              "internet_charge_type": "string",
              "ip_address": "string",
              "status": "string"
            }
          ]
        ]
      },
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
      "in_use": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "include_reservation_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ip_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_addresses": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "isp": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lock_reason": {
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
      "payment_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "segment_instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEipAddressesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEipAddresses), &result)
	return &result
}
