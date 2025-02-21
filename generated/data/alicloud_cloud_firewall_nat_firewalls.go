package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallNatFirewalls = `{
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
              "ali_uid": "number",
              "id": "string",
              "member_uid": "number",
              "nat_gateway_id": "string",
              "nat_gateway_name": "string",
              "nat_route_entry_list": [
                "list",
                [
                  "object",
                  {
                    "destination_cidr": "string",
                    "nexthop_id": "string",
                    "nexthop_type": "string",
                    "route_table_id": "string"
                  }
                ]
              ],
              "proxy_id": "string",
              "proxy_name": "string",
              "strict_mode": "number",
              "vpc_id": "string"
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
        "type": "number"
      },
      "nat_gateway_id": {
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
      "proxy_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "proxy_name": {
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

func AlicloudCloudFirewallNatFirewallsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallNatFirewalls), &result)
	return &result
}
