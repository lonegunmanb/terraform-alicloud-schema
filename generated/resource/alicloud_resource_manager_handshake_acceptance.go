package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudResourceManagerHandshakeAcceptance = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expire_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "handshake_id": {
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
      "invited_account_real_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "master_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "master_account_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "master_account_real_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "modify_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "note": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_directory_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_entity": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_type": {
        "computed": true,
        "description_kind": "plain",
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

func AlicloudResourceManagerHandshakeAcceptanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudResourceManagerHandshakeAcceptance), &result)
	return &result
}
