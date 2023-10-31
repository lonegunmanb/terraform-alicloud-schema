package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCenRouteServices = `{
  "block": {
    "attributes": {
      "access_region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cen_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "host": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "host_region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "host_vpc_id": {
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
      "services": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_region_id": "string",
              "cen_id": "string",
              "cidrs": [
                "list",
                "string"
              ],
              "description": "string",
              "host": "string",
              "host_region_id": "string",
              "host_vpc_id": "string",
              "id": "string",
              "status": "string",
              "update_interval": "string"
            }
          ]
        ]
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

func AlicloudCenRouteServicesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCenRouteServices), &result)
	return &result
}
