package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudNasFileSystem = `{
  "block": {
    "attributes": {
      "capacity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encrypt_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "file_system_type": {
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
      "keytab": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "keytab_md5": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_id": {
        "computed": true,
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_type": {
        "description_kind": "plain",
        "required": true,
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
      "vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "nfs_acl": {
        "block": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "options": {
        "block": {
          "attributes": {
            "enable_oplock": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "recycle_bin": {
        "block": {
          "attributes": {
            "enable_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "reserved_days": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "secondary_size": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "size": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "status": {
              "computed": true,
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
      "smb_acl": {
        "block": {
          "attributes": {
            "enable_anonymous_access": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "encrypt_data": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "home_dir_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "reject_unencrypted_access": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "super_admin_sid": {
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

func AlicloudNasFileSystemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudNasFileSystem), &result)
	return &result
}
