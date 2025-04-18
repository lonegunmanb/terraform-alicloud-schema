package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlsEtl = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
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
      "job_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "configuration": {
        "block": {
          "attributes": {
            "from_time": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "lang": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "logstore": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "parameters": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "script": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "to_time": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "sink": {
              "block": {
                "attributes": {
                  "datasets": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "endpoint": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "logstore": {
                    "description_kind": "plain",
                    "required": true,
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
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
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

func AlicloudSlsEtlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlsEtl), &result)
	return &result
}
