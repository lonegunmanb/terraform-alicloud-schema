package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSnatEntries = `{
  "block": {
    "attributes": {
      "entries": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "snat_entry_id": "string",
              "snat_entry_name": "string",
              "snat_ip": "string",
              "source_cidr": "string",
              "source_vswitch_id": "string",
              "status": "string"
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
      "snat_entry_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snat_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snat_table_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_cidr": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_vswitch_id": {
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

func AlicloudSnatEntriesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSnatEntries), &result)
	return &result
}
