package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNtpStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ntp_status_oper`: Operational Status for the object ntp-status\n\n__PLACEHOLDER__",
		ReadContext: resourceNtpStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ntp_server": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_preferred": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"mode": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"auth": {
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

func resourceNtpStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NtpStatusOperOper := setObjectNtpStatusOperOper(res)
		d.Set("oper", NtpStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNtpStatusOperOper(ret edpt.DataNtpStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"server": setSliceNtpStatusOperOperServer(ret.DtNtpStatusOper.Oper.Server),
		},
	}
}

func setSliceNtpStatusOperOperServer(d []edpt.NtpStatusOperOperServer) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ntp_server"] = item.NtpServer
		in["status"] = item.Status
		in["is_preferred"] = item.Is_preferred
		in["mode"] = item.Mode
		in["auth"] = item.Auth
		result = append(result, in)
	}
	return result
}

func getObjectNtpStatusOperOper(d []interface{}) edpt.NtpStatusOperOper {

	count1 := len(d)
	var ret edpt.NtpStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Server = getSliceNtpStatusOperOperServer(in["server"].([]interface{}))
	}
	return ret
}

func getSliceNtpStatusOperOperServer(d []interface{}) []edpt.NtpStatusOperOperServer {

	count1 := len(d)
	ret := make([]edpt.NtpStatusOperOperServer, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NtpStatusOperOperServer
		oi.NtpServer = in["ntp_server"].(string)
		oi.Status = in["status"].(string)
		oi.Is_preferred = in["is_preferred"].(string)
		oi.Mode = in["mode"].(string)
		oi.Auth = in["auth"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNtpStatusOper(d *schema.ResourceData) edpt.NtpStatusOper {
	var ret edpt.NtpStatusOper

	ret.Oper = getObjectNtpStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
