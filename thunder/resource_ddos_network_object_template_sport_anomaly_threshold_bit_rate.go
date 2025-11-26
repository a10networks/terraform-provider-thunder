package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_template_sport_anomaly_threshold_bit_rate`: Bit rate of a source port entry\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRateCreate,
		UpdateContext: resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRateUpdate,
		ReadContext:   resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRateRead,
		DeleteContext: resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRateDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeInt, Optional: true, Description: "Bit rate of a source port entry",
			},
			"network_object_template_name": {
				Type: schema.TypeString, Required: true, Description: "Network_object_template_name",
			},
		},
	}
}
func resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdBitRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRateRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdBitRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRateRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdBitRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateSportAnomalyThresholdBitRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdBitRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosNetworkObjectTemplateSportAnomalyThresholdBitRate(d *schema.ResourceData) edpt.DdosNetworkObjectTemplateSportAnomalyThresholdBitRate {
	var ret edpt.DdosNetworkObjectTemplateSportAnomalyThresholdBitRate
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	ret.Inst.Network_object_template_name = d.Get("network_object_template_name").(string)
	return ret
}
