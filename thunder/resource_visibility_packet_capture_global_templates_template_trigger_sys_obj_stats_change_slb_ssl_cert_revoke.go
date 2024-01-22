package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_ssl_cert_revoke`: Configure triggers for slb.ssl-cert-revoke object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeDelete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ocsp_chain_status_revoked": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Certificate chain status revoked",
						},
						"ocsp_chain_status_unknown": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Certificate chain status unknown",
						},
						"ocsp_connection_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP connection error",
						},
						"ocsp_uri_not_found": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP URI not found",
						},
						"ocsp_uri_https": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP URI https",
						},
						"ocsp_uri_unsupported": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP URI unsupported",
						},
						"ocsp_response_status_revoked": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP response status revoked",
						},
						"ocsp_response_status_unknown": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP response status unknown",
						},
						"ocsp_cache_status_revoked": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP cache status revoked",
						},
						"ocsp_cache_miss": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP cache miss",
						},
						"ocsp_other_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP other errors",
						},
						"ocsp_response_no_nonce": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP other errors",
						},
						"ocsp_response_nonce_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP other errors",
						},
						"crl_connection_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL connection errors",
						},
						"crl_uri_not_found": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL URI not found",
						},
						"crl_uri_https": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL URI https",
						},
						"crl_uri_unsupported": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL URI unsupported",
						},
						"crl_response_status_revoked": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL response status revoked",
						},
						"crl_response_status_unknown": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL response status unknown",
						},
						"crl_cache_status_revoked": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL cache status revoked",
						},
						"crl_other_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL other errors",
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
						"ocsp_chain_status_revoked": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Certificate chain status revoked",
						},
						"ocsp_chain_status_unknown": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Certificate chain status unknown",
						},
						"ocsp_connection_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP connection error",
						},
						"ocsp_uri_not_found": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP URI not found",
						},
						"ocsp_uri_https": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP URI https",
						},
						"ocsp_uri_unsupported": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP URI unsupported",
						},
						"ocsp_response_status_revoked": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP response status revoked",
						},
						"ocsp_response_status_unknown": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP response status unknown",
						},
						"ocsp_cache_status_revoked": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP cache status revoked",
						},
						"ocsp_cache_miss": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP cache miss",
						},
						"ocsp_other_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP other errors",
						},
						"ocsp_response_no_nonce": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP other errors",
						},
						"ocsp_response_nonce_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP other errors",
						},
						"crl_connection_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL connection errors",
						},
						"crl_uri_not_found": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL URI not found",
						},
						"crl_uri_https": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL URI https",
						},
						"crl_uri_unsupported": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL URI unsupported",
						},
						"crl_response_status_revoked": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL response status revoked",
						},
						"crl_response_status_unknown": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL response status unknown",
						},
						"crl_cache_status_revoked": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL cache status revoked",
						},
						"crl_other_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL other errors",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc2079(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc2079 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc2079
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ocsp_chain_status_revoked = in["ocsp_chain_status_revoked"].(int)
		ret.Ocsp_chain_status_unknown = in["ocsp_chain_status_unknown"].(int)
		ret.Ocsp_connection_error = in["ocsp_connection_error"].(int)
		ret.Ocsp_uri_not_found = in["ocsp_uri_not_found"].(int)
		ret.Ocsp_uri_https = in["ocsp_uri_https"].(int)
		ret.Ocsp_uri_unsupported = in["ocsp_uri_unsupported"].(int)
		ret.Ocsp_response_status_revoked = in["ocsp_response_status_revoked"].(int)
		ret.Ocsp_response_status_unknown = in["ocsp_response_status_unknown"].(int)
		ret.Ocsp_cache_status_revoked = in["ocsp_cache_status_revoked"].(int)
		ret.Ocsp_cache_miss = in["ocsp_cache_miss"].(int)
		ret.Ocsp_other_error = in["ocsp_other_error"].(int)
		ret.Ocsp_response_no_nonce = in["ocsp_response_no_nonce"].(int)
		ret.Ocsp_response_nonce_error = in["ocsp_response_nonce_error"].(int)
		ret.Crl_connection_error = in["crl_connection_error"].(int)
		ret.Crl_uri_not_found = in["crl_uri_not_found"].(int)
		ret.Crl_uri_https = in["crl_uri_https"].(int)
		ret.Crl_uri_unsupported = in["crl_uri_unsupported"].(int)
		ret.Crl_response_status_revoked = in["crl_response_status_revoked"].(int)
		ret.Crl_response_status_unknown = in["crl_response_status_unknown"].(int)
		ret.Crl_cache_status_revoked = in["crl_cache_status_revoked"].(int)
		ret.Crl_other_error = in["crl_other_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate2080(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate2080 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate2080
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Ocsp_chain_status_revoked = in["ocsp_chain_status_revoked"].(int)
		ret.Ocsp_chain_status_unknown = in["ocsp_chain_status_unknown"].(int)
		ret.Ocsp_connection_error = in["ocsp_connection_error"].(int)
		ret.Ocsp_uri_not_found = in["ocsp_uri_not_found"].(int)
		ret.Ocsp_uri_https = in["ocsp_uri_https"].(int)
		ret.Ocsp_uri_unsupported = in["ocsp_uri_unsupported"].(int)
		ret.Ocsp_response_status_revoked = in["ocsp_response_status_revoked"].(int)
		ret.Ocsp_response_status_unknown = in["ocsp_response_status_unknown"].(int)
		ret.Ocsp_cache_status_revoked = in["ocsp_cache_status_revoked"].(int)
		ret.Ocsp_cache_miss = in["ocsp_cache_miss"].(int)
		ret.Ocsp_other_error = in["ocsp_other_error"].(int)
		ret.Ocsp_response_no_nonce = in["ocsp_response_no_nonce"].(int)
		ret.Ocsp_response_nonce_error = in["ocsp_response_nonce_error"].(int)
		ret.Crl_connection_error = in["crl_connection_error"].(int)
		ret.Crl_uri_not_found = in["crl_uri_not_found"].(int)
		ret.Crl_uri_https = in["crl_uri_https"].(int)
		ret.Crl_uri_unsupported = in["crl_uri_unsupported"].(int)
		ret.Crl_response_status_revoked = in["crl_response_status_revoked"].(int)
		ret.Crl_response_status_unknown = in["crl_response_status_unknown"].(int)
		ret.Crl_cache_status_revoked = in["crl_cache_status_revoked"].(int)
		ret.Crl_other_error = in["crl_other_error"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc2079(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate2080(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
