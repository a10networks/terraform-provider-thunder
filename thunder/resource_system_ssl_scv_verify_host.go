package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSslScvVerifyHost() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ssl_scv_verify_host`: To Enable/disable checking claimed identity in server certificate during server certificate validation\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemSslScvVerifyHostCreate,
		UpdateContext: resourceSystemSslScvVerifyHostUpdate,
		ReadContext:   resourceSystemSslScvVerifyHostRead,
		DeleteContext: resourceSystemSslScvVerifyHostDelete,

		Schema: map[string]*schema.Schema{
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable verify host during SCV",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemSslScvVerifyHostCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslScvVerifyHostCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslScvVerifyHost(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSslScvVerifyHostRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemSslScvVerifyHostUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslScvVerifyHostUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslScvVerifyHost(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSslScvVerifyHostRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemSslScvVerifyHostDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslScvVerifyHostDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslScvVerifyHost(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemSslScvVerifyHostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslScvVerifyHostRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslScvVerifyHost(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemSslScvVerifyHost(d *schema.ResourceData) edpt.SystemSslScvVerifyHost {
	var ret edpt.SystemSslScvVerifyHost
	ret.Inst.Disable = d.Get("disable").(int)
	//omit uuid
	return ret
}
