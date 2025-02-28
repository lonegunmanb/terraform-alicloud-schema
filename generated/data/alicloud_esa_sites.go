package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEsaSites = `{
  "block": {
    "attributes": {
      "access_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "coverage": {
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
      "only_enterprise": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "plan_subscribe_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "site_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "site_search_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sites": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_type": "string",
              "coverage": "string",
              "create_time": "string",
              "id": "number",
              "instance_id": "string",
              "modify_time": "string",
              "name_server_list": "string",
              "resource_group_id": "string",
              "site_id": "number",
              "site_name": "string",
              "status": "string"
            }
          ]
        ]
      },
      "status": {
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

func AlicloudEsaSitesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEsaSites), &result)
	return &result
}
