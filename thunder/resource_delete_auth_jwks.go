package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteAuthJwks() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_auth_jwks`: JWK file\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteAuthJwksCreate,
		UpdateContext: resourceDeleteAuthJwksUpdate,
		ReadContext:   resourceDeleteAuthJwksRead,
		DeleteContext: resourceDeleteAuthJwksDelete,

		Schema: map[string]*schema.Schema{
			"jwk_name": {
				Type: schema.TypeString, Optional: true, Description: "Local File Name",
			},
		},
	}
}
func resourceDeleteAuthJwksCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteAuthJwksCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteAuthJwks(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteAuthJwksRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteAuthJwksUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteAuthJwksUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteAuthJwks(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteAuthJwksRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteAuthJwksDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteAuthJwksDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteAuthJwks(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteAuthJwksRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteAuthJwksRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteAuthJwks(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteAuthJwks(d *schema.ResourceData) edpt.DeleteAuthJwks {
	var ret edpt.DeleteAuthJwks
	ret.Inst.JwkName = d.Get("jwk_name").(string)
	return ret
}
