package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudArmsDispatchRule = `{
  "block": {
    "attributes": {
      "dispatch_rule_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dispatch_type": {
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
      "is_recover": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "group_rules": {
        "block": {
          "attributes": {
            "group_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "group_interval": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "group_wait_time": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "grouping_fields": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "repeat_interval": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      },
      "label_match_expression_grid": {
        "block": {
          "block_types": {
            "label_match_expression_groups": {
              "block": {
                "block_types": {
                  "label_match_expressions": {
                    "block": {
                      "attributes": {
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      },
      "notify_rules": {
        "block": {
          "attributes": {
            "notify_channels": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "notify_end_time": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "notify_start_time": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "notify_objects": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "notify_object_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "notify_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudArmsDispatchRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudArmsDispatchRule), &result)
	return &result
}
