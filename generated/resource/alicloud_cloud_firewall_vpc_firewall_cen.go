package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallVpcFirewallCen = `{
  "block": {
    "attributes": {
      "cen_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      },
      "vpc_region": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "local_vpc": {
        "block": {
          "attributes": {
            "attachment_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "attachment_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "defend_cidr_list": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "eni_list": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "eni_id": "string",
                    "eni_private_ip_address": "string"
                  }
                ]
              ]
            },
            "manual_vswitch_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "network_instance_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "network_instance_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "network_instance_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "owner_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "region_no": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "route_mode": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "support_manual_mode": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "transit_router_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "transit_router_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "vpc_cidr_table_list": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "route_entry_list": [
                      "set",
                      [
                        "object",
                        {
                          "destination_cidr": "string",
                          "next_hop_instance_id": "string"
                        }
                      ]
                    ],
                    "route_table_id": "string"
                  }
                ]
              ]
            },
            "vpc_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "vpc_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
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

func AlicloudCloudFirewallVpcFirewallCenSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallVpcFirewallCen), &result)
	return &result
}
