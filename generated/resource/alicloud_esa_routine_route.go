package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaRoutineRoute = `{
  "block": {
    "attributes": {
      "bypass": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "config_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "fallback": {
        "computed": true,
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
      "route_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routine_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sequence": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "site_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
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

func AlicloudEsaRoutineRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaRoutineRoute), &result)
	return &result
}
