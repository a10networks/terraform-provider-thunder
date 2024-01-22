package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslForwardProxy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_ssl_forward_proxy`: SSL forward proxy stats info\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbSslForwardProxyCreate,
		UpdateContext: resourceSlbSslForwardProxyUpdate,
		ReadContext:   resourceSlbSslForwardProxyRead,
		DeleteContext: resourceSlbSslForwardProxyDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'cert_create': Certificates created; 'cert_expr': Certificates expired; 'cert_hit': Certificate cache hits; 'cert_miss': Certificate cache miss; 'conn_bypass': Connections bypassed; 'conn_inspect': Connections inspected; 'bypass-failsafe-ssl-sessions': Bypass Failsafe SSL sessions; 'bypass-sni-sessions': Bypass SNI sessions; 'bypass-client-auth-sessions': Bypass Client Auth sessions; 'failed-in-ssl-handshakes': Failed in SSL handshakes; 'failed-in-crypto-operations': Failed in crypto operations; 'failed-in-tcp': Failed in TCP; 'failed-in-certificate-verification': Failed in Certificate verification; 'failed-in-certificate-signing': Failed in Certificate signing; 'invalid-ocsp-stapling-response': Invalid OCSP Stapling Response; 'revoked-ocsp-response': Revoked OCSP Response; 'unsupported-ssl-version': Unsupported SSL version; 'certificates-in-cache': Certificates in cache; 'connections-failed': Connections failed; 'aflex-bypass': Bypass triggered by aFleX; 'bypass-cert-subject-sessions': Bypass Cert Subject sessions; 'bypass-cert-issuer-sessions': Bypass Cert issuer sessions; 'bypass-cert-san-sessions': Bypass Cert SAN sessions; 'bypass-no-sni-sessions': Bypass NO SNI sessions; 'reset-no-sni-sessions': Reset No SNI sessions; 'bypass-esni-sessions': Bypass ESNI sessions; 'drop-esni-sessions': Drop ESNI sessions; 'bypass-username-sessions': Bypass Username sessions; 'bypass-ad-group-sessions': Bypass AD-group sessions; 'tot_conn_in_buff': Total buffered async connections; 'curr_conn_in_buff': Current buffered async connections; 'async_conn_timeout': SSLi Async connections failures due to timeout; 'async_conn_limit_drop': SSLi Async connections failures due to limit; 'cert_in_cache': Certificates in cache used by HC SSLi App; 'bypass-client-ip-sessions': Bypass Client IP sessions; 'bypass-server-ip-sessions': Bypass Server IP sessions;",
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
func resourceSlbSslForwardProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslForwardProxyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslForwardProxy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSslForwardProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbSslForwardProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslForwardProxyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslForwardProxy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSslForwardProxyRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbSslForwardProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslForwardProxyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslForwardProxy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbSslForwardProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslForwardProxyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslForwardProxy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbSslForwardProxySamplingEnable(d []interface{}) []edpt.SlbSslForwardProxySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbSslForwardProxySamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSslForwardProxySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSslForwardProxy(d *schema.ResourceData) edpt.SlbSslForwardProxy {
	var ret edpt.SlbSslForwardProxy
	ret.Inst.SamplingEnable = getSliceSlbSslForwardProxySamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
