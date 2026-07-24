package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudApigServices = `{
  "block": {
    "attributes": {
      "gateway_id": {
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
      "services": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "addresses": [
                "list",
                "string"
              ],
              "create_timestamp": "number",
              "dns_servers": [
                "list",
                "string"
              ],
              "express_type": "string",
              "gateway_id": "string",
              "health_check_config": [
                "list",
                [
                  "object",
                  {
                    "enable": "bool",
                    "expected_statuses": [
                      "list",
                      "string"
                    ],
                    "healthy_threshold": "number",
                    "http_host": "string",
                    "http_path": "string",
                    "interval": "number",
                    "protocol": "string",
                    "timeout": "number",
                    "unhealthy_threshold": "number"
                  }
                ]
              ],
              "health_status": "string",
              "healthy_panic_threshold": "number",
              "id": "string",
              "namespace": "string",
              "outlier_detection_config": [
                "list",
                [
                  "object",
                  {
                    "base_ejection_time": "number",
                    "enable": "bool",
                    "failure_percentage_minimum_hosts": "number",
                    "failure_percentage_threshold": "number",
                    "interval": "number"
                  }
                ]
              ],
              "outlier_endpoints": [
                "list",
                "string"
              ],
              "ports": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "port": "number",
                    "protocol": "string"
                  }
                ]
              ],
              "protocol": "string",
              "qualifier": "string",
              "resource_group_id": "string",
              "runtime_detail_error_code": "string",
              "runtime_detail_status": "string",
              "service_id": "string",
              "service_name": "string",
              "source_type": "string",
              "unhealthy_endpoints": [
                "list",
                "string"
              ],
              "update_timestamp": "number"
            }
          ]
        ]
      },
      "source_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudApigServicesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudApigServices), &result)
	return &result
}
