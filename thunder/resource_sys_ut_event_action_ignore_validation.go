package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtEventActionIgnoreValidation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_event_action_ignore_validation`: Ignore following layers for validation\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtEventActionIgnoreValidationCreate,
		UpdateContext: resourceSysUtEventActionIgnoreValidationUpdate,
		ReadContext:   resourceSysUtEventActionIgnoreValidationRead,
		DeleteContext: resourceSysUtEventActionIgnoreValidationDelete,

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
			"event_number": {
				Type: schema.TypeString, Required: true, Description: "EventNumber",
			},
			"direction": {
				Type: schema.TypeString, Required: true, Description: "Direction",
			},
		},
	}
}
func resourceSysUtEventActionIgnoreValidationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionIgnoreValidationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionIgnoreValidation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtEventActionIgnoreValidationRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtEventActionIgnoreValidationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionIgnoreValidationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionIgnoreValidation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtEventActionIgnoreValidationRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtEventActionIgnoreValidationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionIgnoreValidationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionIgnoreValidation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtEventActionIgnoreValidationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionIgnoreValidationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionIgnoreValidation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSysUtEventActionIgnoreValidation(d *schema.ResourceData) edpt.SysUtEventActionIgnoreValidation {
	var ret edpt.SysUtEventActionIgnoreValidation
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.L1 = d.Get("l1").(int)
	ret.Inst.L2 = d.Get("l2").(int)
	ret.Inst.L3 = d.Get("l3").(int)
	ret.Inst.L4 = d.Get("l4").(int)
	//omit uuid
	ret.Inst.EventNumber = d.Get("event_number").(string)
	ret.Inst.Direction = d.Get("direction").(string)
	return ret
}
