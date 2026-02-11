package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudImages = `{
  "block": {
    "attributes": {
      "action_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "architecture": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dry_run": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "image_family": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_owner_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "images": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "architecture": "string",
              "creation_time": "string",
              "description": "string",
              "disk_device_mappings": [
                "list",
                [
                  "object",
                  {
                    "device": "string",
                    "size": "string",
                    "snapshot_id": "string"
                  }
                ]
              ],
              "id": "string",
              "image_id": "string",
              "image_owner_alias": "string",
              "image_version": "string",
              "is_copied": "bool",
              "is_self_shared": "string",
              "is_subscribed": "bool",
              "is_support_io_optimized": "bool",
              "name": "string",
              "os_name": "string",
              "os_name_en": "string",
              "os_type": "string",
              "platform": "string",
              "product_code": "string",
              "progress": "string",
              "size": "number",
              "state": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "usage": "string"
            }
          ]
        ]
      },
      "instance_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "is_support_cloud_init": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "is_support_io_optimized": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "most_recent": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "os_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owners": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshot_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
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
      },
      "usage": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudImagesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudImages), &result)
	return &result
}
