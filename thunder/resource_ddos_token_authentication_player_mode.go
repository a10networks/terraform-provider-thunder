package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTokenAuthenticationPlayerMode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_token_authentication_player_mode`: Token Authentication Player Mode\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTokenAuthenticationPlayerModeCreate,
		UpdateContext: resourceDdosTokenAuthenticationPlayerModeUpdate,
		ReadContext:   resourceDdosTokenAuthenticationPlayerModeRead,
		DeleteContext: resourceDdosTokenAuthenticationPlayerModeDelete,

		Schema: map[string]*schema.Schema{
			"mode": {
				Type: schema.TypeString, Optional: true, Default: "many-to-one", Description: "'one-to-one': Only one player talks to one server; 'many-to-one': Many player talk to one server;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosTokenAuthenticationPlayerModeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTokenAuthenticationPlayerModeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTokenAuthenticationPlayerMode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTokenAuthenticationPlayerModeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTokenAuthenticationPlayerModeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTokenAuthenticationPlayerModeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTokenAuthenticationPlayerMode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTokenAuthenticationPlayerModeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTokenAuthenticationPlayerModeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTokenAuthenticationPlayerModeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTokenAuthenticationPlayerMode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTokenAuthenticationPlayerModeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTokenAuthenticationPlayerModeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTokenAuthenticationPlayerMode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTokenAuthenticationPlayerMode(d *schema.ResourceData) edpt.DdosTokenAuthenticationPlayerMode {
	var ret edpt.DdosTokenAuthenticationPlayerMode
	ret.Inst.Mode = d.Get("mode").(string)
	//omit uuid
	return ret
}
