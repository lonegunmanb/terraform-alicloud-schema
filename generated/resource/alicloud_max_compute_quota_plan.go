package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMaxComputeQuotaPlan = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "is_effective": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "nickname": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "plan_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "quota": {
        "block": {
          "block_types": {
            "parameter": {
              "block": {
                "attributes": {
                  "elastic_reserved_cu": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "max_cu": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "min_cu": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "sub_quota_info_list": {
              "block": {
                "attributes": {
                  "nick_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "parameter": {
                    "block": {
                      "attributes": {
                        "elastic_reserved_cu": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "max_cu": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "min_cu": {
                          "description_kind": "plain",
                          "required": true,
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
              "nesting_mode": "set"
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

func AlicloudMaxComputeQuotaPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMaxComputeQuotaPlan), &result)
	return &result
}
