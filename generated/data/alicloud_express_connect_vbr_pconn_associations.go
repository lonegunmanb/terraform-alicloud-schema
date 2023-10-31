package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectVbrPconnAssociations = `{
  "block": {
    "attributes": {
      "associations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "circuit_code": "string",
              "enable_ipv6": "bool",
              "id": "string",
              "local_gateway_ip": "string",
              "local_ipv6_gateway_ip": "string",
              "peer_gateway_ip": "string",
              "peer_ipv6_gateway_ip": "string",
              "peering_ipv6_subnet_mask": "string",
              "peering_subnet_mask": "string",
              "physical_connection_id": "string",
              "status": "string",
              "vbr_id": "string",
              "vlan_id": "number"
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
      "vbr_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudExpressConnectVbrPconnAssociationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectVbrPconnAssociations), &result)
	return &result
}
