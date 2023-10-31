package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCddcDedicatedHostGroup = `{
  "block": {
    "attributes": {
      "allocation_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cpu_allocation_ratio": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "dedicated_host_group_desc": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disk_allocation_ratio": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "engine": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "host_replace_policy": {
        "computed": true,
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
      "mem_allocation_ratio": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "open_permission": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCddcDedicatedHostGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCddcDedicatedHostGroup), &result)
	return &result
}
