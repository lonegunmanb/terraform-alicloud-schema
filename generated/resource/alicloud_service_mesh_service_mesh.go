package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudServiceMeshServiceMesh = `{
  "block": {
    "attributes": {
      "cluster_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "cluster_spec": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customized_prometheus": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "edition": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "force": {
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
      "kubeconfig": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "prometheus_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_mesh_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
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
      "version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "extra_configuration": {
        "block": {
          "attributes": {
            "cr_aggregation_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "load_balancer": {
        "block": {
          "attributes": {
            "api_server_loadbalancer_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "api_server_public_eip": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "pilot_public_eip": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "pilot_public_eip_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pilot_public_loadbalancer_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "mesh_config": {
        "block": {
          "attributes": {
            "customized_zipkin": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_locality_lb": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_ip_ranges": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "outbound_traffic_policy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "prometheus": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "external_url": "string",
                    "use_external": "bool"
                  }
                ]
              ]
            },
            "telemetry": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tracing": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "access_log": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "gateway_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "gateway_lifecycle": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "project": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "sidecar_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "sidecar_lifecycle": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "audit": {
              "block": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "project": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "control_plane_log": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "log_ttl_in_day": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "project": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "kiali": {
              "block": {
                "attributes": {
                  "aggregated_kiali_address": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "auth_strategy": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "custom_prometheus_url": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "distributed_kiali_access_tokens": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "distributed_kiali_addresses": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "integrate_clb": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "kiali_arms_auth_tokens": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kiali_service_annotations": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "url": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "use_populated_arms_prometheus": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "block_types": {
                  "open_id_config": {
                    "block": {
                      "attributes": {
                        "client_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "client_secret": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "issuer_uri": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "scopes": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "ram_oauth_config": {
                    "block": {
                      "attributes": {
                        "redirect_uris": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "server_config": {
                    "block": {
                      "attributes": {
                        "web_fqdn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "web_port": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "web_root": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "web_schema": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "opa": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "limit_cpu": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "limit_memory": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_level": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "request_cpu": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "request_memory": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "pilot": {
              "block": {
                "attributes": {
                  "http10_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "trace_sampling": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "proxy": {
              "block": {
                "attributes": {
                  "cluster_domain": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "limit_cpu": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "limit_memory": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "request_cpu": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "request_memory": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "sidecar_injector": {
              "block": {
                "attributes": {
                  "auto_injection_policy_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_namespaces_by_default": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "limit_cpu": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "limit_memory": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "request_cpu": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "request_memory": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "sidecar_injector_webhook_as_yaml": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "block_types": {
                  "init_cni_configuration": {
                    "block": {
                      "attributes": {
                        "enabled": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "exclude_namespaces": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network": {
        "block": {
          "attributes": {
            "security_group_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "vpc_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vswitche_list": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
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
            },
            "update": {
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

func AlicloudServiceMeshServiceMeshSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudServiceMeshServiceMesh), &result)
	return &result
}
