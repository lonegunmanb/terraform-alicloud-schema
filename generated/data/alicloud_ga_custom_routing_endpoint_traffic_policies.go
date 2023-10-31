package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGaCustomRoutingEndpointTrafficPolicies = `{
  "block": {
    "attributes": {
      "accelerator_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_routing_endpoint_traffic_policies": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "accelerator_id": "string",
              "address": "string",
              "custom_routing_endpoint_traffic_policy_id": "string",
              "endpoint_group_id": "string",
              "endpoint_id": "string",
              "id": "string",
              "listener_id": "string",
              "port_ranges": [
                "list",
                [
                  "object",
                  {
                    "from_port": "number",
                    "to_port": "number"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "endpoint_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_id": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudGaCustomRoutingEndpointTrafficPoliciesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGaCustomRoutingEndpointTrafficPolicies), &result)
	return &result
}
