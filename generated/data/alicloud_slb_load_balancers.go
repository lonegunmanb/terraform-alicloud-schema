package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudSlbLoadBalancers = `{
  "block": {
    "attributes": {
      "address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
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
              "address": "string",
              "address_ip_version": "string",
              "address_type": "string",
              "auto_release_time": "number",
              "backend_servers": [
                "list",
                [
                  "object",
                  {
                    "description": "string",
                    "server_id": "string",
                    "type": "string",
                    "weight": "number"
                  }
                ]
              ],
              "bandwidth": "number",
              "create_time_stamp": "number",
              "delete_protection": "string",
              "end_time": "string",
              "end_time_stamp": "number",
              "id": "string",
              "internet_charge_type": "string",
              "listener_ports_and_protocal": [
                "list",
                [
                  "object",
                  {
                    "listener_port": "number",
                    "listener_protocal": "string"
                  }
                ]
              ],
              "listener_ports_and_protocol": [
                "list",
                [
                  "object",
                  {
                    "description": "string",
                    "forward_port": "number",
                    "listener_forward": "string",
                    "listener_port": "number",
                    "listener_protocol": "string"
                  }
                ]
              ],
              "load_balancer_id": "string",
              "load_balancer_name": "string",
              "load_balancer_spec": "string",
              "master_zone_id": "string",
              "modification_protection_reason": "string",
              "modification_protection_status": "string",
              "network_type": "string",
              "payment_type": "string",
              "region_id_alias": "string",
              "renewal_cyc_unit": "string",
              "renewal_duration": "number",
              "renewal_status": "string",
              "resource_group_id": "string",
              "slave_zone_id": "string",
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
      "internet_charge_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancer_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_availability_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_zone_id": {
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
      "network_type": {
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
      "server_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_intranet_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "slave_availability_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "slave_zone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "slbs": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address": "string",
              "creation_time": "string",
              "id": "string",
              "internet": "bool",
              "master_availability_zone": "string",
              "name": "string",
              "network_type": "string",
              "region_id": "string",
              "slave_availability_zone": "string",
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

func AlicloudSlbLoadBalancersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudSlbLoadBalancers), &result)
	return &result
}
