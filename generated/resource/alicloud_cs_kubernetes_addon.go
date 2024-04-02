package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCsKubernetesAddon = `{
  "block": {
    "attributes": {
      "can_upgrade": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "cleanup_cloud_resources": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "config": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "next_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "required": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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

func AlicloudCsKubernetesAddonSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCsKubernetesAddon), &result)
	return &result
}
