package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallIpsConfig = `{
  "block": {
    "attributes": {
      "basic_rules": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "cti_rules": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_sdl": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "patch_rules": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "rule_class": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "run_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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

func AlicloudCloudFirewallIpsConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallIpsConfig), &result)
	return &result
}
