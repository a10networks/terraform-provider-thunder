package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationSessionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_session_oper`: Operational Status for the object session\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationSessionOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"session_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vport": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"user": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"client_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"domain": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"domain_group": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"created_time": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ttl_in_seconds": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"token_lifetime": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"cmd_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"username": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"vserver": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"partition": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationSessionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSessionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSessionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationSessionOperOper := setObjectAamAuthenticationSessionOperOper(res)
		d.Set("oper", AamAuthenticationSessionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationSessionOperOper(ret edpt.DataAamAuthenticationSessionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_list": setSliceAamAuthenticationSessionOperOperSessionList(ret.DtAamAuthenticationSessionOper.Oper.SessionList),
			"cmd_type":     ret.DtAamAuthenticationSessionOper.Oper.CmdType,
			"username":     ret.DtAamAuthenticationSessionOper.Oper.Username,
			"vserver":      ret.DtAamAuthenticationSessionOper.Oper.Vserver,
			"ipv4":         ret.DtAamAuthenticationSessionOper.Oper.Ipv4,
			"ipv6":         ret.DtAamAuthenticationSessionOper.Oper.Ipv6,
			"partition":    ret.DtAamAuthenticationSessionOper.Oper.Partition,
		},
	}
}

func setSliceAamAuthenticationSessionOperOperSessionList(d []edpt.AamAuthenticationSessionOperOperSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["session_id"] = item.SessionId
		in["type"] = item.Type
		in["vip"] = item.Vip
		in["vport"] = item.Vport
		in["user"] = item.User
		in["client_ip"] = item.ClientIp
		in["domain"] = item.Domain
		in["domain_group"] = item.DomainGroup
		in["created_time"] = item.CreatedTime
		in["ttl_in_seconds"] = item.TtlInSeconds
		in["token_lifetime"] = item.TokenLifetime
		result = append(result, in)
	}
	return result
}

func getObjectAamAuthenticationSessionOperOper(d []interface{}) edpt.AamAuthenticationSessionOperOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationSessionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionList = getSliceAamAuthenticationSessionOperOperSessionList(in["session_list"].([]interface{}))
		ret.CmdType = in["cmd_type"].(string)
		ret.Username = in["username"].(string)
		ret.Vserver = in["vserver"].(string)
		ret.Ipv4 = in["ipv4"].(string)
		ret.Ipv6 = in["ipv6"].(string)
		ret.Partition = in["partition"].(string)
	}
	return ret
}

func getSliceAamAuthenticationSessionOperOperSessionList(d []interface{}) []edpt.AamAuthenticationSessionOperOperSessionList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationSessionOperOperSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationSessionOperOperSessionList
		oi.SessionId = in["session_id"].(int)
		oi.Type = in["type"].(string)
		oi.Vip = in["vip"].(string)
		oi.Vport = in["vport"].(string)
		oi.User = in["user"].(string)
		oi.ClientIp = in["client_ip"].(string)
		oi.Domain = in["domain"].(string)
		oi.DomainGroup = in["domain_group"].(string)
		oi.CreatedTime = in["created_time"].(string)
		oi.TtlInSeconds = in["ttl_in_seconds"].(int)
		oi.TokenLifetime = in["token_lifetime"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationSessionOper(d *schema.ResourceData) edpt.AamAuthenticationSessionOper {
	var ret edpt.AamAuthenticationSessionOper

	ret.Oper = getObjectAamAuthenticationSessionOperOper(d.Get("oper").([]interface{}))
	return ret
}
