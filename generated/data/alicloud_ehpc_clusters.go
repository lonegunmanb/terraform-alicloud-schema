package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEhpcClusters = `{
  "block": {
    "attributes": {
      "clusters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "account_type": "string",
              "application": [
                "list",
                [
                  "object",
                  {
                    "tag": "string"
                  }
                ]
              ],
              "client_version": "string",
              "cluster_id": "string",
              "cluster_name": "string",
              "compute_count": "number",
              "compute_instance_type": "string",
              "create_time": "string",
              "deploy_mode": "string",
              "description": "string",
              "ha_enable": "bool",
              "id": "string",
              "image_id": "string",
              "image_owner_alias": "string",
              "login_count": "number",
              "login_instance_type": "string",
              "manager_count": "number",
              "manager_instance_type": "string",
              "os_tag": "string",
              "post_install_script": [
                "list",
                [
                  "object",
                  {
                    "args": "string",
                    "url": "string"
                  }
                ]
              ],
              "remote_directory": "string",
              "scc_cluster_id": "string",
              "scheduler_type": "string",
              "security_group_id": "string",
              "status": "string",
              "volume_id": "string",
              "volume_mountpoint": "string",
              "volume_protocol": "string",
              "volume_type": "string",
              "vpc_id": "string",
              "vswitch_id": "string",
              "zone_id": "string"
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

func AlicloudEhpcClustersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEhpcClusters), &result)
	return &result
}
