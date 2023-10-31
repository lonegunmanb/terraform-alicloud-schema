package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudDataWorksFolder = `{
  "block": {
    "attributes": {
      "folder_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "folder_path": {
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
      "project_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudDataWorksFolderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudDataWorksFolder), &result)
	return &result
}
