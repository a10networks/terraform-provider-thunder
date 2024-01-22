package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHealthStat() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_health_stat`: Configure health monitor\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbHealthStatCreate,
		UpdateContext: resourceSlbHealthStatUpdate,
		ReadContext:   resourceSlbHealthStatRead,
		DeleteContext: resourceSlbHealthStatDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'num_burst': Number of burst; 'max_jiffie': Maximum number of jiffies; 'min_jiffie': Minimum number of jiffies; 'avg_jiffie': Average number of jiffies; 'open_socket': Number of open sockets; 'open_socket_failed': Number of failed open sockets; 'close_socket': Number of closed sockets; 'connect_failed': Number of failed connections; 'send_packet': Number of packets sent; 'send_packet_failed': Number of packet send failures; 'recv_packet': Number of received packets; 'recv_packet_failed': Number of failed packet receives; 'retry_times': Retry times; 'timeout': Timouet value; 'unexpected_error': Number of unexpected errors; 'conn_imdt_succ': Number of connection immediete success; 'sock_close_before_17': Number of sockets closed before l7; 'sock_close_without_notify': Number of sockets closed without notify; 'curr_health_rate': Current health rate; 'ext_health_rate': External health rate; 'ext_health_rate_val': External health rate value; 'total_number': Total number; 'status_up': Number of status ups; 'status_down': Number of status downs; 'status_unkn': Number of status unknowns; 'status_other': Number of other status; 'running_time': Running time; 'config_health_rate': Config health rate; 'ssl_post_handshake_packet': Number of ssl post handshake packets before client sends request; 'timeout_with_packet': Number of pin timeouts while socket has packets;",
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
func resourceSlbHealthStatCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthStatCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthStat(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHealthStatRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbHealthStatUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthStatUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthStat(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHealthStatRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbHealthStatDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthStatDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthStat(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbHealthStatRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthStatRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthStat(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbHealthStatSamplingEnable(d []interface{}) []edpt.SlbHealthStatSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbHealthStatSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbHealthStatSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbHealthStat(d *schema.ResourceData) edpt.SlbHealthStat {
	var ret edpt.SlbHealthStat
	ret.Inst.SamplingEnable = getSliceSlbHealthStatSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
