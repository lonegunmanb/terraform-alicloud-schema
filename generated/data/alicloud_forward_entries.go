package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudForwardEntries = `{
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
              "external_ip": "string",
              "external_port": "string",
              "forward_entry_id": "string",
              "forward_entry_name": "string",
              "id": "string",
              "internal_ip": "string",
              "internal_port": "string",
              "ip_protocol": "string",
              "name": "string",
              "status": "string"
            }
          ]
        ]
      },
      "external_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "external_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "forward_entry_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "forward_table_id": {
        "description_kind": "plain",
        "required": true,
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
      "internal_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internal_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_protocol": {
        "description_kind": "plain",
        "optional": true,
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

func AlicloudForwardEntriesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudForwardEntries), &result)
	return &result
}
