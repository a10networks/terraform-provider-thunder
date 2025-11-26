package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_template_sport_anomaly_threshold_bit_rate_percentage`: Percentage of source port entry's parent entry\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentageCreate,
		UpdateContext: resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentageUpdate,
		ReadContext:   resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentageRead,
		DeleteContext: resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentageDelete,

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
func resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentageCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentageRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentageUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentageRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentageDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentageRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage(d *schema.ResourceData) edpt.DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage {
	var ret edpt.DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	ret.Inst.Network_object_template_name = d.Get("network_object_template_name").(string)
	return ret
}
