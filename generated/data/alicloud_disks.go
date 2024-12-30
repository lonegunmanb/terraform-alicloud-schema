package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDisks = `{
  "block": {
    "attributes": {
      "additional_attributes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "auto_snapshot_policy_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zone": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delete_auto_snapshot": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "delete_with_instance": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disk_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disk_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disks": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "attached_time": "string",
              "auto_snapshot_policy_id": "string",
              "availability_zone": "string",
              "category": "string",
              "creation_time": "string",
              "delete_auto_snapshot": "bool",
              "delete_with_instance": "bool",
              "description": "string",
              "detached_time": "string",
              "device": "string",
              "disk_id": "string",
              "disk_name": "string",
              "disk_type": "string",
              "enable_auto_snapshot": "bool",
              "enable_automated_snapshot_policy": "bool",
              "encrypted": "string",
              "expiration_time": "string",
              "expired_time": "string",
              "id": "string",
              "image_id": "string",
              "instance_id": "string",
              "iops": "number",
              "iops_read": "number",
              "iops_write": "number",
              "kms_key_id": "string",
              "mount_instance_num": "number",
              "mount_instances": [
                "list",
                [
                  "object",
                  {
                    "attached_time": "string",
                    "device": "string",
                    "instance_id": "string"
                  }
                ]
              ],
              "name": "string",
              "operation_locks": [
                "list",
                [
                  "object",
                  {
                    "lock_reason": "string"
                  }
                ]
              ],
              "payment_type": "string",
              "performance_level": "string",
              "portable": "bool",
              "product_code": "string",
              "region_id": "string",
              "resource_group_id": "string",
              "size": "number",
              "snapshot_id": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "type": "string",
              "zone_id": "string"
            }
          ]
        ]
      },
      "dry_run": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_auto_snapshot": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_automated_snapshot_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_shared": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "encrypted": {
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
      "instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_id": {
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
      "payment_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "portable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "total_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "type": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "operation_locks": {
        "block": {
          "attributes": {
            "lock_reason": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDisksSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDisks), &result)
	return &result
}
