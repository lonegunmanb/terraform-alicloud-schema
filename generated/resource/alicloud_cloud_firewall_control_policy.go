package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudFirewallControlPolicy = `{
  "block": {
    "attributes": {
      "acl_action": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "acl_uuid": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "application_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dest_port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dest_port_group": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dest_port_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "direction": {
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
      "ip_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lang": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "proto": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "release": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_ip": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudFirewallControlPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudFirewallControlPolicy), &result)
	return &result
}
