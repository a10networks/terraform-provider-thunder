package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslForwardProxyCertOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_forward_proxy_cert_oper`: Operational Status for the object ssl-forward-proxy-cert\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslForwardProxyCertOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vserver": {
							Type: schema.TypeString, Optional: true, Description: "virtual server name",
						},
						"port": {
							Type: schema.TypeInt, Optional: true, Description: "Virtual Port",
						},
						"server_ip": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 or IPv6 address of the server",
						},
						"server_port": {
							Type: schema.TypeInt, Optional: true, Description: "Port of the server",
						},
						"server_name": {
							Type: schema.TypeString, Optional: true, Description: "Name of the server",
						},
						"hashed_certificate": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server_ip": {
										Type: schema.TypeString, Optional: true, Description: "Server IP",
									},
									"server_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "Server IPV6",
									},
									"real_port": {
										Type: schema.TypeInt, Optional: true, Description: "Real Port",
									},
									"protocol": {
										Type: schema.TypeString, Optional: true, Description: "protocol",
									},
									"server_name": {
										Type: schema.TypeString, Optional: true, Description: "Server Name",
									},
									"state": {
										Type: schema.TypeString, Optional: true, Description: "State",
									},
									"hit_times": {
										Type: schema.TypeInt, Optional: true, Description: "hit times",
									},
									"idle_time": {
										Type: schema.TypeInt, Optional: true, Description: "Idle time in seconds",
									},
									"timeout_after": {
										Type: schema.TypeInt, Optional: true, Description: "Timeout after (seconds)",
									},
									"expires_after": {
										Type: schema.TypeInt, Optional: true, Description: "Expires after (seconds)",
									},
									"alpn_protocol": {
										Type: schema.TypeInt, Optional: true, Description: "ALPN Protocol (0:None, 1:HTTP, 2:SPDY, 3:HTTP2)",
									},
									"version": {
										Type: schema.TypeInt, Optional: true, Description: "Certificate version",
									},
									"subject": {
										Type: schema.TypeString, Optional: true, Description: "Certificate subject",
									},
									"common_name": {
										Type: schema.TypeString, Optional: true, Description: "Certificate common name",
									},
									"division": {
										Type: schema.TypeString, Optional: true, Description: "Certificate division",
									},
									"locality": {
										Type: schema.TypeString, Optional: true, Description: "Certificate locality",
									},
									"state_province": {
										Type: schema.TypeString, Optional: true, Description: "Certificate state or province",
									},
									"country": {
										Type: schema.TypeString, Optional: true, Description: "Certificate country",
									},
									"subject_alt_name": {
										Type: schema.TypeString, Optional: true, Description: "Certificate subject alternative name",
									},
									"email": {
										Type: schema.TypeString, Optional: true, Description: "Certificate email",
									},
									"start_time": {
										Type: schema.TypeString, Optional: true, Description: "Certificate start time",
									},
									"expire_time": {
										Type: schema.TypeString, Optional: true, Description: "Certificate expiration time",
									},
									"issuer": {
										Type: schema.TypeString, Optional: true, Description: "Certificate issuer",
									},
									"serial": {
										Type: schema.TypeString, Optional: true, Description: "Certificate serial number in hex",
									},
								},
							},
						},
						"hashed_persist_entries": {
							Type: schema.TypeInt, Optional: true, Description: "DELETE method specific option to clear the hashed persistence entries for one vport",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSslForwardProxyCertOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslForwardProxyCertOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslForwardProxyCertOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslForwardProxyCertOperOper := setObjectSlbSslForwardProxyCertOperOper(res)
		d.Set("oper", SlbSslForwardProxyCertOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslForwardProxyCertOperOper(ret edpt.DataSlbSslForwardProxyCertOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vserver":                ret.DtSlbSslForwardProxyCertOper.Oper.Vserver,
			"port":                   ret.DtSlbSslForwardProxyCertOper.Oper.Port,
			"server_ip":              ret.DtSlbSslForwardProxyCertOper.Oper.ServerIp,
			"server_port":            ret.DtSlbSslForwardProxyCertOper.Oper.ServerPort,
			"server_name":            ret.DtSlbSslForwardProxyCertOper.Oper.ServerName,
			"hashed_certificate":     setSliceSlbSslForwardProxyCertOperOperHashedCertificate(ret.DtSlbSslForwardProxyCertOper.Oper.HashedCertificate),
			"hashed_persist_entries": ret.DtSlbSslForwardProxyCertOper.Oper.HashedPersistEntries,
		},
	}
}

func setSliceSlbSslForwardProxyCertOperOperHashedCertificate(d []edpt.SlbSslForwardProxyCertOperOperHashedCertificate) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["server_ip"] = item.ServerIp
		in["server_ipv6"] = item.ServerIpv6
		in["real_port"] = item.RealPort
		in["protocol"] = item.Protocol
		in["server_name"] = item.ServerName
		in["state"] = item.State
		in["hit_times"] = item.HitTimes
		in["idle_time"] = item.IdleTime
		in["timeout_after"] = item.TimeoutAfter
		in["expires_after"] = item.ExpiresAfter
		in["alpn_protocol"] = item.AlpnProtocol
		in["version"] = item.Version
		in["subject"] = item.Subject
		in["common_name"] = item.CommonName
		in["division"] = item.Division
		in["locality"] = item.Locality
		in["state_province"] = item.StateProvince
		in["country"] = item.Country
		in["subject_alt_name"] = item.SubjectAltName
		in["email"] = item.Email
		in["start_time"] = item.StartTime
		in["expire_time"] = item.ExpireTime
		in["issuer"] = item.Issuer
		in["serial"] = item.Serial
		result = append(result, in)
	}
	return result
}

func getObjectSlbSslForwardProxyCertOperOper(d []interface{}) edpt.SlbSslForwardProxyCertOperOper {

	count1 := len(d)
	var ret edpt.SlbSslForwardProxyCertOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Vserver = in["vserver"].(string)
		ret.Port = in["port"].(int)
		ret.ServerIp = in["server_ip"].(string)
		ret.ServerPort = in["server_port"].(int)
		ret.ServerName = in["server_name"].(string)
		ret.HashedCertificate = getSliceSlbSslForwardProxyCertOperOperHashedCertificate(in["hashed_certificate"].([]interface{}))
		ret.HashedPersistEntries = in["hashed_persist_entries"].(int)
	}
	return ret
}

func getSliceSlbSslForwardProxyCertOperOperHashedCertificate(d []interface{}) []edpt.SlbSslForwardProxyCertOperOperHashedCertificate {

	count1 := len(d)
	ret := make([]edpt.SlbSslForwardProxyCertOperOperHashedCertificate, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSslForwardProxyCertOperOperHashedCertificate
		oi.ServerIp = in["server_ip"].(string)
		oi.ServerIpv6 = in["server_ipv6"].(string)
		oi.RealPort = in["real_port"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.ServerName = in["server_name"].(string)
		oi.State = in["state"].(string)
		oi.HitTimes = in["hit_times"].(int)
		oi.IdleTime = in["idle_time"].(int)
		oi.TimeoutAfter = in["timeout_after"].(int)
		oi.ExpiresAfter = in["expires_after"].(int)
		oi.AlpnProtocol = in["alpn_protocol"].(int)
		oi.Version = in["version"].(int)
		oi.Subject = in["subject"].(string)
		oi.CommonName = in["common_name"].(string)
		oi.Division = in["division"].(string)
		oi.Locality = in["locality"].(string)
		oi.StateProvince = in["state_province"].(string)
		oi.Country = in["country"].(string)
		oi.SubjectAltName = in["subject_alt_name"].(string)
		oi.Email = in["email"].(string)
		oi.StartTime = in["start_time"].(string)
		oi.ExpireTime = in["expire_time"].(string)
		oi.Issuer = in["issuer"].(string)
		oi.Serial = in["serial"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSslForwardProxyCertOper(d *schema.ResourceData) edpt.SlbSslForwardProxyCertOper {
	var ret edpt.SlbSslForwardProxyCertOper

	ret.Oper = getObjectSlbSslForwardProxyCertOperOper(d.Get("oper").([]interface{}))
	return ret
}
