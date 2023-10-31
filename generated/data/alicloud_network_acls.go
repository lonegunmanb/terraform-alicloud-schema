package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudNetworkAcls = `{
  "block": {
    "attributes": {
      "acls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "description": "string",
              "egress_acl_entries": [
                "list",
                [
                  "object",
                  {
                    "description": "string",
                    "destination_cidr_ip": "string",
                    "network_acl_entry_name": "string",
                    "policy": "string",
                    "port": "string",
                    "protocol": "string"
                  }
                ]
              ],
              "id": "string",
              "ingress_acl_entries": [
                "list",
                [
                  "object",
                  {
                    "description": "string",
                    "network_acl_entry_name": "string",
                    "policy": "string",
                    "port": "string",
                    "protocol": "string",
                    "source_cidr_ip": "string"
                  }
                ]
              ],
              "network_acl_id": "string",
              "network_acl_name": "string",
              "resources": [
                "list",
                [
                  "object",
                  {
                    "resource_id": "string",
                    "resource_type": "string",
                    "status": "string"
                  }
                ]
              ],
              "status": "string",
              "vpc_id": "string"
            }
          ]
        ]
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
      "name_regex": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "network_acl_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudNetworkAclsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudNetworkAcls), &result)
	return &result
}
