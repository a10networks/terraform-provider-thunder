package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHealthGateway() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_health_gateway`: Configure gateway health-check\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbHealthGatewayCreate,
		UpdateContext: resourceSlbHealthGatewayUpdate,
		ReadContext:   resourceSlbHealthGatewayRead,
		DeleteContext: resourceSlbHealthGatewayDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'total_sent': Number of Total health-check sent; 'total_retry_sent': Number of Total health-check retry sent; 'total_timeout': Number of Total health-check timeout;",
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
func resourceSlbHealthGatewayCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthGatewayCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthGateway(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHealthGatewayRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbHealthGatewayUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthGatewayUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthGateway(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHealthGatewayRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbHealthGatewayDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthGatewayDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthGateway(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbHealthGatewayRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthGatewayRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthGateway(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbHealthGatewaySamplingEnable(d []interface{}) []edpt.SlbHealthGatewaySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbHealthGatewaySamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbHealthGatewaySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbHealthGateway(d *schema.ResourceData) edpt.SlbHealthGateway {
	var ret edpt.SlbHealthGateway
	ret.Inst.SamplingEnable = getSliceSlbHealthGatewaySamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
