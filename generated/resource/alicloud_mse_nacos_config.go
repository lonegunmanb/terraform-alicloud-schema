package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMseNacosConfig = `{
  "block": {
    "attributes": {
      "accept_language": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "app_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "beta_ips": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "data_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "desc": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encrypted_data_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "group": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
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

func AlicloudMseNacosConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMseNacosConfig), &result)
	return &result
}
