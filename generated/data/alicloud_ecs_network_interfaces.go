package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEcsNetworkInterfaces = `{
  "block": {
    "attributes": {
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
      "instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "interfaces": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "associated_public_ip": [
                "list",
                [
                  "object",
                  {
                    "public_ip_address": "string"
                  }
                ]
              ],
              "creation_time": "string",
              "description": "string",
              "id": "string",
              "instance_id": "string",
              "mac": "string",
              "name": "string",
              "network_interface_id": "string",
              "network_interface_name": "string",
              "network_interface_traffic_mode": "string",
              "owner_id": "string",
              "primary_ip_address": "string",
              "private_ip": "string",
              "private_ip_addresses": [
                "list",
                "string"
              ],
              "private_ips": [
                "list",
                "string"
              ],
              "queue_number": "number",
              "resource_group_id": "string",
              "security_group_ids": [
                "list",
                "string"
              ],
              "security_groups": [
                "list",
                "string"
              ],
              "service_id": "number",
              "service_managed": "bool",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "type": "string",
              "vpc_id": "string",
              "vswitch_id": "string",
              "zone_id": "string"
            }
          ]
        ]
      },
      "name": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "network_interface_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "primary_ip_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_ip": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_managed": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vswitch_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudEcsNetworkInterfacesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEcsNetworkInterfaces), &result)
	return &result
}
