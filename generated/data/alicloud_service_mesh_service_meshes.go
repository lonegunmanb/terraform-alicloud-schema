package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudServiceMeshServiceMeshes = `{
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
      "meshes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "clusters": [
                "list",
                "string"
              ],
              "create_time": "string",
              "edition": "string",
              "endpoints": [
                "list",
                [
                  "object",
                  {
                    "intranet_api_server_endpoint": "string",
                    "intranet_pilot_endpoint": "string",
                    "public_api_server_endpoint": "string",
                    "public_pilot_endpoint": "string"
                  }
                ]
              ],
              "error_message": "string",
              "id": "string",
              "istio_operator_version": "string",
              "kube_config": "string",
              "load_balancer": [
                "list",
                [
                  "object",
                  {
                    "api_server_loadbalancer_id": "string",
                    "api_server_public_eip": "bool",
                    "pilot_public_eip": "bool",
                    "pilot_public_loadbalancer_id": "string"
                  }
                ]
              ],
              "mesh_config": [
                "list",
                [
                  "object",
                  {
                    "access_log": [
                      "list",
                      [
                        "object",
                        {
                          "enabled": "bool",
                          "project": "string"
                        }
                      ]
                    ],
                    "audit": [
                      "list",
                      [
                        "object",
                        {
                          "enabled": "bool",
                          "project": "string"
                        }
                      ]
                    ],
                    "control_plane_log": [
                      "list",
                      [
                        "object",
                        {
                          "enabled": "bool",
                          "project": "string"
                        }
                      ]
                    ],
                    "customized_zipkin": "bool",
                    "enable_locality_lb": "bool",
                    "include_ip_ranges": "string",
                    "kiali": [
                      "list",
                      [
                        "object",
                        {
                          "enabled": "bool",
                          "url": "string"
                        }
                      ]
                    ],
                    "opa": [
                      "list",
                      [
                        "object",
                        {
                          "enabled": "bool",
                          "limit_cpu": "string",
                          "limit_memory": "string",
                          "log_level": "string",
                          "request_cpu": "string",
                          "request_memory": "string"
                        }
                      ]
                    ],
                    "outbound_traffic_policy": "string",
                    "pilot": [
                      "list",
                      [
                        "object",
                        {
                          "http10_enabled": "bool",
                          "trace_sampling": "number"
                        }
                      ]
                    ],
                    "prometheus": [
                      "list",
                      [
                        "object",
                        {
                          "external_url": "string",
                          "use_external": "bool"
                        }
                      ]
                    ],
                    "proxy": [
                      "list",
                      [
                        "object",
                        {
                          "cluster_domain": "string",
                          "limit_cpu": "string",
                          "limit_memory": "string",
                          "request_cpu": "string",
                          "request_memory": "string"
                        }
                      ]
                    ],
                    "sidecar_injector": [
                      "list",
                      [
                        "object",
                        {
                          "auto_injection_policy_enabled": "bool",
                          "enable_namespaces_by_default": "bool",
                          "init_cni_configuration": [
                            "list",
                            [
                              "object",
                              {
                                "enabled": "bool",
                                "exclude_namespaces": "string"
                              }
                            ]
                          ],
                          "limit_cpu": "string",
                          "limit_memory": "string",
                          "request_cpu": "string",
                          "request_memory": "string",
                          "sidecar_injector_webhook_as_yaml": "string"
                        }
                      ]
                    ],
                    "telemetry": "bool",
                    "tracing": "bool"
                  }
                ]
              ],
              "network": [
                "list",
                [
                  "object",
                  {
                    "security_group_id": "string",
                    "vpc_id": "string",
                    "vswitche_list": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "service_mesh_id": "string",
              "service_mesh_name": "string",
              "sidecar_version": "string",
              "status": "string",
              "version": "string"
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

func AlicloudServiceMeshServiceMeshesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudServiceMeshServiceMeshes), &result)
	return &result
}
