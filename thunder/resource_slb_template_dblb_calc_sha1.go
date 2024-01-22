package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDblbCalcSha1() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dblb_calc_sha1`: Calculate sha1 value of a cleartext password\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDblbCalcSha1Create,
		UpdateContext: resourceSlbTemplateDblbCalcSha1Update,
		ReadContext:   resourceSlbTemplateDblbCalcSha1Read,
		DeleteContext: resourceSlbTemplateDblbCalcSha1Delete,

		Schema: map[string]*schema.Schema{
			"sha1_value": {
				Type: schema.TypeString, Optional: true, Description: "Cleartext password",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbTemplateDblbCalcSha1Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDblbCalcSha1Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDblbCalcSha1(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDblbCalcSha1Read(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDblbCalcSha1Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDblbCalcSha1Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDblbCalcSha1(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDblbCalcSha1Read(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDblbCalcSha1Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDblbCalcSha1Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDblbCalcSha1(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDblbCalcSha1Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDblbCalcSha1Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDblbCalcSha1(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateDblbCalcSha1(d *schema.ResourceData) edpt.SlbTemplateDblbCalcSha1 {
	var ret edpt.SlbTemplateDblbCalcSha1
	ret.Inst.Sha1Value = d.Get("sha1_value").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
