package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileDnssecDnskeyOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_dnssec_dnskey_oper`: Operational Status for the object dnssec-dnskey\n\n__PLACEHOLDER__",
		CreateContext: resourceFileDnssecDnskeyOperCreate,
		UpdateContext: resourceFileDnssecDnskeyOperUpdate,
		ReadContext:   resourceFileDnssecDnskeyOperRead,
		DeleteContext: resourceFileDnssecDnskeyOperDelete,

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
									"parent": {
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
func resourceFileDnssecDnskeyOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileDnssecDnskeyOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileDnssecDnskeyOper(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileDnssecDnskeyOperRead(ctx, d, meta)
	}
	return diags
}

func resourceFileDnssecDnskeyOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileDnssecDnskeyOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileDnssecDnskeyOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileDnssecDnskeyOperRead(ctx, d, meta)
	}
	return diags
}
func resourceFileDnssecDnskeyOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileDnssecDnskeyOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileDnssecDnskeyOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileDnssecDnskeyOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileDnssecDnskeyOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileDnssecDnskeyOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFileDnssecDnskeyOperOper(d []interface{}) edpt.FileDnssecDnskeyOperOper {

	count1 := len(d)
	var ret edpt.FileDnssecDnskeyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceFileDnssecDnskeyOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceFileDnssecDnskeyOperOperFileList(d []interface{}) []edpt.FileDnssecDnskeyOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.FileDnssecDnskeyOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FileDnssecDnskeyOperOperFileList
		oi.File = in["file"].(string)
		oi.Parent = in["parent"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFileDnssecDnskeyOper(d *schema.ResourceData) edpt.FileDnssecDnskeyOper {
	var ret edpt.FileDnssecDnskeyOper
	ret.Inst.Oper = getObjectFileDnssecDnskeyOperOper(d.Get("oper").([]interface{}))
	return ret
}
