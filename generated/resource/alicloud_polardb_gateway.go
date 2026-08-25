package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPolardbGateway = `{
  "block": {
    "attributes": {
      "auto_renew": {
        "description": "Whether to enable automatic renewal for a subscription gateway.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "create_time": {
        "computed": true,
        "description": "The time when the PolarDB gateway was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "current_version": {
        "computed": true,
        "description": "The current gateway version.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_class": {
        "description": "The specifications of the PolarDB gateway.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_type": {
        "description": "The database engine. Valid values: ` + "`" + `MySQL` + "`" + `, ` + "`" + `PostgreSQL` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the PolarDB gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address": "string",
              "endpoint_id": "string",
              "gateway_id": "string",
              "network_type": "string",
              "port": "number",
              "tunnel_id": "string",
              "vpc_id": "string"
            }
          ]
        ]
      },
      "expire_time": {
        "computed": true,
        "description": "The expiration time of the PolarDB gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "expired": {
        "computed": true,
        "description": "Indicates whether the PolarDB gateway has expired.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "latest_version": {
        "computed": true,
        "description": "The latest available gateway version.",
        "description_kind": "plain",
        "type": "string"
      },
      "modify_time": {
        "computed": true,
        "description": "The time when the PolarDB gateway was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "pay_type": {
        "description": "The billing method. Valid values: ` + "`" + `Postpaid` + "`" + `, ` + "`" + `Prepaid` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "period": {
        "description": "The unit of the subscription duration. Valid values: ` + "`" + `Month` + "`" + `, ` + "`" + `Year` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_id": {
        "computed": true,
        "description": "The region ID of the PolarDB gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "running_version": {
        "computed": true,
        "description": "The running gateway version.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_id": {
        "description": "The ID of the security group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_ip_arrays": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ip_list": "string",
              "name": "string"
            }
          ]
        ]
      },
      "status": {
        "computed": true,
        "description": "The status of the PolarDB gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "used_time": {
        "description": "The subscription duration. Valid values are 1 to 9 for Month and 1 to 3 for Year.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "vpc_id": {
        "description": "The ID of the VPC.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vswitch_id": {
        "description": "The ID of the vSwitch.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zone_id": {
        "description": "The zone ID of the PolarDB gateway.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudPolardbGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPolardbGateway), &result)
	return &result
}
