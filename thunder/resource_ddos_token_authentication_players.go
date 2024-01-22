package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTokenAuthenticationPlayers() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_token_authentication_players`: Token Authentication Player Mode\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTokenAuthenticationPlayersCreate,
		UpdateContext: resourceDdosTokenAuthenticationPlayersUpdate,
		ReadContext:   resourceDdosTokenAuthenticationPlayersRead,
		DeleteContext: resourceDdosTokenAuthenticationPlayersDelete,

		Schema: map[string]*schema.Schema{
			"dst_ip": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"dst_port": {
				Type: schema.TypeInt, Required: true, Description: "",
			},
			"magic_value": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"src_ip": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"src_port": {
				Type: schema.TypeInt, Required: true, Description: "",
			},
		},
	}
}
func resourceDdosTokenAuthenticationPlayersCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTokenAuthenticationPlayersCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTokenAuthenticationPlayers(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTokenAuthenticationPlayersRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTokenAuthenticationPlayersUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTokenAuthenticationPlayersUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTokenAuthenticationPlayers(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTokenAuthenticationPlayersRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTokenAuthenticationPlayersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTokenAuthenticationPlayersDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTokenAuthenticationPlayers(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTokenAuthenticationPlayersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTokenAuthenticationPlayersRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTokenAuthenticationPlayers(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTokenAuthenticationPlayers(d *schema.ResourceData) edpt.DdosTokenAuthenticationPlayers {
	var ret edpt.DdosTokenAuthenticationPlayers
	ret.Inst.DstIp = d.Get("dst_ip").(string)
	ret.Inst.DstPort = d.Get("dst_port").(int)
	ret.Inst.MagicValue = d.Get("magic_value").(int)
	ret.Inst.SrcIp = d.Get("src_ip").(string)
	ret.Inst.SrcPort = d.Get("src_port").(int)
	return ret
}
