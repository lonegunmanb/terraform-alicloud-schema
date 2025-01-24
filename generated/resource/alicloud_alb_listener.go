package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const alicloudAlbListener = `{
  "block": {
    "attributes": {
      "access_log_record_customized_headers_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ca_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "dry_run": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "gzip_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "http2_enabled": {
        "computed": true,
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
      "idle_timeout": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "listener_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "listener_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "listener_protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "load_balancer_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "request_timeout": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "security_policy_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
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
      }
    },
    "block_types": {
      "access_log_tracing_config": {
        "block": {
          "attributes": {
            "tracing_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "tracing_sample": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "tracing_type": {
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
      "acl_config": {
        "block": {
          "attributes": {
            "acl_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "acl_relations": {
              "block": {
                "attributes": {
                  "acl_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
      },
      "ca_certificates": {
        "block": {
          "attributes": {
            "certificate_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "certificates": {
        "block": {
          "attributes": {
            "certificate_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "default_actions": {
        "block": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "forward_group_config": {
              "block": {
                "block_types": {
                  "server_group_tuples": {
                    "block": {
                      "attributes": {
                        "server_group_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "min_items": 1,
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "quic_config": {
        "block": {
          "attributes": {
            "quic_listener_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "quic_upgrade_enabled": {
              "computed": true,
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
      },
      "x_forwarded_for_config": {
        "block": {
          "attributes": {
            "x_forwarded_for_client_cert_client_verify_alias": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "x_forwarded_for_client_cert_client_verify_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "x_forwarded_for_client_cert_finger_print_alias": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "x_forwarded_for_client_cert_finger_print_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "x_forwarded_for_client_cert_issuer_dn_alias": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "x_forwarded_for_client_cert_issuer_dn_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "x_forwarded_for_client_cert_subject_dn_alias": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "x_forwarded_for_client_cert_subject_dn_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "x_forwarded_for_client_source_ips_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "x_forwarded_for_client_source_ips_trusted": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "x_forwarded_for_client_src_port_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "x_forwarded_for_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "x_forwarded_for_host_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "x_forwarded_for_processing_mode": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "x_forwarded_for_proto_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "x_forwarded_for_slb_id_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "x_forwarded_for_slb_port_enabled": {
              "computed": true,
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
      "xforwarded_for_config": {
        "block": {
          "attributes": {
            "xforwardedforclientcert_issuerdnalias": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "xforwardedforclientcert_issuerdnenabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "xforwardedforclientcertclientverifyalias": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "xforwardedforclientcertclientverifyenabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "xforwardedforclientcertfingerprintalias": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "xforwardedforclientcertfingerprintenabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "xforwardedforclientcertsubjectdnalias": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "xforwardedforclientcertsubjectdnenabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "xforwardedforclientsrcportenabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "xforwardedforenabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "xforwardedforprotoenabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "xforwardedforslbidenabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "xforwardedforslbportenabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AlicloudAlbListenerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(alicloudAlbListener), &result)
	return &result
}
