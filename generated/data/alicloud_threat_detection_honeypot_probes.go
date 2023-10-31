package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudThreatDetectionHoneypotProbes = `{
  "block": {
    "attributes": {
      "display_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "output_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "probe_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "probe_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "probes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arp": "bool",
              "control_node_id": "string",
              "display_name": "string",
              "honeypot_bind_list": [
                "list",
                [
                  "object",
                  {
                    "bind_port_list": [
                      "list",
                      [
                        "object",
                        {
                          "bind_port": "bool",
                          "end_port": "number",
                          "fixed": "bool",
                          "start_port": "number",
                          "target_port": "number"
                        }
                      ]
                    ],
                    "honeypot_id": "string"
                  }
                ]
              ],
              "honeypot_probe_id": "string",
              "id": "string",
              "ping": "bool",
              "probe_type": "string",
              "service_ip_list": [
                "list",
                "string"
              ],
              "status": "string",
              "uuid": "string",
              "vpc_id": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudThreatDetectionHoneypotProbesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudThreatDetectionHoneypotProbes), &result)
	return &result
}
