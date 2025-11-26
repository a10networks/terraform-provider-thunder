package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileSystemSyslogOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_system_syslog_oper`: Operational Status for the object syslog\n\n__PLACEHOLDER__",
		CreateContext: resourceFileSystemSyslogOperCreate,
		UpdateContext: resourceFileSystemSyslogOperUpdate,
		ReadContext:   resourceFileSystemSyslogOperRead,
		DeleteContext: resourceFileSystemSyslogOperDelete,

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
func resourceFileSystemSyslogOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemSyslogOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemSyslogOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSystemSyslogOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileSystemSyslogOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemSyslogOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemSyslogOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSystemSyslogOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileSystemSyslogOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemSyslogOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemSyslogOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileSystemSyslogOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemSyslogOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemSyslogOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileSystemSyslogOperOper(d []interface{}) edpt.FileSystemSyslogOperOper {

	count1 := len(d)
	var ret edpt.FileSystemSyslogOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Resp = getSliceFileSystemSyslogOperOperResp(in["resp"].([]interface{}))
	}
	return ret
}

func getSliceFileSystemSyslogOperOperResp(d []interface{}) []edpt.FileSystemSyslogOperOperResp {

	count1 := len(d)
	ret := make([]edpt.FileSystemSyslogOperOperResp, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileSystemSyslogOperOperResp
		oi.Message = in["message"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileSystemSyslogOper(d *schema.ResourceData) edpt.FileSystemSyslogOper {
	var ret edpt.FileSystemSyslogOper
	ret.Inst.Oper = getObjectFileSystemSyslogOperOper(d.Get("oper").([]interface{}))
	return ret
}
