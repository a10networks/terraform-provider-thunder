package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTokenAuthenticationSecretSalt() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_token_authentication_secret_salt`: Token Authentication Secret Salt\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTokenAuthenticationSecretSaltCreate,
		UpdateContext: resourceDdosTokenAuthenticationSecretSaltUpdate,
		ReadContext:   resourceDdosTokenAuthenticationSecretSaltRead,
		DeleteContext: resourceDdosTokenAuthenticationSecretSaltDelete,

		Schema: map[string]*schema.Schema{
			"current_salt": {
				Type: schema.TypeInt, Optional: true, Description: "Current salt value",
			},
			"previous_salt": {
				Type: schema.TypeInt, Optional: true, Description: "Previous salt value",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosTokenAuthenticationSecretSaltCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTokenAuthenticationSecretSaltCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTokenAuthenticationSecretSalt(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTokenAuthenticationSecretSaltRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTokenAuthenticationSecretSaltUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTokenAuthenticationSecretSaltUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTokenAuthenticationSecretSalt(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTokenAuthenticationSecretSaltRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTokenAuthenticationSecretSaltDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTokenAuthenticationSecretSaltDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTokenAuthenticationSecretSalt(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTokenAuthenticationSecretSaltRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTokenAuthenticationSecretSaltRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTokenAuthenticationSecretSalt(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTokenAuthenticationSecretSalt(d *schema.ResourceData) edpt.DdosTokenAuthenticationSecretSalt {
	var ret edpt.DdosTokenAuthenticationSecretSalt
	ret.Inst.CurrentSalt = d.Get("current_salt").(int)
	ret.Inst.PreviousSalt = d.Get("previous_salt").(int)
	//omit uuid
	return ret
}
