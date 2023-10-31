package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEssScalinggroupVserverGroups = `{
  "block": {
    "attributes": {
      "force": {
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
      "scaling_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "vserver_groups": {
        "block": {
          "attributes": {
            "loadbalancer_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "vserver_attributes": {
              "block": {
                "attributes": {
                  "port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "vserver_group_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "weight": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
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

func AlicloudEssScalinggroupVserverGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEssScalinggroupVserverGroups), &result)
	return &result
}
