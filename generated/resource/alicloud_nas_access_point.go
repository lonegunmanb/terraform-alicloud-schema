package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudNasAccessPoint = `{
  "block": {
    "attributes": {
      "access_group": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "access_point_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "access_point_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled_ram": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "file_system_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "root_path": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "posix_user": {
        "block": {
          "attributes": {
            "posix_group_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "posix_secondary_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "number"
              ]
            },
            "posix_user_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "root_path_permission": {
        "block": {
          "attributes": {
            "owner_group_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "owner_user_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "permission": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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

func AlicloudNasAccessPointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudNasAccessPoint), &result)
	return &result
}
