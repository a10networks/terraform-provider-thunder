package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileWafOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_waf_oper`: Operational Status for the object waf\n\n__PLACEHOLDER__",
		CreateContext: resourceFileWafOperCreate,
		UpdateContext: resourceFileWafOperUpdate,
		ReadContext:   resourceFileWafOperRead,
		DeleteContext: resourceFileWafOperDelete,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
func resourceFileWafOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileWafOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileWafOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileWafOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileWafOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileWafOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileWafOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileWafOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileWafOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileWafOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileWafOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileWafOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileWafOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileWafOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileWafOperOper(d []interface{}) edpt.FileWafOperOper {

	count1 := len(d)
	var ret edpt.FileWafOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileWafOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileWafOperOperFileList(d []interface{}) []edpt.FileWafOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileWafOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileWafOperOperFileList
		oi.File = in["file"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileWafOper(d *schema.ResourceData) edpt.FileWafOper {
	var ret edpt.FileWafOper
	ret.Inst.Oper = getObjectFileWafOperOper(d.Get("oper").([]interface{}))
	return ret
}
