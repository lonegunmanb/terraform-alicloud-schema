package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudZones = `{
  "block": {
    "attributes": {
      "available_disk_category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "available_instance_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "available_resource_creation": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "available_slb_address_ip_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "available_slb_address_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
        "type": [
          "list",
          "string"
        ]
      },
      "instance_charge_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "network_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spot_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "available_disk_categories": [
                "list",
                "string"
              ],
              "available_instance_types": [
                "list",
                "string"
              ],
              "available_resource_creation": [
                "list",
                "string"
              ],
              "id": "string",
              "local_name": "string",
              "multi_zone_ids": [
                "list",
                "string"
              ],
              "slb_slave_zone_ids": [
                "list",
                "string"
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudZonesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudZones), &result)
	return &result
}
