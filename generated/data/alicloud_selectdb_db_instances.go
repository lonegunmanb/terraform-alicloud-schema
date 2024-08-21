package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSelectdbDbInstances = `{
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
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cache_size_postpaid": "number",
              "cache_size_prepaid": "number",
              "cluster_count_postpaid": "number",
              "cluster_count_prepaid": "number",
              "cpu_postpaid": "number",
              "cpu_prepaid": "number",
              "db_instance_description": "string",
              "db_instance_id": "string",
              "engine": "string",
              "engine_minor_version": "string",
              "engine_version": "string",
              "gmt_created": "string",
              "gmt_expired": "string",
              "gmt_modified": "string",
              "id": "string",
              "lock_mode": "string",
              "lock_reason": "string",
              "memory_postpaid": "number",
              "memory_prepaid": "number",
              "payment_type": "string",
              "region_id": "string",
              "status": "string",
              "sub_domain": "string",
              "vpc_id": "string",
              "vswitch_id": "string",
              "zone_id": "string"
            }
          ]
        ]
      },
      "output_file": {
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

func AlicloudSelectdbDbInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSelectdbDbInstances), &result)
	return &result
}
