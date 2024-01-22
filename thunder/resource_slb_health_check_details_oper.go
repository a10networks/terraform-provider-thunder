package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHealthCheckDetailsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_health_check_details_oper`: Operational Status for the object health-check-details\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbHealthCheckDetailsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pin_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"process_index": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"health_state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"state_reason": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"monitor_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"received_success": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"received_fail": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"response_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"curr_interval": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"method": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"attr_alias_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"attr_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"half_open": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"send": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"resp_cont": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"force_up": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"url": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"expect_text": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"expect_resp_code": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"expect_text_regex": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"expect_resp_regex_code": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"maintenance_code": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"user": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pass": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"postdata": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"host": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"kerberos_realm": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"kerberos_kdc": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"kerberos_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"snmp_operation": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"community": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"oid": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"domain": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"starttls": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mail_from": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"rcpt_to": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipaddr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dns_qtype": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dns_recurse": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dns_expect_type": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dns_expect": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"transport_proto": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sip_register": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"secret": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"query": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"base_dn": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ldap_ssl": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ldap_tls": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"attr_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"db_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"receive": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"rcv_integer": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"db_row": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"db_column": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pname": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"tcp_only": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"attr_program": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"arguments": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"attr_rpn": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"http_wait_resp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"l4_conn_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"l4_errors": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"avg_rtt": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"curr_rtt": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"avg_tcp_rtt": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"curr_tcp_rtt": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"status_code_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"http_req_sent": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"http_errors": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mac_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbHealthCheckDetailsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthCheckDetailsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthCheckDetailsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbHealthCheckDetailsOperOper := setObjectSlbHealthCheckDetailsOperOper(res)
		d.Set("oper", SlbHealthCheckDetailsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHealthCheckDetailsOperOper(ret edpt.DataSlbHealthCheckDetailsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"pin_id":                 ret.DtSlbHealthCheckDetailsOper.Oper.PinId,
			"process_index":          ret.DtSlbHealthCheckDetailsOper.Oper.ProcessIndex,
			"health_state":           ret.DtSlbHealthCheckDetailsOper.Oper.HealthState,
			"state_reason":           ret.DtSlbHealthCheckDetailsOper.Oper.StateReason,
			"monitor_name":           ret.DtSlbHealthCheckDetailsOper.Oper.MonitorName,
			"received_success":       ret.DtSlbHealthCheckDetailsOper.Oper.ReceivedSuccess,
			"received_fail":          ret.DtSlbHealthCheckDetailsOper.Oper.ReceivedFail,
			"response_timeout":       ret.DtSlbHealthCheckDetailsOper.Oper.ResponseTimeout,
			"curr_interval":          ret.DtSlbHealthCheckDetailsOper.Oper.CurrInterval,
			"method":                 ret.DtSlbHealthCheckDetailsOper.Oper.Method,
			"attr_alias_addr":        ret.DtSlbHealthCheckDetailsOper.Oper.AttrAliasAddr,
			"attr_port":              ret.DtSlbHealthCheckDetailsOper.Oper.AttrPort,
			"half_open":              ret.DtSlbHealthCheckDetailsOper.Oper.HalfOpen,
			"send":                   ret.DtSlbHealthCheckDetailsOper.Oper.Send,
			"resp_cont":              ret.DtSlbHealthCheckDetailsOper.Oper.RespCont,
			"force_up":               ret.DtSlbHealthCheckDetailsOper.Oper.ForceUp,
			"url":                    ret.DtSlbHealthCheckDetailsOper.Oper.Url,
			"expect_text":            ret.DtSlbHealthCheckDetailsOper.Oper.ExpectText,
			"expect_resp_code":       ret.DtSlbHealthCheckDetailsOper.Oper.ExpectRespCode,
			"expect_text_regex":      ret.DtSlbHealthCheckDetailsOper.Oper.ExpectTextRegex,
			"expect_resp_regex_code": ret.DtSlbHealthCheckDetailsOper.Oper.ExpectRespRegexCode,
			"maintenance_code":       ret.DtSlbHealthCheckDetailsOper.Oper.MaintenanceCode,
			"user":                   ret.DtSlbHealthCheckDetailsOper.Oper.User,
			"pass":                   ret.DtSlbHealthCheckDetailsOper.Oper.Pass,
			"postdata":               ret.DtSlbHealthCheckDetailsOper.Oper.Postdata,
			"host":                   ret.DtSlbHealthCheckDetailsOper.Oper.Host,
			"kerberos_realm":         ret.DtSlbHealthCheckDetailsOper.Oper.KerberosRealm,
			"kerberos_kdc":           ret.DtSlbHealthCheckDetailsOper.Oper.KerberosKdc,
			"kerberos_port":          ret.DtSlbHealthCheckDetailsOper.Oper.KerberosPort,
			"snmp_operation":         ret.DtSlbHealthCheckDetailsOper.Oper.SnmpOperation,
			"community":              ret.DtSlbHealthCheckDetailsOper.Oper.Community,
			"oid":                    ret.DtSlbHealthCheckDetailsOper.Oper.Oid,
			"domain":                 ret.DtSlbHealthCheckDetailsOper.Oper.Domain,
			"starttls":               ret.DtSlbHealthCheckDetailsOper.Oper.Starttls,
			"mail_from":              ret.DtSlbHealthCheckDetailsOper.Oper.MailFrom,
			"rcpt_to":                ret.DtSlbHealthCheckDetailsOper.Oper.RcptTo,
			"ipaddr":                 ret.DtSlbHealthCheckDetailsOper.Oper.Ipaddr,
			"dns_qtype":              ret.DtSlbHealthCheckDetailsOper.Oper.DnsQtype,
			"dns_recurse":            ret.DtSlbHealthCheckDetailsOper.Oper.DnsRecurse,
			"dns_expect_type":        ret.DtSlbHealthCheckDetailsOper.Oper.DnsExpectType,
			"dns_expect":             ret.DtSlbHealthCheckDetailsOper.Oper.DnsExpect,
			"transport_proto":        ret.DtSlbHealthCheckDetailsOper.Oper.TransportProto,
			"sip_register":           ret.DtSlbHealthCheckDetailsOper.Oper.SipRegister,
			"secret":                 ret.DtSlbHealthCheckDetailsOper.Oper.Secret,
			"query":                  ret.DtSlbHealthCheckDetailsOper.Oper.Query,
			"base_dn":                ret.DtSlbHealthCheckDetailsOper.Oper.BaseDn,
			"ldap_ssl":               ret.DtSlbHealthCheckDetailsOper.Oper.LdapSsl,
			"ldap_tls":               ret.DtSlbHealthCheckDetailsOper.Oper.LdapTls,
			"attr_type":              ret.DtSlbHealthCheckDetailsOper.Oper.AttrType,
			"db_name":                ret.DtSlbHealthCheckDetailsOper.Oper.DbName,
			"receive":                ret.DtSlbHealthCheckDetailsOper.Oper.Receive,
			"rcv_integer":            ret.DtSlbHealthCheckDetailsOper.Oper.RcvInteger,
			"db_row":                 ret.DtSlbHealthCheckDetailsOper.Oper.DbRow,
			"db_column":              ret.DtSlbHealthCheckDetailsOper.Oper.DbColumn,
			"pname":                  ret.DtSlbHealthCheckDetailsOper.Oper.Pname,
			"tcp_only":               ret.DtSlbHealthCheckDetailsOper.Oper.TcpOnly,
			"attr_program":           ret.DtSlbHealthCheckDetailsOper.Oper.AttrProgram,
			"arguments":              ret.DtSlbHealthCheckDetailsOper.Oper.Arguments,
			"attr_rpn":               ret.DtSlbHealthCheckDetailsOper.Oper.AttrRpn,
			"http_wait_resp":         ret.DtSlbHealthCheckDetailsOper.Oper.HttpWaitResp,
			"l4_conn_num":            ret.DtSlbHealthCheckDetailsOper.Oper.L4ConnNum,
			"l4_errors":              ret.DtSlbHealthCheckDetailsOper.Oper.L4Errors,
			"avg_rtt":                ret.DtSlbHealthCheckDetailsOper.Oper.AvgRtt,
			"curr_rtt":               ret.DtSlbHealthCheckDetailsOper.Oper.CurrRtt,
			"avg_tcp_rtt":            ret.DtSlbHealthCheckDetailsOper.Oper.AvgTcpRtt,
			"curr_tcp_rtt":           ret.DtSlbHealthCheckDetailsOper.Oper.CurrTcpRtt,
			"status_code_rcv":        ret.DtSlbHealthCheckDetailsOper.Oper.StatusCodeRcv,
			"http_req_sent":          ret.DtSlbHealthCheckDetailsOper.Oper.HttpReqSent,
			"http_errors":            ret.DtSlbHealthCheckDetailsOper.Oper.HttpErrors,
			"mac_addr":               ret.DtSlbHealthCheckDetailsOper.Oper.MacAddr,
		},
	}
}

func getObjectSlbHealthCheckDetailsOperOper(d []interface{}) edpt.SlbHealthCheckDetailsOperOper {

	count1 := len(d)
	var ret edpt.SlbHealthCheckDetailsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PinId = in["pin_id"].(int)
		ret.ProcessIndex = in["process_index"].(int)
		ret.HealthState = in["health_state"].(string)
		ret.StateReason = in["state_reason"].(string)
		ret.MonitorName = in["monitor_name"].(string)
		ret.ReceivedSuccess = in["received_success"].(int)
		ret.ReceivedFail = in["received_fail"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.CurrInterval = in["curr_interval"].(int)
		ret.Method = in["method"].(string)
		ret.AttrAliasAddr = in["attr_alias_addr"].(string)
		ret.AttrPort = in["attr_port"].(int)
		ret.HalfOpen = in["half_open"].(int)
		ret.Send = in["send"].(string)
		ret.RespCont = in["resp_cont"].(string)
		ret.ForceUp = in["force_up"].(int)
		ret.Url = in["url"].(string)
		ret.ExpectText = in["expect_text"].(string)
		ret.ExpectRespCode = in["expect_resp_code"].(string)
		ret.ExpectTextRegex = in["expect_text_regex"].(string)
		ret.ExpectRespRegexCode = in["expect_resp_regex_code"].(string)
		ret.MaintenanceCode = in["maintenance_code"].(string)
		ret.User = in["user"].(string)
		ret.Pass = in["pass"].(string)
		ret.Postdata = in["postdata"].(string)
		ret.Host = in["host"].(string)
		ret.KerberosRealm = in["kerberos_realm"].(string)
		ret.KerberosKdc = in["kerberos_kdc"].(string)
		ret.KerberosPort = in["kerberos_port"].(int)
		ret.SnmpOperation = in["snmp_operation"].(int)
		ret.Community = in["community"].(string)
		ret.Oid = in["oid"].(string)
		ret.Domain = in["domain"].(string)
		ret.Starttls = in["starttls"].(int)
		ret.MailFrom = in["mail_from"].(string)
		ret.RcptTo = in["rcpt_to"].(string)
		ret.Ipaddr = in["ipaddr"].(string)
		ret.DnsQtype = in["dns_qtype"].(int)
		ret.DnsRecurse = in["dns_recurse"].(int)
		ret.DnsExpectType = in["dns_expect_type"].(int)
		ret.DnsExpect = in["dns_expect"].(string)
		ret.TransportProto = in["transport_proto"].(int)
		ret.SipRegister = in["sip_register"].(int)
		ret.Secret = in["secret"].(string)
		ret.Query = in["query"].(string)
		ret.BaseDn = in["base_dn"].(string)
		ret.LdapSsl = in["ldap_ssl"].(int)
		ret.LdapTls = in["ldap_tls"].(int)
		ret.AttrType = in["attr_type"].(string)
		ret.DbName = in["db_name"].(string)
		ret.Receive = in["receive"].(string)
		ret.RcvInteger = in["rcv_integer"].(int)
		ret.DbRow = in["db_row"].(int)
		ret.DbColumn = in["db_column"].(int)
		ret.Pname = in["pname"].(string)
		ret.TcpOnly = in["tcp_only"].(int)
		ret.AttrProgram = in["attr_program"].(string)
		ret.Arguments = in["arguments"].(string)
		ret.AttrRpn = in["attr_rpn"].(string)
		ret.HttpWaitResp = in["http_wait_resp"].(int)
		ret.L4ConnNum = in["l4_conn_num"].(int)
		ret.L4Errors = in["l4_errors"].(int)
		ret.AvgRtt = in["avg_rtt"].(int)
		ret.CurrRtt = in["curr_rtt"].(int)
		ret.AvgTcpRtt = in["avg_tcp_rtt"].(int)
		ret.CurrTcpRtt = in["curr_tcp_rtt"].(int)
		ret.StatusCodeRcv = in["status_code_rcv"].(int)
		ret.HttpReqSent = in["http_req_sent"].(int)
		ret.HttpErrors = in["http_errors"].(int)
		ret.MacAddr = in["mac_addr"].(string)
	}
	return ret
}

func dataToEndpointSlbHealthCheckDetailsOper(d *schema.ResourceData) edpt.SlbHealthCheckDetailsOper {
	var ret edpt.SlbHealthCheckDetailsOper

	ret.Oper = getObjectSlbHealthCheckDetailsOperOper(d.Get("oper").([]interface{}))
	return ret
}
