package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlsMetricStore = `{
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
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "hot_ttl": {
        "computed": true,
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
      "infrequent_access_ttl": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "last_modify_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "max_split_shard_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "metric_store_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "shard_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "ttl": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "block_types": {
      "encrypt_conf": {
        "block": {
          "attributes": {
            "enable": {
              "description_kind": "plain",
              "required": true,
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
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cmk_key_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "region_id": {
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

func AlicloudSlsMetricStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlsMetricStore), &result)
	return &result
}
