package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudLogStoreIndex = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_reduce": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "log_reduce_black_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "log_reduce_white_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "logstore": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_text_len": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "field_search": {
        "block": {
          "attributes": {
            "alias": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "case_sensitive": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_analytics": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_chinese": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "token": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "json_keys": {
              "block": {
                "attributes": {
                  "alias": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "doc_value": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "full_text": {
        "block": {
          "attributes": {
            "case_sensitive": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_chinese": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "token": {
              "description_kind": "plain",
              "optional": true,
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
  "version": 0
}`

func AlicloudLogStoreIndexSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudLogStoreIndex), &result)
	return &result
}
