package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDcdnDomains = `{
  "block": {
    "attributes": {
      "change_end_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "change_start_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "check_domain_show": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "domain_search_type": {
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
              "cert_name": "string",
              "cname": "string",
              "description": "string",
              "domain_name": "string",
              "gmt_modified": "string",
              "id": "string",
              "resource_group_id": "string",
              "scope": "string",
              "sources": [
                "list",
                [
                  "object",
                  {
                    "content": "string",
                    "enabled": "string",
                    "port": "number",
                    "priority": "string",
                    "type": "string",
                    "weight": "string"
                  }
                ]
              ],
              "ssl_protocol": "string",
              "ssl_pub": "string",
              "status": "string"
            }
          ]
        ]
      },
      "enable_details": {
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
      "ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name_regex": {
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
      "security_token": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDcdnDomainsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDcdnDomains), &result)
	return &result
}
