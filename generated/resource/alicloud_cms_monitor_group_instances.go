package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCmsMonitorGroupInstances = `{
  "block": {
    "attributes": {
      "group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "instances": {
        "block": {
          "attributes": {
            "category": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "instance_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "instance_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "region_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCmsMonitorGroupInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCmsMonitorGroupInstances), &result)
	return &result
}
