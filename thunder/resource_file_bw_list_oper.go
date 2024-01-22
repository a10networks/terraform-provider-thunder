package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileBwListOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_bw_list_oper`: Operational Status for the object bw-list\n\n__PLACEHOLDER__",
		CreateContext: resourceFileBwListOperCreate,
		UpdateContext: resourceFileBwListOperUpdate,
		ReadContext:   resourceFileBwListOperRead,
		DeleteContext: resourceFileBwListOperDelete,

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
									"url": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ref_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"period": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"update_times": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"parse_errors": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"error_line": {
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
func resourceFileBwListOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileBwListOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileBwListOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileBwListOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileBwListOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileBwListOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileBwListOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileBwListOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileBwListOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileBwListOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileBwListOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileBwListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileBwListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileBwListOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileBwListOperOper(d []interface{}) edpt.FileBwListOperOper {

	count1 := len(d)
	var ret edpt.FileBwListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileBwListOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileBwListOperOperFileList(d []interface{}) []edpt.FileBwListOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileBwListOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileBwListOperOperFileList
		oi.File = in["file"].(string)
		oi.Url = in["url"].(string)
		oi.RefCount = in["ref_count"].(int)
		oi.Period = in["period"].(int)
		oi.UpdateTimes = in["update_times"].(int)
		oi.ParseErrors = in["parse_errors"].(int)
		oi.ErrorLine = in["error_line"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileBwListOper(d *schema.ResourceData) edpt.FileBwListOper {
	var ret edpt.FileBwListOper
	ret.Inst.Oper = getObjectFileBwListOperOper(d.Get("oper").([]interface{}))
	return ret
}
