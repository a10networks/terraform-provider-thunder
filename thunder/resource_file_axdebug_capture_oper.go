package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileAxdebugCaptureOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_axdebug_capture_oper`: Operational Status for the object axdebug-capture\n\n__PLACEHOLDER__",
		CreateContext: resourceFileAxdebugCaptureOperCreate,
		UpdateContext: resourceFileAxdebugCaptureOperUpdate,
		ReadContext:   resourceFileAxdebugCaptureOperRead,
		DeleteContext: resourceFileAxdebugCaptureOperDelete,

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
		},
	}
}
func resourceFileAxdebugCaptureOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAxdebugCaptureOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAxdebugCaptureOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAxdebugCaptureOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileAxdebugCaptureOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAxdebugCaptureOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAxdebugCaptureOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAxdebugCaptureOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileAxdebugCaptureOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAxdebugCaptureOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAxdebugCaptureOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileAxdebugCaptureOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAxdebugCaptureOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAxdebugCaptureOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileAxdebugCaptureOperOper(d []interface{}) edpt.FileAxdebugCaptureOperOper {

	count1 := len(d)
	var ret edpt.FileAxdebugCaptureOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileAxdebugCaptureOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileAxdebugCaptureOperOperFileList(d []interface{}) []edpt.FileAxdebugCaptureOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileAxdebugCaptureOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileAxdebugCaptureOperOperFileList
		oi.File = in["file"].(string)
		oi.UpdateTime = in["update_time"].(string)
		oi.Size = in["size"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileAxdebugCaptureOper(d *schema.ResourceData) edpt.FileAxdebugCaptureOper {
	var ret edpt.FileAxdebugCaptureOper
	ret.Inst.Oper = getObjectFileAxdebugCaptureOperOper(d.Get("oper").([]interface{}))
	return ret
}
