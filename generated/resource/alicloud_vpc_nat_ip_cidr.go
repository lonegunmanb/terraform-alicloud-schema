package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpcNatIpCidr = `{
  "block": {
    "attributes": {
      "dry_run": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nat_gateway_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nat_ip_cidr": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nat_ip_cidr_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nat_ip_cidr_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudVpcNatIpCidrSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpcNatIpCidr), &result)
	return &result
}
