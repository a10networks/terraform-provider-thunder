package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectSportAnomalyThresholdPacketRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_sport_anomaly_threshold_packet_rate`: Packet rate of a source port entry\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectSportAnomalyThresholdPacketRateCreate,
		UpdateContext: resourceDdosNetworkObjectSportAnomalyThresholdPacketRateUpdate,
		ReadContext:   resourceDdosNetworkObjectSportAnomalyThresholdPacketRateRead,
		DeleteContext: resourceDdosNetworkObjectSportAnomalyThresholdPacketRateDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeInt, Optional: true, Description: "Packet rate of a source port entry",
			},
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "ObjectName",
			},
		},
	}
}
func resourceDdosNetworkObjectSportAnomalyThresholdPacketRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdPacketRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdPacketRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSportAnomalyThresholdPacketRateRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectSportAnomalyThresholdPacketRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdPacketRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdPacketRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSportAnomalyThresholdPacketRateRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectSportAnomalyThresholdPacketRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdPacketRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdPacketRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectSportAnomalyThresholdPacketRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdPacketRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdPacketRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosNetworkObjectSportAnomalyThresholdPacketRate(d *schema.ResourceData) edpt.DdosNetworkObjectSportAnomalyThresholdPacketRate {
	var ret edpt.DdosNetworkObjectSportAnomalyThresholdPacketRate
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	ret.Inst.ObjectName = d.Get("object_name").(string)
	return ret
}
