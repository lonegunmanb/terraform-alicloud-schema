package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDcdnIpaDomains = `{
  "block": {
    "attributes": {
      "domain_name": {
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
              "create_time": "string",
              "description": "string",
              "domain_name": "string",
              "id": "string",
              "resource_group_id": "string",
              "scope": "string",
              "sources": [
                "list",
                [
                  "object",
                  {
                    "content": "string",
                    "port": "number",
                    "priority": "string",
                    "type": "string",
                    "weight": "number"
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

func AlicloudDcdnIpaDomainsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDcdnIpaDomains), &result)
	return &result
}
