package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpnConnection = `{
  "block": {
    "attributes": {
      "customer_gateway_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "effect_immediately": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_dpd": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_nat_traversal": {
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
      "local_subnet": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remote_subnet": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpn_gateway_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "bgp_config": {
        "block": {
          "attributes": {
            "enable": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "local_asn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "local_bgp_ip": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tunnel_cidr": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "health_check_config": {
        "block": {
          "attributes": {
            "dip": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enable": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "interval": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "retry": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "sip": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "ike_config": {
        "block": {
          "attributes": {
            "ike_auth_alg": {
              "computed": true,
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
            "ike_local_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
            "ike_remote_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ike_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "psk": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
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
        "max_items": 1,
        "nesting_mode": "list"
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

func AlicloudVpnConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpnConnection), &result)
	return &result
}
