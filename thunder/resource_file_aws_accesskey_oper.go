package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileAwsAccesskeyOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_aws_accesskey_oper`: Operational Status for the object aws-accesskey\n\n__PLACEHOLDER__",
		CreateContext: resourceFileAwsAccesskeyOperCreate,
		UpdateContext: resourceFileAwsAccesskeyOperUpdate,
		ReadContext:   resourceFileAwsAccesskeyOperRead,
		DeleteContext: resourceFileAwsAccesskeyOperDelete,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"file_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"index": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"content": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"comment": {
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
func resourceFileAwsAccesskeyOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAwsAccesskeyOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAwsAccesskeyOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAwsAccesskeyOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileAwsAccesskeyOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAwsAccesskeyOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAwsAccesskeyOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileAwsAccesskeyOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileAwsAccesskeyOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAwsAccesskeyOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAwsAccesskeyOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileAwsAccesskeyOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileAwsAccesskeyOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileAwsAccesskeyOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileAwsAccesskeyOperOper(d []interface{}) edpt.FileAwsAccesskeyOperOper {

	count1 := len(d)
	var ret edpt.FileAwsAccesskeyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileAwsAccesskeyOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileAwsAccesskeyOperOperFileList(d []interface{}) []edpt.FileAwsAccesskeyOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileAwsAccesskeyOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileAwsAccesskeyOperOperFileList
		oi.User = in["user"].(string)
		oi.Index = in["index"].(string)
		oi.Type = in["type"].(string)
		oi.Content = in["content"].(string)
		oi.Comment = in["comment"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileAwsAccesskeyOper(d *schema.ResourceData) edpt.FileAwsAccesskeyOper {
	var ret edpt.FileAwsAccesskeyOper
	ret.Inst.Oper = getObjectFileAwsAccesskeyOperOper(d.Get("oper").([]interface{}))
	return ret
}
