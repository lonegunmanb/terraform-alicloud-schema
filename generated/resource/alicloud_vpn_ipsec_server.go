package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpnIpsecServer = `{
  "block": {
    "attributes": {
      "client_ip_pool": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dry_run": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "effect_immediately": {
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
      "ipsec_server_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "local_subnet": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "psk": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "psk_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "vpn_gateway_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "ike_config": {
        "block": {
          "attributes": {
            "ike_auth_alg": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ike_enc_alg": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ike_lifetime": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ike_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ike_pfs": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ike_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "local_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "remote_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "ipsec_config": {
        "block": {
          "attributes": {
            "ipsec_auth_alg": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ipsec_enc_alg": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ipsec_lifetime": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ipsec_pfs": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
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

func AlicloudVpnIpsecServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpnIpsecServer), &result)
	return &result
}
