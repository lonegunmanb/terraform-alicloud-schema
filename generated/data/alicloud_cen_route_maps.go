package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenRouteMaps = `{
  "block": {
    "attributes": {
      "cen_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cen_region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description_regex": {
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
      "maps": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "as_path_match_mode": "string",
              "cen_id": "string",
              "cen_region_id": "string",
              "cidr_match_mode": "string",
              "community_match_mode": "string",
              "community_operate_mode": "string",
              "description": "string",
              "destination_child_instance_types": [
                "list",
                "string"
              ],
              "destination_cidr_blocks": [
                "list",
                "string"
              ],
              "destination_instance_ids": [
                "list",
                "string"
              ],
              "destination_instance_ids_reverse_match": "bool",
              "destination_route_table_ids": [
                "list",
                "string"
              ],
              "id": "string",
              "map_result": "string",
              "match_asns": [
                "list",
                "string"
              ],
              "match_community_set": [
                "list",
                "string"
              ],
              "next_priority": "number",
              "operate_community_set": [
                "list",
                "string"
              ],
              "preference": "number",
              "prepend_as_path": [
                "list",
                "string"
              ],
              "priority": "number",
              "route_map_id": "string",
              "route_types": [
                "list",
                "string"
              ],
              "source_child_instance_types": [
                "list",
                "string"
              ],
              "source_instance_ids": [
                "list",
                "string"
              ],
              "source_instance_ids_reverse_match": "bool",
              "source_region_ids": [
                "list",
                "string"
              ],
              "source_route_table_ids": [
                "list",
                "string"
              ],
              "status": "string",
              "transmit_direction": "string"
            }
          ]
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transmit_direction": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCenRouteMapsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenRouteMaps), &result)
	return &result
}
