package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlmCreateLicenseRequest() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_glm_create_license_request`: Create a GLM trial or license request\n\n__PLACEHOLDER__",
		CreateContext: resourceGlmCreateLicenseRequestCreate,
		UpdateContext: resourceGlmCreateLicenseRequestUpdate,
		ReadContext:   resourceGlmCreateLicenseRequestRead,
		DeleteContext: resourceGlmCreateLicenseRequestDelete,

		Schema: map[string]*schema.Schema{
			"create_license_request": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Create a GLM trial or license request",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceGlmCreateLicenseRequestCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmCreateLicenseRequestCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmCreateLicenseRequest(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGlmCreateLicenseRequestRead(ctx, d, meta)
	}
	return diags
}

func resourceGlmCreateLicenseRequestUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmCreateLicenseRequestUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmCreateLicenseRequest(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGlmCreateLicenseRequestRead(ctx, d, meta)
	}
	return diags
}
func resourceGlmCreateLicenseRequestDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmCreateLicenseRequestDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmCreateLicenseRequest(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGlmCreateLicenseRequestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmCreateLicenseRequestRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmCreateLicenseRequest(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGlmCreateLicenseRequest(d *schema.ResourceData) edpt.GlmCreateLicenseRequest {
	var ret edpt.GlmCreateLicenseRequest
	ret.Inst.CreateLicenseRequest = d.Get("create_license_request").(int)
	//omit uuid
	return ret
}
