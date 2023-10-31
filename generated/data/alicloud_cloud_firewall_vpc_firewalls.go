package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallVpcFirewalls = `{
  "block": {
    "attributes": {
      "firewalls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bandwidth": "number",
              "connect_type": "string",
              "id": "string",
              "local_vpc": [
                "list",
                [
                  "object",
                  {
                    "eni_id": "string",
                    "eni_private_ip_address": "string",
                    "local_vpc_cidr_table_list": [
                      "list",
                      [
                        "object",
                        {
                          "local_route_entry_list": [
                            "list",
                            [
                              "object",
                              {
                                "local_destination_cidr": "string",
                                "local_next_hop_instance_id": "string"
                              }
                            ]
                          ],
                          "local_route_table_id": "string"
                        }
                      ]
                    ],
                    "region_no": "string",
                    "router_interface_id": "string",
                    "vpc_id": "string",
                    "vpc_name": "string"
                  }
                ]
              ],
              "peer_vpc": [
                "list",
                [
                  "object",
                  {
                    "eni_id": "string",
                    "eni_private_ip_address": "string",
                    "peer_vpc_cidr_table_list": [
                      "list",
                      [
                        "object",
                        {
                          "peer_route_entry_list": [
                            "list",
                            [
                              "object",
                              {
                                "peer_destination_cidr": "string",
                                "peer_next_hop_instance_id": "string"
                              }
                            ]
                          ],
                          "peer_route_table_id": "string"
                        }
                      ]
                    ],
                    "region_no": "string",
                    "router_interface_id": "string",
                    "vpc_id": "string",
                    "vpc_name": "string"
                  }
                ]
              ],
              "region_status": "string",
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
      "region_no": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      },
      "vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudFirewallVpcFirewallsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallVpcFirewalls), &result)
	return &result
}
