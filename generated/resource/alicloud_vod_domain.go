package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudVodDomain = `{
  "block": {
    "attributes": {
      "cert_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "check_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cname": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "gmt_created": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gmt_modified": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_protocol": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ssl_pub": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "top_level_domain": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "weight": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "sources": {
        "block": {
          "attributes": {
            "source_content": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_port": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_priority": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
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

func AlicloudVodDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudVodDomain), &result)
	return &result
}
