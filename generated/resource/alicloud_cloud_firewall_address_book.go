package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallAddressBook = `{
  "block": {
    "attributes": {
      "address_list": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "address_list_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "asset_member_uids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "number"
        ]
      },
      "auto_add_tag_ecs": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_type": {
        "description_kind": "plain",
        "required": true,
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
      "reference_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tag_relation": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "asset_region_resource_types": {
        "block": {
          "attributes": {
            "asset_region_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "resource_type": {
              "block": {
                "block_types": {
                  "ipv4": {
                    "block": {
                      "attributes": {
                        "ai_gateway_eip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "alb_eip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "api_gateway_eip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "bastion_host_egress_ip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "bastion_host_ingress_ip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "bastion_host_ip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "ecs_eip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "ecs_public_ip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "eip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "eni_eip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "ga_eip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "havip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "nat_eip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "nat_public_ip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "nlb_eip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "slb_eip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "slb_public_ip": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "ipv6": {
                    "block": {
                      "attributes": {
                        "ai_gateway_eipv6": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "alb_ipv6": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "api_gateway_eipv6": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "ecs_ipv6": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "eni_eipv6": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "ga_eipv6": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "nlb_ipv6": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "slb_ipv6": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "ecs_tags": {
        "block": {
          "attributes": {
            "tag_key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tag_value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudFirewallAddressBookSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallAddressBook), &result)
	return &result
}
