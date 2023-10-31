package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectVirtualPhysicalConnections = `{
  "block": {
    "attributes": {
      "business_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
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
              "expect_spec": "string",
              "id": "string",
              "line_operator": "string",
              "loa_status": "string",
              "order_mode": "string",
              "parent_physical_connection_ali_uid": "string",
              "parent_physical_connection_id": "string",
              "peer_location": "string",
              "port_number": "string",
              "port_type": "string",
              "redundant_physical_connection_id": "string",
              "resource_group_id": "string",
              "spec": "string",
              "status": "string",
              "virtual_physical_connection_id": "string",
              "virtual_physical_connection_name": "string",
              "virtual_physical_connection_status": "string",
              "vlan_id": "number",
              "vpconn_ali_uid": "string"
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
      "is_confirmed": {
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
      "parent_physical_connection_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "virtual_physical_connection_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "virtual_physical_connection_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vlan_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "number"
        ]
      },
      "vpconn_ali_uid": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudExpressConnectVirtualPhysicalConnectionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectVirtualPhysicalConnections), &result)
	return &result
}
