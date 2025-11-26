package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentage() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_template_sport_anomaly_threshold_packet_rate_percentage`: Percentage of source port entry's parent entry\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentageCreate,
		UpdateContext: resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentageUpdate,
		ReadContext:   resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentageRead,
		DeleteContext: resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentageDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
			},
			"network_object_template_name": {
				Type: schema.TypeString, Required: true, Description: "Network_object_template_name",
			},
		},
	}
}
func resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentageCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentage(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentageRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentageUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentage(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentageRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentageDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentage(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentageRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentage(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentage(d *schema.ResourceData) edpt.DdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentage {
	var ret edpt.DdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentage
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	ret.Inst.Network_object_template_name = d.Get("network_object_template_name").(string)
	return ret
}
