package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslCertRevoke() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_ssl_cert_revoke`: Configure ssl-cert-revoke-stats\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbSslCertRevokeCreate,
		UpdateContext: resourceSlbSslCertRevokeUpdate,
		ReadContext:   resourceSlbSslCertRevokeRead,
		DeleteContext: resourceSlbSslCertRevokeDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'ocsp_stapling_response_good': OCSP stapling response good; 'ocsp_chain_status_good': Certificate chain status good; 'ocsp_chain_status_revoked': Certificate chain status revoked; 'ocsp_chain_status_unknown': Certificate chain status unknown; 'ocsp_request': OCSP requests; 'ocsp_response': OCSP responses; 'ocsp_connection_error': OCSP connection error; 'ocsp_uri_not_found': OCSP URI not found; 'ocsp_uri_https': Log OCSP URI https; 'ocsp_uri_unsupported': OCSP URI unsupported; 'ocsp_response_status_good': OCSP response status good; 'ocsp_response_status_revoked': OCSP response status revoked; 'ocsp_response_status_unknown': OCSP response status unknown; 'ocsp_cache_status_good': OCSP cache status good; 'ocsp_cache_status_revoked': OCSP cache status revoked; 'ocsp_cache_miss': OCSP cache miss; 'ocsp_cache_expired': OCSP cache expired; 'ocsp_other_error': Log OCSP other errors; 'ocsp_response_no_nonce': Log OCSP other errors; 'ocsp_response_nonce_error': Log OCSP other errors; 'crl_request': CRL requests; 'crl_response': CRL responses; 'crl_connection_error': CRL connection errors; 'crl_uri_not_found': CRL URI not found; 'crl_uri_https': CRL URI https; 'crl_uri_unsupported': CRL URI unsupported; 'crl_response_status_good': CRL response status good; 'crl_response_status_revoked': CRL response status revoked; 'crl_response_status_unknown': CRL response status unknown; 'crl_cache_status_good': CRL cache status good; 'crl_cache_status_revoked': CRL cache status revoked; 'crl_other_error': CRL other errors;",
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
func resourceSlbSslCertRevokeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslCertRevokeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslCertRevoke(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSslCertRevokeRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbSslCertRevokeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslCertRevokeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslCertRevoke(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSslCertRevokeRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbSslCertRevokeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslCertRevokeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslCertRevoke(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbSslCertRevokeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslCertRevokeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslCertRevoke(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbSslCertRevokeSamplingEnable(d []interface{}) []edpt.SlbSslCertRevokeSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbSslCertRevokeSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSslCertRevokeSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSslCertRevoke(d *schema.ResourceData) edpt.SlbSslCertRevoke {
	var ret edpt.SlbSslCertRevoke
	ret.Inst.SamplingEnable = getSliceSlbSslCertRevokeSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
