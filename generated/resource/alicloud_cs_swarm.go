package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCsSwarm = `{
  "block": {
    "attributes": {
      "agent_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cidr_block": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "disk_category": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disk_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "is_outdated": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "need_slb": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "node_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "nodes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "eip": "string",
              "id": "string",
              "name": "string",
              "private_ip": "string",
              "status": "string"
            }
          ]
        ]
      },
      "password": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "release_eip": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "security_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "size": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "slb_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCsSwarmSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCsSwarm), &result)
	return &result
}
