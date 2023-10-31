package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcsImagePipelines = `{
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
      "name": {
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
      "pipelines": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "add_account": [
                "list",
                "string"
              ],
              "base_image": "string",
              "base_image_type": "string",
              "build_content": "string",
              "creation_time": "string",
              "delete_instance_on_failure": "bool",
              "description": "string",
              "id": "string",
              "image_name": "string",
              "image_pipeline_id": "string",
              "instance_type": "string",
              "internet_max_bandwidth_out": "number",
              "name": "string",
              "resource_group_id": "string",
              "system_disk_size": "number",
              "tags": [
                "map",
                "string"
              ],
              "to_region_id": [
                "list",
                "string"
              ],
              "vswitch_id": "string"
            }
          ]
        ]
      },
      "resource_group_id": {
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

func AlicloudEcsImagePipelinesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcsImagePipelines), &result)
	return &result
}
