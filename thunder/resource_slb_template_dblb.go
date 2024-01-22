package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDblb() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dblb`: DBLB template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDblbCreate,
		UpdateContext: resourceSlbTemplateDblbUpdate,
		ReadContext:   resourceSlbTemplateDblbRead,
		DeleteContext: resourceSlbTemplateDblbDelete,

		Schema: map[string]*schema.Schema{
			"calc_sha1": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sha1_value": {
							Type: schema.TypeString, Optional: true, Description: "Cleartext password",
						},
					},
				},
			},
			"class_list": {
				Type: schema.TypeString, Optional: true, Description: "Specify user/password string class list (Class list name)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DBLB template name",
			},
			"server_version": {
				Type: schema.TypeString, Optional: true, Description: "'MSSQL2008': MSSQL server 2008 or 2008 R2; 'MSSQL2012': MSSQL server 2012; 'MySQL': MySQL server (any version);",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbTemplateDblbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDblbCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDblb(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDblbRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDblbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDblbUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDblb(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDblbRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDblbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDblbDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDblb(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDblbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDblbRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDblb(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbTemplateDblbCalcSha11415(d []interface{}) edpt.SlbTemplateDblbCalcSha11415 {

	count1 := len(d)
	var ret edpt.SlbTemplateDblbCalcSha11415
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Sha1Value = in["sha1_value"].(string)
	}
	return ret
}

func dataToEndpointSlbTemplateDblb(d *schema.ResourceData) edpt.SlbTemplateDblb {
	var ret edpt.SlbTemplateDblb
	ret.Inst.CalcSha1 = getObjectSlbTemplateDblbCalcSha11415(d.Get("calc_sha1").([]interface{}))
	ret.Inst.ClassList = d.Get("class_list").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.ServerVersion = d.Get("server_version").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
