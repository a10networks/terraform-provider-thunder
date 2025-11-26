package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectSportAnomalyThresholdIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_sport_anomaly_threshold_ip`: Configure anomaly thresholds applied to source port entries of an IPv4 address\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectSportAnomalyThresholdIpCreate,
		UpdateContext: resourceDdosNetworkObjectSportAnomalyThresholdIpUpdate,
		ReadContext:   resourceDdosNetworkObjectSportAnomalyThresholdIpRead,
		DeleteContext: resourceDdosNetworkObjectSportAnomalyThresholdIpDelete,

		Schema: map[string]*schema.Schema{
			"bit_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Bit rate of a source port entry",
			},
			"bit_rate_percentage": {
				Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
			},
			"bit_rate_percentage_str": {
				Type: schema.TypeString, Required: true, Description: "'bit-rate-percentage': Percentage of source port entry's parent entry;",
			},
			"bit_rate_str": {
				Type: schema.TypeString, Required: true, Description: "'bit-rate': Bit rate of a source port entry;",
			},
			"ip_addr": {
				Type: schema.TypeString, Required: true, Description: "Override threshold",
			},
			"ip_sport_bit_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Bit rate of a source port entry",
			},
			"ip_sport_bit_rate_percentage": {
				Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
			},
			"ip_sport_bit_rate_percentage_str": {
				Type: schema.TypeString, Required: true, Description: "'bit-rate-percentage': Percentage of source port entry's parent entry;",
			},
			"ip_sport_bit_rate_str": {
				Type: schema.TypeString, Required: true, Description: "'bit-rate': Bit rate of a source port entry;",
			},
			"ip_sport_packet_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Packet rate of a source port entry",
			},
			"ip_sport_packet_rate_percentage": {
				Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
			},
			"ip_sport_packet_rate_percentage_str": {
				Type: schema.TypeString, Required: true, Description: "'packet-rate-percentage': Percentage of source port entry's parent entry;",
			},
			"ip_sport_packet_rate_str": {
				Type: schema.TypeString, Required: true, Description: "'packet-rate': Packet rate of a source port entry;",
			},
			"packet_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Packet rate of a source port entry",
			},
			"packet_rate_percentage": {
				Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
			},
			"packet_rate_percentage_str": {
				Type: schema.TypeString, Required: true, Description: "'packet-rate-percentage': Percentage of source port entry's parent entry;",
			},
			"packet_rate_str": {
				Type: schema.TypeString, Required: true, Description: "'packet-rate': Packet rate of a source port entry;",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'udp': UDP port; 'tcp': TCP Port;",
			},
			"sport_num": {
				Type: schema.TypeInt, Required: true, Description: "Source port number",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "ObjectName",
			},
		},
	}
}
func resourceDdosNetworkObjectSportAnomalyThresholdIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSportAnomalyThresholdIpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectSportAnomalyThresholdIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSportAnomalyThresholdIpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectSportAnomalyThresholdIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectSportAnomalyThresholdIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosNetworkObjectSportAnomalyThresholdIp(d *schema.ResourceData) edpt.DdosNetworkObjectSportAnomalyThresholdIp {
	var ret edpt.DdosNetworkObjectSportAnomalyThresholdIp
	ret.Inst.BitRate = d.Get("bit_rate").(int)
	ret.Inst.BitRatePercentage = d.Get("bit_rate_percentage").(int)
	ret.Inst.BitRatePercentageStr = d.Get("bit_rate_percentage_str").(string)
	ret.Inst.BitRateStr = d.Get("bit_rate_str").(string)
	ret.Inst.IpAddr = d.Get("ip_addr").(string)
	ret.Inst.IpSportBitRate = d.Get("ip_sport_bit_rate").(int)
	ret.Inst.IpSportBitRatePercentage = d.Get("ip_sport_bit_rate_percentage").(int)
	ret.Inst.IpSportBitRatePercentageStr = d.Get("ip_sport_bit_rate_percentage_str").(string)
	ret.Inst.IpSportBitRateStr = d.Get("ip_sport_bit_rate_str").(string)
	ret.Inst.IpSportPacketRate = d.Get("ip_sport_packet_rate").(int)
	ret.Inst.IpSportPacketRatePercentage = d.Get("ip_sport_packet_rate_percentage").(int)
	ret.Inst.IpSportPacketRatePercentageStr = d.Get("ip_sport_packet_rate_percentage_str").(string)
	ret.Inst.IpSportPacketRateStr = d.Get("ip_sport_packet_rate_str").(string)
	ret.Inst.PacketRate = d.Get("packet_rate").(int)
	ret.Inst.PacketRatePercentage = d.Get("packet_rate_percentage").(int)
	ret.Inst.PacketRatePercentageStr = d.Get("packet_rate_percentage_str").(string)
	ret.Inst.PacketRateStr = d.Get("packet_rate_str").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SportNum = d.Get("sport_num").(int)
	//omit uuid
	ret.Inst.ObjectName = d.Get("object_name").(string)
	return ret
}
