package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudExpressConnectVbrPconnAssociation = `{
  "block": {
    "attributes": {
      "circuit_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_ipv6": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "local_gateway_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "local_ipv6_gateway_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "peer_gateway_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "peer_ipv6_gateway_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "peering_ipv6_subnet_mask": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "peering_subnet_mask": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "physical_connection_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vbr_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vlan_id": {
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

func AlicloudExpressConnectVbrPconnAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudExpressConnectVbrPconnAssociation), &result)
	return &result
}
