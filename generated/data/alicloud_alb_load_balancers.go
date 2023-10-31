package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlbLoadBalancers = `{
  "block": {
    "attributes": {
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
              "access_log_config": [
                "list",
                [
                  "object",
                  {
                    "log_project": "string",
                    "log_store": "string"
                  }
                ]
              ],
              "address_allocated_mode": "string",
              "address_type": "string",
              "bandwidth_package_id": "string",
              "create_time": "string",
              "deletion_protection_config": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool",
                    "enabled_time": "string"
                  }
                ]
              ],
              "dns_name": "string",
              "id": "string",
              "load_balancer_billing_config": [
                "list",
                [
                  "object",
                  {
                    "pay_type": "string"
                  }
                ]
              ],
              "load_balancer_business_status": "string",
              "load_balancer_bussiness_status": "string",
              "load_balancer_edition": "string",
              "load_balancer_id": "string",
              "load_balancer_name": "string",
              "load_balancer_operation_locks": [
                "list",
                [
                  "object",
                  {
                    "lock_reason": "string",
                    "lock_type": "string"
                  }
                ]
              ],
              "modification_protection_config": [
                "list",
                [
                  "object",
                  {
                    "reason": "string",
                    "status": "string"
                  }
                ]
              ],
              "resource_group_id": "string",
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
                    "load_balancer_addresses": [
                      "list",
                      [
                        "object",
                        {
                          "address": "string"
                        }
                      ]
                    ],
                    "vswitch_id": "string",
                    "zone_id": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "load_balancer_business_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancer_bussiness_status": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancer_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "load_balancer_name": {
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
      "vpc_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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

func AlicloudAlbLoadBalancersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlbLoadBalancers), &result)
	return &result
}
