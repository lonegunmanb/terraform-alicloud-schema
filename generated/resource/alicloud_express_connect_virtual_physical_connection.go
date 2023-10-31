package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectVirtualPhysicalConnection = `{
  "block": {
    "attributes": {
      "access_point_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ad_location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bandwidth": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "business_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "circuit_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dry_run": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enabled_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "end_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expect_spec": {
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
      "line_operator": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "loa_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "order_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent_physical_connection_ali_uid": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "parent_physical_connection_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "port_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "port_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "redundant_physical_connection_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spec": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "virtual_physical_connection_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "virtual_physical_connection_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vlan_id": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "vpconn_ali_uid": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudExpressConnectVirtualPhysicalConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectVirtualPhysicalConnection), &result)
	return &result
}
