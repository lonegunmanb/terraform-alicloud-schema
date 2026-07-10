package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDrdsPolardbxInstances = `{
  "block": {
    "attributes": {
      "description_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "descriptions": {
        "computed": true,
        "description_kind": "plain",
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cn_class": "string",
              "cn_node_count": "number",
              "create_time": "string",
              "description": "string",
              "dn_class": "string",
              "dn_node_count": "number",
              "engine_version": "string",
              "id": "string",
              "network_type": "string",
              "payment_type": "string",
              "polardbx_instance_id": "string",
              "primary_zone": "string",
              "region_id": "string",
              "resource_group_id": "string",
              "secondary_zone": "string",
              "status": "string",
              "storage_type": "string",
              "tertiary_zone": "string",
              "topology_type": "string",
              "vpc_id": "string",
              "zone_id": "string"
            }
          ]
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "total_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDrdsPolardbxInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDrdsPolardbxInstances), &result)
	return &result
}
