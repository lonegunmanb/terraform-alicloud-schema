package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpnGatewayVpnAttachments = `{
  "block": {
    "attributes": {
      "attachments": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "attach_type": "string",
              "bgp_config": [
                "list",
                [
                  "object",
                  {
                    "local_asn": "number",
                    "local_bgp_ip": "string",
                    "status": "string",
                    "tunnel_cidr": "string"
                  }
                ]
              ],
              "connection_status": "string",
              "create_time": "string",
              "customer_gateway_id": "string",
              "effect_immediately": "bool",
              "enable_dpd": "bool",
              "enable_nat_traversal": "bool",
              "enable_tunnels_bgp": "bool",
              "health_check_config": [
                "list",
                [
                  "object",
                  {
                    "dip": "string",
                    "enable": "bool",
                    "interval": "number",
                    "policy": "string",
                    "retry": "number",
                    "sip": "string",
                    "status": "string"
                  }
                ]
              ],
              "id": "string",
              "ike_config": [
                "list",
                [
                  "object",
                  {
                    "ike_auth_alg": "string",
                    "ike_enc_alg": "string",
                    "ike_lifetime": "number",
                    "ike_mode": "string",
                    "ike_pfs": "string",
                    "ike_version": "string",
                    "local_id": "string",
                    "psk": "string",
                    "remote_id": "string"
                  }
                ]
              ],
              "internet_ip": "string",
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
              "network_type": "string",
              "remote_subnet": "string",
              "resource_group_id": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "tunnel_options_specification": [
                "set",
                [
                  "object",
                  {
                    "customer_gateway_id": "string",
                    "enable_dpd": "bool",
                    "enable_nat_traversal": "bool",
                    "internet_ip": "string",
                    "role": "string",
                    "state": "string",
                    "status": "string",
                    "tunnel_bgp_config": [
                      "list",
                      [
                        "object",
                        {
                          "bgp_status": "string",
                          "local_asn": "number",
                          "local_bgp_ip": "string",
                          "peer_asn": "string",
                          "peer_bgp_ip": "string",
                          "tunnel_cidr": "string"
                        }
                      ]
                    ],
                    "tunnel_id": "string",
                    "tunnel_ike_config": [
                      "set",
                      [
                        "object",
                        {
                          "ike_auth_alg": "string",
                          "ike_enc_alg": "string",
                          "ike_lifetime": "number",
                          "ike_mode": "string",
                          "ike_pfs": "string",
                          "ike_version": "string",
                          "local_id": "string",
                          "psk": "string",
                          "remote_id": "string"
                        }
                      ]
                    ],
                    "tunnel_index": "number",
                    "tunnel_ipsec_config": [
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
                    "zone_no": "string"
                  }
                ]
              ],
              "vpn_attachment_name": "string",
              "vpn_connection_id": "string"
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

func AlicloudVpnGatewayVpnAttachmentsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpnGatewayVpnAttachments), &result)
	return &result
}
