package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnablePassword() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_enable_password`: Modify enable password parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceEnablePasswordCreate,
		UpdateContext: resourceEnablePasswordUpdate,
		ReadContext:   resourceEnablePasswordRead,
		DeleteContext: resourceEnablePasswordDelete,

		Schema: map[string]*schema.Schema{
			"encrypted": {
				Type: schema.TypeString, Optional: true, Description: "Specific an ENCRYPTED password string (The ENCRYPTED password string)",
			},
			"password": {
				Type: schema.TypeString, Optional: true, Description: "The password",
			},
		},
	}
}
func resourceEnablePasswordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnablePasswordCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnablePassword(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnablePasswordRead(ctx, d, meta)
	}
	return diags
}

func resourceEnablePasswordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnablePasswordUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnablePassword(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnablePasswordRead(ctx, d, meta)
	}
	return diags
}
func resourceEnablePasswordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnablePasswordDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnablePassword(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceEnablePasswordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnablePasswordRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnablePassword(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointEnablePassword(d *schema.ResourceData) edpt.EnablePassword {
	var ret edpt.EnablePassword
	ret.Inst.Encrypted = d.Get("encrypted").(string)
	ret.Inst.Password = d.Get("password").(string)
	return ret
}
