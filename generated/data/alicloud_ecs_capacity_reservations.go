package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcsCapacityReservations = `{
  "block": {
    "attributes": {
      "capacity_reservation_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
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
      "instance_type": {
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
      "platform": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reservations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "capacity_reservation_id": "string",
              "capacity_reservation_name": "string",
              "description": "string",
              "end_time": "string",
              "end_time_type": "string",
              "id": "string",
              "instance_amount": "string",
              "instance_type": "string",
              "match_criteria": "string",
              "payment_type": "string",
              "platform": "string",
              "resource_group_id": "string",
              "start_time": "string",
              "start_time_type": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "time_slot": "string",
              "zone_ids": [
                "list",
                "string"
              ]
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

func AlicloudEcsCapacityReservationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcsCapacityReservations), &result)
	return &result
}
