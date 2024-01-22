package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryStatistics() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_web_category_statistics`: Statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceWebCategoryStatisticsCreate,
		UpdateContext: resourceWebCategoryStatisticsUpdate,
		ReadContext:   resourceWebCategoryStatisticsRead,
		DeleteContext: resourceWebCategoryStatisticsDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'db-lookup': db-lookup; 'cloud-cache-lookup': cloud-cache-lookup; 'cloud-lookup': cloud-lookup; 'rtu-lookup': rtu-lookup; 'lookup-latency': lookup-latency; 'db-mem': db-mem; 'rtu-cache-mem': rtu-cache-mem; 'lookup-cache-mem': lookup-cache-mem;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceWebCategoryStatisticsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryStatisticsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryStatistics(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebCategoryStatisticsRead(ctx, d, meta)
	}
	return diags
}

func resourceWebCategoryStatisticsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryStatisticsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryStatistics(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebCategoryStatisticsRead(ctx, d, meta)
	}
	return diags
}
func resourceWebCategoryStatisticsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryStatisticsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryStatistics(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceWebCategoryStatisticsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryStatisticsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryStatistics(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceWebCategoryStatisticsSamplingEnable(d []interface{}) []edpt.WebCategoryStatisticsSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.WebCategoryStatisticsSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryStatisticsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointWebCategoryStatistics(d *schema.ResourceData) edpt.WebCategoryStatistics {
	var ret edpt.WebCategoryStatistics
	ret.Inst.SamplingEnable = getSliceWebCategoryStatisticsSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
