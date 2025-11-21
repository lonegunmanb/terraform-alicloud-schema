package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaOriginPool = `{
  "block": {
    "attributes": {
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "origin_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "origin_pool_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "site_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "origins": {
        "block": {
          "attributes": {
            "address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "header": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "origin_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "weight": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "auth_conf": {
              "block": {
                "attributes": {
                  "access_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "auth_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "region": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secret_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "version": {
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

func AlicloudEsaOriginPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaOriginPool), &result)
	return &result
}
