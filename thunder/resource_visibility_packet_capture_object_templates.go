package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplates() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates`: Configure object packet capture templates for T2 counters\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesDelete,

		Schema: map[string]*schema.Schema{
			"aam_aaa_policy_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"aam_auth_captcha_inst_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"parse_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total JSON Response Parse Failure",
									},
									"json_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure JSON Response",
									},
									"attr_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Attibute Check Failure",
									},
									"timeout_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout",
									},
									"other_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Other Error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"parse_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total JSON Response Parse Failure",
									},
									"json_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure JSON Response",
									},
									"attr_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Attibute Check Failure",
									},
									"timeout_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout",
									},
									"other_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Other Error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"aam_auth_logon_http_ins_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"spn_krb_faiure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SPN Kerberos Failure",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"spn_krb_faiure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SPN Kerberos Failure",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"aam_auth_relay_form_inst_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"invalid_srv_rsp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid Server Response",
									},
									"post_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for POST Failed",
									},
									"invalid_cred": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid Credential",
									},
									"bad_req": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Request",
									},
									"not_fnd": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Not Found",
									},
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Internal Server Error",
									},
									"other_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Other Error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"invalid_srv_rsp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid Server Response",
									},
									"post_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for POST Failed",
									},
									"invalid_cred": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid Credential",
									},
									"bad_req": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Request",
									},
									"not_fnd": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Not Found",
									},
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Internal Server Error",
									},
									"other_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Other Error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"aam_auth_relay_hbase_inst_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"no_creds": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Credential",
									},
									"bad_req": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Request",
									},
									"unauth": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unauthorized",
									},
									"forbidden": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Forbidden",
									},
									"not_found": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Not Found",
									},
									"server_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Internal Server Error",
									},
									"unavailable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Service Unavailable",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"no_creds": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Credential",
									},
									"bad_req": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Request",
									},
									"unauth": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unauthorized",
									},
									"forbidden": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Forbidden",
									},
									"not_found": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Not Found",
									},
									"server_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Internal Server Error",
									},
									"unavailable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Service Unavailable",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"aam_auth_relay_ntlm_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Failure",
									},
									"buffer_alloc_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Buffer Allocation Failure",
									},
									"encoding_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Encoding Failure",
									},
									"insert_header_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Insert Header Failure",
									},
									"parse_header_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse Header Failure",
									},
									"internal_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Internal Error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Failure",
									},
									"buffer_alloc_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Buffer Allocation Failure",
									},
									"encoding_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Encoding Failure",
									},
									"insert_header_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Insert Header Failure",
									},
									"parse_header_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse Header Failure",
									},
									"internal_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Internal Error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"aam_auth_relay_ws_fed_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Failure",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Failure",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"aam_auth_saml_id_prov_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"md_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Metadata Update Fail Count",
									},
									"acs_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ACS Fail Count",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"md_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Metadata Update Fail Count",
									},
									"acs_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ACS Fail Count",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"aam_auth_saml_service_prov_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acs_authz_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SAML Single-Sign-On Authorization Fail",
									},
									"acs_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SAML Single-Sign-On Error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"acs_authz_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SAML Single-Sign-On Authorization Fail",
									},
									"acs_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SAML Single-Sign-On Error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"aam_auth_server_ldap_inst_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"admin_bind_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Admin Bind Failure",
									},
									"bind_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User Bind Failure",
									},
									"search_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Search Failure",
									},
									"authorize_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Authorization Failure",
									},
									"timeout_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Timeout",
									},
									"other_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Other Error",
									},
									"ssl_session_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TLS/SSL Session Failure",
									},
									"pw_change_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Password change failure",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"admin_bind_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Admin Bind Failure",
									},
									"bind_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User Bind Failure",
									},
									"search_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Search Failure",
									},
									"authorize_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Authorization Failure",
									},
									"timeout_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Timeout",
									},
									"other_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Other Error",
									},
									"ssl_session_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TLS/SSL Session Failure",
									},
									"pw_change_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Password change failure",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"aam_auth_server_ocsp_inst_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"timeout": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Timeout",
									},
									"fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Handle OCSP response failed",
									},
									"stapling_timeout": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP Stapling Timeout",
									},
									"stapling_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Handle OCSP response failed",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"timeout": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Timeout",
									},
									"fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Handle OCSP response failed",
									},
									"stapling_timeout": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP Stapling Timeout",
									},
									"stapling_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Handle OCSP response failed",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"aam_auth_server_rad_inst_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authen_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Authentication Failure",
									},
									"authorize_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Authorization Failure",
									},
									"timeout_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Timeout",
									},
									"other_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Other Error",
									},
									"accounting_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Accounting Failure",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"authen_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Authentication Failure",
									},
									"authorize_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Authorization Failure",
									},
									"timeout_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Timeout",
									},
									"other_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Other Error",
									},
									"accounting_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Accounting Failure",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"aam_auth_server_win_inst_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"krb_timeout_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos Timeout",
									},
									"krb_other_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos Other Error",
									},
									"krb_pw_expiry": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos password expiry",
									},
									"krb_pw_change_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos password change failure",
									},
									"ntlm_proto_nego_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Protocol Negotiation Failure",
									},
									"ntlm_session_setup_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Session Setup Failure",
									},
									"ntlm_prepare_req_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Prepare Request Error",
									},
									"ntlm_auth_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Authentication Failure",
									},
									"ntlm_timeout_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Timeout",
									},
									"ntlm_other_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Other Error",
									},
									"krb_validate_kdc_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos KDC Validation Failure",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"krb_timeout_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos Timeout",
									},
									"krb_other_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos Other Error",
									},
									"krb_pw_expiry": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos password expiry",
									},
									"krb_pw_change_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos password change failure",
									},
									"ntlm_proto_nego_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Protocol Negotiation Failure",
									},
									"ntlm_session_setup_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Session Setup Failure",
									},
									"ntlm_prepare_req_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Prepare Request Error",
									},
									"ntlm_auth_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Authentication Failure",
									},
									"ntlm_timeout_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Timeout",
									},
									"ntlm_other_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Other Error",
									},
									"krb_validate_kdc_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos KDC Validation Failure",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"aam_auth_service_group_mem_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"curr_conn_overflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Current connection counter overflow count",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"curr_conn_overflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Current connection counter overflow count",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"aam_auth_service_group_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server_selection_fail_reset": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resets sent out for Service selection failure",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"server_selection_fail_reset": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resets sent out for Service selection failure",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"aam_jwt_authorization_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"jwt_authorize_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Authorize Failure",
									},
									"jwt_missing_token": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Missing Token",
									},
									"jwt_missing_claim": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Missing Claim",
									},
									"jwt_token_expired": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Token Expired",
									},
									"jwt_signature_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Signature Failure",
									},
									"jwt_other_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Other Error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"jwt_authorize_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Authorize Failure",
									},
									"jwt_missing_token": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Missing Token",
									},
									"jwt_missing_claim": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Missing Claim",
									},
									"jwt_token_expired": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Token Expired",
									},
									"jwt_signature_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Signature Failure",
									},
									"jwt_other_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for JWT Other Error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"cgnv6_dns64_vs_port_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"es_total_failure_actions": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total failure actions",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"es_total_failure_actions": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total failure actions",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"cgnv6_encap_domain_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inbound_addr_port_validation_failed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv4 Destination Address Port Validation Failed",
									},
									"inbound_rev_lookup_failed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv4 Reverse Route Lookup Failed",
									},
									"inbound_dest_unreachable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv6 Destination Address Unreachable",
									},
									"outbound_addr_validation_failed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv6 Source Address Validation Failed",
									},
									"outbound_rev_lookup_failed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv6 Reverse Route Lookup Failed",
									},
									"outbound_dest_unreachable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv4 Destination Address Unreachable",
									},
									"packet_mtu_exceeded": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Exceeded MTU",
									},
									"interface_not_configured": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Interfaces not Configured Dropped",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"inbound_addr_port_validation_failed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv4 Destination Address Port Validation Failed",
									},
									"inbound_rev_lookup_failed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv4 Reverse Route Lookup Failed",
									},
									"inbound_dest_unreachable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv6 Destination Address Unreachable",
									},
									"outbound_addr_validation_failed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv6 Source Address Validation Failed",
									},
									"outbound_rev_lookup_failed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv6 Reverse Route Lookup Failed",
									},
									"outbound_dest_unreachable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv4 Destination Address Unreachable",
									},
									"packet_mtu_exceeded": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Exceeded MTU",
									},
									"interface_not_configured": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Interfaces not Configured Dropped",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"cgnv6_map_trans_domain_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inbound_addr_port_validation_failed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv4 Destination Address Port Validation Failed",
									},
									"inbound_rev_lookup_failed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv4 Reverse Route Lookup Failed",
									},
									"inbound_dest_unreachable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv6 Destination Address Unreachable",
									},
									"outbound_addr_validation_failed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv6 Source Address Validation Failed",
									},
									"outbound_rev_lookup_failed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv6 Reverse Route Lookup Failed",
									},
									"outbound_dest_unreachable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv4 Destination Address Unreachable",
									},
									"packet_mtu_exceeded": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Exceeded MTU",
									},
									"interface_not_configured": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Interfaces not Configured Dropped",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"inbound_addr_port_validation_failed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv4 Destination Address Port Validation Failed",
									},
									"inbound_rev_lookup_failed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv4 Reverse Route Lookup Failed",
									},
									"inbound_dest_unreachable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv6 Destination Address Unreachable",
									},
									"outbound_addr_validation_failed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv6 Source Address Validation Failed",
									},
									"outbound_rev_lookup_failed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv6 Reverse Route Lookup Failed",
									},
									"outbound_dest_unreachable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv4 Destination Address Unreachable",
									},
									"packet_mtu_exceeded": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Exceeded MTU",
									},
									"interface_not_configured": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Interfaces not Configured Dropped",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"cgnv6_serv_group_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server_selection_fail_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Service selection fail drop",
									},
									"server_selection_fail_reset": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Service selection fail reset",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"server_selection_fail_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Service selection fail drop",
									},
									"server_selection_fail_reset": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Service selection fail reset",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"dns_vport_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dnsrrl_total_dropped": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dns rrl drop",
									},
									"total_filter_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for query filter drop",
									},
									"total_max_query_len_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for query too long drop",
									},
									"rcode_notimpl_receive": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for response rcode type error receive",
									},
									"rcode_notimpl_response": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for rcode type error response",
									},
									"gslb_query_bad": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for gslb query bad",
									},
									"gslb_response_bad": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for gslb response bad",
									},
									"total_dns_filter_type_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters Total DNS Filter Type Drop",
									},
									"total_dns_filter_class_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters Total DNS Filter Class Drop",
									},
									"dns_filter_type_a_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type A Drop",
									},
									"dns_filter_type_aaaa_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type AAAA Drop",
									},
									"dns_filter_type_cname_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type CNAME Drop",
									},
									"dns_filter_type_mx_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type MX Drop",
									},
									"dns_filter_type_ns_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type NS Drop",
									},
									"dns_filter_type_srv_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type SRV Drop",
									},
									"dns_filter_type_ptr_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type PTR Drop",
									},
									"dns_filter_type_soa_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type SOA Drop",
									},
									"dns_filter_type_txt_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type TXT Drop",
									},
									"dns_filter_type_any_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type Any Drop",
									},
									"dns_filter_type_others_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type OTHERS Drop",
									},
									"dns_filter_class_internet_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class INTERNET Drop",
									},
									"dns_filter_class_chaos_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class CHAOS Drop",
									},
									"dns_filter_class_hesiod_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class HESIOD Drop",
									},
									"dns_filter_class_none_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class NONE Drop",
									},
									"dns_filter_class_any_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class ANY Drop",
									},
									"dns_filter_class_others_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class OTHER Drop",
									},
									"dns_rpz_action_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS RPZ Action Drop",
									},
									"dnsrrl_bad_fqdn": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS Response-Rate-Limiting Bad FQDN",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"dnsrrl_total_dropped": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dns rrl drop",
									},
									"total_filter_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for query filter drop",
									},
									"total_max_query_len_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for query too long drop",
									},
									"rcode_notimpl_receive": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for response rcode type error receive",
									},
									"rcode_notimpl_response": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for rcode type error response",
									},
									"gslb_query_bad": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for gslb query bad",
									},
									"gslb_response_bad": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for gslb response bad",
									},
									"total_dns_filter_type_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters Total DNS Filter Type Drop",
									},
									"total_dns_filter_class_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters Total DNS Filter Class Drop",
									},
									"dns_filter_type_a_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type A Drop",
									},
									"dns_filter_type_aaaa_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type AAAA Drop",
									},
									"dns_filter_type_cname_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type CNAME Drop",
									},
									"dns_filter_type_mx_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type MX Drop",
									},
									"dns_filter_type_ns_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type NS Drop",
									},
									"dns_filter_type_srv_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type SRV Drop",
									},
									"dns_filter_type_ptr_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type PTR Drop",
									},
									"dns_filter_type_soa_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type SOA Drop",
									},
									"dns_filter_type_txt_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type TXT Drop",
									},
									"dns_filter_type_any_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type Any Drop",
									},
									"dns_filter_type_others_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type OTHERS Drop",
									},
									"dns_filter_class_internet_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class INTERNET Drop",
									},
									"dns_filter_class_chaos_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class CHAOS Drop",
									},
									"dns_filter_class_hesiod_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class HESIOD Drop",
									},
									"dns_filter_class_none_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class NONE Drop",
									},
									"dns_filter_class_any_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class ANY Drop",
									},
									"dns_filter_class_others_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class OTHER Drop",
									},
									"dns_rpz_action_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS RPZ Action Drop",
									},
									"dnsrrl_bad_fqdn": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS Response-Rate-Limiting Bad FQDN",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"fw_server_port_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"es_resp_400": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 400",
									},
									"es_resp_500": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 500",
									},
									"es_resp_invalid_http": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total non-http response",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"es_resp_400": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 400",
									},
									"es_resp_500": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 500",
									},
									"es_resp_invalid_http": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total non-http response",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"fw_service_group_mem_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"curr_conn_overflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Current connection counter overflow count",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"curr_conn_overflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Current connection counter overflow count",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"fw_service_group_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server_selection_fail_reset": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Service selection fail reset",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"server_selection_fail_reset": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Service selection fail reset",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"imap_vport_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"svrsel_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
									},
									"no_route": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no route failure",
									},
									"snat_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for source nat failure",
									},
									"line_too_long": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for line too long",
									},
									"invalid_start_line": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid start line",
									},
									"cant_find_pasv": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find pasv",
									},
									"smp_create_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for smp create fail",
									},
									"data_server_conn_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data svr conn fail",
									},
									"data_send_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data send fail",
									},
									"cant_find_epsv": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find epsv",
									},
									"auth_unsupported": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported auth",
									},
									"unsupported_pbsz_value": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PBSZ",
									},
									"unsupported_prot_value": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PROT",
									},
									"bad_sequence": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Sequence",
									},
									"rsv_persist_conn_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel Persist fail",
									},
									"smp_v6_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv6 fail",
									},
									"smp_v4_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv4 fail",
									},
									"insert_tuple_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel insert tuple fail",
									},
									"cl_est_err": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client EST state erro",
									},
									"ser_connecting_err": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv CTNG state error",
									},
									"server_response_err": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv RESP state error",
									},
									"cl_request_err": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client RQ state error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"svrsel_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
									},
									"no_route": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no route failure",
									},
									"snat_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for source nat failure",
									},
									"line_too_long": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for line too long",
									},
									"invalid_start_line": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid start line",
									},
									"cant_find_pasv": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find pasv",
									},
									"smp_create_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for smp create fail",
									},
									"data_server_conn_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data svr conn fail",
									},
									"data_send_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data send fail",
									},
									"cant_find_epsv": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find epsv",
									},
									"auth_unsupported": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported auth",
									},
									"unsupported_pbsz_value": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PBSZ",
									},
									"unsupported_prot_value": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PROT",
									},
									"bad_sequence": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Sequence",
									},
									"rsv_persist_conn_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel Persist fail",
									},
									"smp_v6_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv6 fail",
									},
									"smp_v4_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv4 fail",
									},
									"insert_tuple_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel insert tuple fail",
									},
									"cl_est_err": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client EST state erro",
									},
									"ser_connecting_err": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv CTNG state error",
									},
									"server_response_err": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv RESP state error",
									},
									"cl_request_err": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client RQ state error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"interface_ethernet_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"input_errors": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Input errors",
									},
									"crc": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRC",
									},
									"runts": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Runts",
									},
									"giants": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Giants",
									},
									"output_errors": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Output errors",
									},
									"collisions": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Collisions",
									},
									"giants_output": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Output Giants",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"input_errors": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Input errors",
									},
									"crc": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRC",
									},
									"runts": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Runts",
									},
									"giants": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Giants",
									},
									"output_errors": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Output errors",
									},
									"collisions": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Collisions",
									},
									"giants_output": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Output Giants",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"interface_tunnel_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"num_rx_err_pkts": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for received error packets",
									},
									"num_tx_err_pkts": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sent error packets",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"num_rx_err_pkts": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for received error packets",
									},
									"num_tx_err_pkts": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sent error packets",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"netflow_monitor_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"nat44_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT44 Flow Records Failed",
									},
									"nat64_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT64 Flow Records Failed",
									},
									"dslite_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Dslite Flow Records Failed",
									},
									"session_event_nat44_records_sent_failur": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Nat44 Session Event Records Failed",
									},
									"session_event_nat64_records_sent_failur": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Nat64 Session Event Records Falied",
									},
									"session_event_dslite_records_sent_failu": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Dslite Session Event Records Failed",
									},
									"session_event_fw4_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for FW4 Session Event Records Failed",
									},
									"session_event_fw6_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for FW6 Session Event Records Failed",
									},
									"port_mapping_nat44_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Mapping Nat44 Event Records Failed",
									},
									"port_mapping_nat64_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Mapping Nat64 Event Records Failed",
									},
									"port_mapping_dslite_records_sent_failur": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Mapping Dslite Event Records failed",
									},
									"netflow_v5_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Netflow v5 Records Failed",
									},
									"netflow_v5_ext_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Netflow v5 Ext Records Failed",
									},
									"port_batching_nat44_records_sent_failur": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching Nat44 Records Failed",
									},
									"port_batching_nat64_records_sent_failur": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching Nat64 Records Failed",
									},
									"port_batching_dslite_records_sent_failu": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching Dslite Records Failed",
									},
									"port_batching_v2_nat44_records_sent_fai": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching V2 Nat44 Records Failed",
									},
									"port_batching_v2_nat64_records_sent_fai": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching V2 Nat64 Records Failed",
									},
									"port_batching_v2_dslite_records_sent_fa": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching V2 Dslite Records Falied",
									},
									"custom_session_event_nat44_creation_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Session Creation Records Failed",
									},
									"custom_session_event_nat64_creation_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Session Creation Records Failed",
									},
									"custom_session_event_dslite_creation_re": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Session Creation Records Failed",
									},
									"custom_session_event_nat44_deletion_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Session Deletion Records Failed",
									},
									"custom_session_event_nat64_deletion_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Session Deletion Records Failed",
									},
									"custom_session_event_dslite_deletion_re": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Session Deletion Records Failed",
									},
									"custom_session_event_fw4_creation_recor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW4 Session Creation Records Failed",
									},
									"custom_session_event_fw6_creation_recor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW6 Session Creation Records Failed",
									},
									"custom_session_event_fw4_deletion_recor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW4 Session Deletion Records Failed",
									},
									"custom_session_event_fw6_deletion_recor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW6 Session Deletion Records Failed",
									},
									"custom_deny_reset_event_fw4_records_sen": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW4 Deny/Reset Event Records Failed",
									},
									"custom_deny_reset_event_fw6_records_sen": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW6 Deny/Reset Event Records Failed",
									},
									"custom_port_mapping_nat44_creation_reco": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Map Creation Records Failed",
									},
									"custom_port_mapping_nat64_creation_reco": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Map Creation Records Failed",
									},
									"custom_port_mapping_dslite_creation_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Map Creation Records Failed",
									},
									"custom_port_mapping_nat44_deletion_reco": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Map Deletion Records Failed",
									},
									"custom_port_mapping_nat64_deletion_reco": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Map Deletion Records Failed",
									},
									"custom_port_mapping_dslite_deletion_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Map Deletion Records Failed",
									},
									"custom_port_batching_nat44_creation_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch Creation Records Failed",
									},
									"custom_port_batching_nat64_creation_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch Creation Records Failed",
									},
									"custom_port_batching_dslite_creation_re": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch Creation Records Failed",
									},
									"custom_port_batching_nat44_deletion_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch Deletion Records Failed",
									},
									"custom_port_batching_nat64_deletion_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch Deletion Records Failed",
									},
									"custom_port_batching_dslite_deletion_re": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch Deletion Records Failed",
									},
									"custom_port_batching_v2_nat44_creation_": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch V2 Creation Records Failed",
									},
									"custom_port_batching_v2_nat64_creation_": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch V2 Creation Records Failed",
									},
									"custom_port_batching_v2_dslite_creation": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch V2 Creation Records Failed",
									},
									"custom_port_batching_v2_nat44_deletion_": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch V2 Deletion Records Failed",
									},
									"custom_port_batching_v2_nat64_deletion_": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch V2 Deletion Records Failed",
									},
									"custom_port_batching_v2_dslite_deletion": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch V2 Deletion Records Failed",
									},
									"custom_gtp_c_tunnel_event_records_sent_": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP C Tunnel Records Sent Failure",
									},
									"custom_gtp_u_tunnel_event_records_sent_": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP U Tunnel Records Sent Failure",
									},
									"custom_gtp_deny_event_records_sent_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP Deny Records Sent Failure",
									},
									"custom_gtp_info_event_records_sent_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP Info Records Sent Failure",
									},
									"custom_fw_iddos_entry_created_records_s": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW iDDoS Entry Created Records Sent Failure",
									},
									"custom_fw_iddos_entry_deleted_records_s": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW iDDoS Entry Deleted Records Sent Failure",
									},
									"custom_fw_sesn_limit_exceeded_records_s": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW Session Limit Exceeded Records Sent Failure",
									},
									"custom_nat_iddos_l3_entry_created_recor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom NAT iDDoS L3 Entry Created Records Sent Failure",
									},
									"custom_nat_iddos_l3_entry_deleted_recor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom NAT iDDoS L3 Entry Deleted Records Sent Failure",
									},
									"custom_nat_iddos_l4_entry_created_recor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom NAT iDDoS L4 Entry Created Records Sent Failure",
									},
									"custom_nat_iddos_l4_entry_deleted_recor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom NAT iDDoS L4 Entry Deleted Records Sent Failure",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"nat44_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT44 Flow Records Failed",
									},
									"nat64_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT64 Flow Records Failed",
									},
									"dslite_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Dslite Flow Records Failed",
									},
									"session_event_nat44_records_sent_failur": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Nat44 Session Event Records Failed",
									},
									"session_event_nat64_records_sent_failur": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Nat64 Session Event Records Falied",
									},
									"session_event_dslite_records_sent_failu": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Dslite Session Event Records Failed",
									},
									"session_event_fw4_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for FW4 Session Event Records Failed",
									},
									"session_event_fw6_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for FW6 Session Event Records Failed",
									},
									"port_mapping_nat44_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Mapping Nat44 Event Records Failed",
									},
									"port_mapping_nat64_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Mapping Nat64 Event Records Failed",
									},
									"port_mapping_dslite_records_sent_failur": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Mapping Dslite Event Records failed",
									},
									"netflow_v5_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Netflow v5 Records Failed",
									},
									"netflow_v5_ext_records_sent_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Netflow v5 Ext Records Failed",
									},
									"port_batching_nat44_records_sent_failur": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching Nat44 Records Failed",
									},
									"port_batching_nat64_records_sent_failur": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching Nat64 Records Failed",
									},
									"port_batching_dslite_records_sent_failu": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching Dslite Records Failed",
									},
									"port_batching_v2_nat44_records_sent_fai": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching V2 Nat44 Records Failed",
									},
									"port_batching_v2_nat64_records_sent_fai": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching V2 Nat64 Records Failed",
									},
									"port_batching_v2_dslite_records_sent_fa": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Batching V2 Dslite Records Falied",
									},
									"custom_session_event_nat44_creation_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Session Creation Records Failed",
									},
									"custom_session_event_nat64_creation_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Session Creation Records Failed",
									},
									"custom_session_event_dslite_creation_re": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Session Creation Records Failed",
									},
									"custom_session_event_nat44_deletion_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Session Deletion Records Failed",
									},
									"custom_session_event_nat64_deletion_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Session Deletion Records Failed",
									},
									"custom_session_event_dslite_deletion_re": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Session Deletion Records Failed",
									},
									"custom_session_event_fw4_creation_recor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW4 Session Creation Records Failed",
									},
									"custom_session_event_fw6_creation_recor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW6 Session Creation Records Failed",
									},
									"custom_session_event_fw4_deletion_recor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW4 Session Deletion Records Failed",
									},
									"custom_session_event_fw6_deletion_recor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW6 Session Deletion Records Failed",
									},
									"custom_deny_reset_event_fw4_records_sen": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW4 Deny/Reset Event Records Failed",
									},
									"custom_deny_reset_event_fw6_records_sen": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW6 Deny/Reset Event Records Failed",
									},
									"custom_port_mapping_nat44_creation_reco": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Map Creation Records Failed",
									},
									"custom_port_mapping_nat64_creation_reco": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Map Creation Records Failed",
									},
									"custom_port_mapping_dslite_creation_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Map Creation Records Failed",
									},
									"custom_port_mapping_nat44_deletion_reco": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Map Deletion Records Failed",
									},
									"custom_port_mapping_nat64_deletion_reco": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Map Deletion Records Failed",
									},
									"custom_port_mapping_dslite_deletion_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Map Deletion Records Failed",
									},
									"custom_port_batching_nat44_creation_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch Creation Records Failed",
									},
									"custom_port_batching_nat64_creation_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch Creation Records Failed",
									},
									"custom_port_batching_dslite_creation_re": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch Creation Records Failed",
									},
									"custom_port_batching_nat44_deletion_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch Deletion Records Failed",
									},
									"custom_port_batching_nat64_deletion_rec": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch Deletion Records Failed",
									},
									"custom_port_batching_dslite_deletion_re": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch Deletion Records Failed",
									},
									"custom_port_batching_v2_nat44_creation_": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch V2 Creation Records Failed",
									},
									"custom_port_batching_v2_nat64_creation_": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch V2 Creation Records Failed",
									},
									"custom_port_batching_v2_dslite_creation": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch V2 Creation Records Failed",
									},
									"custom_port_batching_v2_nat44_deletion_": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat44 Port Batch V2 Deletion Records Failed",
									},
									"custom_port_batching_v2_nat64_deletion_": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Nat64 Port Batch V2 Deletion Records Failed",
									},
									"custom_port_batching_v2_dslite_deletion": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom Dslite Port Batch V2 Deletion Records Failed",
									},
									"custom_gtp_c_tunnel_event_records_sent_": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP C Tunnel Records Sent Failure",
									},
									"custom_gtp_u_tunnel_event_records_sent_": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP U Tunnel Records Sent Failure",
									},
									"custom_gtp_deny_event_records_sent_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP Deny Records Sent Failure",
									},
									"custom_gtp_info_event_records_sent_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom GTP Info Records Sent Failure",
									},
									"custom_fw_iddos_entry_created_records_s": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW iDDoS Entry Created Records Sent Failure",
									},
									"custom_fw_iddos_entry_deleted_records_s": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW iDDoS Entry Deleted Records Sent Failure",
									},
									"custom_fw_sesn_limit_exceeded_records_s": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom FW Session Limit Exceeded Records Sent Failure",
									},
									"custom_nat_iddos_l3_entry_created_recor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom NAT iDDoS L3 Entry Created Records Sent Failure",
									},
									"custom_nat_iddos_l3_entry_deleted_recor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom NAT iDDoS L3 Entry Deleted Records Sent Failure",
									},
									"custom_nat_iddos_l4_entry_created_recor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom NAT iDDoS L4 Entry Created Records Sent Failure",
									},
									"custom_nat_iddos_l4_entry_deleted_recor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Custom NAT iDDoS L4 Entry Deleted Records Sent Failure",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"pop3_vport_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"svrsel_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
									},
									"no_route": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no route failure",
									},
									"snat_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for source nat failure",
									},
									"line_too_long": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for line too long",
									},
									"invalid_start_line": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid start line",
									},
									"unsupported_command": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported cmd",
									},
									"bad_sequence": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Sequence",
									},
									"rsv_persist_conn_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel Persist fail",
									},
									"smp_v6_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv6 fail",
									},
									"smp_v4_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv4 fail",
									},
									"insert_tuple_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel insert tuple fail",
									},
									"cl_est_err": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client EST state erro",
									},
									"ser_connecting_err": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv CTNG state error",
									},
									"server_response_err": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv RESP state error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"svrsel_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
									},
									"no_route": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no route failure",
									},
									"snat_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for source nat failure",
									},
									"line_too_long": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for line too long",
									},
									"invalid_start_line": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid start line",
									},
									"unsupported_command": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported cmd",
									},
									"bad_sequence": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Sequence",
									},
									"rsv_persist_conn_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel Persist fail",
									},
									"smp_v6_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv6 fail",
									},
									"smp_v4_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv4 fail",
									},
									"insert_tuple_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel insert tuple fail",
									},
									"cl_est_err": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client EST state erro",
									},
									"ser_connecting_err": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv CTNG state error",
									},
									"server_response_err": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv RESP state error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"rule_set_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"unmatched_drops": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unmatched drops counter",
									},
									"deny": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Denied counter",
									},
									"reset": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Reset counter",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"unmatched_drops": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unmatched drops counter",
									},
									"deny": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Denied counter",
									},
									"reset": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Reset counter",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"slb_port_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"es_resp_300": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 300",
									},
									"es_resp_400": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 400",
									},
									"es_resp_500": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 500",
									},
									"resp_3xx": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 3xx",
									},
									"resp_4xx": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 4xx",
									},
									"resp_5xx": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 5xx",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"es_resp_300": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 300",
									},
									"es_resp_400": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 400",
									},
									"es_resp_500": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 500",
									},
									"resp_3xx": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 3xx",
									},
									"resp_4xx": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 4xx",
									},
									"resp_5xx": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 5xx",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"slb_templ_cache_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"nc_req_header": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheNcReqHeader, help nc_req_header",
									},
									"nc_res_header": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheNcResHeader, help nc_res_header",
									},
									"rv_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheRvFailure, help rv_failure",
									},
									"content_toobig": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheContentToobig, help content_toobig",
									},
									"content_toosmall": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheContentToosmall, help content_toosmall",
									},
									"entry_create_failures": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheEntryCreateFailures, help entry_create_failures",
									},
									"header_save_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for header_save_error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"nc_req_header": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheNcReqHeader, help nc_req_header",
									},
									"nc_res_header": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheNcResHeader, help nc_res_header",
									},
									"rv_failure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheRvFailure, help rv_failure",
									},
									"content_toobig": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheContentToobig, help content_toobig",
									},
									"content_toosmall": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheContentToosmall, help content_toosmall",
									},
									"entry_create_failures": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheEntryCreateFailures, help entry_create_failures",
									},
									"header_save_error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for header_save_error",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"slb_vport_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"total_mf_dns_pkts": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total MF DNS packets",
									},
									"es_total_failure_actions": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total failure actions",
									},
									"compression_miss_no_client": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Compression miss no client",
									},
									"compression_miss_template_exclusion": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Compression miss template exclusion",
									},
									"loc_deny": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Geo-location Deny count",
									},
									"dnsrrl_total_dropped": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS Response-Rate-Limiting Total Responses Dropped",
									},
									"dnsrrl_bad_fqdn": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS Response-Rate-Limiting Bad FQDN",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"total_mf_dns_pkts": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total MF DNS packets",
									},
									"es_total_failure_actions": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total failure actions",
									},
									"compression_miss_no_client": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Compression miss no client",
									},
									"compression_miss_template_exclusion": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Compression miss template exclusion",
									},
									"loc_deny": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Geo-location Deny count",
									},
									"dnsrrl_total_dropped": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS Response-Rate-Limiting Total Responses Dropped",
									},
									"dnsrrl_bad_fqdn": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS Response-Rate-Limiting Bad FQDN",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"smtp_vport_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"no_proxy": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No proxy error",
									},
									"parse_req_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse request failure",
									},
									"server_select_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
									},
									"forward_req_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Forward request failure",
									},
									"forward_req_data_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Forward REQ data failure",
									},
									"snat_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT failure",
									},
									"send_client_service_not_ready": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Sent client serv-not-rdy",
									},
									"recv_server_unknow_reply_code": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Recv server unknown-code",
									},
									"read_request_line_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Read request line fail",
									},
									"get_all_headers_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Get all headers fail",
									},
									"too_many_headers": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Too many headers",
									},
									"line_too_long": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Line too long",
									},
									"line_extend_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Line extend fail",
									},
									"line_table_extend_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Table extend fail",
									},
									"parse_request_line_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse request line fail",
									},
									"insert_resonse_line_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Ins response line fail",
									},
									"remove_resonse_line_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Del response line fail",
									},
									"parse_resonse_line_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse response line fail",
									},
									"server_starttls_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server side STARTTLS fail",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"no_proxy": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No proxy error",
									},
									"parse_req_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse request failure",
									},
									"server_select_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
									},
									"forward_req_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Forward request failure",
									},
									"forward_req_data_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Forward REQ data failure",
									},
									"snat_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT failure",
									},
									"send_client_service_not_ready": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Sent client serv-not-rdy",
									},
									"recv_server_unknow_reply_code": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Recv server unknown-code",
									},
									"read_request_line_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Read request line fail",
									},
									"get_all_headers_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Get all headers fail",
									},
									"too_many_headers": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Too many headers",
									},
									"line_too_long": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Line too long",
									},
									"line_extend_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Line extend fail",
									},
									"line_table_extend_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Table extend fail",
									},
									"parse_request_line_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse request line fail",
									},
									"insert_resonse_line_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Ins response line fail",
									},
									"remove_resonse_line_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Del response line fail",
									},
									"parse_resonse_line_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse response line fail",
									},
									"server_starttls_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server side STARTTLS fail",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"templ_gtp_plcy_tmpl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
						},
						"capture_config": {
							Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"trigger_stats_severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
									},
									"error_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
									},
									"error_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
									},
									"drop_alert": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
									},
									"drop_warning": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
									},
									"drop_critical": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_inc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"drop_vld_gtp_ie_repeat_count_exceed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP repeated IE count exceeded",
									},
									"drop_vld_reserved_field_set": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Reserved Header Field Set",
									},
									"drop_vld_tunnel_id_flag": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Tunnel Header Flag Not Set",
									},
									"drop_vld_invalid_flow_label_v0": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Invalid Flow Label in GTPv0-C Header",
									},
									"drop_vld_invalid_teid": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Invalid TEID Value",
									},
									"drop_vld_out_of_state": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Out Of State GTP Message",
									},
									"drop_vld_mandatory_information_element": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Mandatory IE Not Present",
									},
									"drop_vld_mandatory_ie_in_grouped_ie": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Mandatory IE in Grouped IE Not Present",
									},
									"drop_vld_out_of_order_ie": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv1-C Message Out of Order IE",
									},
									"drop_vld_out_of_state_ie": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Unexpected IE Present in Message",
									},
									"drop_vld_reserved_information_element": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Reserved IE Field Present",
									},
									"drop_vld_version_not_supported": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Invalid GTP version",
									},
									"drop_vld_message_length": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Message Length Exceeded",
									},
									"drop_vld_cross_layer_correlation": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Cross Layer IP Address Mismatch",
									},
									"drop_vld_country_code_mismatch": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Country Code Mismatch in IMSI and MSISDN",
									},
									"drop_vld_gtp_u_spoofed_source_address": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-U IP Address Spoofed",
									},
									"drop_vld_gtp_bearer_count_exceed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP Bearer count exceeded max (11)",
									},
									"drop_vld_gtp_v2_wrong_lbi_create_bearer": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV2-C Wrong LBI in Create Bearer Request",
									},
									"drop_vld_v0_reserved_message_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv0-C Reserved Message Drop",
									},
									"drop_vld_v1_reserved_message_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv1-C Reserved Message Drop",
									},
									"drop_vld_v2_reserved_message_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv2-C Reserved Message Drop",
									},
									"drop_vld_invalid_pkt_len_piggyback": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Piggyback message invalid packet length",
									},
									"drop_vld_sanity_failed_piggyback": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: piggyback message anomaly failed",
									},
									"drop_vld_sequence_num_correlation": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-C Sequence number Mismatch",
									},
									"drop_vld_gtpv0_seqnum_buffer_full": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV0-C conn Sequence number Buffer Full",
									},
									"drop_vld_gtpv1_seqnum_buffer_full": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV1-C conn Sequence number Buffer Full",
									},
									"drop_vld_gtpv2_seqnum_buffer_full": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV2-C conn Sequence number Buffer Full",
									},
									"drop_vld_gtp_invalid_imsi_len_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-C Invalid IMSI Length Drop",
									},
									"drop_vld_gtp_invalid_apn_len_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-C Invalid APN Length Drop",
									},
									"drop_vld_protocol_flag_unset": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Protocol flag in Header Field not Set",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trigger_stats_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"threshold_exceeded_by": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
									},
									"drop_vld_gtp_ie_repeat_count_exceed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP repeated IE count exceeded",
									},
									"drop_vld_reserved_field_set": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Reserved Header Field Set",
									},
									"drop_vld_tunnel_id_flag": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Tunnel Header Flag Not Set",
									},
									"drop_vld_invalid_flow_label_v0": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Invalid Flow Label in GTPv0-C Header",
									},
									"drop_vld_invalid_teid": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Invalid TEID Value",
									},
									"drop_vld_out_of_state": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Out Of State GTP Message",
									},
									"drop_vld_mandatory_information_element": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Mandatory IE Not Present",
									},
									"drop_vld_mandatory_ie_in_grouped_ie": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Mandatory IE in Grouped IE Not Present",
									},
									"drop_vld_out_of_order_ie": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv1-C Message Out of Order IE",
									},
									"drop_vld_out_of_state_ie": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Unexpected IE Present in Message",
									},
									"drop_vld_reserved_information_element": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Reserved IE Field Present",
									},
									"drop_vld_version_not_supported": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Invalid GTP version",
									},
									"drop_vld_message_length": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Message Length Exceeded",
									},
									"drop_vld_cross_layer_correlation": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Cross Layer IP Address Mismatch",
									},
									"drop_vld_country_code_mismatch": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Country Code Mismatch in IMSI and MSISDN",
									},
									"drop_vld_gtp_u_spoofed_source_address": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-U IP Address Spoofed",
									},
									"drop_vld_gtp_bearer_count_exceed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP Bearer count exceeded max (11)",
									},
									"drop_vld_gtp_v2_wrong_lbi_create_bearer": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV2-C Wrong LBI in Create Bearer Request",
									},
									"drop_vld_v0_reserved_message_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv0-C Reserved Message Drop",
									},
									"drop_vld_v1_reserved_message_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv1-C Reserved Message Drop",
									},
									"drop_vld_v2_reserved_message_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPv2-C Reserved Message Drop",
									},
									"drop_vld_invalid_pkt_len_piggyback": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Piggyback message invalid packet length",
									},
									"drop_vld_sanity_failed_piggyback": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: piggyback message anomaly failed",
									},
									"drop_vld_sequence_num_correlation": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-C Sequence number Mismatch",
									},
									"drop_vld_gtpv0_seqnum_buffer_full": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV0-C conn Sequence number Buffer Full",
									},
									"drop_vld_gtpv1_seqnum_buffer_full": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV1-C conn Sequence number Buffer Full",
									},
									"drop_vld_gtpv2_seqnum_buffer_full": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTPV2-C conn Sequence number Buffer Full",
									},
									"drop_vld_gtp_invalid_imsi_len_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-C Invalid IMSI Length Drop",
									},
									"drop_vld_gtp_invalid_apn_len_drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: GTP-C Invalid APN Length Drop",
									},
									"drop_vld_protocol_flag_unset": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Validation Drop: Protocol flag in Header Field not Set",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityPacketCaptureObjectTemplatesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplates(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplates(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplates(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplates(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Error = in["error"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ParseFail = in["parse_fail"].(int)
		ret.JsonFail = in["json_fail"].(int)
		ret.AttrFail = in["attr_fail"].(int)
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.ParseFail = in["parse_fail"].(int)
		ret.JsonFail = in["json_fail"].(int)
		ret.AttrFail = in["attr_fail"].(int)
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Spn_krb_faiure = in["spn_krb_faiure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Spn_krb_faiure = in["spn_krb_faiure"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Invalid_srv_rsp = in["invalid_srv_rsp"].(int)
		ret.Post_fail = in["post_fail"].(int)
		ret.Invalid_cred = in["invalid_cred"].(int)
		ret.Bad_req = in["bad_req"].(int)
		ret.Not_fnd = in["not_fnd"].(int)
		ret.Error = in["error"].(int)
		ret.Other_error = in["other_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Invalid_srv_rsp = in["invalid_srv_rsp"].(int)
		ret.Post_fail = in["post_fail"].(int)
		ret.Invalid_cred = in["invalid_cred"].(int)
		ret.Bad_req = in["bad_req"].(int)
		ret.Not_fnd = in["not_fnd"].(int)
		ret.Error = in["error"].(int)
		ret.Other_error = in["other_error"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NoCreds = in["no_creds"].(int)
		ret.BadReq = in["bad_req"].(int)
		ret.Unauth = in["unauth"].(int)
		ret.Forbidden = in["forbidden"].(int)
		ret.NotFound = in["not_found"].(int)
		ret.ServerError = in["server_error"].(int)
		ret.Unavailable = in["unavailable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.NoCreds = in["no_creds"].(int)
		ret.BadReq = in["bad_req"].(int)
		ret.Unauth = in["unauth"].(int)
		ret.Forbidden = in["forbidden"].(int)
		ret.NotFound = in["not_found"].(int)
		ret.ServerError = in["server_error"].(int)
		ret.Unavailable = in["unavailable"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Failure = in["failure"].(int)
		ret.BufferAllocFail = in["buffer_alloc_fail"].(int)
		ret.EncodingFail = in["encoding_fail"].(int)
		ret.InsertHeaderFail = in["insert_header_fail"].(int)
		ret.ParseHeaderFail = in["parse_header_fail"].(int)
		ret.InternalError = in["internal_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Failure = in["failure"].(int)
		ret.BufferAllocFail = in["buffer_alloc_fail"].(int)
		ret.EncodingFail = in["encoding_fail"].(int)
		ret.InsertHeaderFail = in["insert_header_fail"].(int)
		ret.ParseHeaderFail = in["parse_header_fail"].(int)
		ret.InternalError = in["internal_error"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Failure = in["failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Failure = in["failure"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MdFail = in["md_fail"].(int)
		ret.AcsFail = in["acs_fail"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.MdFail = in["md_fail"].(int)
		ret.AcsFail = in["acs_fail"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AcsAuthzFail = in["acs_authz_fail"].(int)
		ret.AcsError = in["acs_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.AcsAuthzFail = in["acs_authz_fail"].(int)
		ret.AcsError = in["acs_error"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AdminBindFailure = in["admin_bind_failure"].(int)
		ret.BindFailure = in["bind_failure"].(int)
		ret.SearchFailure = in["search_failure"].(int)
		ret.AuthorizeFailure = in["authorize_failure"].(int)
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
		ret.SslSessionFailure = in["ssl_session_failure"].(int)
		ret.Pw_change_failure = in["pw_change_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.AdminBindFailure = in["admin_bind_failure"].(int)
		ret.BindFailure = in["bind_failure"].(int)
		ret.SearchFailure = in["search_failure"].(int)
		ret.AuthorizeFailure = in["authorize_failure"].(int)
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
		ret.SslSessionFailure = in["ssl_session_failure"].(int)
		ret.Pw_change_failure = in["pw_change_failure"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Timeout = in["timeout"].(int)
		ret.Fail = in["fail"].(int)
		ret.StaplingTimeout = in["stapling_timeout"].(int)
		ret.StaplingFail = in["stapling_fail"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Timeout = in["timeout"].(int)
		ret.Fail = in["fail"].(int)
		ret.StaplingTimeout = in["stapling_timeout"].(int)
		ret.StaplingFail = in["stapling_fail"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authen_failure = in["authen_failure"].(int)
		ret.Authorize_failure = in["authorize_failure"].(int)
		ret.Timeout_error = in["timeout_error"].(int)
		ret.Other_error = in["other_error"].(int)
		ret.AccountingFailure = in["accounting_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Authen_failure = in["authen_failure"].(int)
		ret.Authorize_failure = in["authorize_failure"].(int)
		ret.Timeout_error = in["timeout_error"].(int)
		ret.Other_error = in["other_error"].(int)
		ret.AccountingFailure = in["accounting_failure"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Krb_timeout_error = in["krb_timeout_error"].(int)
		ret.Krb_other_error = in["krb_other_error"].(int)
		ret.Krb_pw_expiry = in["krb_pw_expiry"].(int)
		ret.Krb_pw_change_failure = in["krb_pw_change_failure"].(int)
		ret.Ntlm_proto_nego_failure = in["ntlm_proto_nego_failure"].(int)
		ret.Ntlm_session_setup_failure = in["ntlm_session_setup_failure"].(int)
		ret.Ntlm_prepare_req_error = in["ntlm_prepare_req_error"].(int)
		ret.Ntlm_auth_failure = in["ntlm_auth_failure"].(int)
		ret.Ntlm_timeout_error = in["ntlm_timeout_error"].(int)
		ret.Ntlm_other_error = in["ntlm_other_error"].(int)
		ret.Krb_validate_kdc_failure = in["krb_validate_kdc_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Krb_timeout_error = in["krb_timeout_error"].(int)
		ret.Krb_other_error = in["krb_other_error"].(int)
		ret.Krb_pw_expiry = in["krb_pw_expiry"].(int)
		ret.Krb_pw_change_failure = in["krb_pw_change_failure"].(int)
		ret.Ntlm_proto_nego_failure = in["ntlm_proto_nego_failure"].(int)
		ret.Ntlm_session_setup_failure = in["ntlm_session_setup_failure"].(int)
		ret.Ntlm_prepare_req_error = in["ntlm_prepare_req_error"].(int)
		ret.Ntlm_auth_failure = in["ntlm_auth_failure"].(int)
		ret.Ntlm_timeout_error = in["ntlm_timeout_error"].(int)
		ret.Ntlm_other_error = in["ntlm_other_error"].(int)
		ret.Krb_validate_kdc_failure = in["krb_validate_kdc_failure"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_conn_overflow = in["curr_conn_overflow"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Curr_conn_overflow = in["curr_conn_overflow"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Server_selection_fail_reset = in["server_selection_fail_reset"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Server_selection_fail_reset = in["server_selection_fail_reset"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.JwtAuthorizeFailure = in["jwt_authorize_failure"].(int)
		ret.JwtMissingToken = in["jwt_missing_token"].(int)
		ret.JwtMissingClaim = in["jwt_missing_claim"].(int)
		ret.JwtTokenExpired = in["jwt_token_expired"].(int)
		ret.JwtSignatureFailure = in["jwt_signature_failure"].(int)
		ret.JwtOtherError = in["jwt_other_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.JwtAuthorizeFailure = in["jwt_authorize_failure"].(int)
		ret.JwtMissingToken = in["jwt_missing_token"].(int)
		ret.JwtMissingClaim = in["jwt_missing_claim"].(int)
		ret.JwtTokenExpired = in["jwt_token_expired"].(int)
		ret.JwtSignatureFailure = in["jwt_signature_failure"].(int)
		ret.JwtOtherError = in["jwt_other_error"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Es_total_failure_actions = in["es_total_failure_actions"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Es_total_failure_actions = in["es_total_failure_actions"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Inbound_addr_port_validation_failed = in["inbound_addr_port_validation_failed"].(int)
		ret.Inbound_rev_lookup_failed = in["inbound_rev_lookup_failed"].(int)
		ret.Inbound_dest_unreachable = in["inbound_dest_unreachable"].(int)
		ret.Outbound_addr_validation_failed = in["outbound_addr_validation_failed"].(int)
		ret.Outbound_rev_lookup_failed = in["outbound_rev_lookup_failed"].(int)
		ret.Outbound_dest_unreachable = in["outbound_dest_unreachable"].(int)
		ret.Packet_mtu_exceeded = in["packet_mtu_exceeded"].(int)
		ret.Interface_not_configured = in["interface_not_configured"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Inbound_addr_port_validation_failed = in["inbound_addr_port_validation_failed"].(int)
		ret.Inbound_rev_lookup_failed = in["inbound_rev_lookup_failed"].(int)
		ret.Inbound_dest_unreachable = in["inbound_dest_unreachable"].(int)
		ret.Outbound_addr_validation_failed = in["outbound_addr_validation_failed"].(int)
		ret.Outbound_rev_lookup_failed = in["outbound_rev_lookup_failed"].(int)
		ret.Outbound_dest_unreachable = in["outbound_dest_unreachable"].(int)
		ret.Packet_mtu_exceeded = in["packet_mtu_exceeded"].(int)
		ret.Interface_not_configured = in["interface_not_configured"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Inbound_addr_port_validation_failed = in["inbound_addr_port_validation_failed"].(int)
		ret.Inbound_rev_lookup_failed = in["inbound_rev_lookup_failed"].(int)
		ret.Inbound_dest_unreachable = in["inbound_dest_unreachable"].(int)
		ret.Outbound_addr_validation_failed = in["outbound_addr_validation_failed"].(int)
		ret.Outbound_rev_lookup_failed = in["outbound_rev_lookup_failed"].(int)
		ret.Outbound_dest_unreachable = in["outbound_dest_unreachable"].(int)
		ret.Packet_mtu_exceeded = in["packet_mtu_exceeded"].(int)
		ret.Interface_not_configured = in["interface_not_configured"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Inbound_addr_port_validation_failed = in["inbound_addr_port_validation_failed"].(int)
		ret.Inbound_rev_lookup_failed = in["inbound_rev_lookup_failed"].(int)
		ret.Inbound_dest_unreachable = in["inbound_dest_unreachable"].(int)
		ret.Outbound_addr_validation_failed = in["outbound_addr_validation_failed"].(int)
		ret.Outbound_rev_lookup_failed = in["outbound_rev_lookup_failed"].(int)
		ret.Outbound_dest_unreachable = in["outbound_dest_unreachable"].(int)
		ret.Packet_mtu_exceeded = in["packet_mtu_exceeded"].(int)
		ret.Interface_not_configured = in["interface_not_configured"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Server_selection_fail_drop = in["server_selection_fail_drop"].(int)
		ret.Server_selection_fail_reset = in["server_selection_fail_reset"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Server_selection_fail_drop = in["server_selection_fail_drop"].(int)
		ret.Server_selection_fail_reset = in["server_selection_fail_reset"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesDns_vportTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesDns_vportTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesDns_vportTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesDns_vportTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dnsrrl_total_dropped = in["dnsrrl_total_dropped"].(int)
		ret.Total_filter_drop = in["total_filter_drop"].(int)
		ret.Total_max_query_len_drop = in["total_max_query_len_drop"].(int)
		ret.Rcode_notimpl_receive = in["rcode_notimpl_receive"].(int)
		ret.Rcode_notimpl_response = in["rcode_notimpl_response"].(int)
		ret.Gslb_query_bad = in["gslb_query_bad"].(int)
		ret.Gslb_response_bad = in["gslb_response_bad"].(int)
		ret.Total_dns_filter_type_drop = in["total_dns_filter_type_drop"].(int)
		ret.Total_dns_filter_class_drop = in["total_dns_filter_class_drop"].(int)
		ret.Dns_filter_type_a_drop = in["dns_filter_type_a_drop"].(int)
		ret.Dns_filter_type_aaaa_drop = in["dns_filter_type_aaaa_drop"].(int)
		ret.Dns_filter_type_cname_drop = in["dns_filter_type_cname_drop"].(int)
		ret.Dns_filter_type_mx_drop = in["dns_filter_type_mx_drop"].(int)
		ret.Dns_filter_type_ns_drop = in["dns_filter_type_ns_drop"].(int)
		ret.Dns_filter_type_srv_drop = in["dns_filter_type_srv_drop"].(int)
		ret.Dns_filter_type_ptr_drop = in["dns_filter_type_ptr_drop"].(int)
		ret.Dns_filter_type_soa_drop = in["dns_filter_type_soa_drop"].(int)
		ret.Dns_filter_type_txt_drop = in["dns_filter_type_txt_drop"].(int)
		ret.Dns_filter_type_any_drop = in["dns_filter_type_any_drop"].(int)
		ret.Dns_filter_type_others_drop = in["dns_filter_type_others_drop"].(int)
		ret.Dns_filter_class_internet_drop = in["dns_filter_class_internet_drop"].(int)
		ret.Dns_filter_class_chaos_drop = in["dns_filter_class_chaos_drop"].(int)
		ret.Dns_filter_class_hesiod_drop = in["dns_filter_class_hesiod_drop"].(int)
		ret.Dns_filter_class_none_drop = in["dns_filter_class_none_drop"].(int)
		ret.Dns_filter_class_any_drop = in["dns_filter_class_any_drop"].(int)
		ret.Dns_filter_class_others_drop = in["dns_filter_class_others_drop"].(int)
		ret.Dns_rpz_action_drop = in["dns_rpz_action_drop"].(int)
		ret.Dnsrrl_bad_fqdn = in["dnsrrl_bad_fqdn"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Dnsrrl_total_dropped = in["dnsrrl_total_dropped"].(int)
		ret.Total_filter_drop = in["total_filter_drop"].(int)
		ret.Total_max_query_len_drop = in["total_max_query_len_drop"].(int)
		ret.Rcode_notimpl_receive = in["rcode_notimpl_receive"].(int)
		ret.Rcode_notimpl_response = in["rcode_notimpl_response"].(int)
		ret.Gslb_query_bad = in["gslb_query_bad"].(int)
		ret.Gslb_response_bad = in["gslb_response_bad"].(int)
		ret.Total_dns_filter_type_drop = in["total_dns_filter_type_drop"].(int)
		ret.Total_dns_filter_class_drop = in["total_dns_filter_class_drop"].(int)
		ret.Dns_filter_type_a_drop = in["dns_filter_type_a_drop"].(int)
		ret.Dns_filter_type_aaaa_drop = in["dns_filter_type_aaaa_drop"].(int)
		ret.Dns_filter_type_cname_drop = in["dns_filter_type_cname_drop"].(int)
		ret.Dns_filter_type_mx_drop = in["dns_filter_type_mx_drop"].(int)
		ret.Dns_filter_type_ns_drop = in["dns_filter_type_ns_drop"].(int)
		ret.Dns_filter_type_srv_drop = in["dns_filter_type_srv_drop"].(int)
		ret.Dns_filter_type_ptr_drop = in["dns_filter_type_ptr_drop"].(int)
		ret.Dns_filter_type_soa_drop = in["dns_filter_type_soa_drop"].(int)
		ret.Dns_filter_type_txt_drop = in["dns_filter_type_txt_drop"].(int)
		ret.Dns_filter_type_any_drop = in["dns_filter_type_any_drop"].(int)
		ret.Dns_filter_type_others_drop = in["dns_filter_type_others_drop"].(int)
		ret.Dns_filter_class_internet_drop = in["dns_filter_class_internet_drop"].(int)
		ret.Dns_filter_class_chaos_drop = in["dns_filter_class_chaos_drop"].(int)
		ret.Dns_filter_class_hesiod_drop = in["dns_filter_class_hesiod_drop"].(int)
		ret.Dns_filter_class_none_drop = in["dns_filter_class_none_drop"].(int)
		ret.Dns_filter_class_any_drop = in["dns_filter_class_any_drop"].(int)
		ret.Dns_filter_class_others_drop = in["dns_filter_class_others_drop"].(int)
		ret.Dns_rpz_action_drop = in["dns_rpz_action_drop"].(int)
		ret.Dnsrrl_bad_fqdn = in["dnsrrl_bad_fqdn"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesFwServerPortTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesFwServerPortTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesFwServerPortTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesFwServerPortTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Es_resp_400 = in["es_resp_400"].(int)
		ret.Es_resp_500 = in["es_resp_500"].(int)
		ret.Es_resp_invalid_http = in["es_resp_invalid_http"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Es_resp_400 = in["es_resp_400"].(int)
		ret.Es_resp_500 = in["es_resp_500"].(int)
		ret.Es_resp_invalid_http = in["es_resp_invalid_http"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_conn_overflow = in["curr_conn_overflow"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Curr_conn_overflow = in["curr_conn_overflow"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Server_selection_fail_reset = in["server_selection_fail_reset"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Server_selection_fail_reset = in["server_selection_fail_reset"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesImapVportTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Invalid_start_line = in["invalid_start_line"].(int)
		ret.Cant_find_pasv = in["cant_find_pasv"].(int)
		ret.Smp_create_fail = in["smp_create_fail"].(int)
		ret.Data_server_conn_fail = in["data_server_conn_fail"].(int)
		ret.Data_send_fail = in["data_send_fail"].(int)
		ret.Cant_find_epsv = in["cant_find_epsv"].(int)
		ret.Auth_unsupported = in["auth_unsupported"].(int)
		ret.Unsupported_pbsz_value = in["unsupported_pbsz_value"].(int)
		ret.Unsupported_prot_value = in["unsupported_prot_value"].(int)
		ret.Bad_sequence = in["bad_sequence"].(int)
		ret.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		ret.Smp_v6_fail = in["smp_v6_fail"].(int)
		ret.Smp_v4_fail = in["smp_v4_fail"].(int)
		ret.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		ret.Cl_est_err = in["cl_est_err"].(int)
		ret.Ser_connecting_err = in["ser_connecting_err"].(int)
		ret.Server_response_err = in["server_response_err"].(int)
		ret.Cl_request_err = in["cl_request_err"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Invalid_start_line = in["invalid_start_line"].(int)
		ret.Cant_find_pasv = in["cant_find_pasv"].(int)
		ret.Smp_create_fail = in["smp_create_fail"].(int)
		ret.Data_server_conn_fail = in["data_server_conn_fail"].(int)
		ret.Data_send_fail = in["data_send_fail"].(int)
		ret.Cant_find_epsv = in["cant_find_epsv"].(int)
		ret.Auth_unsupported = in["auth_unsupported"].(int)
		ret.Unsupported_pbsz_value = in["unsupported_pbsz_value"].(int)
		ret.Unsupported_prot_value = in["unsupported_prot_value"].(int)
		ret.Bad_sequence = in["bad_sequence"].(int)
		ret.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		ret.Smp_v6_fail = in["smp_v6_fail"].(int)
		ret.Smp_v4_fail = in["smp_v4_fail"].(int)
		ret.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		ret.Cl_est_err = in["cl_est_err"].(int)
		ret.Ser_connecting_err = in["ser_connecting_err"].(int)
		ret.Server_response_err = in["server_response_err"].(int)
		ret.Cl_request_err = in["cl_request_err"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Input_errors = in["input_errors"].(int)
		ret.Crc = in["crc"].(int)
		ret.Runts = in["runts"].(int)
		ret.Giants = in["giants"].(int)
		ret.Output_errors = in["output_errors"].(int)
		ret.Collisions = in["collisions"].(int)
		ret.Giants_output = in["giants_output"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Input_errors = in["input_errors"].(int)
		ret.Crc = in["crc"].(int)
		ret.Runts = in["runts"].(int)
		ret.Giants = in["giants"].(int)
		ret.Output_errors = in["output_errors"].(int)
		ret.Collisions = in["collisions"].(int)
		ret.Giants_output = in["giants_output"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NumRxErrPkts = in["num_rx_err_pkts"].(int)
		ret.NumTxErrPkts = in["num_tx_err_pkts"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.NumRxErrPkts = in["num_rx_err_pkts"].(int)
		ret.NumTxErrPkts = in["num_tx_err_pkts"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Nat44RecordsSentFailure = in["nat44_records_sent_failure"].(int)
		ret.Nat64RecordsSentFailure = in["nat64_records_sent_failure"].(int)
		ret.DsliteRecordsSentFailure = in["dslite_records_sent_failure"].(int)
		ret.SessionEventNat44RecordsSentFailur = in["session_event_nat44_records_sent_failur"].(int)
		ret.SessionEventNat64RecordsSentFailur = in["session_event_nat64_records_sent_failur"].(int)
		ret.SessionEventDsliteRecordsSentFailu = in["session_event_dslite_records_sent_failu"].(int)
		ret.SessionEventFw4RecordsSentFailure = in["session_event_fw4_records_sent_failure"].(int)
		ret.SessionEventFw6RecordsSentFailure = in["session_event_fw6_records_sent_failure"].(int)
		ret.PortMappingNat44RecordsSentFailure = in["port_mapping_nat44_records_sent_failure"].(int)
		ret.PortMappingNat64RecordsSentFailure = in["port_mapping_nat64_records_sent_failure"].(int)
		ret.PortMappingDsliteRecordsSentFailur = in["port_mapping_dslite_records_sent_failur"].(int)
		ret.NetflowV5RecordsSentFailure = in["netflow_v5_records_sent_failure"].(int)
		ret.NetflowV5ExtRecordsSentFailure = in["netflow_v5_ext_records_sent_failure"].(int)
		ret.PortBatchingNat44RecordsSentFailur = in["port_batching_nat44_records_sent_failur"].(int)
		ret.PortBatchingNat64RecordsSentFailur = in["port_batching_nat64_records_sent_failur"].(int)
		ret.PortBatchingDsliteRecordsSentFailu = in["port_batching_dslite_records_sent_failu"].(int)
		ret.PortBatchingV2Nat44RecordsSentFai = in["port_batching_v2_nat44_records_sent_fai"].(int)
		ret.PortBatchingV2Nat64RecordsSentFai = in["port_batching_v2_nat64_records_sent_fai"].(int)
		ret.PortBatchingV2DsliteRecordsSentFa = in["port_batching_v2_dslite_records_sent_fa"].(int)
		ret.CustomSessionEventNat44CreationRec = in["custom_session_event_nat44_creation_rec"].(int)
		ret.CustomSessionEventNat64CreationRec = in["custom_session_event_nat64_creation_rec"].(int)
		ret.CustomSessionEventDsliteCreationRe = in["custom_session_event_dslite_creation_re"].(int)
		ret.CustomSessionEventNat44DeletionRec = in["custom_session_event_nat44_deletion_rec"].(int)
		ret.CustomSessionEventNat64DeletionRec = in["custom_session_event_nat64_deletion_rec"].(int)
		ret.CustomSessionEventDsliteDeletionRe = in["custom_session_event_dslite_deletion_re"].(int)
		ret.CustomSessionEventFw4CreationRecor = in["custom_session_event_fw4_creation_recor"].(int)
		ret.CustomSessionEventFw6CreationRecor = in["custom_session_event_fw6_creation_recor"].(int)
		ret.CustomSessionEventFw4DeletionRecor = in["custom_session_event_fw4_deletion_recor"].(int)
		ret.CustomSessionEventFw6DeletionRecor = in["custom_session_event_fw6_deletion_recor"].(int)
		ret.CustomDenyResetEventFw4RecordsSen = in["custom_deny_reset_event_fw4_records_sen"].(int)
		ret.CustomDenyResetEventFw6RecordsSen = in["custom_deny_reset_event_fw6_records_sen"].(int)
		ret.CustomPortMappingNat44CreationReco = in["custom_port_mapping_nat44_creation_reco"].(int)
		ret.CustomPortMappingNat64CreationReco = in["custom_port_mapping_nat64_creation_reco"].(int)
		ret.CustomPortMappingDsliteCreationRec = in["custom_port_mapping_dslite_creation_rec"].(int)
		ret.CustomPortMappingNat44DeletionReco = in["custom_port_mapping_nat44_deletion_reco"].(int)
		ret.CustomPortMappingNat64DeletionReco = in["custom_port_mapping_nat64_deletion_reco"].(int)
		ret.CustomPortMappingDsliteDeletionRec = in["custom_port_mapping_dslite_deletion_rec"].(int)
		ret.CustomPortBatchingNat44CreationRec = in["custom_port_batching_nat44_creation_rec"].(int)
		ret.CustomPortBatchingNat64CreationRec = in["custom_port_batching_nat64_creation_rec"].(int)
		ret.CustomPortBatchingDsliteCreationRe = in["custom_port_batching_dslite_creation_re"].(int)
		ret.CustomPortBatchingNat44DeletionRec = in["custom_port_batching_nat44_deletion_rec"].(int)
		ret.CustomPortBatchingNat64DeletionRec = in["custom_port_batching_nat64_deletion_rec"].(int)
		ret.CustomPortBatchingDsliteDeletionRe = in["custom_port_batching_dslite_deletion_re"].(int)
		ret.CustomPortBatchingV2Nat44Creation = in["custom_port_batching_v2_nat44_creation_"].(int)
		ret.CustomPortBatchingV2Nat64Creation = in["custom_port_batching_v2_nat64_creation_"].(int)
		ret.CustomPortBatchingV2DsliteCreation = in["custom_port_batching_v2_dslite_creation"].(int)
		ret.CustomPortBatchingV2Nat44Deletion = in["custom_port_batching_v2_nat44_deletion_"].(int)
		ret.CustomPortBatchingV2Nat64Deletion = in["custom_port_batching_v2_nat64_deletion_"].(int)
		ret.CustomPortBatchingV2DsliteDeletion = in["custom_port_batching_v2_dslite_deletion"].(int)
		ret.CustomGtpCTunnelEventRecordsSent = in["custom_gtp_c_tunnel_event_records_sent_"].(int)
		ret.CustomGtpUTunnelEventRecordsSent = in["custom_gtp_u_tunnel_event_records_sent_"].(int)
		ret.CustomGtpDenyEventRecordsSentFail = in["custom_gtp_deny_event_records_sent_fail"].(int)
		ret.CustomGtpInfoEventRecordsSentFail = in["custom_gtp_info_event_records_sent_fail"].(int)
		ret.CustomFwIddosEntryCreatedRecordsS = in["custom_fw_iddos_entry_created_records_s"].(int)
		ret.CustomFwIddosEntryDeletedRecordsS = in["custom_fw_iddos_entry_deleted_records_s"].(int)
		ret.CustomFwSesnLimitExceededRecordsS = in["custom_fw_sesn_limit_exceeded_records_s"].(int)
		ret.CustomNatIddosL3EntryCreatedRecor = in["custom_nat_iddos_l3_entry_created_recor"].(int)
		ret.CustomNatIddosL3EntryDeletedRecor = in["custom_nat_iddos_l3_entry_deleted_recor"].(int)
		ret.CustomNatIddosL4EntryCreatedRecor = in["custom_nat_iddos_l4_entry_created_recor"].(int)
		ret.CustomNatIddosL4EntryDeletedRecor = in["custom_nat_iddos_l4_entry_deleted_recor"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Nat44RecordsSentFailure = in["nat44_records_sent_failure"].(int)
		ret.Nat64RecordsSentFailure = in["nat64_records_sent_failure"].(int)
		ret.DsliteRecordsSentFailure = in["dslite_records_sent_failure"].(int)
		ret.SessionEventNat44RecordsSentFailur = in["session_event_nat44_records_sent_failur"].(int)
		ret.SessionEventNat64RecordsSentFailur = in["session_event_nat64_records_sent_failur"].(int)
		ret.SessionEventDsliteRecordsSentFailu = in["session_event_dslite_records_sent_failu"].(int)
		ret.SessionEventFw4RecordsSentFailure = in["session_event_fw4_records_sent_failure"].(int)
		ret.SessionEventFw6RecordsSentFailure = in["session_event_fw6_records_sent_failure"].(int)
		ret.PortMappingNat44RecordsSentFailure = in["port_mapping_nat44_records_sent_failure"].(int)
		ret.PortMappingNat64RecordsSentFailure = in["port_mapping_nat64_records_sent_failure"].(int)
		ret.PortMappingDsliteRecordsSentFailur = in["port_mapping_dslite_records_sent_failur"].(int)
		ret.NetflowV5RecordsSentFailure = in["netflow_v5_records_sent_failure"].(int)
		ret.NetflowV5ExtRecordsSentFailure = in["netflow_v5_ext_records_sent_failure"].(int)
		ret.PortBatchingNat44RecordsSentFailur = in["port_batching_nat44_records_sent_failur"].(int)
		ret.PortBatchingNat64RecordsSentFailur = in["port_batching_nat64_records_sent_failur"].(int)
		ret.PortBatchingDsliteRecordsSentFailu = in["port_batching_dslite_records_sent_failu"].(int)
		ret.PortBatchingV2Nat44RecordsSentFai = in["port_batching_v2_nat44_records_sent_fai"].(int)
		ret.PortBatchingV2Nat64RecordsSentFai = in["port_batching_v2_nat64_records_sent_fai"].(int)
		ret.PortBatchingV2DsliteRecordsSentFa = in["port_batching_v2_dslite_records_sent_fa"].(int)
		ret.CustomSessionEventNat44CreationRec = in["custom_session_event_nat44_creation_rec"].(int)
		ret.CustomSessionEventNat64CreationRec = in["custom_session_event_nat64_creation_rec"].(int)
		ret.CustomSessionEventDsliteCreationRe = in["custom_session_event_dslite_creation_re"].(int)
		ret.CustomSessionEventNat44DeletionRec = in["custom_session_event_nat44_deletion_rec"].(int)
		ret.CustomSessionEventNat64DeletionRec = in["custom_session_event_nat64_deletion_rec"].(int)
		ret.CustomSessionEventDsliteDeletionRe = in["custom_session_event_dslite_deletion_re"].(int)
		ret.CustomSessionEventFw4CreationRecor = in["custom_session_event_fw4_creation_recor"].(int)
		ret.CustomSessionEventFw6CreationRecor = in["custom_session_event_fw6_creation_recor"].(int)
		ret.CustomSessionEventFw4DeletionRecor = in["custom_session_event_fw4_deletion_recor"].(int)
		ret.CustomSessionEventFw6DeletionRecor = in["custom_session_event_fw6_deletion_recor"].(int)
		ret.CustomDenyResetEventFw4RecordsSen = in["custom_deny_reset_event_fw4_records_sen"].(int)
		ret.CustomDenyResetEventFw6RecordsSen = in["custom_deny_reset_event_fw6_records_sen"].(int)
		ret.CustomPortMappingNat44CreationReco = in["custom_port_mapping_nat44_creation_reco"].(int)
		ret.CustomPortMappingNat64CreationReco = in["custom_port_mapping_nat64_creation_reco"].(int)
		ret.CustomPortMappingDsliteCreationRec = in["custom_port_mapping_dslite_creation_rec"].(int)
		ret.CustomPortMappingNat44DeletionReco = in["custom_port_mapping_nat44_deletion_reco"].(int)
		ret.CustomPortMappingNat64DeletionReco = in["custom_port_mapping_nat64_deletion_reco"].(int)
		ret.CustomPortMappingDsliteDeletionRec = in["custom_port_mapping_dslite_deletion_rec"].(int)
		ret.CustomPortBatchingNat44CreationRec = in["custom_port_batching_nat44_creation_rec"].(int)
		ret.CustomPortBatchingNat64CreationRec = in["custom_port_batching_nat64_creation_rec"].(int)
		ret.CustomPortBatchingDsliteCreationRe = in["custom_port_batching_dslite_creation_re"].(int)
		ret.CustomPortBatchingNat44DeletionRec = in["custom_port_batching_nat44_deletion_rec"].(int)
		ret.CustomPortBatchingNat64DeletionRec = in["custom_port_batching_nat64_deletion_rec"].(int)
		ret.CustomPortBatchingDsliteDeletionRe = in["custom_port_batching_dslite_deletion_re"].(int)
		ret.CustomPortBatchingV2Nat44Creation = in["custom_port_batching_v2_nat44_creation_"].(int)
		ret.CustomPortBatchingV2Nat64Creation = in["custom_port_batching_v2_nat64_creation_"].(int)
		ret.CustomPortBatchingV2DsliteCreation = in["custom_port_batching_v2_dslite_creation"].(int)
		ret.CustomPortBatchingV2Nat44Deletion = in["custom_port_batching_v2_nat44_deletion_"].(int)
		ret.CustomPortBatchingV2Nat64Deletion = in["custom_port_batching_v2_nat64_deletion_"].(int)
		ret.CustomPortBatchingV2DsliteDeletion = in["custom_port_batching_v2_dslite_deletion"].(int)
		ret.CustomGtpCTunnelEventRecordsSent = in["custom_gtp_c_tunnel_event_records_sent_"].(int)
		ret.CustomGtpUTunnelEventRecordsSent = in["custom_gtp_u_tunnel_event_records_sent_"].(int)
		ret.CustomGtpDenyEventRecordsSentFail = in["custom_gtp_deny_event_records_sent_fail"].(int)
		ret.CustomGtpInfoEventRecordsSentFail = in["custom_gtp_info_event_records_sent_fail"].(int)
		ret.CustomFwIddosEntryCreatedRecordsS = in["custom_fw_iddos_entry_created_records_s"].(int)
		ret.CustomFwIddosEntryDeletedRecordsS = in["custom_fw_iddos_entry_deleted_records_s"].(int)
		ret.CustomFwSesnLimitExceededRecordsS = in["custom_fw_sesn_limit_exceeded_records_s"].(int)
		ret.CustomNatIddosL3EntryCreatedRecor = in["custom_nat_iddos_l3_entry_created_recor"].(int)
		ret.CustomNatIddosL3EntryDeletedRecor = in["custom_nat_iddos_l3_entry_deleted_recor"].(int)
		ret.CustomNatIddosL4EntryCreatedRecor = in["custom_nat_iddos_l4_entry_created_recor"].(int)
		ret.CustomNatIddosL4EntryDeletedRecor = in["custom_nat_iddos_l4_entry_deleted_recor"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesPop3VportTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesPop3VportTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesPop3VportTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesPop3VportTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Invalid_start_line = in["invalid_start_line"].(int)
		ret.Unsupported_command = in["unsupported_command"].(int)
		ret.Bad_sequence = in["bad_sequence"].(int)
		ret.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		ret.Smp_v6_fail = in["smp_v6_fail"].(int)
		ret.Smp_v4_fail = in["smp_v4_fail"].(int)
		ret.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		ret.Cl_est_err = in["cl_est_err"].(int)
		ret.Ser_connecting_err = in["ser_connecting_err"].(int)
		ret.Server_response_err = in["server_response_err"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Invalid_start_line = in["invalid_start_line"].(int)
		ret.Unsupported_command = in["unsupported_command"].(int)
		ret.Bad_sequence = in["bad_sequence"].(int)
		ret.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		ret.Smp_v6_fail = in["smp_v6_fail"].(int)
		ret.Smp_v4_fail = in["smp_v4_fail"].(int)
		ret.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		ret.Cl_est_err = in["cl_est_err"].(int)
		ret.Ser_connecting_err = in["ser_connecting_err"].(int)
		ret.Server_response_err = in["server_response_err"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesRuleSetTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesRuleSetTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesRuleSetTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesRuleSetTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UnmatchedDrops = in["unmatched_drops"].(int)
		ret.Deny = in["deny"].(int)
		ret.Reset = in["reset"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.UnmatchedDrops = in["unmatched_drops"].(int)
		ret.Deny = in["deny"].(int)
		ret.Reset = in["reset"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesSlbPortTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesSlbPortTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesSlbPortTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesSlbPortTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Es_resp_300 = in["es_resp_300"].(int)
		ret.Es_resp_400 = in["es_resp_400"].(int)
		ret.Es_resp_500 = in["es_resp_500"].(int)
		ret.Resp3xx = in["resp_3xx"].(int)
		ret.Resp4xx = in["resp_4xx"].(int)
		ret.Resp5xx = in["resp_5xx"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Es_resp_300 = in["es_resp_300"].(int)
		ret.Es_resp_400 = in["es_resp_400"].(int)
		ret.Es_resp_500 = in["es_resp_500"].(int)
		ret.Resp3xx = in["resp_3xx"].(int)
		ret.Resp4xx = in["resp_4xx"].(int)
		ret.Resp5xx = in["resp_5xx"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Nc_req_header = in["nc_req_header"].(int)
		ret.Nc_res_header = in["nc_res_header"].(int)
		ret.Rv_failure = in["rv_failure"].(int)
		ret.Content_toobig = in["content_toobig"].(int)
		ret.Content_toosmall = in["content_toosmall"].(int)
		ret.Entry_create_failures = in["entry_create_failures"].(int)
		ret.Header_save_error = in["header_save_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Nc_req_header = in["nc_req_header"].(int)
		ret.Nc_res_header = in["nc_res_header"].(int)
		ret.Rv_failure = in["rv_failure"].(int)
		ret.Content_toobig = in["content_toobig"].(int)
		ret.Content_toosmall = in["content_toosmall"].(int)
		ret.Entry_create_failures = in["entry_create_failures"].(int)
		ret.Header_save_error = in["header_save_error"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesSlbVportTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesSlbVportTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesSlbVportTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesSlbVportTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total_mf_dns_pkts = in["total_mf_dns_pkts"].(int)
		ret.Es_total_failure_actions = in["es_total_failure_actions"].(int)
		ret.Compression_miss_no_client = in["compression_miss_no_client"].(int)
		ret.Compression_miss_template_exclusion = in["compression_miss_template_exclusion"].(int)
		ret.Loc_deny = in["loc_deny"].(int)
		ret.Dnsrrl_total_dropped = in["dnsrrl_total_dropped"].(int)
		ret.Dnsrrl_bad_fqdn = in["dnsrrl_bad_fqdn"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Total_mf_dns_pkts = in["total_mf_dns_pkts"].(int)
		ret.Es_total_failure_actions = in["es_total_failure_actions"].(int)
		ret.Compression_miss_no_client = in["compression_miss_no_client"].(int)
		ret.Compression_miss_template_exclusion = in["compression_miss_template_exclusion"].(int)
		ret.Loc_deny = in["loc_deny"].(int)
		ret.Dnsrrl_total_dropped = in["dnsrrl_total_dropped"].(int)
		ret.Dnsrrl_bad_fqdn = in["dnsrrl_bad_fqdn"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesSmtpVportTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesSmtpVportTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesSmtpVportTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.No_proxy = in["no_proxy"].(int)
		ret.Parse_req_fail = in["parse_req_fail"].(int)
		ret.Server_select_fail = in["server_select_fail"].(int)
		ret.Forward_req_fail = in["forward_req_fail"].(int)
		ret.Forward_req_data_fail = in["forward_req_data_fail"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Send_client_service_not_ready = in["send_client_service_not_ready"].(int)
		ret.Recv_server_unknow_reply_code = in["recv_server_unknow_reply_code"].(int)
		ret.Read_request_line_fail = in["read_request_line_fail"].(int)
		ret.Get_all_headers_fail = in["get_all_headers_fail"].(int)
		ret.Too_many_headers = in["too_many_headers"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Line_extend_fail = in["line_extend_fail"].(int)
		ret.Line_table_extend_fail = in["line_table_extend_fail"].(int)
		ret.Parse_request_line_fail = in["parse_request_line_fail"].(int)
		ret.Insert_resonse_line_fail = in["insert_resonse_line_fail"].(int)
		ret.Remove_resonse_line_fail = in["remove_resonse_line_fail"].(int)
		ret.Parse_resonse_line_fail = in["parse_resonse_line_fail"].(int)
		ret.Server_starttls_fail = in["server_starttls_fail"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.No_proxy = in["no_proxy"].(int)
		ret.Parse_req_fail = in["parse_req_fail"].(int)
		ret.Server_select_fail = in["server_select_fail"].(int)
		ret.Forward_req_fail = in["forward_req_fail"].(int)
		ret.Forward_req_data_fail = in["forward_req_data_fail"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Send_client_service_not_ready = in["send_client_service_not_ready"].(int)
		ret.Recv_server_unknow_reply_code = in["recv_server_unknow_reply_code"].(int)
		ret.Read_request_line_fail = in["read_request_line_fail"].(int)
		ret.Get_all_headers_fail = in["get_all_headers_fail"].(int)
		ret.Too_many_headers = in["too_many_headers"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Line_extend_fail = in["line_extend_fail"].(int)
		ret.Line_table_extend_fail = in["line_table_extend_fail"].(int)
		ret.Parse_request_line_fail = in["parse_request_line_fail"].(int)
		ret.Insert_resonse_line_fail = in["insert_resonse_line_fail"].(int)
		ret.Remove_resonse_line_fail = in["remove_resonse_line_fail"].(int)
		ret.Parse_resonse_line_fail = in["parse_resonse_line_fail"].(int)
		ret.Server_starttls_fail = in["server_starttls_fail"].(int)
		//omit uuid
	}
	return ret
}

func getSliceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplList(d []interface{}) []edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplList
		oi.Name = in["name"].(string)
		oi.CaptureConfig = in["capture_config"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsSeverity(in["trigger_stats_severity"].([]interface{}))
		oi.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsInc(in["trigger_stats_inc"].([]interface{}))
		oi.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsRate(in["trigger_stats_rate"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsSeverity(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsSeverity {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsInc(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsInc {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsInc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DropVldGtpIeRepeatCountExceed = in["drop_vld_gtp_ie_repeat_count_exceed"].(int)
		ret.DropVldReservedFieldSet = in["drop_vld_reserved_field_set"].(int)
		ret.DropVldTunnelIdFlag = in["drop_vld_tunnel_id_flag"].(int)
		ret.DropVldInvalidFlowLabelV0 = in["drop_vld_invalid_flow_label_v0"].(int)
		ret.DropVldInvalidTeid = in["drop_vld_invalid_teid"].(int)
		ret.DropVldOutOfState = in["drop_vld_out_of_state"].(int)
		ret.DropVldMandatoryInformationElement = in["drop_vld_mandatory_information_element"].(int)
		ret.DropVldMandatoryIeInGroupedIe = in["drop_vld_mandatory_ie_in_grouped_ie"].(int)
		ret.DropVldOutOfOrderIe = in["drop_vld_out_of_order_ie"].(int)
		ret.DropVldOutOfStateIe = in["drop_vld_out_of_state_ie"].(int)
		ret.DropVldReservedInformationElement = in["drop_vld_reserved_information_element"].(int)
		ret.DropVldVersionNotSupported = in["drop_vld_version_not_supported"].(int)
		ret.DropVldMessageLength = in["drop_vld_message_length"].(int)
		ret.DropVldCrossLayerCorrelation = in["drop_vld_cross_layer_correlation"].(int)
		ret.DropVldCountryCodeMismatch = in["drop_vld_country_code_mismatch"].(int)
		ret.DropVldGtpUSpoofedSourceAddress = in["drop_vld_gtp_u_spoofed_source_address"].(int)
		ret.DropVldGtpBearerCountExceed = in["drop_vld_gtp_bearer_count_exceed"].(int)
		ret.DropVldGtpV2WrongLbiCreateBearer = in["drop_vld_gtp_v2_wrong_lbi_create_bearer"].(int)
		ret.DropVldV0ReservedMessageDrop = in["drop_vld_v0_reserved_message_drop"].(int)
		ret.DropVldV1ReservedMessageDrop = in["drop_vld_v1_reserved_message_drop"].(int)
		ret.DropVldV2ReservedMessageDrop = in["drop_vld_v2_reserved_message_drop"].(int)
		ret.DropVldInvalidPktLenPiggyback = in["drop_vld_invalid_pkt_len_piggyback"].(int)
		ret.DropVldSanityFailedPiggyback = in["drop_vld_sanity_failed_piggyback"].(int)
		ret.DropVldSequenceNumCorrelation = in["drop_vld_sequence_num_correlation"].(int)
		ret.DropVldGtpv0SeqnumBufferFull = in["drop_vld_gtpv0_seqnum_buffer_full"].(int)
		ret.DropVldGtpv1SeqnumBufferFull = in["drop_vld_gtpv1_seqnum_buffer_full"].(int)
		ret.DropVldGtpv2SeqnumBufferFull = in["drop_vld_gtpv2_seqnum_buffer_full"].(int)
		ret.DropVldGtpInvalidImsiLenDrop = in["drop_vld_gtp_invalid_imsi_len_drop"].(int)
		ret.DropVldGtpInvalidApnLenDrop = in["drop_vld_gtp_invalid_apn_len_drop"].(int)
		ret.DropVldProtocolFlagUnset = in["drop_vld_protocol_flag_unset"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsRate(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsRate {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsRate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.DropVldGtpIeRepeatCountExceed = in["drop_vld_gtp_ie_repeat_count_exceed"].(int)
		ret.DropVldReservedFieldSet = in["drop_vld_reserved_field_set"].(int)
		ret.DropVldTunnelIdFlag = in["drop_vld_tunnel_id_flag"].(int)
		ret.DropVldInvalidFlowLabelV0 = in["drop_vld_invalid_flow_label_v0"].(int)
		ret.DropVldInvalidTeid = in["drop_vld_invalid_teid"].(int)
		ret.DropVldOutOfState = in["drop_vld_out_of_state"].(int)
		ret.DropVldMandatoryInformationElement = in["drop_vld_mandatory_information_element"].(int)
		ret.DropVldMandatoryIeInGroupedIe = in["drop_vld_mandatory_ie_in_grouped_ie"].(int)
		ret.DropVldOutOfOrderIe = in["drop_vld_out_of_order_ie"].(int)
		ret.DropVldOutOfStateIe = in["drop_vld_out_of_state_ie"].(int)
		ret.DropVldReservedInformationElement = in["drop_vld_reserved_information_element"].(int)
		ret.DropVldVersionNotSupported = in["drop_vld_version_not_supported"].(int)
		ret.DropVldMessageLength = in["drop_vld_message_length"].(int)
		ret.DropVldCrossLayerCorrelation = in["drop_vld_cross_layer_correlation"].(int)
		ret.DropVldCountryCodeMismatch = in["drop_vld_country_code_mismatch"].(int)
		ret.DropVldGtpUSpoofedSourceAddress = in["drop_vld_gtp_u_spoofed_source_address"].(int)
		ret.DropVldGtpBearerCountExceed = in["drop_vld_gtp_bearer_count_exceed"].(int)
		ret.DropVldGtpV2WrongLbiCreateBearer = in["drop_vld_gtp_v2_wrong_lbi_create_bearer"].(int)
		ret.DropVldV0ReservedMessageDrop = in["drop_vld_v0_reserved_message_drop"].(int)
		ret.DropVldV1ReservedMessageDrop = in["drop_vld_v1_reserved_message_drop"].(int)
		ret.DropVldV2ReservedMessageDrop = in["drop_vld_v2_reserved_message_drop"].(int)
		ret.DropVldInvalidPktLenPiggyback = in["drop_vld_invalid_pkt_len_piggyback"].(int)
		ret.DropVldSanityFailedPiggyback = in["drop_vld_sanity_failed_piggyback"].(int)
		ret.DropVldSequenceNumCorrelation = in["drop_vld_sequence_num_correlation"].(int)
		ret.DropVldGtpv0SeqnumBufferFull = in["drop_vld_gtpv0_seqnum_buffer_full"].(int)
		ret.DropVldGtpv1SeqnumBufferFull = in["drop_vld_gtpv1_seqnum_buffer_full"].(int)
		ret.DropVldGtpv2SeqnumBufferFull = in["drop_vld_gtpv2_seqnum_buffer_full"].(int)
		ret.DropVldGtpInvalidImsiLenDrop = in["drop_vld_gtp_invalid_imsi_len_drop"].(int)
		ret.DropVldGtpInvalidApnLenDrop = in["drop_vld_gtp_invalid_apn_len_drop"].(int)
		ret.DropVldProtocolFlagUnset = in["drop_vld_protocol_flag_unset"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureObjectTemplates(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplates {
	var ret edpt.VisibilityPacketCaptureObjectTemplates
	ret.Inst.AamAaaPolicyTmplList = getSliceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplList(d.Get("aam_aaa_policy_tmpl_list").([]interface{}))
	ret.Inst.AamAuthCaptchaInstTmplList = getSliceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplList(d.Get("aam_auth_captcha_inst_tmpl_list").([]interface{}))
	ret.Inst.AamAuthLogonHttpInsTmplList = getSliceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplList(d.Get("aam_auth_logon_http_ins_tmpl_list").([]interface{}))
	ret.Inst.AamAuthRelayFormInstTmplList = getSliceVisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplList(d.Get("aam_auth_relay_form_inst_tmpl_list").([]interface{}))
	ret.Inst.AamAuthRelayHbaseInstTmplList = getSliceVisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplList(d.Get("aam_auth_relay_hbase_inst_tmpl_list").([]interface{}))
	ret.Inst.AamAuthRelayNtlmTmplList = getSliceVisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplList(d.Get("aam_auth_relay_ntlm_tmpl_list").([]interface{}))
	ret.Inst.AamAuthRelayWsFedTmplList = getSliceVisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplList(d.Get("aam_auth_relay_ws_fed_tmpl_list").([]interface{}))
	ret.Inst.AamAuthSamlIdProvTmplList = getSliceVisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplList(d.Get("aam_auth_saml_id_prov_tmpl_list").([]interface{}))
	ret.Inst.AamAuthSamlServiceProvTmplList = getSliceVisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplList(d.Get("aam_auth_saml_service_prov_tmpl_list").([]interface{}))
	ret.Inst.AamAuthServerLdapInstTmplList = getSliceVisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplList(d.Get("aam_auth_server_ldap_inst_tmpl_list").([]interface{}))
	ret.Inst.AamAuthServerOcspInstTmplList = getSliceVisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplList(d.Get("aam_auth_server_ocsp_inst_tmpl_list").([]interface{}))
	ret.Inst.AamAuthServerRadInstTmplList = getSliceVisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplList(d.Get("aam_auth_server_rad_inst_tmpl_list").([]interface{}))
	ret.Inst.AamAuthServerWinInstTmplList = getSliceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplList(d.Get("aam_auth_server_win_inst_tmpl_list").([]interface{}))
	ret.Inst.AamAuthServiceGroupMemTmplList = getSliceVisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplList(d.Get("aam_auth_service_group_mem_tmpl_list").([]interface{}))
	ret.Inst.AamAuthServiceGroupTmplList = getSliceVisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplList(d.Get("aam_auth_service_group_tmpl_list").([]interface{}))
	ret.Inst.AamJwtAuthorizationTmplList = getSliceVisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplList(d.Get("aam_jwt_authorization_tmpl_list").([]interface{}))
	ret.Inst.Cgnv6Dns64VsPortTmplList = getSliceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplList(d.Get("cgnv6_dns64_vs_port_tmpl_list").([]interface{}))
	ret.Inst.Cgnv6EncapDomainTmplList = getSliceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplList(d.Get("cgnv6_encap_domain_tmpl_list").([]interface{}))
	ret.Inst.Cgnv6MapTransDomainTmplList = getSliceVisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplList(d.Get("cgnv6_map_trans_domain_tmpl_list").([]interface{}))
	ret.Inst.Cgnv6ServGroupTmplList = getSliceVisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplList(d.Get("cgnv6_serv_group_tmpl_list").([]interface{}))
	ret.Inst.Dns_vportTmplList = getSliceVisibilityPacketCaptureObjectTemplatesDns_vportTmplList(d.Get("dns_vport_tmpl_list").([]interface{}))
	ret.Inst.FwServerPortTmplList = getSliceVisibilityPacketCaptureObjectTemplatesFwServerPortTmplList(d.Get("fw_server_port_tmpl_list").([]interface{}))
	ret.Inst.FwServiceGroupMemTmplList = getSliceVisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplList(d.Get("fw_service_group_mem_tmpl_list").([]interface{}))
	ret.Inst.FwServiceGroupTmplList = getSliceVisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplList(d.Get("fw_service_group_tmpl_list").([]interface{}))
	ret.Inst.ImapVportTmplList = getSliceVisibilityPacketCaptureObjectTemplatesImapVportTmplList(d.Get("imap_vport_tmpl_list").([]interface{}))
	ret.Inst.InterfaceEthernetTmplList = getSliceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplList(d.Get("interface_ethernet_tmpl_list").([]interface{}))
	ret.Inst.InterfaceTunnelTmplList = getSliceVisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplList(d.Get("interface_tunnel_tmpl_list").([]interface{}))
	ret.Inst.NetflowMonitorTmplList = getSliceVisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplList(d.Get("netflow_monitor_tmpl_list").([]interface{}))
	ret.Inst.Pop3VportTmplList = getSliceVisibilityPacketCaptureObjectTemplatesPop3VportTmplList(d.Get("pop3_vport_tmpl_list").([]interface{}))
	ret.Inst.RuleSetTmplList = getSliceVisibilityPacketCaptureObjectTemplatesRuleSetTmplList(d.Get("rule_set_tmpl_list").([]interface{}))
	ret.Inst.SlbPortTmplList = getSliceVisibilityPacketCaptureObjectTemplatesSlbPortTmplList(d.Get("slb_port_tmpl_list").([]interface{}))
	ret.Inst.SlbTemplCacheTmplList = getSliceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplList(d.Get("slb_templ_cache_tmpl_list").([]interface{}))
	ret.Inst.SlbVportTmplList = getSliceVisibilityPacketCaptureObjectTemplatesSlbVportTmplList(d.Get("slb_vport_tmpl_list").([]interface{}))
	ret.Inst.SmtpVportTmplList = getSliceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplList(d.Get("smtp_vport_tmpl_list").([]interface{}))
	ret.Inst.TemplGtpPlcyTmplList = getSliceVisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplList(d.Get("templ_gtp_plcy_tmpl_list").([]interface{}))
	//omit uuid
	return ret
}
