package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudHbaseInstanceTypes = `{
  "block": {
    "attributes": {
      "charge_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "core_instance_types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "category": "string",
              "cpu_size": "number",
              "engine": "string",
              "instance_type": "string",
              "max_core_count": "number",
              "mem_size": "number",
              "storage_type": "string",
              "version": "string",
              "zone": "string"
            }
          ]
        ]
      },
      "disk_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine": {
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
        "type": [
          "list",
          "string"
        ]
      },
      "instance_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_instance_types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cpu_size": "number",
              "instance_type": "string",
              "mem_size": "number"
            }
          ]
        ]
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cpu_size": "number",
              "mem_size": "number",
              "value": "string"
            }
          ]
        ]
      },
      "version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudHbaseInstanceTypesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudHbaseInstanceTypes), &result)
	return &result
}
