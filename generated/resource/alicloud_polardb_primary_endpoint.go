package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPolardbPrimaryEndpoint = `{
  "block": {
    "attributes": {
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
      "port": {
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

func AlicloudPolardbPrimaryEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPolardbPrimaryEndpoint), &result)
	return &result
}
