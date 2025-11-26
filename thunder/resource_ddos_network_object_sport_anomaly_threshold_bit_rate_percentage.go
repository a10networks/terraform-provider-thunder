package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectSportAnomalyThresholdBitRatePercentage() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_sport_anomaly_threshold_bit_rate_percentage`: Percentage of source port entry's parent entry\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectSportAnomalyThresholdBitRatePercentageCreate,
		UpdateContext: resourceDdosNetworkObjectSportAnomalyThresholdBitRatePercentageUpdate,
		ReadContext:   resourceDdosNetworkObjectSportAnomalyThresholdBitRatePercentageRead,
		DeleteContext: resourceDdosNetworkObjectSportAnomalyThresholdBitRatePercentageDelete,

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
func resourceDdosNetworkObjectSportAnomalyThresholdBitRatePercentageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdBitRatePercentageCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdBitRatePercentage(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSportAnomalyThresholdBitRatePercentageRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectSportAnomalyThresholdBitRatePercentageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdBitRatePercentageUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdBitRatePercentage(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSportAnomalyThresholdBitRatePercentageRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectSportAnomalyThresholdBitRatePercentageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdBitRatePercentageDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdBitRatePercentage(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectSportAnomalyThresholdBitRatePercentageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdBitRatePercentageRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdBitRatePercentage(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosNetworkObjectSportAnomalyThresholdBitRatePercentage(d *schema.ResourceData) edpt.DdosNetworkObjectSportAnomalyThresholdBitRatePercentage {
	var ret edpt.DdosNetworkObjectSportAnomalyThresholdBitRatePercentage
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	ret.Inst.ObjectName = d.Get("object_name").(string)
	return ret
}
