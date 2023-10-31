package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlbAcls = `{
  "block": {
    "attributes": {
      "acls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "entry_list": [
                "list",
                [
                  "object",
                  {
                    "comment": "string",
                    "entry": "string"
                  }
                ]
              ],
              "id": "string",
              "ip_version": "string",
              "name": "string",
              "related_listeners": [
                "list",
                [
                  "object",
                  {
                    "acl_type": "string",
                    "frontend_port": "number",
                    "load_balancer_id": "string",
                    "protocol": "string"
                  }
                ]
              ],
              "resource_group_id": "string",
              "tags": [
                "map",
                "string"
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
      },
      "resource_group_id": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSlbAclsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlbAcls), &result)
	return &result
}
