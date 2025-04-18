package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMongodbReplicaSetRole = `{
  "block": {
    "attributes": {
      "connection_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "connection_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_id": {
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
      "network_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replica_set_role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMongodbReplicaSetRoleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMongodbReplicaSetRole), &result)
	return &result
}
