package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcdDesktopTypes = `{
  "block": {
    "attributes": {
      "cpu_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "gpu_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "instance_type_family": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "memory_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cpu_count": "string",
              "data_disk_size": "string",
              "desktop_type_id": "string",
              "gpu_count": "number",
              "gpu_spec": "string",
              "id": "string",
              "instance_type_family": "string",
              "memory_size": "string",
              "status": "string",
              "system_disk_size": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEcdDesktopTypesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcdDesktopTypes), &result)
	return &result
}
