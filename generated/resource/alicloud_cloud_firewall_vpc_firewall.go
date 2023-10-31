package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallVpcFirewall = `{
  "block": {
    "attributes": {
      "bandwidth": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "connect_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "member_uid": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_firewall_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_firewall_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "local_vpc": {
        "block": {
          "attributes": {
            "eni_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "eni_private_ip_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "region_no": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "router_interface_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "vpc_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vpc_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "local_vpc_cidr_table_list": {
              "block": {
                "attributes": {
                  "local_route_table_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "local_route_entry_list": {
                    "block": {
                      "attributes": {
                        "local_destination_cidr": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "local_next_hop_instance_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
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
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "peer_vpc": {
        "block": {
          "attributes": {
            "eni_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "eni_private_ip_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "region_no": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "router_interface_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "vpc_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vpc_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "peer_vpc_cidr_table_list": {
              "block": {
                "attributes": {
                  "peer_route_table_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "peer_route_entry_list": {
                    "block": {
                      "attributes": {
                        "peer_destination_cidr": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "peer_next_hop_instance_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
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
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func AlicloudCloudFirewallVpcFirewallSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallVpcFirewall), &result)
	return &result
}
