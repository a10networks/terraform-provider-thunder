package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAdminDetailOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_admin_detail_oper`: Operational Status for the object admin-detail\n\n__PLACEHOLDER__",
		ReadContext: resourceAdminDetailOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"priviledge": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"partition": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"access_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"gui_role": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"trusted_host": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"lock_status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"lock_time": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"unlock_time": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"password_type": {
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

func resourceAdminDetailOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminDetailOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminDetailOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AdminDetailOperOper := setObjectAdminDetailOperOper(res)
		d.Set("oper", AdminDetailOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAdminDetailOperOper(ret edpt.DataAdminDetailOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"user_list": setSliceAdminDetailOperOperUserList(ret.DtAdminDetailOper.Oper.UserList),
		},
	}
}

func setSliceAdminDetailOperOperUserList(d []edpt.AdminDetailOperOperUserList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["user_name"] = item.User_name
		in["status"] = item.Status
		in["priviledge"] = item.Priviledge
		in["partition"] = item.Partition
		in["access_type"] = item.Access_type
		in["gui_role"] = item.Gui_role
		in["trusted_host"] = item.Trusted_host
		in["lock_status"] = item.Lock_status
		in["lock_time"] = item.Lock_time
		in["unlock_time"] = item.Unlock_time
		in["password_type"] = item.Password_type
		result = append(result, in)
	}
	return result
}

func getObjectAdminDetailOperOper(d []interface{}) edpt.AdminDetailOperOper {

	count1 := len(d)
	var ret edpt.AdminDetailOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UserList = getSliceAdminDetailOperOperUserList(in["user_list"].([]interface{}))
	}
	return ret
}

func getSliceAdminDetailOperOperUserList(d []interface{}) []edpt.AdminDetailOperOperUserList {

	count1 := len(d)
	ret := make([]edpt.AdminDetailOperOperUserList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AdminDetailOperOperUserList
		oi.User_name = in["user_name"].(string)
		oi.Status = in["status"].(string)
		oi.Priviledge = in["priviledge"].(string)
		oi.Partition = in["partition"].(string)
		oi.Access_type = in["access_type"].(string)
		oi.Gui_role = in["gui_role"].(string)
		oi.Trusted_host = in["trusted_host"].(string)
		oi.Lock_status = in["lock_status"].(string)
		oi.Lock_time = in["lock_time"].(string)
		oi.Unlock_time = in["unlock_time"].(string)
		oi.Password_type = in["password_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAdminDetailOper(d *schema.ResourceData) edpt.AdminDetailOper {
	var ret edpt.AdminDetailOper

	ret.Oper = getObjectAdminDetailOperOper(d.Get("oper").([]interface{}))
	return ret
}
