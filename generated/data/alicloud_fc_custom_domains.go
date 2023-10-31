package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudFcCustomDomains = `{
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
              "account_id": "string",
              "api_version": "string",
              "cert_config": [
                "list",
                [
                  "object",
                  {
                    "cert_name": "string",
                    "certificate": "string"
                  }
                ]
              ],
              "created_time": "string",
              "domain_name": "string",
              "id": "string",
              "last_modified_time": "string",
              "protocol": "string",
              "route_config": [
                "list",
                [
                  "object",
                  {
                    "function_name": "string",
                    "methods": [
                      "list",
                      "string"
                    ],
                    "path": "string",
                    "qualifier": "string",
                    "service_name": "string"
                  }
                ]
              ]
            }
          ]
        ]
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudFcCustomDomainsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudFcCustomDomains), &result)
	return &result
}
