package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenRouteEntries = `{
  "block": {
    "attributes": {
      "cidr_block": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "entries": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cidr_block": "string",
              "conflicts": [
                "list",
                [
                  "object",
                  {
                    "cidr_block": "string",
                    "instance_id": "string",
                    "instance_type": "string",
                    "region_id": "string",
                    "status": "string"
                  }
                ]
              ],
              "next_hop_id": "string",
              "next_hop_type": "string",
              "operational_mode": "bool",
              "publish_status": "string",
              "route_table_id": "string",
              "route_type": "string"
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_table_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCenRouteEntriesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenRouteEntries), &result)
	return &result
}
