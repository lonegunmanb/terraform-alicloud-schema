package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEipanycastAnycastEipAddresses = `{
  "block": {
    "attributes": {
      "addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ali_uid": "number",
              "anycast_eip_address_name": "string",
              "anycast_eip_bind_info_list": [
                "list",
                [
                  "object",
                  {
                    "bind_instance_id": "string",
                    "bind_instance_region_id": "string",
                    "bind_instance_type": "string",
                    "bind_time": "string"
                  }
                ]
              ],
              "anycast_id": "string",
              "bandwidth": "number",
              "bid": "string",
              "business_status": "string",
              "description": "string",
              "id": "string",
              "internet_charge_type": "string",
              "ip_address": "string",
              "payment_type": "string",
              "service_location": "string",
              "status": "string"
            }
          ]
        ]
      },
      "anycast_eip_address_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bind_instance_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "business_status": {
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
      "internet_charge_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_address": {
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
      "service_location": {
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

func AlicloudEipanycastAnycastEipAddressesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEipanycastAnycastEipAddresses), &result)
	return &result
}
