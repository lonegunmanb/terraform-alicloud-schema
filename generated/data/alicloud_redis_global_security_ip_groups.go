package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudRedisGlobalSecurityIpGroups = `{
  "block": {
    "attributes": {
      "engine": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "global_security_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "global_ip_group_name": "string",
              "global_ip_list": "string",
              "global_security_group_id": "string",
              "id": "string",
              "region_id": "string"
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

func AlicloudRedisGlobalSecurityIpGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudRedisGlobalSecurityIpGroups), &result)
	return &result
}
