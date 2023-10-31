package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSslVpnServer = `{
  "block": {
    "attributes": {
      "cipher": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "client_ip_pool": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "compress": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "connections": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internet_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "local_subnet": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_connections": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "protocol": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpn_gateway_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudSslVpnServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSslVpnServer), &result)
	return &result
}
