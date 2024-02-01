package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpnConnection = `{
  "block": {
    "attributes": {
      "auto_config_route": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "customer_gateway_id": {
        "description_kind": "plain",
        "optional": true,
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
      "enable_tunnels_bgp": {
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
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_type": {
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
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "vpn_connection_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
            "status": {
              "computed": true,
              "description_kind": "plain",
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
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ike_lifetime": {
              "computed": true,
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
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ike_pfs": {
              "computed": true,
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
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "psk": {
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
      "ipsec_config": {
        "block": {
          "attributes": {
            "ipsec_auth_alg": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ipsec_enc_alg": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ipsec_lifetime": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ipsec_pfs": {
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
      },
      "tunnel_options_specification": {
        "block": {
          "attributes": {
            "customer_gateway_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enable_dpd": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_nat_traversal": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "internet_ip": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "role": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tunnel_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "zone_no": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "tunnel_bgp_config": {
              "block": {
                "attributes": {
                  "bgp_status": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "local_asn": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "local_bgp_ip": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "peer_asn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "peer_bgp_ip": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tunnel_cidr": {
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
            "tunnel_ike_config": {
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
                  "psk": {
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "tunnel_ipsec_config": {
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
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
