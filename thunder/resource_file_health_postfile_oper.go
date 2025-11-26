package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileHealthPostfileOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_health_postfile_oper`: Operational Status for the object health-postfile\n\n__PLACEHOLDER__",
		CreateContext: resourceFileHealthPostfileOperCreate,
		UpdateContext: resourceFileHealthPostfileOperUpdate,
		ReadContext:   resourceFileHealthPostfileOperRead,
		DeleteContext: resourceFileHealthPostfileOperDelete,

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
func resourceFileHealthPostfileOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileHealthPostfileOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileHealthPostfileOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileHealthPostfileOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileHealthPostfileOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileHealthPostfileOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileHealthPostfileOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileHealthPostfileOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileHealthPostfileOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileHealthPostfileOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileHealthPostfileOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileHealthPostfileOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileHealthPostfileOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileHealthPostfileOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileHealthPostfileOperOper(d []interface{}) edpt.FileHealthPostfileOperOper {

	count1 := len(d)
	var ret edpt.FileHealthPostfileOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileHealthPostfileOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileHealthPostfileOperOperFileList(d []interface{}) []edpt.FileHealthPostfileOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileHealthPostfileOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileHealthPostfileOperOperFileList
		oi.File = in["file"].(string)
		oi.Partition = in["partition"].(string)
		oi.AllPartitions = in["all_partitions"].(int)
		oi.Content = in["content"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileHealthPostfileOper(d *schema.ResourceData) edpt.FileHealthPostfileOper {
	var ret edpt.FileHealthPostfileOper
	ret.Inst.Oper = getObjectFileHealthPostfileOperOper(d.Get("oper").([]interface{}))
	return ret
}
