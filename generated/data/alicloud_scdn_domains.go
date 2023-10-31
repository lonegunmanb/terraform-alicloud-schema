package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudScdnDomains = `{
  "block": {
    "attributes": {
      "domains": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cert_infos": [
                "list",
                [
                  "object",
                  {
                    "cert_name": "string",
                    "cert_type": "string",
                    "ssl_protocol": "string",
                    "ssl_pub": "string"
                  }
                ]
              ],
              "cname": "string",
              "create_time": "string",
              "description": "string",
              "domain_name": "string",
              "gmt_modified": "string",
              "id": "string",
              "resource_group_id": "string",
              "sources": [
                "list",
                [
                  "object",
                  {
                    "content": "string",
                    "enabled": "string",
                    "port": "number",
                    "priority": "string",
                    "type": "string"
                  }
                ]
              ],
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

func AlicloudScdnDomainsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudScdnDomains), &result)
	return &result
}
