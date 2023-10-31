package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPvtzRules = `{
  "block": {
    "attributes": {
      "endpoint_id": {
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
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bind_vpcs": [
                "list",
                [
                  "object",
                  {
                    "region_id": "string",
                    "region_name": "string",
                    "vpc_id": "string",
                    "vpc_name": "string"
                  }
                ]
              ],
              "create_time": "string",
              "endpoint_id": "string",
              "endpoint_name": "string",
              "forward_ips": [
                "list",
                [
                  "object",
                  {
                    "ip": "string",
                    "port": "number"
                  }
                ]
              ],
              "id": "string",
              "rule_id": "string",
              "rule_name": "string",
              "type": "string",
              "zone_name": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudPvtzRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPvtzRules), &result)
	return &result
}
