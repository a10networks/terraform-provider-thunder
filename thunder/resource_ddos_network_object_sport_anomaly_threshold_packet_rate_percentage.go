package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectSportAnomalyThresholdPacketRatePercentage() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_sport_anomaly_threshold_packet_rate_percentage`: Percentage of source port entry's parent entry\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectSportAnomalyThresholdPacketRatePercentageCreate,
		UpdateContext: resourceDdosNetworkObjectSportAnomalyThresholdPacketRatePercentageUpdate,
		ReadContext:   resourceDdosNetworkObjectSportAnomalyThresholdPacketRatePercentageRead,
		DeleteContext: resourceDdosNetworkObjectSportAnomalyThresholdPacketRatePercentageDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
			},
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "ObjectName",
			},
		},
	}
}
func resourceDdosNetworkObjectSportAnomalyThresholdPacketRatePercentageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdPacketRatePercentageCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdPacketRatePercentage(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSportAnomalyThresholdPacketRatePercentageRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectSportAnomalyThresholdPacketRatePercentageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdPacketRatePercentageUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdPacketRatePercentage(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSportAnomalyThresholdPacketRatePercentageRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectSportAnomalyThresholdPacketRatePercentageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdPacketRatePercentageDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdPacketRatePercentage(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectSportAnomalyThresholdPacketRatePercentageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdPacketRatePercentageRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdPacketRatePercentage(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosNetworkObjectSportAnomalyThresholdPacketRatePercentage(d *schema.ResourceData) edpt.DdosNetworkObjectSportAnomalyThresholdPacketRatePercentage {
	var ret edpt.DdosNetworkObjectSportAnomalyThresholdPacketRatePercentage
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	ret.Inst.ObjectName = d.Get("object_name").(string)
	return ret
}
