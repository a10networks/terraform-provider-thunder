package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemFpgaCoreCrc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_fpga_core_crc`: FPGA Core CRC check Global Commands\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemFpgaCoreCrcCreate,
		UpdateContext: resourceSystemFpgaCoreCrcUpdate,
		ReadContext:   resourceSystemFpgaCoreCrcRead,
		DeleteContext: resourceSystemFpgaCoreCrcDelete,

		Schema: map[string]*schema.Schema{
			"monitor_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable FPGA Core CRC error monitoring and act on it",
			},
			"reboot_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system reboot if system encounters FPGA Core CRC error",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemFpgaCoreCrcCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemFpgaCoreCrcCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemFpgaCoreCrc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemFpgaCoreCrcRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemFpgaCoreCrcUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemFpgaCoreCrcUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemFpgaCoreCrc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemFpgaCoreCrcRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemFpgaCoreCrcDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemFpgaCoreCrcDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemFpgaCoreCrc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemFpgaCoreCrcRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemFpgaCoreCrcRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemFpgaCoreCrc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemFpgaCoreCrc(d *schema.ResourceData) edpt.SystemFpgaCoreCrc {
	var ret edpt.SystemFpgaCoreCrc
	ret.Inst.MonitorDisable = d.Get("monitor_disable").(int)
	ret.Inst.RebootEnable = d.Get("reboot_enable").(int)
	//omit uuid
	return ret
}
