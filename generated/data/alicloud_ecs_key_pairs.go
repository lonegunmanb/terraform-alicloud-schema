package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcsKeyPairs = `{
  "block": {
    "attributes": {
      "finger_print": {
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
      "key_pairs": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "finger_print": "string",
              "id": "string",
              "instances": [
                "list",
                [
                  "object",
                  {
                    "availability_zone": "string",
                    "description": "string",
                    "image_id": "string",
                    "instance_id": "string",
                    "instance_name": "string",
                    "instance_type": "string",
                    "key_name": "string",
                    "private_ip": "string",
                    "public_ip": "string",
                    "region_id": "string",
                    "status": "string",
                    "vswitch_id": "string"
                  }
                ]
              ],
              "key_name": "string",
              "key_pair_name": "string",
              "resource_group_id": "string",
              "tags": [
                "map",
                "string"
              ]
            }
          ]
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
      "pairs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "finger_print": "string",
              "id": "string",
              "instances": [
                "list",
                [
                  "object",
                  {
                    "availability_zone": "string",
                    "description": "string",
                    "image_id": "string",
                    "instance_id": "string",
                    "instance_name": "string",
                    "instance_type": "string",
                    "key_name": "string",
                    "private_ip": "string",
                    "public_ip": "string",
                    "region_id": "string",
                    "status": "string",
                    "vswitch_id": "string"
                  }
                ]
              ],
              "key_name": "string",
              "key_pair_name": "string",
              "resource_group_id": "string",
              "tags": [
                "map",
                "string"
              ]
            }
          ]
        ]
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

func AlicloudEcsKeyPairsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcsKeyPairs), &result)
	return &result
}
