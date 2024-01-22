package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateIcmpV4TypeOther() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_icmp_v4_type_other`: Specify other type\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateIcmpV4TypeOtherCreate,
		UpdateContext: resourceDdosTemplateIcmpV4TypeOtherUpdate,
		ReadContext:   resourceDdosTemplateIcmpV4TypeOtherRead,
		DeleteContext: resourceDdosTemplateIcmpV4TypeOtherDelete,

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
func resourceDdosTemplateIcmpV4TypeOtherCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV4TypeOtherCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV4TypeOther(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateIcmpV4TypeOtherRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateIcmpV4TypeOtherUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV4TypeOtherUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV4TypeOther(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateIcmpV4TypeOtherRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateIcmpV4TypeOtherDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV4TypeOtherDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV4TypeOther(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateIcmpV4TypeOtherRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV4TypeOtherRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV4TypeOther(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTemplateIcmpV4TypeOther(d *schema.ResourceData) edpt.DdosTemplateIcmpV4TypeOther {
	var ret edpt.DdosTemplateIcmpV4TypeOther
	ret.Inst.TypeOtherDeny = d.Get("type_other_deny").(int)
	ret.Inst.TypeOtherRate = d.Get("type_other_rate").(int)
	//omit uuid
	ret.Inst.IcmpTmplName = d.Get("icmp_tmpl_name").(string)
	return ret
}
