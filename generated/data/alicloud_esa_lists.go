package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaLists = `{
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
      "lists": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "id": "string",
              "kind": "string",
              "length": "number",
              "list_id": "string",
              "name": "string",
              "update_time": "string"
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
      }
    },
    "block_types": {
      "query_args": {
        "block": {
          "attributes": {
            "desc": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "description_like": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id_like": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "item_like": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kind": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name_item_like": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name_like": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "order_by": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEsaListsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaLists), &result)
	return &result
}
