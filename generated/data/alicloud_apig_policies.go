package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudApigPolicies = `{
  "block": {
    "attributes": {
      "attach_resource_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "attach_resource_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "environment_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gateway_id": {
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policies": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "attach_resource_ids": [
                "list",
                "string"
              ],
              "attach_resource_type": "string",
              "environment_id": "string",
              "gateway_id": "string",
              "id": "string",
              "policy_attachment_id": "string",
              "policy_class_id": "string",
              "policy_class_name": "string",
              "policy_config": "string",
              "policy_id": "string",
              "policy_name": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudApigPoliciesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudApigPolicies), &result)
	return &result
}
