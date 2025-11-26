package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMfaAuth() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_mfa_auth`: MFA authentication API\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemMfaAuthCreate,
		UpdateContext: resourceSystemMfaAuthUpdate,
		ReadContext:   resourceSystemMfaAuthRead,
		DeleteContext: resourceSystemMfaAuthDelete,

		Schema: map[string]*schema.Schema{
			"second_factor": {
				Type: schema.TypeString, Optional: true, Description: "Input second factor paramter",
			},
			"username": {
				Type: schema.TypeString, Optional: true, Description: "Username for MFA validation",
			},
		},
	}
}
func resourceSystemMfaAuthCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMfaAuthCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMfaAuth(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMfaAuthRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemMfaAuthUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMfaAuthUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMfaAuth(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMfaAuthRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemMfaAuthDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMfaAuthDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMfaAuth(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemMfaAuthRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMfaAuthRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMfaAuth(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemMfaAuth(d *schema.ResourceData) edpt.SystemMfaAuth {
	var ret edpt.SystemMfaAuth
	ret.Inst.SecondFactor = d.Get("second_factor").(string)
	ret.Inst.Username = d.Get("username").(string)
	return ret
}
