package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGaCustomRoutingEndpointGroupDestination = `{
  "block": {
    "attributes": {
      "accelerator_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_routing_endpoint_group_destination_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "from_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "listener_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "protocols": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "to_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
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

func AlicloudGaCustomRoutingEndpointGroupDestinationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGaCustomRoutingEndpointGroupDestination), &result)
	return &result
}
