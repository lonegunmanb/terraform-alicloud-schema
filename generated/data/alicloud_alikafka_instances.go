package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlikafkaInstances = `{
  "block": {
    "attributes": {
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
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allowed_list": [
                "list",
                [
                  "object",
                  {
                    "deploy_type": "string",
                    "internet_list": [
                      "list",
                      [
                        "object",
                        {
                          "allowed_ip_list": [
                            "list",
                            "string"
                          ],
                          "port_range": "string"
                        }
                      ]
                    ],
                    "vpc_list": [
                      "list",
                      [
                        "object",
                        {
                          "allowed_ip_list": [
                            "list",
                            "string"
                          ],
                          "port_range": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "config": "string",
              "create_time": "string",
              "deploy_type": "number",
              "disk_size": "number",
              "disk_type": "number",
              "domain_endpoint": "string",
              "eip_max": "number",
              "end_point": "string",
              "expired_time": "number",
              "id": "string",
              "io_max": "number",
              "msg_retain": "number",
              "name": "string",
              "paid_type": "string",
              "partition_num": "number",
              "sasl_domain_endpoint": "string",
              "security_group": "string",
              "service_status": "number",
              "service_version": "string",
              "spec_type": "string",
              "ssl_domain_endpoint": "string",
              "ssl_end_point": "string",
              "tags": [
                "map",
                "string"
              ],
              "topic_quota": "number",
              "upgrade_service_detail_info": [
                "list",
                [
                  "object",
                  {
                    "current2_open_source_version": "string"
                  }
                ]
              ],
              "vpc_id": "string",
              "vswitch_id": "string",
              "zone_id": "string"
            }
          ]
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudAlikafkaInstancesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlikafkaInstances), &result)
	return &result
}
