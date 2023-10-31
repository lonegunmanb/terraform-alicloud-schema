package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectPhysicalConnection = `{
  "block": {
    "attributes": {
      "access_point_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "bandwidth": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "circuit_code": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_location": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "physical_connection_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "redundant_physical_connection_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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

func AlicloudExpressConnectPhysicalConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectPhysicalConnection), &result)
	return &result
}
