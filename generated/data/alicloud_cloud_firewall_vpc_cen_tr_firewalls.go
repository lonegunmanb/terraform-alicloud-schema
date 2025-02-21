package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallVpcCenTrFirewalls = `{
  "block": {
    "attributes": {
      "cen_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "current_page": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "firewall_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "firewall_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "firewall_switch_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "firewalls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cen_id": "string",
              "cen_name": "string",
              "firewall_id": "string",
              "firewall_name": "string",
              "firewall_switch_status": "string",
              "id": "string",
              "ips_config": [
                "list",
                [
                  "object",
                  {
                    "basic_rules": "number",
                    "enable_all_patch": "number",
                    "run_mode": "number"
                  }
                ]
              ],
              "precheck_status": "string",
              "region_no": "string",
              "region_status": "string",
              "result_code": "string",
              "route_mode": "string",
              "transit_router_id": "string"
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
      "region_no": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_router_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudFirewallVpcCenTrFirewallsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallVpcCenTrFirewalls), &result)
	return &result
}
