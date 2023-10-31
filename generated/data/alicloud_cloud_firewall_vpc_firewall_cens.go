package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallVpcFirewallCens = `{
  "block": {
    "attributes": {
      "cen_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cens": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cen_id": "string",
              "connect_type": "string",
              "id": "string",
              "local_vpc": [
                "list",
                [
                  "object",
                  {
                    "attachment_id": "string",
                    "attachment_name": "string",
                    "defend_cidr_list": [
                      "list",
                      "string"
                    ],
                    "eni_list": [
                      "list",
                      [
                        "object",
                        {
                          "eni_id": "string",
                          "eni_private_ip_address": "string"
                        }
                      ]
                    ],
                    "manual_vswitch_id": "string",
                    "network_instance_id": "string",
                    "network_instance_name": "string",
                    "network_instance_type": "string",
                    "owner_id": "string",
                    "region_no": "string",
                    "route_mode": "string",
                    "support_manual_mode": "string",
                    "transit_router_id": "string",
                    "transit_router_type": "string",
                    "vpc_cidr_table_list": [
                      "list",
                      [
                        "object",
                        {
                          "route_entry_list": [
                            "list",
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
                    ],
                    "vpc_id": "string",
                    "vpc_name": "string"
                  }
                ]
              ],
              "network_instance_id": "string",
              "status": "string",
              "vpc_firewall_id": "string",
              "vpc_firewall_name": "string"
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
      "network_instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "vpc_firewall_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_firewall_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudFirewallVpcFirewallCensSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallVpcFirewallCens), &result)
	return &result
}
