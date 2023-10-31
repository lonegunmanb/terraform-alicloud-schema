package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVpcPublicIpAddressPoolCidrBlocks = `{
  "block": {
    "attributes": {
      "blocks": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cidr_block": "string",
              "create_time": "string",
              "id": "string",
              "public_ip_address_pool_id": "string",
              "status": "string",
              "total_ip_num": "number",
              "used_ip_num": "number"
            }
          ]
        ]
      },
      "cidr_block": {
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_ip_address_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudVpcPublicIpAddressPoolCidrBlocksSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVpcPublicIpAddressPoolCidrBlocks), &result)
	return &result
}
