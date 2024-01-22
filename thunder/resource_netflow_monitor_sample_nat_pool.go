package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowMonitorSampleNatPool() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_netflow_monitor_sample_nat_pool`: Select nat pool to monitor\n\n__PLACEHOLDER__",
		CreateContext: resourceNetflowMonitorSampleNatPoolCreate,
		UpdateContext: resourceNetflowMonitorSampleNatPoolUpdate,
		ReadContext:   resourceNetflowMonitorSampleNatPoolRead,
		DeleteContext: resourceNetflowMonitorSampleNatPoolDelete,

		Schema: map[string]*schema.Schema{
			"pool_name": {
				Type: schema.TypeString, Required: true, Description: "Name of nat pool",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceNetflowMonitorSampleNatPoolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorSampleNatPoolCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorSampleNatPool(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorSampleNatPoolRead(ctx, d, meta)
	}
	return diags
}

func resourceNetflowMonitorSampleNatPoolUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorSampleNatPoolUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorSampleNatPool(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorSampleNatPoolRead(ctx, d, meta)
	}
	return diags
}
func resourceNetflowMonitorSampleNatPoolDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorSampleNatPoolDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorSampleNatPool(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetflowMonitorSampleNatPoolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorSampleNatPoolRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorSampleNatPool(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetflowMonitorSampleNatPool(d *schema.ResourceData) edpt.NetflowMonitorSampleNatPool {
	var ret edpt.NetflowMonitorSampleNatPool
	ret.Inst.PoolName = d.Get("pool_name").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
