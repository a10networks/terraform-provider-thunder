package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileHealthExternalOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_health_external_oper`: Operational Status for the object health-external\n\n__PLACEHOLDER__",
		CreateContext: resourceFileHealthExternalOperCreate,
		UpdateContext: resourceFileHealthExternalOperUpdate,
		ReadContext:   resourceFileHealthExternalOperRead,
		DeleteContext: resourceFileHealthExternalOperDelete,

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
									"desc": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"partition": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"all_partitions": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"content": {
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
func resourceFileHealthExternalOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileHealthExternalOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileHealthExternalOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileHealthExternalOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileHealthExternalOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileHealthExternalOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileHealthExternalOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileHealthExternalOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileHealthExternalOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileHealthExternalOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileHealthExternalOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileHealthExternalOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileHealthExternalOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileHealthExternalOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileHealthExternalOperOper(d []interface{}) edpt.FileHealthExternalOperOper {

	count1 := len(d)
	var ret edpt.FileHealthExternalOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileHealthExternalOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileHealthExternalOperOperFileList(d []interface{}) []edpt.FileHealthExternalOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileHealthExternalOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileHealthExternalOperOperFileList
		oi.File = in["file"].(string)
		oi.Desc = in["desc"].(string)
		oi.Partition = in["partition"].(string)
		oi.AllPartitions = in["all_partitions"].(int)
		oi.Content = in["content"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileHealthExternalOper(d *schema.ResourceData) edpt.FileHealthExternalOper {
	var ret edpt.FileHealthExternalOper
	ret.Inst.Oper = getObjectFileHealthExternalOperOper(d.Get("oper").([]interface{}))
	return ret
}
