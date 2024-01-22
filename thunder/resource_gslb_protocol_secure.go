package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbProtocolSecure() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_protocol_secure`: GSLB Secure operation\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbProtocolSecureCreate,
		UpdateContext: resourceGslbProtocolSecureUpdate,
		ReadContext:   resourceGslbProtocolSecureRead,
		DeleteContext: resourceGslbProtocolSecureDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Secure; 'disable': Disable Secure (default); 'enable-fallback': Fall back to non-secure if fail;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceGslbProtocolSecureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolSecureCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocolSecure(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbProtocolSecureRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbProtocolSecureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolSecureUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocolSecure(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbProtocolSecureRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbProtocolSecureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolSecureDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocolSecure(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbProtocolSecureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolSecureRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocolSecure(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbProtocolSecure(d *schema.ResourceData) edpt.GslbProtocolSecure {
	var ret edpt.GslbProtocolSecure
	ret.Inst.Action = d.Get("action").(string)
	//omit uuid
	return ret
}
