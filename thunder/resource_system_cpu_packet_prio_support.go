package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCpuPacketPrioSupport() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_cpu_packet_prio_support`: Enable/Disable CPU packet prioritization support\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemCpuPacketPrioSupportCreate,
		UpdateContext: resourceSystemCpuPacketPrioSupportUpdate,
		ReadContext:   resourceSystemCpuPacketPrioSupportRead,
		DeleteContext: resourceSystemCpuPacketPrioSupportDelete,

		Schema: map[string]*schema.Schema{
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable CPU packet prioritization Support",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable CPU packet prioritization Support",
			},
		},
	}
}
func resourceSystemCpuPacketPrioSupportCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuPacketPrioSupportCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuPacketPrioSupport(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemCpuPacketPrioSupportRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemCpuPacketPrioSupportUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuPacketPrioSupportUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuPacketPrioSupport(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemCpuPacketPrioSupportRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemCpuPacketPrioSupportDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuPacketPrioSupportDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuPacketPrioSupport(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemCpuPacketPrioSupportRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuPacketPrioSupportRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuPacketPrioSupport(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemCpuPacketPrioSupport(d *schema.ResourceData) edpt.SystemCpuPacketPrioSupport {
	var ret edpt.SystemCpuPacketPrioSupport
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	return ret
}
