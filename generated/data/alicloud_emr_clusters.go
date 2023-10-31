package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudEmrClusters = `{
  "block": {
    "attributes": {
      "cluster_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_type_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "clusters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_info": [
                "list",
                [
                  "object",
                  {
                    "zk_links": [
                      "list",
                      [
                        "object",
                        {
                          "link": "string",
                          "port": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "auto_scaling_allowed": "bool",
              "auto_scaling_by_load_allowed": "bool",
              "auto_scaling_enable": "bool",
              "auto_scaling_spot_with_limit_allowed": "bool",
              "bootstrap_action_list": [
                "list",
                [
                  "object",
                  {
                    "arg": "string",
                    "name": "string",
                    "path": "string"
                  }
                ]
              ],
              "bootstrap_failed": "bool",
              "cluster_id": "string",
              "cluster_name": "string",
              "create_resource": "string",
              "create_time": "string",
              "create_type": "string",
              "deposit_type": "string",
              "eas_enable": "bool",
              "expired_time": "string",
              "extra_info": "string",
              "has_uncompleted_order": "bool",
              "high_availability_enable": "bool",
              "host_group_list": [
                "list",
                [
                  "object",
                  {
                    "band_width": "string",
                    "charge_type": "string",
                    "cpu_core": "number",
                    "disk_capacity": "number",
                    "disk_count": "number",
                    "disk_type": "string",
                    "host_group_change_type": "string",
                    "host_group_id": "string",
                    "host_group_name": "string",
                    "host_group_type": "string",
                    "instance_type": "string",
                    "memory_capacity": "number",
                    "node_count": "number",
                    "nodes": [
                      "list",
                      [
                        "object",
                        {
                          "create_time": "string",
                          "disk_infos": [
                            "list",
                            [
                              "object",
                              {
                                "device": "string",
                                "disk_id": "string",
                                "disk_name": "string",
                                "size": "number",
                                "type": "string"
                              }
                            ]
                          ],
                          "emr_expired_time": "string",
                          "expired_time": "string",
                          "inner_ip": "string",
                          "instance_id": "string",
                          "pub_ip": "string",
                          "status": "string",
                          "support_ipv6": "bool",
                          "zone_id": "string"
                        }
                      ]
                    ],
                    "period": "string"
                  }
                ]
              ],
              "host_pool_info": [
                "list",
                [
                  "object",
                  {
                    "hp_biz_id": "string",
                    "hp_name": "string"
                  }
                ]
              ],
              "id": "string",
              "image_id": "string",
              "local_meta_db": "bool",
              "machine_type": "string",
              "meta_store_type": "string",
              "net_type": "string",
              "payment_type": "string",
              "period": "number",
              "relate_cluster_info": [
                "list",
                [
                  "object",
                  {
                    "cluster_id": "string",
                    "cluster_name": "string",
                    "cluster_type": "string",
                    "status": "string"
                  }
                ]
              ],
              "resize_disk_enable": "bool",
              "running_time": "number",
              "security_group_id": "string",
              "security_group_name": "string",
              "software_info": [
                "list",
                [
                  "object",
                  {
                    "cluster_type": "string",
                    "emr_ver": "string",
                    "softwares": [
                      "list",
                      [
                        "object",
                        {
                          "display_name": "string",
                          "name": "string",
                          "only_display": "bool",
                          "start_tpe": "number",
                          "version": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "start_time": "string",
              "status": "string",
              "stop_time": "string",
              "tags": [
                "map",
                "string"
              ],
              "type": "string",
              "user_defined_emr_ecs_role": "string",
              "user_id": "string",
              "vpc_id": "string",
              "vswitch_id": "string",
              "zone_id": "string"
            }
          ]
        ]
      },
      "create_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "deposit_type": {
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
      "is_desc": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "machine_type": {
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
      "resource_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
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

func AlicloudEmrClustersSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudEmrClusters), &result)
	return &result
}
