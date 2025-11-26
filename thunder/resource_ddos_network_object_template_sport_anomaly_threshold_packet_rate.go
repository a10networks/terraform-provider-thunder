package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_template_sport_anomaly_threshold_packet_rate`: Packet rate of a source port entry\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRateCreate,
		UpdateContext: resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRateUpdate,
		ReadContext:   resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRateRead,
		DeleteContext: resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRateDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeInt, Optional: true, Description: "Packet rate of a source port entry",
			},
			"network_object_template_name": {
				Type: schema.TypeString, Required: true, Description: "Network_object_template_name",
			},
		},
	}
}
func resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdPacketRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRateRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdPacketRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRateRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdPacketRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateSportAnomalyThresholdPacketRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdPacketRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdPacketRate(d *schema.ResourceData) edpt.DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate {
	var ret edpt.DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	ret.Inst.Network_object_template_name = d.Get("network_object_template_name").(string)
	return ret
}
