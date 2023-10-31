package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudLogStore = `{
  "block": {
    "attributes": {
      "append_meta": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_split": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_web_tracking": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "hot_ttl": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_split_shard_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "retention_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "shard_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "shards": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "begin_key": "string",
              "end_key": "string",
              "id": "number",
              "status": "string"
            }
          ]
        ]
      },
      "telemetry_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "encrypt_conf": {
        "block": {
          "attributes": {
            "enable": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "encrypt_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "user_cmk_info": {
              "block": {
                "attributes": {
                  "arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "cmk_key_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "region_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
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
            "read": {
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

func AlicloudLogStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudLogStore), &result)
	return &result
}
