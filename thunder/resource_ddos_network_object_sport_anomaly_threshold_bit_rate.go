package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectSportAnomalyThresholdBitRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_sport_anomaly_threshold_bit_rate`: Bit rate of a source port entry\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectSportAnomalyThresholdBitRateCreate,
		UpdateContext: resourceDdosNetworkObjectSportAnomalyThresholdBitRateUpdate,
		ReadContext:   resourceDdosNetworkObjectSportAnomalyThresholdBitRateRead,
		DeleteContext: resourceDdosNetworkObjectSportAnomalyThresholdBitRateDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeInt, Optional: true, Description: "Bit rate of a source port entry",
			},
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "ObjectName",
			},
		},
	}
}
func resourceDdosNetworkObjectSportAnomalyThresholdBitRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdBitRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdBitRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSportAnomalyThresholdBitRateRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectSportAnomalyThresholdBitRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdBitRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdBitRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSportAnomalyThresholdBitRateRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectSportAnomalyThresholdBitRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdBitRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdBitRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectSportAnomalyThresholdBitRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSportAnomalyThresholdBitRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSportAnomalyThresholdBitRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosNetworkObjectSportAnomalyThresholdBitRate(d *schema.ResourceData) edpt.DdosNetworkObjectSportAnomalyThresholdBitRate {
	var ret edpt.DdosNetworkObjectSportAnomalyThresholdBitRate
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	ret.Inst.ObjectName = d.Get("object_name").(string)
	return ret
}
