package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaRewriteUrlRule = `{
  "block": {
    "attributes": {
      "config_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "query_string": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rewrite_query_string_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rewrite_uri_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sequence": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "site_id": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "site_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "uri": {
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

func AlicloudEsaRewriteUrlRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaRewriteUrlRule), &result)
	return &result
}
