package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCddcDedicatedHostGroups = `{
  "block": {
    "attributes": {
      "engine": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allocation_policy": "string",
              "bastion_instance_id": "string",
              "cpu_allocate_ration": "number",
              "cpu_allocated_amount": "number",
              "cpu_allocation_ratio": "number",
              "create_time": "string",
              "dedicated_host_count_group_by_host_type": [
                "list",
                [
                  "object",
                  {
                    "place_holder": "string"
                  }
                ]
              ],
              "dedicated_host_group_desc": "string",
              "dedicated_host_group_id": "string",
              "deploy_type": "string",
              "disk_allocate_ration": "number",
              "disk_allocated_amount": "number",
              "disk_allocation_ratio": "number",
              "disk_used_amount": "number",
              "disk_utility": "number",
              "engine": "string",
              "host_number": "number",
              "host_replace_policy": "string",
              "id": "string",
              "instance_number": "number",
              "mem_allocate_ration": "number",
              "mem_allocated_amount": "number",
              "mem_allocation_ratio": "number",
              "mem_used_amount": "number",
              "mem_utility": "number",
              "text": "string",
              "vpc_id": "string",
              "zone_id_list": [
                "list",
                [
                  "object",
                  {
                    "zone_id_list": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCddcDedicatedHostGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCddcDedicatedHostGroups), &result)
	return &result
}
