package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDcdnWafRule = `{
  "block": {
    "attributes": {
      "action": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cc_status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cn_region_list": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "defense_scene": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "effect": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gmt_modified": {
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
      "other_region_list": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "regular_rules": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "regular_types": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "remote_addr": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "rule_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scenes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "waf_group_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "conditions": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "op_value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sub_key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "rate_limit": {
        "block": {
          "attributes": {
            "interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "sub_key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "threshold": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ttl": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "status": {
              "block": {
                "attributes": {
                  "code": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "count": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ratio": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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

func AlicloudDcdnWafRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDcdnWafRule), &result)
	return &result
}
