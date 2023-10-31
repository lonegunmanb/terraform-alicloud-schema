package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPvtzZones = `{
  "block": {
    "attributes": {
      "enable_details": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "keyword": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
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
      "query_region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "query_vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "search_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bind_vpcs": [
                "list",
                [
                  "object",
                  {
                    "region_id": "string",
                    "region_name": "string",
                    "vpc_id": "string",
                    "vpc_name": "string"
                  }
                ]
              ],
              "create_timestamp": "number",
              "creation_time": "string",
              "id": "string",
              "is_ptr": "bool",
              "name": "string",
              "proxy_pattern": "string",
              "record_count": "number",
              "remark": "string",
              "resource_group_id": "string",
              "slave_dns": "bool",
              "update_time": "string",
              "update_timestamp": "number",
              "zone_id": "string",
              "zone_name": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudPvtzZonesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPvtzZones), &result)
	return &result
}
