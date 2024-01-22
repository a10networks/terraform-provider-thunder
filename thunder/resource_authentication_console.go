package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAuthenticationConsole() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_authentication_console`: Configure console authentication type\n\n__PLACEHOLDER__",
		CreateContext: resourceAuthenticationConsoleCreate,
		UpdateContext: resourceAuthenticationConsoleUpdate,
		ReadContext:   resourceAuthenticationConsoleRead,
		DeleteContext: resourceAuthenticationConsoleDelete,

		Schema: map[string]*schema.Schema{
			"type_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "The login authentication type",
						},
						"console_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAuthenticationConsoleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAuthenticationConsoleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAuthenticationConsole(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAuthenticationConsoleRead(ctx, d, meta)
	}
	return diags
}

func resourceAuthenticationConsoleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAuthenticationConsoleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAuthenticationConsole(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAuthenticationConsoleRead(ctx, d, meta)
	}
	return diags
}
func resourceAuthenticationConsoleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAuthenticationConsoleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAuthenticationConsole(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAuthenticationConsoleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAuthenticationConsoleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAuthenticationConsole(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAuthenticationConsoleTypeCfg(d []interface{}) edpt.AuthenticationConsoleTypeCfg {

	count1 := len(d)
	var ret edpt.AuthenticationConsoleTypeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Type = in["type"].(int)
		ret.ConsoleType = in["console_type"].(string)
	}
	return ret
}

func dataToEndpointAuthenticationConsole(d *schema.ResourceData) edpt.AuthenticationConsole {
	var ret edpt.AuthenticationConsole
	ret.Inst.TypeCfg = getObjectAuthenticationConsoleTypeCfg(d.Get("type_cfg").([]interface{}))
	//omit uuid
	return ret
}
