package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileSystemVarlogOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_system_varlog_oper`: Operational Status for the object varlog\n\n__PLACEHOLDER__",
		CreateContext: resourceFileSystemVarlogOperCreate,
		UpdateContext: resourceFileSystemVarlogOperUpdate,
		ReadContext:   resourceFileSystemVarlogOperRead,
		DeleteContext: resourceFileSystemVarlogOperDelete,

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
func resourceFileSystemVarlogOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemVarlogOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemVarlogOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSystemVarlogOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileSystemVarlogOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemVarlogOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemVarlogOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSystemVarlogOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileSystemVarlogOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemVarlogOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemVarlogOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileSystemVarlogOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemVarlogOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemVarlogOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileSystemVarlogOperOper(d []interface{}) edpt.FileSystemVarlogOperOper {

	count1 := len(d)
	var ret edpt.FileSystemVarlogOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Resp = getSliceFileSystemVarlogOperOperResp(in["resp"].([]interface{}))
	}
	return ret
}

func getSliceFileSystemVarlogOperOperResp(d []interface{}) []edpt.FileSystemVarlogOperOperResp {

	count1 := len(d)
	ret := make([]edpt.FileSystemVarlogOperOperResp, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileSystemVarlogOperOperResp
		oi.Message = in["message"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileSystemVarlogOper(d *schema.ResourceData) edpt.FileSystemVarlogOper {
	var ret edpt.FileSystemVarlogOper
	ret.Inst.Oper = getObjectFileSystemVarlogOperOper(d.Get("oper").([]interface{}))
	return ret
}
