package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDbfsInstances = `{
  "block": {
    "attributes": {
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
              "attach_node_number": "number",
              "category": "string",
              "create_time": "string",
              "dbfs_cluster_id": "string",
              "ecs_list": [
                "list",
                [
                  "object",
                  {
                    "ecs_id": "string"
                  }
                ]
              ],
              "enable_raid": "bool",
              "encryption": "bool",
              "id": "string",
              "instance_id": "string",
              "instance_name": "string",
              "kms_key_id": "string",
              "payment_type": "string",
              "performance_level": "string",
              "raid_stripe_unit_number": "string",
              "size": "number",
              "status": "string",
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

func AlicloudDbfsInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDbfsInstances), &result)
	return &result
}
