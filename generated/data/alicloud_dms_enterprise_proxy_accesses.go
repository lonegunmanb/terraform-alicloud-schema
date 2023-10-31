package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDmsEnterpriseProxyAccesses = `{
  "block": {
    "attributes": {
      "accesses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_id": "string",
              "access_secret": "string",
              "create_time": "string",
              "id": "string",
              "indep_account": "string",
              "instance_id": "string",
              "origin_info": "string",
              "proxy_access_id": "string",
              "proxy_id": "string",
              "user_id": "string",
              "user_name": "string",
              "user_uid": "string"
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "proxy_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDmsEnterpriseProxyAccessesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDmsEnterpriseProxyAccesses), &result)
	return &result
}
