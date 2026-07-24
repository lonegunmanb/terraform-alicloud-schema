package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudApigGateways = `{
  "block": {
    "attributes": {
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "gateway_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gateway_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gateways": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_from": "string",
              "create_time": "number",
              "environments": [
                "list",
                [
                  "object",
                  {
                    "alias": "string",
                    "environment_id": "string",
                    "name": "string"
                  }
                ]
              ],
              "expire_time": "number",
              "gateway_edition": "string",
              "gateway_id": "string",
              "gateway_name": "string",
              "gateway_type": "string",
              "id": "string",
              "load_balancers": [
                "list",
                [
                  "object",
                  {
                    "address": "string",
                    "address_ip_version": "string",
                    "address_type": "string",
                    "gateway_default": "bool",
                    "ipv4_addresses": [
                      "list",
                      "string"
                    ],
                    "ipv6_addresses": [
                      "list",
                      "string"
                    ],
                    "load_balancer_id": "string",
                    "mode": "string",
                    "ports": [
                      "list",
                      [
                        "object",
                        {
                          "port": "number",
                          "protocol": "string"
                        }
                      ]
                    ],
                    "status": "string",
                    "type": "string"
                  }
                ]
              ],
              "payment_type": "string",
              "resource_group_id": "string",
              "security_group": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "security_group_id": "string"
                  }
                ]
              ],
              "spec": "string",
              "status": "string",
              "sub_domain_infos": [
                "list",
                [
                  "object",
                  {
                    "domain_id": "string",
                    "name": "string",
                    "network_type": "string",
                    "protocol": "string"
                  }
                ]
              ],
              "tags": [
                "map",
                "string"
              ],
              "target_version": "string",
              "update_time": "number",
              "version": "string",
              "vpc": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "vpc_id": "string"
                  }
                ]
              ],
              "vswitch": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "vswitch_id": "string"
                  }
                ]
              ],
              "zones": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "vswitch_id": "string",
                    "zone_id": "string"
                  }
                ]
              ]
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudApigGatewaysSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudApigGateways), &result)
	return &result
}
