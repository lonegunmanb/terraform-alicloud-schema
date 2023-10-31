package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudOpenSearchAppGroups = `{
  "block": {
    "attributes": {
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "app_group_id": "string",
              "app_group_name": "string",
              "charge_way": "number",
              "commodity_code": "string",
              "create_time": "number",
              "current_version": "string",
              "description": "string",
              "domain": "string",
              "expire_on": "string",
              "first_rank_algo_deployment_id": "number",
              "has_pending_quota_review_task": "number",
              "id": "string",
              "instance_id": "string",
              "lock_mode": "string",
              "locked_by_expiration": "number",
              "payment_type": "string",
              "pending_second_rank_algo_deployment_id": "number",
              "processing_order_id": "string",
              "produced": "number",
              "project_id": "string",
              "quota": [
                "list",
                [
                  "object",
                  {
                    "compute_resource": "string",
                    "doc_size": "string",
                    "spec": "string"
                  }
                ]
              ],
              "resource_group_id": "string",
              "second_rank_algo_deployment_id": "number",
              "status": "string",
              "switched_time": "number",
              "type": "string"
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
      "instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
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
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudOpenSearchAppGroupsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudOpenSearchAppGroups), &result)
	return &result
}
