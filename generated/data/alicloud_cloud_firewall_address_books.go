package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallAddressBooks = `{
  "block": {
    "attributes": {
      "books": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address_list": [
                "list",
                "string"
              ],
              "address_list_count": "number",
              "asset_member_uids": [
                "list",
                "number"
              ],
              "asset_region_resource_types": [
                "list",
                [
                  "object",
                  {
                    "asset_region_id": "string",
                    "resource_type": [
                      "list",
                      [
                        "object",
                        {
                          "ipv4": [
                            "list",
                            [
                              "object",
                              {
                                "ai_gateway_eip": "bool",
                                "alb_eip": "bool",
                                "api_gateway_eip": "bool",
                                "bastion_host_egress_ip": "bool",
                                "bastion_host_ingress_ip": "bool",
                                "bastion_host_ip": "bool",
                                "ecs_eip": "bool",
                                "ecs_public_ip": "bool",
                                "eip": "bool",
                                "eni_eip": "bool",
                                "ga_eip": "bool",
                                "havip": "bool",
                                "nat_eip": "bool",
                                "nat_public_ip": "bool",
                                "nlb_eip": "bool",
                                "slb_eip": "bool",
                                "slb_public_ip": "bool"
                              }
                            ]
                          ],
                          "ipv6": [
                            "list",
                            [
                              "object",
                              {
                                "ai_gateway_eipv6": "bool",
                                "alb_ipv6": "bool",
                                "api_gateway_eipv6": "bool",
                                "ecs_ipv6": "bool",
                                "eni_eipv6": "bool",
                                "ga_eipv6": "bool",
                                "nlb_ipv6": "bool",
                                "slb_ipv6": "bool"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "auto_add_tag_ecs": "number",
              "description": "string",
              "ecs_tags": [
                "set",
                [
                  "object",
                  {
                    "tag_key": "string",
                    "tag_value": "string"
                  }
                ]
              ],
              "group_name": "string",
              "group_type": "string",
              "group_uuid": "string",
              "id": "string",
              "reference_count": "number",
              "tag_relation": "string"
            }
          ]
        ]
      },
      "group_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudFirewallAddressBooksSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallAddressBooks), &result)
	return &result
}
