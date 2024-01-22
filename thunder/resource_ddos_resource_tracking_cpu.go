package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosResourceTrackingCpu() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_resource_tracking_cpu`: DDoS CPU Usage Per-Dst Tracking\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosResourceTrackingCpuCreate,
		UpdateContext: resourceDdosResourceTrackingCpuUpdate,
		ReadContext:   resourceDdosResourceTrackingCpuRead,
		DeleteContext: resourceDdosResourceTrackingCpuDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable CPU usage tracking per dst object (default: disabled)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosResourceTrackingCpuCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosResourceTrackingCpuCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosResourceTrackingCpu(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosResourceTrackingCpuRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosResourceTrackingCpuUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosResourceTrackingCpuUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosResourceTrackingCpu(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosResourceTrackingCpuRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosResourceTrackingCpuDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosResourceTrackingCpuDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosResourceTrackingCpu(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosResourceTrackingCpuRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosResourceTrackingCpuRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosResourceTrackingCpu(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosResourceTrackingCpu(d *schema.ResourceData) edpt.DdosResourceTrackingCpu {
	var ret edpt.DdosResourceTrackingCpu
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
