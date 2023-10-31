package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudKmsSecret = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dkms_instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_automatic_rotation": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "encryption_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "extended_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "force_delete_without_recovery": {
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
      "planned_delete_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "recovery_window_in_days": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "rotation_interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secret_data": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "secret_data_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secret_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secret_type": {
        "computed": true,
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
      "version_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "version_stages": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
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

func AlicloudKmsSecretSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudKmsSecret), &result)
	return &result
}
