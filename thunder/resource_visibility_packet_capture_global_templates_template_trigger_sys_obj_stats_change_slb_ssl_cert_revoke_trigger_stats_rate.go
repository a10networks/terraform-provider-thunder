package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_ssl_cert_revoke_trigger_stats_rate`: Configure stats to trigger packet capture on increment rate\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRateCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRateUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRateRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRateDelete,

		Schema: map[string]*schema.Schema{
			"crl_cache_status_revoked": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL cache status revoked",
			},
			"crl_connection_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL connection errors",
			},
			"crl_other_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL other errors",
			},
			"crl_response_status_revoked": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL response status revoked",
			},
			"crl_response_status_unknown": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL response status unknown",
			},
			"crl_uri_https": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL URI https",
			},
			"crl_uri_not_found": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL URI not found",
			},
			"crl_uri_unsupported": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL URI unsupported",
			},
			"duration": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
			},
			"ocsp_cache_miss": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP cache miss",
			},
			"ocsp_cache_status_revoked": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP cache status revoked",
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
			"ocsp_other_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP other errors",
			},
			"ocsp_response_no_nonce": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP other errors",
			},
			"ocsp_response_nonce_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP other errors",
			},
			"ocsp_response_status_revoked": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP response status revoked",
			},
			"ocsp_response_status_unknown": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP response status unknown",
			},
			"ocsp_uri_https": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP URI https",
			},
			"ocsp_uri_not_found": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP URI not found",
			},
			"ocsp_uri_unsupported": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP URI unsupported",
			},
			"threshold_exceeded_by": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate
	ret.Inst.Crl_cache_status_revoked = d.Get("crl_cache_status_revoked").(int)
	ret.Inst.Crl_connection_error = d.Get("crl_connection_error").(int)
	ret.Inst.Crl_other_error = d.Get("crl_other_error").(int)
	ret.Inst.Crl_response_status_revoked = d.Get("crl_response_status_revoked").(int)
	ret.Inst.Crl_response_status_unknown = d.Get("crl_response_status_unknown").(int)
	ret.Inst.Crl_uri_https = d.Get("crl_uri_https").(int)
	ret.Inst.Crl_uri_not_found = d.Get("crl_uri_not_found").(int)
	ret.Inst.Crl_uri_unsupported = d.Get("crl_uri_unsupported").(int)
	ret.Inst.Duration = d.Get("duration").(int)
	ret.Inst.Ocsp_cache_miss = d.Get("ocsp_cache_miss").(int)
	ret.Inst.Ocsp_cache_status_revoked = d.Get("ocsp_cache_status_revoked").(int)
	ret.Inst.Ocsp_chain_status_revoked = d.Get("ocsp_chain_status_revoked").(int)
	ret.Inst.Ocsp_chain_status_unknown = d.Get("ocsp_chain_status_unknown").(int)
	ret.Inst.Ocsp_connection_error = d.Get("ocsp_connection_error").(int)
	ret.Inst.Ocsp_other_error = d.Get("ocsp_other_error").(int)
	ret.Inst.Ocsp_response_no_nonce = d.Get("ocsp_response_no_nonce").(int)
	ret.Inst.Ocsp_response_nonce_error = d.Get("ocsp_response_nonce_error").(int)
	ret.Inst.Ocsp_response_status_revoked = d.Get("ocsp_response_status_revoked").(int)
	ret.Inst.Ocsp_response_status_unknown = d.Get("ocsp_response_status_unknown").(int)
	ret.Inst.Ocsp_uri_https = d.Get("ocsp_uri_https").(int)
	ret.Inst.Ocsp_uri_not_found = d.Get("ocsp_uri_not_found").(int)
	ret.Inst.Ocsp_uri_unsupported = d.Get("ocsp_uri_unsupported").(int)
	ret.Inst.ThresholdExceededBy = d.Get("threshold_exceeded_by").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
