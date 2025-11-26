package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileSystemA10logOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_system_a10log_oper`: Operational Status for the object a10log\n\n__PLACEHOLDER__",
		CreateContext: resourceFileSystemA10logOperCreate,
		UpdateContext: resourceFileSystemA10logOperUpdate,
		ReadContext:   resourceFileSystemA10logOperRead,
		DeleteContext: resourceFileSystemA10logOperDelete,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resp": {
							Type: schema.TypeList, Optional: true, Description: "",
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
func resourceFileSystemA10logOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemA10logOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemA10logOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSystemA10logOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileSystemA10logOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemA10logOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemA10logOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSystemA10logOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileSystemA10logOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemA10logOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemA10logOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileSystemA10logOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemA10logOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemA10logOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileSystemA10logOperOper(d []interface{}) edpt.FileSystemA10logOperOper {

	count1 := len(d)
	var ret edpt.FileSystemA10logOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Resp = getSliceFileSystemA10logOperOperResp(in["resp"].([]interface{}))
	}
	return ret
}

func getSliceFileSystemA10logOperOperResp(d []interface{}) []edpt.FileSystemA10logOperOperResp {

	count1 := len(d)
	ret := make([]edpt.FileSystemA10logOperOperResp, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileSystemA10logOperOperResp
		oi.Message = in["message"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileSystemA10logOper(d *schema.ResourceData) edpt.FileSystemA10logOper {
	var ret edpt.FileSystemA10logOper
	ret.Inst.Oper = getObjectFileSystemA10logOperOper(d.Get("oper").([]interface{}))
	return ret
}
