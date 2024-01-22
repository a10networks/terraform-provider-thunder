package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateIcmpV6TypeOther() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_icmp_v6_type_other`: Specify other type\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateIcmpV6TypeOtherCreate,
		UpdateContext: resourceDdosTemplateIcmpV6TypeOtherUpdate,
		ReadContext:   resourceDdosTemplateIcmpV6TypeOtherRead,
		DeleteContext: resourceDdosTemplateIcmpV6TypeOtherDelete,

		Schema: map[string]*schema.Schema{
			"type_other_deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Deny all other type",
			},
			"type_other_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Specify rate with other type",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"icmp_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "IcmpTmplName",
			},
		},
	}
}
func resourceDdosTemplateIcmpV6TypeOtherCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV6TypeOtherCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV6TypeOther(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateIcmpV6TypeOtherRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateIcmpV6TypeOtherUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV6TypeOtherUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV6TypeOther(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateIcmpV6TypeOtherRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateIcmpV6TypeOtherDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV6TypeOtherDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV6TypeOther(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateIcmpV6TypeOtherRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV6TypeOtherRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV6TypeOther(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTemplateIcmpV6TypeOther(d *schema.ResourceData) edpt.DdosTemplateIcmpV6TypeOther {
	var ret edpt.DdosTemplateIcmpV6TypeOther
	ret.Inst.TypeOtherDeny = d.Get("type_other_deny").(int)
	ret.Inst.TypeOtherRate = d.Get("type_other_rate").(int)
	//omit uuid
	ret.Inst.IcmpTmplName = d.Get("icmp_tmpl_name").(string)
	return ret
}
