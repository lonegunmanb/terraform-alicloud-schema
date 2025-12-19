package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaSite = `{
  "block": {
    "attributes": {
      "access_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "add_client_geolocation_header": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "add_real_client_ip_header": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cache_architecture_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cache_reserve_enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cache_reserve_instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "case_insensitive": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "coverage": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cross_border_optimization": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "development_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flatten_mode": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipv6_enable": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_region": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "paused": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "seo_bypass": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "site_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "site_name_exclusive": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "site_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tag_name": {
        "description_kind": "plain",
        "optional": true,
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
      "version_management": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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

func AlicloudEsaSiteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaSite), &result)
	return &result
}
