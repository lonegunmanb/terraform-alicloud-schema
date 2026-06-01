package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudPolardbAicluster = `{
  "block": {
    "attributes": {
      "auto_renew": {
        "description": "Whether to enable auto-renewal.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_use_coupon": {
        "description": "Whether to automatically use coupons. Default: ` + "`" + `true` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "db_cluster_description": {
        "description": "The description of the AI DB cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_cluster_id": {
        "description": "The ID of the associated DB cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_node_class": {
        "description": "The DB node class of the AI cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "extension": {
        "description": "The extension type. Valid values: ` + "`" + `maas` + "`" + `, ` + "`" + `custom` + "`" + `.",
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
      "inference_engine": {
        "description": "The inference engine. Valid values: ` + "`" + `sglang` + "`" + `, ` + "`" + `vllm` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kube_type": {
        "description": "The type of the Kubernetes cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_name": {
        "description": "The model name. Example: ` + "`" + `Qwen3.5-9B` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_type": {
        "computed": true,
        "description": "The model type. Example: ` + "`" + `public` + "`" + `.",
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
        "description": "The subscription period.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "promotion_code": {
        "description": "The promotion code. If not specified, the default coupon is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_id": {
        "description": "The region ID of the AI cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_group_id": {
        "description": "The security group ID.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the AI cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "used_time": {
        "description": "The subscription duration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_id": {
        "description": "The VPC ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vswitch_id": {
        "description": "The vSwitch ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zone_id": {
        "description": "The zone ID of the AI cluster.",
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

func AlicloudPolardbAiclusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudPolardbAicluster), &result)
	return &result
}
