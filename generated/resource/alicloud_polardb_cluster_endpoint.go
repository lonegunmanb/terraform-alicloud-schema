package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPolardbClusterEndpoint = `{
  "block": {
    "attributes": {
      "auto_add_new_nodes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connection_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_endpoint_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_endpoint_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_config": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "endpoint_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "net_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nodes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "read_write_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_auto_rotate": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_certificate_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ssl_connection_string": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ssl_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_expire_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudPolardbClusterEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPolardbClusterEndpoint), &result)
	return &result
}
