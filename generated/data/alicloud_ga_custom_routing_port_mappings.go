package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGaCustomRoutingPortMappings = `{
  "block": {
    "attributes": {
      "accelerator_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "custom_routing_port_mappings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "accelerator_id": "string",
              "accelerator_port": "number",
              "destination_socket_address": [
                "list",
                [
                  "object",
                  {
                    "ip_address": "string",
                    "port": "number"
                  }
                ]
              ],
              "endpoint_group_id": "string",
              "endpoint_group_region": "string",
              "endpoint_id": "string",
              "listener_id": "string",
              "protocols": [
                "list",
                "string"
              ],
              "status": "string",
              "vswitch": "string"
            }
          ]
        ]
      },
      "endpoint_group_id": {
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
      "listener_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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

func AlicloudGaCustomRoutingPortMappingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGaCustomRoutingPortMappings), &result)
	return &result
}
