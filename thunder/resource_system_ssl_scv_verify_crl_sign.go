package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSslScvVerifyCrlSign() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ssl_scv_verify_crl_sign`: To Enable/disable verifying CRL signature during server certificate validation\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemSslScvVerifyCrlSignCreate,
		UpdateContext: resourceSystemSslScvVerifyCrlSignUpdate,
		ReadContext:   resourceSystemSslScvVerifyCrlSignRead,
		DeleteContext: resourceSystemSslScvVerifyCrlSignDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable verify CRL signature during SCV",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemSslScvVerifyCrlSignCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslScvVerifyCrlSignCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslScvVerifyCrlSign(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSslScvVerifyCrlSignRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemSslScvVerifyCrlSignUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslScvVerifyCrlSignUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslScvVerifyCrlSign(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSslScvVerifyCrlSignRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemSslScvVerifyCrlSignDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslScvVerifyCrlSignDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslScvVerifyCrlSign(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemSslScvVerifyCrlSignRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslScvVerifyCrlSignRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslScvVerifyCrlSign(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemSslScvVerifyCrlSign(d *schema.ResourceData) edpt.SystemSslScvVerifyCrlSign {
	var ret edpt.SystemSslScvVerifyCrlSign
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
