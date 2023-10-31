package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcdBundles = `{
  "block": {
    "attributes": {
      "bundle_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "bundle_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bundles": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bundle_id": "string",
              "bundle_name": "string",
              "bundle_type": "string",
              "description": "string",
              "desktop_type": "string",
              "desktop_type_attribute": [
                "list",
                [
                  "object",
                  {
                    "cpu_count": "number",
                    "gpu_count": "string",
                    "gpu_spec": "string",
                    "memory_size": "string"
                  }
                ]
              ],
              "disks": [
                "list",
                [
                  "object",
                  {
                    "disk_size": "string",
                    "disk_type": "string"
                  }
                ]
              ],
              "id": "string",
              "image_id": "string",
              "os_type": "string"
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

func AlicloudEcdBundlesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcdBundles), &result)
	return &result
}
