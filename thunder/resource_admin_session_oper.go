package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAdminSessionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_admin_session_oper`: Operational Status for the object admin-session\n\n__PLACEHOLDER__",
		ReadContext: resourceAdminSessionOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sid": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"start_time": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"src_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"partition": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"authen": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"role": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cfg_mode": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"priv": {
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

func resourceAdminSessionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminSessionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminSessionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AdminSessionOperOper := setObjectAdminSessionOperOper(res)
		d.Set("oper", AdminSessionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAdminSessionOperOper(ret edpt.DataAdminSessionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_list": setSliceAdminSessionOperOperSessionList(ret.DtAdminSessionOper.Oper.SessionList),
		},
	}
}

func setSliceAdminSessionOperOperSessionList(d []edpt.AdminSessionOperOperSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["sid"] = item.Sid
		in["name"] = item.Name
		in["start_time"] = item.Start_time
		in["src_ip"] = item.Src_ip
		in["type"] = item.Type
		in["partition"] = item.Partition
		in["authen"] = item.Authen
		in["role"] = item.Role
		in["cfg_mode"] = item.Cfg_mode
		in["priv"] = item.Priv
		result = append(result, in)
	}
	return result
}

func getObjectAdminSessionOperOper(d []interface{}) edpt.AdminSessionOperOper {

	count1 := len(d)
	var ret edpt.AdminSessionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionList = getSliceAdminSessionOperOperSessionList(in["session_list"].([]interface{}))
	}
	return ret
}

func getSliceAdminSessionOperOperSessionList(d []interface{}) []edpt.AdminSessionOperOperSessionList {

	count1 := len(d)
	ret := make([]edpt.AdminSessionOperOperSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AdminSessionOperOperSessionList
		oi.Sid = in["sid"].(string)
		oi.Name = in["name"].(string)
		oi.Start_time = in["start_time"].(string)
		oi.Src_ip = in["src_ip"].(string)
		oi.Type = in["type"].(string)
		oi.Partition = in["partition"].(string)
		oi.Authen = in["authen"].(string)
		oi.Role = in["role"].(string)
		oi.Cfg_mode = in["cfg_mode"].(string)
		oi.Priv = in["priv"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAdminSessionOper(d *schema.ResourceData) edpt.AdminSessionOper {
	var ret edpt.AdminSessionOper

	ret.Oper = getObjectAdminSessionOperOper(d.Get("oper").([]interface{}))
	return ret
}
