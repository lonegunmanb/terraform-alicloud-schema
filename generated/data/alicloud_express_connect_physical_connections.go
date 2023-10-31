package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectPhysicalConnections = `{
  "block": {
    "attributes": {
      "connections": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_point_id": "string",
              "ad_location": "string",
              "bandwidth": "string",
              "business_status": "string",
              "circuit_code": "string",
              "create_time": "string",
              "description": "string",
              "enabled_time": "string",
              "end_time": "string",
              "has_reservation_data": "string",
              "id": "string",
              "line_operator": "string",
              "loa_status": "string",
              "payment_type": "string",
              "peer_location": "string",
              "physical_connection_id": "string",
              "physical_connection_name": "string",
              "port_number": "string",
              "port_type": "string",
              "redundant_physical_connection_id": "string",
              "reservation_active_time": "string",
              "reservation_internet_charge_type": "string",
              "reservation_order_type": "string",
              "spec": "string",
              "status": "string",
              "type": "string"
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

func AlicloudExpressConnectPhysicalConnectionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectPhysicalConnections), &result)
	return &result
}
