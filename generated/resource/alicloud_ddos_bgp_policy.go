package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDdosBgpPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "content": {
        "block": {
          "attributes": {
            "black_ip_list_expire_at": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "enable_defense": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_drop_icmp": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_intelligence": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "intelligence_level": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "reflect_block_udp_port_list": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            },
            "region_block_country_list": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            },
            "region_block_province_list": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            },
            "whiten_gfbr_nets": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "finger_print_rule_list": {
              "block": {
                "attributes": {
                  "dst_port_end": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "dst_port_start": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "finger_print_rule_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "match_action": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "max_pkt_len": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "min_pkt_len": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "offset": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "payload_bytes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "protocol": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "rate_value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "seq_no": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "src_port_end": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "src_port_start": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "layer4_rule_list": {
              "block": {
                "attributes": {
                  "action": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "limited": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "match": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "method": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "priority": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "condition_list": {
                    "block": {
                      "attributes": {
                        "arg": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "depth": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "position": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "port_rule_list": {
              "block": {
                "attributes": {
                  "dst_port_end": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "dst_port_start": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "match_action": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port_rule_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "protocol": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "seq_no": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "src_port_end": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "src_port_start": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "source_block_list": {
              "block": {
                "attributes": {
                  "block_expire_seconds": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "every_seconds": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "exceed_limit_times": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "source_limit": {
              "block": {
                "attributes": {
                  "bps": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "pps": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "syn_bps": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "syn_pps": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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

func AlicloudDdosBgpPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDdosBgpPolicy), &result)
	return &result
}
