package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenInterRegionTrafficQosQueues = `{
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
      "queues": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dscps": [
                "list",
                "string"
              ],
              "id": "string",
              "inter_region_traffic_qos_queue_description": "string",
              "inter_region_traffic_qos_queue_id": "string",
              "inter_region_traffic_qos_queue_name": "string",
              "remain_bandwidth_percent": "number",
              "status": "string",
              "traffic_qos_policy_id": "string"
            }
          ]
        ]
      },
      "traffic_qos_policy_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCenInterRegionTrafficQosQueuesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenInterRegionTrafficQosQueues), &result)
	return &result
}
