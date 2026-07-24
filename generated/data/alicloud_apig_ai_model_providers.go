package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudApigAiModelProviders = `{
  "block": {
    "attributes": {
      "gateway_id": {
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "providers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bound_services": [
                "list",
                [
                  "object",
                  {
                    "express_type": "string",
                    "group_name": "string",
                    "name": "string",
                    "namespace": "string",
                    "pai_workspace_id": "string",
                    "pai_workspace_name": "string",
                    "qualifier": "string",
                    "service_id": "string",
                    "source_type": "string",
                    "status": "string"
                  }
                ]
              ],
              "display_name": "string",
              "gateway_id": "string",
              "id": "string",
              "model_cards": [
                "list",
                [
                  "object",
                  {
                    "gateway_id": "string",
                    "model_card_id": "string",
                    "model_name": "string",
                    "model_provider": "string",
                    "source": "string",
                    "update_time": "string"
                  }
                ]
              ],
              "model_count": "number",
              "model_provider": "string",
              "model_provider_id": "string",
              "source": "string",
              "update_time": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudApigAiModelProvidersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudApigAiModelProviders), &result)
	return &result
}
