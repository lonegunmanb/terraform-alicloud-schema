package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcdDesktopGroup = `{
  "block": {
    "attributes": {
      "allow_auto_setup": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "number"
      },
      "allow_buffer_count": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "number"
      },
      "bundle_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "comments": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cpu": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creator": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_disk_category": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_disk_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "desktop_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "directory_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "directory_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "end_user_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "expired_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gpu_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "gpu_spec": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "keep_duration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_desktops_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "memory": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "min_desktops_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "office_site_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "office_site_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "office_site_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "own_bundle_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pay_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_group_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "res_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "scale_strategy_id": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "system_disk_category": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "system_disk_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
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

func AlicloudEcdDesktopGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcdDesktopGroup), &result)
	return &result
}
