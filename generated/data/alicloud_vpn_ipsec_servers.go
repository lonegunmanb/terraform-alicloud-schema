package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpnIpsecServers = `{
  "block": {
    "attributes": {
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
      "ipsec_server_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "servers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "client_ip_pool": "string",
              "create_time": "string",
              "effect_immediately": "bool",
              "id": "string",
              "idaas_instance_id": "string",
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
              "ipsec_server_id": "string",
              "ipsec_server_name": "string",
              "local_subnet": "string",
              "max_connections": "number",
              "multi_factor_auth_enabled": "bool",
              "online_client_count": "number",
              "psk": "string",
              "psk_enabled": "bool",
              "vpn_gateway_id": "string"
            }
          ]
        ]
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

func AlicloudVpnIpsecServersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpnIpsecServers), &result)
	return &result
}
