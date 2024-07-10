package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudGpdbRemoteAdbDataSource = `{
  "block": {
    "attributes": {
      "data_source_name": {
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
      "local_database": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "local_db_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "manager_user_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "manager_user_password": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "remote_adb_data_source_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "remote_database": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "remote_db_instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_password": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
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

func AlicloudGpdbRemoteAdbDataSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudGpdbRemoteAdbDataSource), &result)
	return &result
}
