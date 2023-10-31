package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSecurityGroupRule = `{
  "block": {
    "attributes": {
      "cidr_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
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
      "ip_protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipv6_cidr_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nic_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port_range": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "prefix_list_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "priority": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "security_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_group_owner_account": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_security_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description": "Type of rule, ingress (inbound) or egress (outbound).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSecurityGroupRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSecurityGroupRule), &result)
	return &result
}
