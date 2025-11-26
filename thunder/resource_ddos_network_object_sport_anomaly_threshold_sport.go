package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectSportAnomalyThresholdSport() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_sport_anomaly_threshold_sport`: Configure anomaly thresholds to a specified source port entry of a network IP entry\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectSportAnomalyThresholdSportCreate,
		UpdateContext: resourceDdosNetworkObjectSportAnomalyThresholdSportUpdate,
		ReadContext:   resourceDdosNetworkObjectSportAnomalyThresholdSportRead,
		DeleteContext: resourceDdosNetworkObjectSportAnomalyThresholdSportDelete,

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
				Type: schema.TypeInt, Required: true, Description: "Port Number",
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
func resourceDdosNetworkObjectSportAnomalyThresholdSportCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdSportCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdSport(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSportAnomalyThresholdSportRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectSportAnomalyThresholdSportUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdSportUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdSport(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSportAnomalyThresholdSportRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectSportAnomalyThresholdSportDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdSportDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdSport(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectSportAnomalyThresholdSportRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdSportRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdSport(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosNetworkObjectSportAnomalyThresholdSport(d *schema.ResourceData) edpt.DdosNetworkObjectSportAnomalyThresholdSport {
	var ret edpt.DdosNetworkObjectSportAnomalyThresholdSport
	ret.Inst.BitRate = d.Get("bit_rate").(int)
	ret.Inst.BitRatePercentage = d.Get("bit_rate_percentage").(int)
	ret.Inst.BitRatePercentageStr = d.Get("bit_rate_percentage_str").(string)
	ret.Inst.BitRateStr = d.Get("bit_rate_str").(string)
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
