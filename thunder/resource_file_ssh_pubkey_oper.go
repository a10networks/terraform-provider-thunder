package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileSshPubkeyOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_ssh_pubkey_oper`: Operational Status for the object ssh-pubkey\n\n__PLACEHOLDER__",
		CreateContext: resourceFileSshPubkeyOperCreate,
		UpdateContext: resourceFileSshPubkeyOperUpdate,
		ReadContext:   resourceFileSshPubkeyOperRead,
		DeleteContext: resourceFileSshPubkeyOperDelete,

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
func resourceFileSshPubkeyOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSshPubkeyOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSshPubkeyOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSshPubkeyOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileSshPubkeyOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSshPubkeyOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSshPubkeyOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSshPubkeyOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileSshPubkeyOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSshPubkeyOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSshPubkeyOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileSshPubkeyOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSshPubkeyOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSshPubkeyOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileSshPubkeyOperOper(d []interface{}) edpt.FileSshPubkeyOperOper {

	count1 := len(d)
	var ret edpt.FileSshPubkeyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileSshPubkeyOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileSshPubkeyOperOperFileList(d []interface{}) []edpt.FileSshPubkeyOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileSshPubkeyOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileSshPubkeyOperOperFileList
		oi.User = in["user"].(string)
		oi.Index = in["index"].(string)
		oi.Type = in["type"].(string)
		oi.Content = in["content"].(string)
		oi.Comment = in["comment"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileSshPubkeyOper(d *schema.ResourceData) edpt.FileSshPubkeyOper {
	var ret edpt.FileSshPubkeyOper
	ret.Inst.Oper = getObjectFileSshPubkeyOperOper(d.Get("oper").([]interface{}))
	return ret
}
