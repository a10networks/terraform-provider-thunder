package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtTemplateIgnoreValidation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_template_ignore_validation`: Ignore following layers for validation\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtTemplateIgnoreValidationCreate,
		UpdateContext: resourceSysUtTemplateIgnoreValidationUpdate,
		ReadContext:   resourceSysUtTemplateIgnoreValidationRead,
		DeleteContext: resourceSysUtTemplateIgnoreValidationDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Skip validation",
			},
			"l1": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dont validate TX descriptor. This includes Tx port, Len & vlan",
			},
			"l2": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dont validate L2 header",
			},
			"l3": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dont validate L3 header",
			},
			"l4": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dont validate L4 header",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSysUtTemplateIgnoreValidationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateIgnoreValidationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateIgnoreValidation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateIgnoreValidationRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtTemplateIgnoreValidationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateIgnoreValidationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateIgnoreValidation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateIgnoreValidationRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtTemplateIgnoreValidationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateIgnoreValidationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateIgnoreValidation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtTemplateIgnoreValidationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateIgnoreValidationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateIgnoreValidation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSysUtTemplateIgnoreValidation(d *schema.ResourceData) edpt.SysUtTemplateIgnoreValidation {
	var ret edpt.SysUtTemplateIgnoreValidation
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.L1 = d.Get("l1").(int)
	ret.Inst.L2 = d.Get("l2").(int)
	ret.Inst.L3 = d.Get("l3").(int)
	ret.Inst.L4 = d.Get("l4").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
