package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallVpcCenTrFirewall = `{
  "block": {
    "attributes": {
      "cen_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "firewall_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "firewall_eni_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_eni_vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "firewall_subnet_cidr": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "firewall_vpc_attachment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_vpc_cidr": {
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
      "region_no": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "route_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tr_attachment_master_cidr": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tr_attachment_master_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tr_attachment_slave_cidr": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tr_attachment_slave_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_router_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
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

func AlicloudCloudFirewallVpcCenTrFirewallSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallVpcCenTrFirewall), &result)
	return &result
}
