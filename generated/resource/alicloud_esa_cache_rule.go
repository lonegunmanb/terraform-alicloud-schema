package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaCacheRule = `{
  "block": {
    "attributes": {
      "additional_cacheable_ports": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "browser_cache_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "browser_cache_ttl": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bypass_cache": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cache_deception_armor": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cache_reserve_eligibility": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cache_rule_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "check_presence_cookie": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "check_presence_header": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "edge_cache_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "edge_cache_ttl": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "edge_status_code_cache_ttl": {
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
      "include_cookie": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "include_header": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "query_string": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "query_string_mode": {
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
      "serve_stale": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "sort_query_string_for_cache": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_device_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_geo": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_language": {
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

func AlicloudEsaCacheRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaCacheRule), &result)
	return &result
}
