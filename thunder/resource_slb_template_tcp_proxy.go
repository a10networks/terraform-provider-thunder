package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateTcpProxy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_tcp_proxy`: TCP Proxy\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateTcpProxyCreate,
		UpdateContext: resourceSlbTemplateTcpProxyUpdate,
		ReadContext:   resourceSlbTemplateTcpProxyRead,
		DeleteContext: resourceSlbTemplateTcpProxyDelete,

		Schema: map[string]*schema.Schema{
			"ack_aggressiveness": {
				Type: schema.TypeString, Optional: true, Default: "low", Description: "'low': Delayed ACK; 'medium': Delayed ACK, with ACK on each packet with PUSH flag; 'high': ACK on each packet;",
			},
			"alive_if_active": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "keep connection alive if active traffic",
			},
			"backend_wscale": {
				Type: schema.TypeInt, Optional: true, Description: "The TCP window scale used for the server side, default is off (number)",
			},
			"del_session_on_server_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete session if the server/port goes down (either disabled/hm down)",
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "send reset to client when server is disabled",
			},
			"disable_abc": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Appropriate Byte Counting RFC 3465 Disabled, default is enabled (Appropriate Byte Counting (ABC) is enabled by default)",
			},
			"disable_sack": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "disable Selective Ack Option",
			},
			"disable_tcp_timestamps": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "disable TCP Timestamps Option",
			},
			"disable_window_scale": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "disable TCP Window-Scale Option",
			},
			"down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "send reset to client when server is down",
			},
			"dynamic_buffer_allocation": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Optimally adjust the transmit and receive buffer sizes of TCP proxy while keeping their sum constant",
			},
			"early_retransmit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure the Early-Retransmit Algorithm (RFC 5827) (Early-Retransmit is disabled by default)",
			},
			"fin_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "FIN timeout (sec), default is disabled (number)",
			},
			"force_delete_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "The maximum time that a session can stay in the system before being deleted, default is off (number (second))",
			},
			"force_delete_timeout_100ms": {
				Type: schema.TypeInt, Optional: true, Description: "The maximum time that a session can stay in the system before being deleted, default is off (number in 100ms)",
			},
			"half_close_idle_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "TCP Half Close Idle Timeout (sec), default is off (cmd is deprecated, use fin-timeout instead) (number)",
			},
			"half_open_idle_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "TCP Half Open Idle Timeout (sec), default is off (number)",
			},
			"idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 600, Description: "Idle Timeout (Interval of 60 seconds), default is 600 (idle timeout in second, default 600)",
			},
			"init_cwnd": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "The initial congestion control window size (packets), default is 10 (init-cwnd in packets, default 10)",
			},
			"initial_window_size": {
				Type: schema.TypeInt, Optional: true, Description: "Set the initial window size, default is off (number)",
			},
			"insert_client_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Insert client ip into TCP option",
			},
			"invalid_rate_limit": {
				Type: schema.TypeInt, Optional: true, Default: 500, Description: "Invalid Packet Response Rate Limit (ms), default is 500 (number default 500 challenges)",
			},
			"keepalive_interval": {
				Type: schema.TypeInt, Optional: true, Description: "Interval between keepalive probes (sec), default is off (number (seconds))",
			},
			"keepalive_probes": {
				Type: schema.TypeInt, Optional: true, Description: "Number of keepalive probes sent, default is off",
			},
			"limited_slowstart": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "RFC 3742 Limited Slow-Start for TCP with Large Congestion Windows (number)",
			},
			"maxburst": {
				Type: schema.TypeInt, Optional: true, Default: 25, Description: "The max packet count sent per transmission event (number)",
			},
			"min_rto": {
				Type: schema.TypeInt, Optional: true, Default: 200, Description: "The minmum retransmission timeout, default is 200ms (number)",
			},
			"mss": {
				Type: schema.TypeInt, Optional: true, Default: 1460, Description: "Responding MSS to use if client MSS is large, default is off (number)",
			},
			"nagle": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Nagle Algorithm",
			},
			"naked_ack_on_handshake": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send naked ack before data during 3-way handshake",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "TCP Proxy Template Name",
			},
			"proxy_header": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proxy_header_action": {
							Type: schema.TypeString, Optional: true, Description: "'insert': Insert proxy header;",
						},
						"version": {
							Type: schema.TypeString, Optional: true, Description: "'v1': version 1; 'v2': version 2;",
						},
					},
				},
			},
			"psh_flag_optimization": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Optimized PSH Flag Use",
			},
			"qos": {
				Type: schema.TypeInt, Optional: true, Description: "QOS level (number)",
			},
			"reassembly_limit": {
				Type: schema.TypeInt, Optional: true, Default: 25, Description: "The reassembly queuing limit, default is 25 segments (number)",
			},
			"reassembly_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "The reassembly timeout, default is 30sec (number)",
			},
			"receive_buffer": {
				Type: schema.TypeInt, Optional: true, Default: 200000, Description: "TCP Receive Buffer (default 200k) (number default 200000 bytes)",
			},
			"reno": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Reno Congestion Control Algorithm",
			},
			"reset_fwd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "send reset to server if error happens",
			},
			"reset_rev": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "send reset to client if error happens",
			},
			"retransmit_retries": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Number of Retries for Retransmit, default is 5",
			},
			"server_down_action": {
				Type: schema.TypeString, Optional: true, Description: "'FIN': FIN Connection; 'RST': Reset Connection;",
			},
			"syn_retries": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "SYN Retry Numbers, default is 5",
			},
			"timewait": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Timewait Threshold (sec), default 5 (number)",
			},
			"transmit_buffer": {
				Type: schema.TypeInt, Optional: true, Default: 200000, Description: "TCP Transmit Buffer (default 200k) (number default 200000 bytes)",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbTemplateTcpProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateTcpProxyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateTcpProxy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateTcpProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateTcpProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateTcpProxyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateTcpProxy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateTcpProxyRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateTcpProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateTcpProxyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateTcpProxy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateTcpProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateTcpProxyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateTcpProxy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbTemplateTcpProxyProxyHeader(d []interface{}) edpt.SlbTemplateTcpProxyProxyHeader {

	count1 := len(d)
	var ret edpt.SlbTemplateTcpProxyProxyHeader
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProxyHeaderAction = in["proxy_header_action"].(string)
		ret.Version = in["version"].(string)
	}
	return ret
}

func dataToEndpointSlbTemplateTcpProxy(d *schema.ResourceData) edpt.SlbTemplateTcpProxy {
	var ret edpt.SlbTemplateTcpProxy
	ret.Inst.AckAggressiveness = d.Get("ack_aggressiveness").(string)
	ret.Inst.AliveIfActive = d.Get("alive_if_active").(int)
	ret.Inst.BackendWscale = d.Get("backend_wscale").(int)
	ret.Inst.DelSessionOnServerDown = d.Get("del_session_on_server_down").(int)
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.DisableAbc = d.Get("disable_abc").(int)
	ret.Inst.DisableSack = d.Get("disable_sack").(int)
	ret.Inst.DisableTcpTimestamps = d.Get("disable_tcp_timestamps").(int)
	ret.Inst.DisableWindowScale = d.Get("disable_window_scale").(int)
	ret.Inst.Down = d.Get("down").(int)
	ret.Inst.DynamicBufferAllocation = d.Get("dynamic_buffer_allocation").(int)
	ret.Inst.EarlyRetransmit = d.Get("early_retransmit").(int)
	ret.Inst.FinTimeout = d.Get("fin_timeout").(int)
	ret.Inst.ForceDeleteTimeout = d.Get("force_delete_timeout").(int)
	ret.Inst.ForceDeleteTimeout100ms = d.Get("force_delete_timeout_100ms").(int)
	ret.Inst.HalfCloseIdleTimeout = d.Get("half_close_idle_timeout").(int)
	ret.Inst.HalfOpenIdleTimeout = d.Get("half_open_idle_timeout").(int)
	ret.Inst.IdleTimeout = d.Get("idle_timeout").(int)
	ret.Inst.InitCwnd = d.Get("init_cwnd").(int)
	ret.Inst.InitialWindowSize = d.Get("initial_window_size").(int)
	ret.Inst.InsertClientIp = d.Get("insert_client_ip").(int)
	ret.Inst.InvalidRateLimit = d.Get("invalid_rate_limit").(int)
	ret.Inst.KeepaliveInterval = d.Get("keepalive_interval").(int)
	ret.Inst.KeepaliveProbes = d.Get("keepalive_probes").(int)
	ret.Inst.LimitedSlowstart = d.Get("limited_slowstart").(int)
	ret.Inst.Maxburst = d.Get("maxburst").(int)
	ret.Inst.MinRto = d.Get("min_rto").(int)
	ret.Inst.Mss = d.Get("mss").(int)
	ret.Inst.Nagle = d.Get("nagle").(int)
	ret.Inst.NakedAckOnHandshake = d.Get("naked_ack_on_handshake").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.ProxyHeader = getObjectSlbTemplateTcpProxyProxyHeader(d.Get("proxy_header").([]interface{}))
	ret.Inst.PshFlagOptimization = d.Get("psh_flag_optimization").(int)
	ret.Inst.Qos = d.Get("qos").(int)
	ret.Inst.ReassemblyLimit = d.Get("reassembly_limit").(int)
	ret.Inst.ReassemblyTimeout = d.Get("reassembly_timeout").(int)
	ret.Inst.ReceiveBuffer = d.Get("receive_buffer").(int)
	ret.Inst.Reno = d.Get("reno").(int)
	ret.Inst.ResetFwd = d.Get("reset_fwd").(int)
	ret.Inst.ResetRev = d.Get("reset_rev").(int)
	ret.Inst.RetransmitRetries = d.Get("retransmit_retries").(int)
	ret.Inst.ServerDownAction = d.Get("server_down_action").(string)
	ret.Inst.SynRetries = d.Get("syn_retries").(int)
	ret.Inst.Timewait = d.Get("timewait").(int)
	ret.Inst.TransmitBuffer = d.Get("transmit_buffer").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
