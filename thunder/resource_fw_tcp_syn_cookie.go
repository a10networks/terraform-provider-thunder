package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTcpSynCookie() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_tcp_syn_cookie`: Configure Firewall Syn-Cookie Protection\n\n__PLACEHOLDER__",
		CreateContext: resourceFwTcpSynCookieCreate,
		UpdateContext: resourceFwTcpSynCookieUpdate,
		ReadContext:   resourceFwTcpSynCookieRead,
		DeleteContext: resourceFwTcpSynCookieDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'syn_ack_sent': SYN cookie SYN ACK sent; 'verification_passed': SYN cookie verification passed; 'verification_failed': SYN cookie verification failed; 'conn_setup_failed': SYN cookie connection setup failed;",
						},
					},
				},
			},
			"syn_cookie_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Firewall Syn-Cookie Protection",
			},
			"syn_cookie_on_threshold": {
				Type: schema.TypeInt, Optional: true, Description: "on-threshold for Syn-cookie (Decimal number)",
			},
			"syn_cookie_on_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 120, Description: "on-timeout for Syn-cookie (Timeout in seconds, default is 120 seconds (2 minutes))",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwTcpSynCookieCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpSynCookieCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpSynCookie(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTcpSynCookieRead(ctx, d, meta)
	}
	return diags
}

func resourceFwTcpSynCookieUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpSynCookieUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpSynCookie(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTcpSynCookieRead(ctx, d, meta)
	}
	return diags
}
func resourceFwTcpSynCookieDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpSynCookieDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpSynCookie(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwTcpSynCookieRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpSynCookieRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpSynCookie(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwTcpSynCookieSamplingEnable(d []interface{}) []edpt.FwTcpSynCookieSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.FwTcpSynCookieSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTcpSynCookieSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwTcpSynCookie(d *schema.ResourceData) edpt.FwTcpSynCookie {
	var ret edpt.FwTcpSynCookie
	ret.Inst.SamplingEnable = getSliceFwTcpSynCookieSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SynCookieEnable = d.Get("syn_cookie_enable").(int)
	ret.Inst.SynCookieOnThreshold = d.Get("syn_cookie_on_threshold").(int)
	ret.Inst.SynCookieOnTimeout = d.Get("syn_cookie_on_timeout").(int)
	//omit uuid
	return ret
}
