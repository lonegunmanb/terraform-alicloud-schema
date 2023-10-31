package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDdoscooPort = `{
  "block": {
    "attributes": {
      "backend_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "frontend_port": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "frontend_protocol": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "real_servers": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDdoscooPortSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDdoscooPort), &result)
	return &result
}
