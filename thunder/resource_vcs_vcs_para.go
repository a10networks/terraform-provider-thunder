package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsVcsPara() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vcs_vcs_para`: Virtual Chassis System Paramter\n\n__PLACEHOLDER__",
		CreateContext: resourceVcsVcsParaCreate,
		UpdateContext: resourceVcsVcsParaUpdate,
		ReadContext:   resourceVcsVcsParaRead,
		DeleteContext: resourceVcsVcsParaDelete,

		Schema: map[string]*schema.Schema{
			"chassis_id": {
				Type: schema.TypeInt, Optional: true, Description: "Chassis ID",
			},
			"config_info": {
				Type: schema.TypeString, Optional: true, Description: "Configuration information (Configuration tag)",
			},
			"config_seq": {
				Type: schema.TypeString, Optional: true, Description: "Configuration sequence number",
			},
			"dead_interval": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "The node will be considered dead as lack of hearbeats after this time (in unit of second, 10 by default)",
			},
			"dead_interval_mseconds": {
				Type: schema.TypeInt, Optional: true, Description: "The node will be considered dead as lack of hearbeats after this time (milisecond) (in unit of msecond, default is 0)",
			},
			"failure_retry_count_value": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "0-255, default is 2",
			},
			"floating_ip_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"floating_ip": {
							Type: schema.TypeString, Optional: true, Description: "Floating IP address (IPv4 address)",
						},
						"floating_ip_mask": {
							Type: schema.TypeString, Optional: true, Description: "Netmask",
						},
					},
				},
			},
			"floating_ipv6_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"floating_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "Floating IPv6 address",
						},
					},
				},
			},
			"force_wait_interval": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "The node will wait the specified time interval before it start aVCS (in unit of second (default is 5))",
			},
			"forever": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "VCS retry forever if fails to join the chassis",
			},
			"memory_stat_interval": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Interval of aVCS memory statistics record (minutes)",
			},
			"multicast_ip": {
				Type: schema.TypeString, Optional: true, Default: "224.0.1.210", Description: "Multicast (group) IP address (Multicast IP address)",
			},
			"multicast_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Multicast (group) IPv6 address (Multicast IPv6 address)",
			},
			"multicast_port": {
				Type: schema.TypeInt, Optional: true, Default: 41217, Description: "Port used in multicast communication (Port number)",
			},
			"size": {
				Type: schema.TypeInt, Optional: true, Description: "file size (MBytes) to transmit to monitor the TCP channel",
			},
			"slog_level": {
				Type: schema.TypeInt, Optional: true, Default: 7, Description: "Set the level of slog for aVCS",
			},
			"slog_method": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set the print method of slog for aVCS",
			},
			"speed_limit": {
				Type: schema.TypeInt, Optional: true, Description: "speed (KByte/s) limitation for the transmit monitor",
			},
			"ssl_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSL",
			},
			"tcp_channel_monitor": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable vBlade TCP channel monitor",
			},
			"time_interval": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "how long between heartbeats (in unit of second, default is 3)",
			},
			"time_interval_mseconds": {
				Type: schema.TypeInt, Optional: true, Description: "how long between heartbeats (mseconds) (in unit of milisecond, default is 0)",
			},
			"transmit_fragment_size": {
				Type: schema.TypeInt, Optional: true, Default: 6000, Description: "Set the fragment size (KByte) of the aVCS transmit",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVcsVcsParaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVcsParaCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVcsPara(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsVcsParaRead(ctx, d, meta)
	}
	return diags
}

func resourceVcsVcsParaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVcsParaUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVcsPara(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsVcsParaRead(ctx, d, meta)
	}
	return diags
}
func resourceVcsVcsParaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVcsParaDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVcsPara(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVcsVcsParaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVcsParaRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVcsPara(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVcsVcsParaFloatingIpCfg(d []interface{}) []edpt.VcsVcsParaFloatingIpCfg {

	count1 := len(d)
	ret := make([]edpt.VcsVcsParaFloatingIpCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsVcsParaFloatingIpCfg
		oi.FloatingIp = in["floating_ip"].(string)
		oi.FloatingIpMask = in["floating_ip_mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVcsVcsParaFloatingIpv6Cfg(d []interface{}) []edpt.VcsVcsParaFloatingIpv6Cfg {

	count1 := len(d)
	ret := make([]edpt.VcsVcsParaFloatingIpv6Cfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsVcsParaFloatingIpv6Cfg
		oi.FloatingIpv6 = in["floating_ipv6"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVcsVcsPara(d *schema.ResourceData) edpt.VcsVcsPara {
	var ret edpt.VcsVcsPara
	ret.Inst.ChassisId = d.Get("chassis_id").(int)
	ret.Inst.ConfigInfo = d.Get("config_info").(string)
	ret.Inst.ConfigSeq = d.Get("config_seq").(string)
	ret.Inst.DeadInterval = d.Get("dead_interval").(int)
	ret.Inst.DeadIntervalMseconds = d.Get("dead_interval_mseconds").(int)
	ret.Inst.FailureRetryCountValue = d.Get("failure_retry_count_value").(int)
	ret.Inst.FloatingIpCfg = getSliceVcsVcsParaFloatingIpCfg(d.Get("floating_ip_cfg").([]interface{}))
	ret.Inst.FloatingIpv6Cfg = getSliceVcsVcsParaFloatingIpv6Cfg(d.Get("floating_ipv6_cfg").([]interface{}))
	ret.Inst.ForceWaitInterval = d.Get("force_wait_interval").(int)
	ret.Inst.Forever = d.Get("forever").(int)
	ret.Inst.MemoryStatInterval = d.Get("memory_stat_interval").(int)
	ret.Inst.MulticastIp = d.Get("multicast_ip").(string)
	ret.Inst.MulticastIpv6 = d.Get("multicast_ipv6").(string)
	ret.Inst.MulticastPort = d.Get("multicast_port").(int)
	ret.Inst.Size = d.Get("size").(int)
	ret.Inst.SlogLevel = d.Get("slog_level").(int)
	ret.Inst.SlogMethod = d.Get("slog_method").(int)
	ret.Inst.Speed_limit = d.Get("speed_limit").(int)
	ret.Inst.SslEnable = d.Get("ssl_enable").(int)
	ret.Inst.TcpChannelMonitor = d.Get("tcp_channel_monitor").(int)
	ret.Inst.TimeInterval = d.Get("time_interval").(int)
	ret.Inst.TimeIntervalMseconds = d.Get("time_interval_mseconds").(int)
	ret.Inst.TransmitFragmentSize = d.Get("transmit_fragment_size").(int)
	//omit uuid
	return ret
}
