package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCaptureConfigOperOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_capture_config_oper_oper`: Operational Status for the object capture-config-oper\n\n__PLACEHOLDER__",
		ReadContext: resourceCaptureConfigOperOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"filesize_kbyte": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pkt_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"max_pkt_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"has_file_history": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"max_filesize_kbyte": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"snaplen": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCaptureConfigOperOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCaptureConfigOperOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCaptureConfigOperOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		CaptureConfigOperOperOper := setObjectCaptureConfigOperOperOper(res)
		d.Set("oper", CaptureConfigOperOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCaptureConfigOperOperOper(ret edpt.DataCaptureConfigOperOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"name":               ret.DtCaptureConfigOperOper.Oper.Name,
			"action":             ret.DtCaptureConfigOperOper.Oper.Action,
			"status":             ret.DtCaptureConfigOperOper.Oper.Status,
			"filesize_kbyte":     ret.DtCaptureConfigOperOper.Oper.Filesize_kbyte,
			"pkt_count":          ret.DtCaptureConfigOperOper.Oper.Pkt_count,
			"max_pkt_count":      ret.DtCaptureConfigOperOper.Oper.Max_pkt_count,
			"has_file_history":   ret.DtCaptureConfigOperOper.Oper.Has_file_history,
			"max_filesize_kbyte": ret.DtCaptureConfigOperOper.Oper.Max_filesize_kbyte,
			"snaplen":            ret.DtCaptureConfigOperOper.Oper.Snaplen,
			"filter":             ret.DtCaptureConfigOperOper.Oper.Filter,
		},
	}
}

func getObjectCaptureConfigOperOperOper(d []interface{}) edpt.CaptureConfigOperOperOper {

	count1 := len(d)
	var ret edpt.CaptureConfigOperOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.Action = in["action"].(string)
		ret.Status = in["status"].(string)
		ret.Filesize_kbyte = in["filesize_kbyte"].(int)
		ret.Pkt_count = in["pkt_count"].(int)
		ret.Max_pkt_count = in["max_pkt_count"].(int)
		ret.Has_file_history = in["has_file_history"].(int)
		ret.Max_filesize_kbyte = in["max_filesize_kbyte"].(int)
		ret.Snaplen = in["snaplen"].(int)
		ret.Filter = in["filter"].(string)
	}
	return ret
}

func dataToEndpointCaptureConfigOperOper(d *schema.ResourceData) edpt.CaptureConfigOperOper {
	var ret edpt.CaptureConfigOperOper

	ret.Oper = getObjectCaptureConfigOperOperOper(d.Get("oper").([]interface{}))
	return ret
}
