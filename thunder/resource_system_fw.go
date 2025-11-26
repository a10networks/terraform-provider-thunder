package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemFw() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_fw`: Configure system parameters for firewall\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemFwCreate,
		UpdateContext: resourceSystemFwUpdate,
		ReadContext:   resourceSystemFwRead,
		DeleteContext: resourceSystemFwDelete,

		Schema: map[string]*schema.Schema{
			"application_flow": {
				Type: schema.TypeInt, Optional: true, Description: "Number of flows",
			},
			"application_mempool": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable application memory pool",
			},
			"basic_dpi_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable basic dpi",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemFwCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemFwCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemFw(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemFwRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemFwUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemFwUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemFw(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemFwRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemFwDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemFwDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemFw(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemFwRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemFwRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemFw(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemFw(d *schema.ResourceData) edpt.SystemFw {
	var ret edpt.SystemFw
	ret.Inst.ApplicationFlow = d.Get("application_flow").(int)
	ret.Inst.ApplicationMempool = d.Get("application_mempool").(int)
	ret.Inst.BasicDpiEnable = d.Get("basic_dpi_enable").(int)
	//omit uuid
	return ret
}
