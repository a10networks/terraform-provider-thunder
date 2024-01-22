package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileStartupConfigOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_startup_config_oper`: Operational Status for the object startup-config\n\n__PLACEHOLDER__",
		CreateContext: resourceFileStartupConfigOperCreate,
		UpdateContext: resourceFileStartupConfigOperUpdate,
		ReadContext:   resourceFileStartupConfigOperRead,
		DeleteContext: resourceFileStartupConfigOperDelete,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dirty": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"current_startup_config": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pri_startup_config": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sec_startup_config": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"file_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"profile_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"update_time": {
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
func resourceFileStartupConfigOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileStartupConfigOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileStartupConfigOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileStartupConfigOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileStartupConfigOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileStartupConfigOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileStartupConfigOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileStartupConfigOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileStartupConfigOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileStartupConfigOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileStartupConfigOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileStartupConfigOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileStartupConfigOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileStartupConfigOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileStartupConfigOperOper(d []interface{}) edpt.FileStartupConfigOperOper {

	count1 := len(d)
	var ret edpt.FileStartupConfigOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dirty = in["dirty"].(int)
		ret.CurrentStartupConfig = in["current_startup_config"].(string)
		ret.PriStartupConfig = in["pri_startup_config"].(string)
		ret.SecStartupConfig = in["sec_startup_config"].(string)
		ret.FileList = getSliceFileStartupConfigOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileStartupConfigOperOperFileList(d []interface{}) []edpt.FileStartupConfigOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileStartupConfigOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileStartupConfigOperOperFileList
		oi.ProfileName = in["profile_name"].(string)
		oi.Size = in["size"].(int)
		oi.UpdateTime = in["update_time"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileStartupConfigOper(d *schema.ResourceData) edpt.FileStartupConfigOper {
	var ret edpt.FileStartupConfigOper
	ret.Inst.Oper = getObjectFileStartupConfigOperOper(d.Get("oper").([]interface{}))
	return ret
}
