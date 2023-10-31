package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlidnsRecords = `{
  "block": {
    "attributes": {
      "direction": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "key_word": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "line": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "order_by": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "records": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "domain_name": "string",
              "id": "string",
              "line": "string",
              "locked": "bool",
              "priority": "number",
              "record_id": "string",
              "remark": "string",
              "rr": "string",
              "status": "string",
              "ttl": "number",
              "type": "string",
              "value": "string"
            }
          ]
        ]
      },
      "rr_key_word": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rr_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "search_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type_key_word": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "value_key_word": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "value_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudAlidnsRecordsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlidnsRecords), &result)
	return &result
}
