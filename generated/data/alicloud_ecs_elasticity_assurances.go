package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcsElasticityAssurances = `{
  "block": {
    "attributes": {
      "assurances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allocated_resources": [
                "list",
                [
                  "object",
                  {
                    "instance_type": "string",
                    "total_amount": "number",
                    "used_amount": "number",
                    "zone_id": "string"
                  }
                ]
              ],
              "description": "string",
              "elasticity_assurance_id": "string",
              "end_time": "string",
              "id": "string",
              "instance_charge_type": "string",
              "private_pool_options_id": "string",
              "private_pool_options_match_criteria": "string",
              "private_pool_options_name": "string",
              "resource_group_id": "string",
              "start_time": "string",
              "start_time_type": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "total_assurance_times": "string",
              "used_assurance_times": "number"
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_pool_options_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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

func AlicloudEcsElasticityAssurancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcsElasticityAssurances), &result)
	return &result
}
