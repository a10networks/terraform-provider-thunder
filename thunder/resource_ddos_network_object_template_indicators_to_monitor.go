package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectTemplateIndicatorsToMonitor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_template_indicators_to_monitor`: Configure monitoring of indicators for anomaly based on ONLY adaptive threshold\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectTemplateIndicatorsToMonitorCreate,
		UpdateContext: resourceDdosNetworkObjectTemplateIndicatorsToMonitorUpdate,
		ReadContext:   resourceDdosNetworkObjectTemplateIndicatorsToMonitorRead,
		DeleteContext: resourceDdosNetworkObjectTemplateIndicatorsToMonitorDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Required: true, Description: "",
			},
			"monitor_bit_rate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward bit rate",
			},
			"monitor_fin_rate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "FIN packet rate",
			},
			"monitor_flow_count": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Flow count",
			},
			"monitor_icmp_pkt_rate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "ICMP packet rate",
			},
			"monitor_pkt_rate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward packet rate",
			},
			"monitor_rev_bit_rate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reverse bit rate",
			},
			"monitor_rev_pkt_rate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reverse packet rate",
			},
			"monitor_rst_rate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "RST packet rate",
			},
			"monitor_syn_rate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SYN packet rate",
			},
			"monitor_tcp_pkt_rate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "TCP packet rate",
			},
			"monitor_udp_pkt_rate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "UDP packet rate",
			},
			"monitor_undiscovered_pkt_rate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Undiscovered forward packet rate",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"network_object_template_name": {
				Type: schema.TypeString, Required: true, Description: "Network_object_template_name",
			},
		},
	}
}
func resourceDdosNetworkObjectTemplateIndicatorsToMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateIndicatorsToMonitorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateIndicatorsToMonitor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectTemplateIndicatorsToMonitorRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectTemplateIndicatorsToMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateIndicatorsToMonitorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateIndicatorsToMonitor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectTemplateIndicatorsToMonitorRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectTemplateIndicatorsToMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateIndicatorsToMonitorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateIndicatorsToMonitor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectTemplateIndicatorsToMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateIndicatorsToMonitorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateIndicatorsToMonitor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosNetworkObjectTemplateIndicatorsToMonitor(d *schema.ResourceData) edpt.DdosNetworkObjectTemplateIndicatorsToMonitor {
	var ret edpt.DdosNetworkObjectTemplateIndicatorsToMonitor
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.MonitorBitRate = d.Get("monitor_bit_rate").(int)
	ret.Inst.MonitorFinRate = d.Get("monitor_fin_rate").(int)
	ret.Inst.MonitorFlowCount = d.Get("monitor_flow_count").(int)
	ret.Inst.MonitorIcmpPktRate = d.Get("monitor_icmp_pkt_rate").(int)
	ret.Inst.MonitorPktRate = d.Get("monitor_pkt_rate").(int)
	ret.Inst.MonitorRevBitRate = d.Get("monitor_rev_bit_rate").(int)
	ret.Inst.MonitorRevPktRate = d.Get("monitor_rev_pkt_rate").(int)
	ret.Inst.MonitorRstRate = d.Get("monitor_rst_rate").(int)
	ret.Inst.MonitorSynRate = d.Get("monitor_syn_rate").(int)
	ret.Inst.MonitorTcpPktRate = d.Get("monitor_tcp_pkt_rate").(int)
	ret.Inst.MonitorUdpPktRate = d.Get("monitor_udp_pkt_rate").(int)
	ret.Inst.MonitorUndiscoveredPktRate = d.Get("monitor_undiscovered_pkt_rate").(int)
	//omit uuid
	ret.Inst.Network_object_template_name = d.Get("network_object_template_name").(string)
	return ret
}
