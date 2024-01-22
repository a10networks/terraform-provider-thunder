package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMfaValidationType() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_mfa_validation_type`: Define a 2FA management validation parameter\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemMfaValidationTypeCreate,
		UpdateContext: resourceSystemMfaValidationTypeUpdate,
		ReadContext:   resourceSystemMfaValidationTypeRead,
		DeleteContext: resourceSystemMfaValidationTypeDelete,

		Schema: map[string]*schema.Schema{
			"ca_cert": {
				Type: schema.TypeString, Optional: true, Description: "Configure CA Certificate",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemMfaValidationTypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMfaValidationTypeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMfaValidationType(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMfaValidationTypeRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemMfaValidationTypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMfaValidationTypeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMfaValidationType(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMfaValidationTypeRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemMfaValidationTypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMfaValidationTypeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMfaValidationType(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemMfaValidationTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMfaValidationTypeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMfaValidationType(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemMfaValidationType(d *schema.ResourceData) edpt.SystemMfaValidationType {
	var ret edpt.SystemMfaValidationType
	ret.Inst.CaCert = d.Get("ca_cert").(string)
	//omit uuid
	return ret
}
