package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudApiGatewayBackend = `{
  "block": {
    "attributes": {
      "backend_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "backend_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_event_bridge_service_linked_role": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudApiGatewayBackendSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudApiGatewayBackend), &result)
	return &result
}
