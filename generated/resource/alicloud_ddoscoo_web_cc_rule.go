package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDdoscooWebCcRule = `{
  "block": {
    "attributes": {
      "domain": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "rule_detail": {
        "block": {
          "attributes": {
            "action": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "condition": {
              "block": {
                "attributes": {
                  "content": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "field": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "header_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "match_method": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "set"
            },
            "rate_limit": {
              "block": {
                "attributes": {
                  "interval": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "sub_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "threshold": {
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
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "statistics": {
              "block": {
                "attributes": {
                  "field": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "header_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mode": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "status_code": {
              "block": {
                "attributes": {
                  "code": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "count_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "ratio_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "use_ratio": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
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
        "min_items": 1,
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

func AlicloudDdoscooWebCcRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDdoscooWebCcRule), &result)
	return &result
}
