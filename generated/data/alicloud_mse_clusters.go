package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudMseClusters = `{
  "block": {
    "attributes": {
      "cluster_alias_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "clusters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "acl_id": "string",
              "app_version": "string",
              "cluster_id": "string",
              "cluster_name": "string",
              "cluster_type": "string",
              "cpu": "number",
              "health_status": "string",
              "id": "string",
              "init_cost_time": "number",
              "instance_count": "number",
              "instance_id": "string",
              "instance_models": [
                "list",
                [
                  "object",
                  {
                    "health_status": "string",
                    "instance_type": "string",
                    "internet_ip": "string",
                    "ip": "string",
                    "pod_name": "string",
                    "role": "string",
                    "single_tunnel_vip": "string",
                    "vip": "string"
                  }
                ]
              ],
              "internet_address": "string",
              "internet_domain": "string",
              "internet_port": "string",
              "intranet_address": "string",
              "intranet_domain": "string",
              "intranet_port": "string",
              "memory_capacity": "number",
              "pay_info": "string",
              "pub_network_flow": "string",
              "status": "string"
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
      "request_pars": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudMseClustersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudMseClusters), &result)
	return &result
}
