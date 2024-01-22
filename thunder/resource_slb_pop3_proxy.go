package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPop3Proxy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_pop3_proxy`: Configure POP3 Proxy global\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbPop3ProxyCreate,
		UpdateContext: resourceSlbPop3ProxyUpdate,
		ReadContext:   resourceSlbPop3ProxyRead,
		DeleteContext: resourceSlbPop3ProxyDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'num': Num; 'curr': Current proxy conns; 'total': Total proxy conns; 'svrsel_fail': Server selection failure; 'no_route': no route failure; 'snat_fail': source nat failure; 'line_too_long': line too long; 'line_mem_freed': request line freed; 'invalid_start_line': invalid start line; 'stls': stls cmd; 'request_dont_care': other cmd; 'unsupported_command': Unsupported cmd; 'bad_sequence': Bad Sequence; 'rsv_persist_conn_fail': Serv Sel Persist fail; 'smp_v6_fail': Serv Sel SMPv6 fail; 'smp_v4_fail': Serv Sel SMPv4 fail; 'insert_tuple_fail': Serv Sel insert tuple fail; 'cl_est_err': Client EST state erro; 'ser_connecting_err': Serv CTNG state error; 'server_response_err': Serv RESP state error; 'cl_request_err': Client RQ state error; 'request': Total POP3 Request; 'control_to_ssl': Control chn ssl;",
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
func resourceSlbPop3ProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPop3ProxyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPop3Proxy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbPop3ProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbPop3ProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPop3ProxyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPop3Proxy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbPop3ProxyRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbPop3ProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPop3ProxyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPop3Proxy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbPop3ProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPop3ProxyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPop3Proxy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbPop3ProxySamplingEnable(d []interface{}) []edpt.SlbPop3ProxySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbPop3ProxySamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbPop3ProxySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbPop3Proxy(d *schema.ResourceData) edpt.SlbPop3Proxy {
	var ret edpt.SlbPop3Proxy
	ret.Inst.SamplingEnable = getSliceSlbPop3ProxySamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
