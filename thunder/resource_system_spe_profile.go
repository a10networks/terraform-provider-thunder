package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSpeProfile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_spe_profile`: Set Security Policy Engine Profile\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemSpeProfileCreate,
		UpdateContext: resourceSystemSpeProfileUpdate,
		ReadContext:   resourceSystemSpeProfileRead,
		DeleteContext: resourceSystemSpeProfileDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "ipv4-ipv6", Description: "'ipv4-only': Enable IPv4 HW forward entries only; 'ipv6-only': Enable IPv6 HW forward entries only; 'ipv4-ipv6': Enable Both IPv4/IPv6 HW forward entries (shared);",
			},
		},
	}
}
func resourceSystemSpeProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSpeProfileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSpeProfile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSpeProfileRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemSpeProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSpeProfileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSpeProfile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSpeProfileRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemSpeProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSpeProfileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSpeProfile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemSpeProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSpeProfileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSpeProfile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemSpeProfile(d *schema.ResourceData) edpt.SystemSpeProfile {
	var ret edpt.SystemSpeProfile
	ret.Inst.Action = d.Get("action").(string)
	return ret
}
