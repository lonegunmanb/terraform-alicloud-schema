package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionClientUserDefineRule = `{
  "block": {
    "attributes": {
      "action_type": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "client_user_define_rule_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cmdline": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "file_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hash": {
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
      "ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "new_file_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parent_cmdline": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parent_proc_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "platform": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "port_str": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "proc_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "registry_content": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "registry_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
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

func AlicloudThreatDetectionClientUserDefineRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionClientUserDefineRule), &result)
	return &result
}
