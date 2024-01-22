package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugAuth() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_auth`: Debug authentication\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugAuthCreate,
		UpdateContext: resourceDebugAuthUpdate,
		ReadContext:   resourceDebugAuthRead,
		DeleteContext: resourceDebugAuthDelete,

		Schema: map[string]*schema.Schema{
			"authd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable detail authd log",
			},
			"client_addr": {
				Type: schema.TypeString, Optional: true, Description: "Filter SAML logs by client IP address",
			},
			"level": {
				Type: schema.TypeString, Optional: true, Description: "'1': Diagnose Problems; '2': Detail packet flow;",
			},
			"saml": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable debug SAML authentication logs",
			},
			"saml_sp": {
				Type: schema.TypeString, Optional: true, Description: "Filter SAML logs by SAML service provider name (SAML SP name)",
			},
			"username": {
				Type: schema.TypeString, Optional: true, Description: "Show the logs of specific username (User name)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"virtual_server": {
				Type: schema.TypeString, Optional: true, Description: "Show the logs of specific virtual-server (Virtual-server name)",
			},
		},
	}
}
func resourceDebugAuthCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugAuthCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugAuth(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugAuthRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugAuthUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugAuthUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugAuth(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugAuthRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugAuthDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugAuthDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugAuth(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugAuthRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugAuthRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugAuth(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugAuth(d *schema.ResourceData) edpt.DebugAuth {
	var ret edpt.DebugAuth
	ret.Inst.Authd = d.Get("authd").(int)
	ret.Inst.ClientAddr = d.Get("client_addr").(string)
	ret.Inst.Level = d.Get("level").(string)
	ret.Inst.Saml = d.Get("saml").(int)
	ret.Inst.SamlSp = d.Get("saml_sp").(string)
	ret.Inst.Username = d.Get("username").(string)
	//omit uuid
	ret.Inst.VirtualServer = d.Get("virtual_server").(string)
	return ret
}
