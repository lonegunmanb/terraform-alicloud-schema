package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGaEndpointGroupIpAddressCidrBlocks = `{
  "block": {
    "attributes": {
      "accelerator_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_group_ip_address_cidr_blocks": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "endpoint_group_region": "string",
              "ip_address_cidr_blocks": [
                "list",
                "string"
              ],
              "status": "string"
            }
          ]
        ]
      },
      "endpoint_group_region": {
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

func AlicloudGaEndpointGroupIpAddressCidrBlocksSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGaEndpointGroupIpAddressCidrBlocks), &result)
	return &result
}
