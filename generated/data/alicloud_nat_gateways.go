package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudNatGateways = `{
  "block": {
    "attributes": {
      "dry_run": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "gateways": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "business_status": "string",
              "deletion_protection": "bool",
              "description": "string",
              "ecs_metric_enabled": "bool",
              "expired_time": "string",
              "forward_table_ids": [
                "list",
                "string"
              ],
              "id": "string",
              "internet_charge_type": "string",
              "ip_lists": [
                "list",
                "string"
              ],
              "name": "string",
              "nat_gateway_id": "string",
              "nat_gateway_name": "string",
              "nat_type": "string",
              "network_type": "string",
              "payment_type": "string",
              "resource_group_id": "string",
              "snat_table_ids": [
                "list",
                "string"
              ],
              "spec": "string",
              "specification": "string",
              "status": "string",
              "tags": [
                "map",
                "string"
              ],
              "vpc_id": "string",
              "vswitch_id": "string"
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
      "nat_gateway_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nat_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "page_number": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "page_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "payment_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "specification": {
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
      "total_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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

func AlicloudNatGatewaysSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudNatGateways), &result)
	return &result
}
