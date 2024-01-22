package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowMonitorSampleVe() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_netflow_monitor_sample_ve`: select VE interface to monitor\n\n__PLACEHOLDER__",
		CreateContext: resourceNetflowMonitorSampleVeCreate,
		UpdateContext: resourceNetflowMonitorSampleVeUpdate,
		ReadContext:   resourceNetflowMonitorSampleVeRead,
		DeleteContext: resourceNetflowMonitorSampleVeDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ve_num": {
				Type: schema.TypeInt, Required: true, Description: "VE interface number",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceNetflowMonitorSampleVeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorSampleVeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorSampleVe(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorSampleVeRead(ctx, d, meta)
	}
	return diags
}

func resourceNetflowMonitorSampleVeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorSampleVeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorSampleVe(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorSampleVeRead(ctx, d, meta)
	}
	return diags
}
func resourceNetflowMonitorSampleVeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorSampleVeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorSampleVe(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetflowMonitorSampleVeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorSampleVeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorSampleVe(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetflowMonitorSampleVe(d *schema.ResourceData) edpt.NetflowMonitorSampleVe {
	var ret edpt.NetflowMonitorSampleVe
	//omit uuid
	ret.Inst.VeNum = d.Get("ve_num").(int)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
