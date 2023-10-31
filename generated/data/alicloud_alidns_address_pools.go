package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlidnsAddressPools = `{
  "block": {
    "attributes": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      "pools": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address": [
                "list",
                [
                  "object",
                  {
                    "address": "string",
                    "attribute_info": "string",
                    "lba_weight": "number",
                    "mode": "string",
                    "remark": "string"
                  }
                ]
              ],
              "address_pool_id": "string",
              "address_pool_name": "string",
              "create_time": "string",
              "create_timestamp": "string",
              "id": "string",
              "instance_id": "string",
              "lba_strategy": "string",
              "monitor_config_id": "string",
              "monitor_status": "string",
              "type": "string",
              "update_time": "string",
              "update_timestamp": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudAlidnsAddressPoolsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlidnsAddressPools), &result)
	return &result
}
