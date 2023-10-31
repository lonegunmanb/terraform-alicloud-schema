package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEssScalingConfigurations = `{
  "block": {
    "attributes": {
      "configurations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "creation_time": "string",
              "credit_specification": "string",
              "data_disks": [
                "list",
                [
                  "object",
                  {
                    "category": "string",
                    "delete_with_instance": "bool",
                    "device": "string",
                    "performance_level": "string",
                    "size": "number",
                    "snapshot_id": "string"
                  }
                ]
              ],
              "host_name": "string",
              "id": "string",
              "image_id": "string",
              "instance_name": "string",
              "instance_type": "string",
              "internet_charge_type": "string",
              "internet_max_bandwidth_in": "number",
              "internet_max_bandwidth_out": "number",
              "lifecycle_state": "string",
              "name": "string",
              "scaling_group_id": "string",
              "security_group_id": "string",
              "spot_price_limit": [
                "list",
                [
                  "object",
                  {
                    "instance_type": "string",
                    "price_limit": "number"
                  }
                ]
              ],
              "spot_strategy": "string",
              "system_disk_category": "string",
              "system_disk_performance_level": "string",
              "system_disk_size": "number"
            }
          ]
        ]
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
      "scaling_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEssScalingConfigurationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEssScalingConfigurations), &result)
	return &result
}
