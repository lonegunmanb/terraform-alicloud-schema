package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallPrivateDns = `{
  "block": {
    "attributes": {
      "access_instance_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "access_instance_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_name_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "firewall_type": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_protocol": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "member_uid": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "primary_dns": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "primary_vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "primary_vswitch_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_dns_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region_no": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "standby_dns": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "standby_vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "standby_vswitch_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
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

func AlicloudCloudFirewallPrivateDnsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallPrivateDns), &result)
	return &result
}
