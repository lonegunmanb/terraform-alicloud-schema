package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudConfigCompliancePacks = `{
  "block": {
    "attributes": {
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "packs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "account_id": "string",
              "compliance_pack_id": "string",
              "compliance_pack_name": "string",
              "compliance_pack_template_id": "string",
              "config_rules": [
                "list",
                [
                  "object",
                  {
                    "config_rule_id": "string",
                    "config_rule_parameters": [
                      "list",
                      [
                        "object",
                        {
                          "parameter_name": "string",
                          "parameter_value": "string",
                          "required": "bool"
                        }
                      ]
                    ],
                    "managed_rule_identifier": "string"
                  }
                ]
              ],
              "description": "string",
              "id": "string",
              "risk_level": "number",
              "status": "string"
            }
          ]
        ]
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

func AlicloudConfigCompliancePacksSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudConfigCompliancePacks), &result)
	return &result
}
