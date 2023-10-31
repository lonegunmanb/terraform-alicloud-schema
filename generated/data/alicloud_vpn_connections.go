package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpnConnections = `{
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
              "create_time": "string",
              "customer_gateway_id": "string",
              "effect_immediately": "bool",
              "enable_dpd": "bool",
              "enable_nat_traversal": "bool",
              "id": "string",
              "ike_config": [
                "list",
                [
                  "object",
                  {
                    "ike_auth_alg": "string",
                    "ike_enc_alg": "string",
                    "ike_lifetime": "number",
                    "ike_local_id": "string",
                    "ike_mode": "string",
                    "ike_pfs": "string",
                    "ike_remote_id": "string",
                    "ike_version": "string",
                    "psk": "string"
                  }
                ]
              ],
              "ipsec_config": [
                "list",
                [
                  "object",
                  {
                    "ipsec_auth_alg": "string",
                    "ipsec_enc_alg": "string",
                    "ipsec_lifetime": "number",
                    "ipsec_pfs": "string"
                  }
                ]
              ],
              "local_subnet": "string",
              "name": "string",
              "remote_subnet": "string",
              "status": "string",
              "vco_health_check": [
                "list",
                [
                  "object",
                  {
                    "dip": "string",
                    "enable": "string",
                    "interval": "number",
                    "retry": "number",
                    "sip": "string",
                    "status": "string"
                  }
                ]
              ],
              "vpn_bgp_config": [
                "list",
                [
                  "object",
                  {
                    "auth_key": "string",
                    "local_asn": "number",
                    "local_bgp_ip": "string",
                    "peer_asn": "number",
                    "peer_bgp_ip": "string",
                    "status": "string",
                    "tunnel_cidr": "string"
                  }
                ]
              ],
              "vpn_gateway_id": "string"
            }
          ]
        ]
      },
      "customer_gateway_id": {
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
      "vpn_gateway_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudVpnConnectionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpnConnections), &result)
	return &result
}
