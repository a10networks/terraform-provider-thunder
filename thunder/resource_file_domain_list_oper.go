package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileDomainListOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_domain_list_oper`: Operational Status for the object domain-list\n\n__PLACEHOLDER__",
		CreateContext: resourceFileDomainListOperCreate,
		UpdateContext: resourceFileDomainListOperUpdate,
		ReadContext:   resourceFileDomainListOperRead,
		DeleteContext: resourceFileDomainListOperDelete,

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
									"type_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"total_entry_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"binding_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"displayed_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}
func resourceFileDomainListOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileDomainListOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileDomainListOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileDomainListOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileDomainListOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileDomainListOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileDomainListOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileDomainListOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileDomainListOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileDomainListOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileDomainListOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileDomainListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileDomainListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileDomainListOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileDomainListOperOper(d []interface{}) edpt.FileDomainListOperOper {

	count1 := len(d)
	var ret edpt.FileDomainListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileDomainListOperOperFileList(in["file_list"].([]interface{}))
		ret.DisplayedCount = in["displayed_count"].(int)
	}
	return ret
}

func getSliceFileDomainListOperOperFileList(d []interface{}) []edpt.FileDomainListOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileDomainListOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileDomainListOperOperFileList
		oi.File = in["file"].(string)
		oi.TypeStr = in["type_str"].(string)
		oi.TotalEntryNum = in["total_entry_num"].(int)
		oi.BindingNum = in["binding_num"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileDomainListOper(d *schema.ResourceData) edpt.FileDomainListOper {
	var ret edpt.FileDomainListOper
	ret.Inst.Oper = getObjectFileDomainListOperOper(d.Get("oper").([]interface{}))
	return ret
}
