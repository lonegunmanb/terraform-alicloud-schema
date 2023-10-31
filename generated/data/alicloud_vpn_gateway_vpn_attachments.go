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
              "bgp_config": [
                "list",
                [
                  "object",
                  {
                    "local_asn": "string",
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
                    "ike_lifetime": "string",
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
                    "ipsec_lifetime": "string",
                    "ipsec_pfs": "string"
                  }
                ]
              ],
              "local_subnet": "string",
              "network_type": "string",
              "remote_subnet": "string",
              "status": "string",
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
      },
      "vpn_gateway_id": {
        "deprecated": true,
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
