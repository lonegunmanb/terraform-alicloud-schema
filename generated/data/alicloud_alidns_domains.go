package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlidnsDomains = `{
  "block": {
    "attributes": {
      "ali_domain": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "domain_name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domains": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ali_domain": "bool",
              "available_ttls": [
                "list",
                "number"
              ],
              "dns_servers": [
                "list",
                "string"
              ],
              "domain_id": "string",
              "domain_name": "string",
              "group_id": "string",
              "group_name": "string",
              "id": "string",
              "in_black_hole": "bool",
              "in_clean": "bool",
              "instance_id": "string",
              "line_type": "string",
              "min_ttl": "number",
              "puny_code": "string",
              "record_line_tree_json": "string",
              "record_lines": [
                "list",
                [
                  "object",
                  {
                    "father_code": "string",
                    "line_code": "string",
                    "line_display_name": "string",
                    "line_name": "string"
                  }
                ]
              ],
              "region_lines": "bool",
              "remark": "string",
              "resource_group_id": "string",
              "slave_dns": "bool",
              "tags": [
                "map",
                "string"
              ],
              "version_code": "string",
              "version_name": "string"
            }
          ]
        ]
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "group_name_regex": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_word": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "search_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "starmark": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "version_code": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudAlidnsDomainsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlidnsDomains), &result)
	return &result
}
