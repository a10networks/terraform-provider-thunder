package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileTechsupportOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_techsupport_oper`: Operational Status for the object techsupport\n\n__PLACEHOLDER__",
		CreateContext: resourceFileTechsupportOperCreate,
		UpdateContext: resourceFileTechsupportOperUpdate,
		ReadContext:   resourceFileTechsupportOperRead,
		DeleteContext: resourceFileTechsupportOperDelete,

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
									"update_time": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"status": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"message": {
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
func resourceFileTechsupportOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTechsupportOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTechsupportOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileTechsupportOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileTechsupportOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTechsupportOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTechsupportOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileTechsupportOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileTechsupportOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTechsupportOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTechsupportOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileTechsupportOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTechsupportOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTechsupportOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileTechsupportOperOper(d []interface{}) edpt.FileTechsupportOperOper {

	count1 := len(d)
	var ret edpt.FileTechsupportOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileTechsupportOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileTechsupportOperOperFileList(d []interface{}) []edpt.FileTechsupportOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileTechsupportOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileTechsupportOperOperFileList
		oi.File = in["file"].(string)
		oi.UpdateTime = in["update_time"].(string)
		oi.Size = in["size"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFileTechsupportOperStatus(d []interface{}) edpt.FileTechsupportOperStatus {

	count1 := len(d)
	var ret edpt.FileTechsupportOperStatus
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectFileTechsupportOperStatusOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectFileTechsupportOperStatusOper(d []interface{}) edpt.FileTechsupportOperStatusOper {

	count1 := len(d)
	var ret edpt.FileTechsupportOperStatusOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Message = in["message"].(string)
	}
	return ret
}

func dataToEndpointFileTechsupportOper(d *schema.ResourceData) edpt.FileTechsupportOper {
	var ret edpt.FileTechsupportOper
	ret.Inst.Oper = getObjectFileTechsupportOperOper(d.Get("oper").([]interface{}))
	ret.Inst.Status = getObjectFileTechsupportOperStatus(d.Get("status").([]interface{}))
	return ret
}
