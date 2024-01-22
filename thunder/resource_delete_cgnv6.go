package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteCgnv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_cgnv6`: Delete cgnv6 related files\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteCgnv6Create,
		UpdateContext: resourceDeleteCgnv6Update,
		ReadContext:   resourceDeleteCgnv6Read,
		DeleteContext: resourceDeleteCgnv6Delete,

		Schema: map[string]*schema.Schema{
			"filename": {
				Type: schema.TypeString, Optional: true, Description: "Specify the port-mapping-file to be deleted",
			},
			"fixed_nat": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete fixed-nat port-mapping-file",
			},
			"lw_4o6_binding_table_validation_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "LW-4O6 Binding Table validation log File",
			},
			"lw_filename": {
				Type: schema.TypeString, Optional: true, Description: "LW-4O6 Binding Table Validation Log File Name",
			},
		},
	}
}
func resourceDeleteCgnv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteCgnv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteCgnv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteCgnv6Read(ctx, d, meta)
	}
	return diags
}

func resourceDeleteCgnv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteCgnv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteCgnv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteCgnv6Read(ctx, d, meta)
	}
	return diags
}
func resourceDeleteCgnv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteCgnv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteCgnv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteCgnv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteCgnv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteCgnv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteCgnv6(d *schema.ResourceData) edpt.DeleteCgnv6 {
	var ret edpt.DeleteCgnv6
	ret.Inst.Filename = d.Get("filename").(string)
	ret.Inst.FixedNat = d.Get("fixed_nat").(int)
	ret.Inst.Lw4o6BindingTableValidationLog = d.Get("lw_4o6_binding_table_validation_log").(int)
	ret.Inst.LwFilename = d.Get("lw_filename").(string)
	return ret
}
