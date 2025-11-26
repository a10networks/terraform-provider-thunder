package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileTemplateOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_template_oper`: Operational Status for the object template\n\n__PLACEHOLDER__",
		CreateContext: resourceFileTemplateOperCreate,
		UpdateContext: resourceFileTemplateOperUpdate,
		ReadContext:   resourceFileTemplateOperRead,
		DeleteContext: resourceFileTemplateOperDelete,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"single_act_upgrade_support": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"file_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"file": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
func resourceFileTemplateOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTemplateOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTemplateOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileTemplateOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileTemplateOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTemplateOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTemplateOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileTemplateOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileTemplateOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTemplateOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTemplateOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileTemplateOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTemplateOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTemplateOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileTemplateOperOper(d []interface{}) edpt.FileTemplateOperOper {

	count1 := len(d)
	var ret edpt.FileTemplateOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SingleActUpgradeSupport = in["single_act_upgrade_support"].(string)
		ret.FileList = getSliceFileTemplateOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileTemplateOperOperFileList(d []interface{}) []edpt.FileTemplateOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileTemplateOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileTemplateOperOperFileList
		oi.File = in["file"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileTemplateOper(d *schema.ResourceData) edpt.FileTemplateOper {
	var ret edpt.FileTemplateOper
	ret.Inst.Oper = getObjectFileTemplateOperOper(d.Get("oper").([]interface{}))
	return ret
}
