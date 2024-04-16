package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbrPolicy = `{
  "block": {
    "attributes": {
      "create_time": {
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
      "policy_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "rules": {
        "block": {
          "attributes": {
            "archive_days": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "backup_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "keep_latest_snapshots": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "replication_region_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "retention": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "rule_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "rule_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "schedule": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vault_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "retention_rules": {
              "block": {
                "attributes": {
                  "advanced_retention_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "retention": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
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

func AlicloudHbrPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbrPolicy), &result)
	return &result
}
