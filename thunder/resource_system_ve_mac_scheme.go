package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemVeMacScheme() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ve_mac_scheme`: VE MAC allocation scheme\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemVeMacSchemeCreate,
		UpdateContext: resourceSystemVeMacSchemeUpdate,
		ReadContext:   resourceSystemVeMacSchemeRead,
		DeleteContext: resourceSystemVeMacSchemeDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ve_mac_scheme_val": {
				Type: schema.TypeString, Optional: true, Default: "hash-based", Description: "'hash-based': Hash-based using the VE number; 'round-robin': Round Robin scheme; 'system-mac': Use system MAC address;",
			},
		},
	}
}
func resourceSystemVeMacSchemeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemVeMacSchemeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemVeMacScheme(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemVeMacSchemeRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemVeMacSchemeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemVeMacSchemeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemVeMacScheme(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemVeMacSchemeRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemVeMacSchemeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemVeMacSchemeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemVeMacScheme(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemVeMacSchemeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemVeMacSchemeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemVeMacScheme(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemVeMacScheme(d *schema.ResourceData) edpt.SystemVeMacScheme {
	var ret edpt.SystemVeMacScheme
	//omit uuid
	ret.Inst.VeMacSchemeVal = d.Get("ve_mac_scheme_val").(string)
	return ret
}
