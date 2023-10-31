package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMaxcomputeProjects = `{
  "block": {
    "attributes": {
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
      "projects": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "comment": "string",
              "default_quota": "string",
              "id": "string",
              "ip_white_list": [
                "list",
                [
                  "object",
                  {
                    "ip_list": "string",
                    "vpc_ip_list": "string"
                  }
                ]
              ],
              "owner": "string",
              "project_name": "string",
              "properties": [
                "list",
                [
                  "object",
                  {
                    "allow_full_scan": "bool",
                    "enable_decimal2": "bool",
                    "encryption": [
                      "list",
                      [
                        "object",
                        {
                          "algorithm": "string",
                          "enable": "bool",
                          "key": "string"
                        }
                      ]
                    ],
                    "retention_days": "string",
                    "sql_metering_max": "string",
                    "table_lifecycle": [
                      "list",
                      [
                        "object",
                        {
                          "type": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "timezone": "string",
                    "type_system": "string"
                  }
                ]
              ],
              "security_properties": [
                "list",
                [
                  "object",
                  {
                    "enable_download_privilege": "bool",
                    "label_security": "bool",
                    "object_creator_has_access_permission": "bool",
                    "object_creator_has_grant_permission": "bool",
                    "project_protection": [
                      "list",
                      [
                        "object",
                        {
                          "exception_policy": "string",
                          "protected": "bool"
                        }
                      ]
                    ],
                    "using_acl": "bool",
                    "using_policy": "bool"
                  }
                ]
              ],
              "status": "string",
              "type": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMaxcomputeProjectsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMaxcomputeProjects), &result)
	return &result
}
