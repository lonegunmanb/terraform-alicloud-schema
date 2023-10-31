package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbaseInstances = `{
  "block": {
    "attributes": {
      "availability_zone": {
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
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backup_status": "string",
              "core_disk_size": "number",
              "core_disk_type": "string",
              "core_instance_type": "string",
              "core_node_count": "number",
              "created_time": "string",
              "deletion_protection": "bool",
              "engine": "string",
              "engine_version": "string",
              "expire_time": "string",
              "id": "string",
              "master_instance_type": "string",
              "master_node_count": "number",
              "name": "string",
              "network_type": "string",
              "pay_type": "string",
              "region_id": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "vpc_id": "string",
              "vswitch_id": "string",
              "zone_id": "string"
            }
          ]
        ]
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
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudHbaseInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbaseInstances), &result)
	return &result
}
