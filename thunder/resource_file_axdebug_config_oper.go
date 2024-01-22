package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileAxdebugConfigOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_axdebug_config_oper`: Operational Status for the object axdebug-config\n\n__PLACEHOLDER__",
		CreateContext: resourceFileAxdebugConfigOperCreate,
		UpdateContext: resourceFileAxdebugConfigOperUpdate,
		ReadContext:   resourceFileAxdebugConfigOperRead,
		DeleteContext: resourceFileAxdebugConfigOperDelete,

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
func resourceFileAxdebugConfigOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAxdebugConfigOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAxdebugConfigOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAxdebugConfigOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileAxdebugConfigOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAxdebugConfigOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAxdebugConfigOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAxdebugConfigOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileAxdebugConfigOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAxdebugConfigOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAxdebugConfigOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileAxdebugConfigOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAxdebugConfigOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAxdebugConfigOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileAxdebugConfigOperOper(d []interface{}) edpt.FileAxdebugConfigOperOper {

	count1 := len(d)
	var ret edpt.FileAxdebugConfigOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileAxdebugConfigOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileAxdebugConfigOperOperFileList(d []interface{}) []edpt.FileAxdebugConfigOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileAxdebugConfigOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileAxdebugConfigOperOperFileList
		oi.File = in["file"].(string)
		oi.UpdateTime = in["update_time"].(string)
		oi.Size = in["size"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileAxdebugConfigOper(d *schema.ResourceData) edpt.FileAxdebugConfigOper {
	var ret edpt.FileAxdebugConfigOper
	ret.Inst.Oper = getObjectFileAxdebugConfigOperOper(d.Get("oper").([]interface{}))
	return ret
}
