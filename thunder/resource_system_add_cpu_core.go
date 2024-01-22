package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAddCpuCore() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_add_cpu_core`: Add CPU cores\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemAddCpuCoreCreate,
		UpdateContext: resourceSystemAddCpuCoreUpdate,
		ReadContext:   resourceSystemAddCpuCoreRead,
		DeleteContext: resourceSystemAddCpuCoreDelete,

		Schema: map[string]*schema.Schema{
			"core_index": {
				Type: schema.TypeInt, Optional: true, Description: "core index to be added (Specify core index)",
			},
		},
	}
}
func resourceSystemAddCpuCoreCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAddCpuCoreCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAddCpuCore(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemAddCpuCoreRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemAddCpuCoreUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAddCpuCoreUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAddCpuCore(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemAddCpuCoreRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemAddCpuCoreDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAddCpuCoreDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAddCpuCore(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemAddCpuCoreRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAddCpuCoreRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAddCpuCore(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemAddCpuCore(d *schema.ResourceData) edpt.SystemAddCpuCore {
	var ret edpt.SystemAddCpuCore
	ret.Inst.CoreIndex = d.Get("core_index").(int)
	return ret
}
