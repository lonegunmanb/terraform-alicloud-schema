package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPaiWorkspaceModel = `{
  "block": {
    "attributes": {
      "accessibility": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "extra_info": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_doc": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "model_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "order_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "origin": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "task": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "workspace_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "labels": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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

func AlicloudPaiWorkspaceModelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPaiWorkspaceModel), &result)
	return &result
}
