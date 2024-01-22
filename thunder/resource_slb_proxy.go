package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbProxy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_proxy`: Configure Proxy Global\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbProxyCreate,
		UpdateContext: resourceSlbProxyUpdate,
		ReadContext:   resourceSlbProxyRead,
		DeleteContext: resourceSlbProxyDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'num': Num; 'tcp_event': TCP stack event; 'est_event': Connection established; 'data_event': Data received; 'client_fin': Client FIN; 'server_fin': Server FIN; 'wbuf_event': Ready to send data; 'err_event': Error occured; 'no_mem': No memory; 'client_rst': Client RST; 'server_rst': Server RST; 'queue_depth_over_limit': Queue depth over limit; 'event_failed': Event failed; 'conn_not_exist': Conn not exist; 'service_alloc_cb': Service alloc callback; 'service_alloc_cb_failed': Service alloc callback failed; 'service_free_cb': Service free callback; 'service_free_cb_failed': Service free callback failed; 'est_cb_failed': App EST callback failed; 'data_cb_failed': App DATA callback failed; 'wbuf_cb_failed': App WBUF callback failed; 'err_cb_failed': App ERR callback failed; 'start_server_conn': Start server conn; 'start_server_conn_succ': Success; 'start_server_conn_no_route': No route to server; 'start_server_conn_fail_mem': No memory; 'start_server_conn_fail_snat': Failed Source NAT; 'start_server_conn_fail_persist': Fail Persistence; 'start_server_conn_fail_server': Fail Server issue; 'start_server_conn_fail_tuple': Fail Tuple Issue; 'line_too_long': Line too long;",
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
func resourceSlbProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbProxyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbProxy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbProxyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbProxy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbProxyRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbProxyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbProxy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbProxyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbProxy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbProxySamplingEnable(d []interface{}) []edpt.SlbProxySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbProxySamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbProxySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbProxy(d *schema.ResourceData) edpt.SlbProxy {
	var ret edpt.SlbProxy
	ret.Inst.SamplingEnable = getSliceSlbProxySamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
