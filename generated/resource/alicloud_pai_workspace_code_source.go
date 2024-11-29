package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPaiWorkspaceCodeSource = `{
  "block": {
    "attributes": {
      "accessibility": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "code_branch": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "code_commit": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "code_repo": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "code_repo_access_token": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "code_repo_user_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "display_name": {
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
      "mount_path": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "workspace_id": {
        "description_kind": "plain",
        "required": true,
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

func AlicloudPaiWorkspaceCodeSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPaiWorkspaceCodeSource), &result)
	return &result
}
