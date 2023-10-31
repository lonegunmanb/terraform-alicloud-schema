package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudCloudStorageGatewayGateways = `{
  "block": {
    "attributes": {
      "gateways": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "activated_time": "string",
              "buy_url": "string",
              "category": "string",
              "create_time": "string",
              "description": "string",
              "ecs_instance_id": "string",
              "expire_status": "number",
              "expired_time": "string",
              "gateway_class": "string",
              "gateway_id": "string",
              "gateway_name": "string",
              "gateway_version": "string",
              "id": "string",
              "inner_ip": "string",
              "ip": "string",
              "is_release_after_expiration": "bool",
              "location": "string",
              "payment_type": "string",
              "public_network_bandwidth": "number",
              "renew_url": "string",
              "status": "string",
              "storage_bundle_id": "string",
              "task_id": "string",
              "type": "string",
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
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_bundle_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "total_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudCloudStorageGatewayGatewaysSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudCloudStorageGatewayGateways), &result)
	return &result
}
