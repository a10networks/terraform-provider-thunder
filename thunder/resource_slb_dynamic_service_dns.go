package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbDynamicServiceDns() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_dynamic_service_dns`: Dynamic-service DNS Statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbDynamicServiceDnsCreate,
		UpdateContext: resourceSlbDynamicServiceDnsUpdate,
		ReadContext:   resourceSlbDynamicServiceDnsRead,
		DeleteContext: resourceSlbDynamicServiceDnsDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'n_query': Number of queries; 'n_query_err': Number of query errors; 'n_query_drop': Number of queries dropped due to full queue; 'n_resp': Number of responses; 'n_resp_err': Number of response errors; 'n_resp_f_formerr': Number of response failure with rcode FormErr; 'n_resp_f_servfail': Number of response failure with rcode ServFail; 'n_resp_f_nxdomain': Number of response failure with rcode NXDomain; 'n_resp_f_notimp': Number of response failure with rcode NotImp; 'n_resp_f_refused': Number of response failure with rcode Refused; 'n_resp_f_yxdomain': Number of response failure with rcode YXDomain; 'n_resp_f_yxrrset': Number of response failure with rcode YXRRSet; 'n_resp_f_nxrrset': Number of response failure with rcode NXRRSet; 'n_resp_f_notauth': Number of response failure with rcode NotAuth; 'n_resp_f_notzone': Number of response failure with rcode NotZone; 'n_resp_f_dsotypeni': Number of response failure with rcode DSOTYPENI; 'n_resp_f_badvers': Number of response failure with rcode BADVERS; 'n_resp_f_badkey': Number of response failure with rcode BADKEY; 'n_resp_f_badtime': Number of response failure with rcode BADTIME; 'n_resp_f_badmode': Number of response failure with rcode BADMODE; 'n_resp_f_badname': Number of response failure with rcode BADNAME; 'n_resp_f_badalg': Number of response failure with rcode BADALG; 'n_resp_f_badtrunc': Number of response failure with rcode BADTRUNC; 'n_resp_f_badcookie': Number of response failure with rcode BADCOOKIE; 'n_resp_f_invalid': Number of response failure with invalid rcode;",
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
func resourceSlbDynamicServiceDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDynamicServiceDnsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDynamicServiceDns(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbDynamicServiceDnsRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbDynamicServiceDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDynamicServiceDnsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDynamicServiceDns(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbDynamicServiceDnsRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbDynamicServiceDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDynamicServiceDnsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDynamicServiceDns(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbDynamicServiceDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDynamicServiceDnsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDynamicServiceDns(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbDynamicServiceDnsSamplingEnable(d []interface{}) []edpt.SlbDynamicServiceDnsSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbDynamicServiceDnsSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbDynamicServiceDnsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbDynamicServiceDns(d *schema.ResourceData) edpt.SlbDynamicServiceDns {
	var ret edpt.SlbDynamicServiceDns
	ret.Inst.SamplingEnable = getSliceSlbDynamicServiceDnsSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
