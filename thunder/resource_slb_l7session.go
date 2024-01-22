package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbL7session() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_l7session`: Configure l7session\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbL7sessionCreate,
		UpdateContext: resourceSlbL7sessionUpdate,
		ReadContext:   resourceSlbL7sessionRead,
		DeleteContext: resourceSlbL7sessionDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'start_server_conn_succ': Start Server Conn Success; 'conn_not_exist': Conn does not exist; 'data_event': Data event from TCP; 'client_fin': FIN from client; 'server_fin': FIN from server; 'wbuf_event': Wbuf event from TCP; 'wbuf_cb_failed': Wbuf event callback failed; 'err_event': Err event from TCP; 'err_cb_failed': Err event callback failed; 'server_conn_failed': Server connection failed; 'client_rst': RST from client; 'server_rst': RST from server; 'client_rst_req': RST from client - request; 'client_rst_connecting': RST from client - connecting; 'client_rst_connected': RST from client - connected; 'client_rst_rsp': RST from client - response; 'server_rst_req': RST from server - request; 'server_rst_connecting': RST from server - connecting; 'server_rst_connected': RST from server - connected; 'server_rst_rsp': RST from server - response; 'proxy_v1_connection': counter for Proxy v1 connection; 'proxy_v2_connection': counter for Proxy v2 connection; 'curr_proxy': Curr proxy conn; 'curr_proxy_client': Curr proxy conn - client; 'curr_proxy_server': Curr proxy conn - server; 'curr_proxy_es': Curr proxy conn - ES; 'total_proxy': Total proxy conn; 'total_proxy_client': Total proxy conn - client; 'total_proxy_server': Total proxy conn - server; 'total_proxy_es': Total proxy conn - ES; 'server_select_fail': Server selection fail; 'est_event': Est event from TCP; 'est_cb_failed': Est event callback fail; 'data_cb_failed': Data event callback fail; 'hps_fwdreq_fail': Fwd req fail; 'hps_fwdreq_fail_buff': Fwd req fail - buff; 'hps_fwdreq_fail_rport': Fwd req fail - rport; 'hps_fwdreq_fail_route': Fwd req fail - route; 'hps_fwdreq_fail_persist': Fwd req fail - persist; 'hps_fwdreq_fail_server': Fwd req fail - server; 'hps_fwdreq_fail_tuple': Fwd req fail - tuple; 'udp_data_event': Data event from UDP;",
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
func resourceSlbL7sessionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbL7sessionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbL7session(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbL7sessionRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbL7sessionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbL7sessionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbL7session(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbL7sessionRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbL7sessionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbL7sessionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbL7session(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbL7sessionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbL7sessionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbL7session(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbL7sessionSamplingEnable(d []interface{}) []edpt.SlbL7sessionSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbL7sessionSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbL7sessionSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbL7session(d *schema.ResourceData) edpt.SlbL7session {
	var ret edpt.SlbL7session
	ret.Inst.SamplingEnable = getSliceSlbL7sessionSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
