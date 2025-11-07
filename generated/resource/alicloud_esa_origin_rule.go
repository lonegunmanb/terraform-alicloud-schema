package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaOriginRule = `{
  "block": {
    "attributes": {
      "config_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "dns_record": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "follow302_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "follow302_max_tries": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "follow302_retain_args": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "follow302_retain_header": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "follow302_target_host": {
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
      "origin_host": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "origin_http_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "origin_https_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "origin_mtls": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "origin_read_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "origin_scheme": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "origin_sni": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "origin_verify": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "range": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "range_chunk_size": {
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

func AlicloudEsaOriginRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaOriginRule), &result)
	return &result
}
