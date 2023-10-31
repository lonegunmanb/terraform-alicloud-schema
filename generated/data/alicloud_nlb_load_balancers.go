package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudNlbLoadBalancers = `{
  "block": {
    "attributes": {
      "address_ip_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "address_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "balancers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address_ip_version": "string",
              "address_type": "string",
              "bandwidth_package_id": "string",
              "create_time": "string",
              "cross_zone_enabled": "bool",
              "dns_name": "string",
              "id": "string",
              "ipv6_address_type": "string",
              "load_balancer_business_status": "string",
              "load_balancer_id": "string",
              "load_balancer_name": "string",
              "load_balancer_type": "string",
              "operation_locks": [
                "list",
                [
                  "object",
                  {
                    "lock_reason": "string",
                    "lock_type": "string"
                  }
                ]
              ],
              "resource_group_id": "string",
              "security_group_ids": [
                "list",
                "string"
              ],
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "vpc_id": "string",
              "zone_mappings": [
                "list",
                [
                  "object",
                  {
                    "allocation_id": "string",
                    "eni_id": "string",
                    "ipv6_address": "string",
                    "private_ipv4_address": "string",
                    "public_ipv4_address": "string",
                    "vswitch_id": "string",
                    "zone_id": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "dns_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "ipv6_address_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancer_business_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancer_names": {
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
      "vpc_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudNlbLoadBalancersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudNlbLoadBalancers), &result)
	return &result
}
