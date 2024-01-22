package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbVirtualServerPortStats51() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_virtual_server_port_stats_imap_vport`: Statistics for the object port\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbVirtualServerPortStats51Create,
		UpdateContext: resourceSlbVirtualServerPortStats51Update,
		ReadContext:   resourceSlbVirtualServerPortStats51Read,
		DeleteContext: resourceSlbVirtualServerPortStats51Delete,

		Schema: map[string]*schema.Schema{
			"port_number": {
				Type: schema.TypeInt, Required: true, Description: "Port",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': TCP LB service; 'udp': UDP Port; 'others': for no tcp/udp protocol, do IP load balancing; 'diameter': diameter port; 'dns-tcp': DNS service over TCP; 'dns-udp': DNS service over UDP; 'fast-http': Fast HTTP Port; 'fix': FIX Port; 'ftp': File Transfer Protocol Port; 'ftp-proxy': ftp proxy port; 'http': HTTP Port; 'https': HTTPS port; 'imap': imap proxy port; 'mlb': Message based load balancing; 'mms': Microsoft Multimedia Service Port; 'mysql': mssql port; 'mssql': mssql; 'pop3': pop3 proxy port; 'radius': RADIUS Port; 'rtsp': Real Time Streaming Protocol Port; 'sip': Session initiation protocol over UDP; 'sip-tcp': Session initiation protocol over TCP; 'sips': Session initiation protocol over TLS; 'smpp-tcp': SMPP service over TCP; 'spdy': spdy port; 'spdys': spdys port; 'smtp': SMTP Port; 'mqtt': MQTT Port; 'mqtts': MQTTS Port; 'ssl-proxy': Generic SSL proxy; 'ssli': SSL insight; 'ssh': SSH Port; 'tcp-proxy': Generic TCP proxy; 'tftp': TFTP Port; 'fast-fix': Fast FIX port; 'http-over-quic': HTTP3-over-quic port;",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"imap_vport": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"num": {
										Type: schema.TypeInt, Optional: true, Description: "Num",
									},
									"curr": {
										Type: schema.TypeInt, Optional: true, Description: "Current proxy conns",
									},
									"total": {
										Type: schema.TypeInt, Optional: true, Description: "Total proxy conns",
									},
									"svrsel_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Server selection failure",
									},
									"no_route": {
										Type: schema.TypeInt, Optional: true, Description: "no route failure",
									},
									"snat_fail": {
										Type: schema.TypeInt, Optional: true, Description: "source nat failure",
									},
									"feat": {
										Type: schema.TypeInt, Optional: true, Description: "feat packet",
									},
									"cc": {
										Type: schema.TypeInt, Optional: true, Description: "clear ctrl port packet",
									},
									"data_ssl": {
										Type: schema.TypeInt, Optional: true, Description: "data ssl force",
									},
									"line_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "line too long",
									},
									"line_mem_freed": {
										Type: schema.TypeInt, Optional: true, Description: "request line freed",
									},
									"invalid_start_line": {
										Type: schema.TypeInt, Optional: true, Description: "invalid start line",
									},
									"auth_tls": {
										Type: schema.TypeInt, Optional: true, Description: "auth tls cmd",
									},
									"prot": {
										Type: schema.TypeInt, Optional: true, Description: "prot cmd",
									},
									"pbsz": {
										Type: schema.TypeInt, Optional: true, Description: "pbsz cmd",
									},
									"pasv": {
										Type: schema.TypeInt, Optional: true, Description: "pasv cmd",
									},
									"port": {
										Type: schema.TypeInt, Optional: true, Description: "port cmd",
									},
									"request_dont_care": {
										Type: schema.TypeInt, Optional: true, Description: "other cmd",
									},
									"client_auth_tls": {
										Type: schema.TypeInt, Optional: true, Description: "client auth tls",
									},
									"cant_find_pasv": {
										Type: schema.TypeInt, Optional: true, Description: "cant find pasv",
									},
									"pasv_addr_ne_server": {
										Type: schema.TypeInt, Optional: true, Description: "psv addr not equal to svr",
									},
									"smp_create_fail": {
										Type: schema.TypeInt, Optional: true, Description: "smp create fail",
									},
									"data_server_conn_fail": {
										Type: schema.TypeInt, Optional: true, Description: "data svr conn fail",
									},
									"data_send_fail": {
										Type: schema.TypeInt, Optional: true, Description: "data send fail",
									},
									"epsv": {
										Type: schema.TypeInt, Optional: true, Description: "epsv command",
									},
									"cant_find_epsv": {
										Type: schema.TypeInt, Optional: true, Description: "cant find epsv",
									},
									"data_curr": {
										Type: schema.TypeInt, Optional: true, Description: "Current Data Proxy",
									},
									"data_total": {
										Type: schema.TypeInt, Optional: true, Description: "Total Data Proxy",
									},
									"auth_unsupported": {
										Type: schema.TypeInt, Optional: true, Description: "Unsupported auth",
									},
									"adat": {
										Type: schema.TypeInt, Optional: true, Description: "adat cmd",
									},
									"unsupported_pbsz_value": {
										Type: schema.TypeInt, Optional: true, Description: "Unsupported PBSZ",
									},
									"unsupported_prot_value": {
										Type: schema.TypeInt, Optional: true, Description: "Unsupported PROT",
									},
									"unsupported_command": {
										Type: schema.TypeInt, Optional: true, Description: "Unsupported cmd",
									},
									"control_to_clear": {
										Type: schema.TypeInt, Optional: true, Description: "Control chn clear txt",
									},
									"control_to_ssl": {
										Type: schema.TypeInt, Optional: true, Description: "Control chn ssl",
									},
									"bad_sequence": {
										Type: schema.TypeInt, Optional: true, Description: "Bad Sequence",
									},
									"rsv_persist_conn_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Serv Sel Persist fail",
									},
									"smp_v6_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Serv Sel SMPv6 fail",
									},
									"smp_v4_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Serv Sel SMPv4 fail",
									},
									"insert_tuple_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Serv Sel insert tuple fail",
									},
									"cl_est_err": {
										Type: schema.TypeInt, Optional: true, Description: "Client EST state erro",
									},
									"ser_connecting_err": {
										Type: schema.TypeInt, Optional: true, Description: "Serv CTNG state error",
									},
									"server_response_err": {
										Type: schema.TypeInt, Optional: true, Description: "Serv RESP state error",
									},
									"cl_request_err": {
										Type: schema.TypeInt, Optional: true, Description: "Client RQ state error",
									},
									"data_conn_start_err": {
										Type: schema.TypeInt, Optional: true, Description: "Data Start state error",
									},
									"data_serv_connecting_err": {
										Type: schema.TypeInt, Optional: true, Description: "Data Serv CTNG error",
									},
									"data_serv_connected_err": {
										Type: schema.TypeInt, Optional: true, Description: "Data Serv CTED error",
									},
									"request": {
										Type: schema.TypeInt, Optional: true, Description: "Total IMAP Requests",
									},
									"capability": {
										Type: schema.TypeInt, Optional: true, Description: "Capability cmd",
									},
									"start_tls": {
										Type: schema.TypeInt, Optional: true, Description: "Total Start TLS cmd",
									},
									"login": {
										Type: schema.TypeInt, Optional: true, Description: "Total Login cmd",
									},
									"realloc_error": {
										Type: schema.TypeInt, Optional: true, Description: "Realloc error",
									},
									"alloc_error": {
										Type: schema.TypeInt, Optional: true, Description: "Alloc error",
									},
									"boundary_error": {
										Type: schema.TypeInt, Optional: true, Description: "Boundary error",
									},
									"negative_error": {
										Type: schema.TypeInt, Optional: true, Description: "Negative error",
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbVirtualServerPortStats51Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats51Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats51(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortStats51Read(ctx, d, meta)
	}
	return diags
}

func resourceSlbVirtualServerPortStats51Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats51Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats51(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortStats51Read(ctx, d, meta)
	}
	return diags
}
func resourceSlbVirtualServerPortStats51Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats51Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats51(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbVirtualServerPortStats51Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats51Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats51(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbVirtualServerPortStats51Stats(d []interface{}) edpt.SlbVirtualServerPortStats51Stats {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats51Stats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Imap_vport = getObjectSlbVirtualServerPortStats51StatsImap_vport(in["imap_vport"].([]interface{}))
	}
	return ret
}

func getObjectSlbVirtualServerPortStats51StatsImap_vport(d []interface{}) edpt.SlbVirtualServerPortStats51StatsImap_vport {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats51StatsImap_vport
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Num = in["num"].(int)
		ret.Curr = in["curr"].(int)
		ret.Total = in["total"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Feat = in["feat"].(int)
		ret.Cc = in["cc"].(int)
		ret.Data_ssl = in["data_ssl"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Line_mem_freed = in["line_mem_freed"].(int)
		ret.Invalid_start_line = in["invalid_start_line"].(int)
		ret.Auth_tls = in["auth_tls"].(int)
		ret.Prot = in["prot"].(int)
		ret.Pbsz = in["pbsz"].(int)
		ret.Pasv = in["pasv"].(int)
		ret.Port = in["port"].(int)
		ret.Request_dont_care = in["request_dont_care"].(int)
		ret.Client_auth_tls = in["client_auth_tls"].(int)
		ret.Cant_find_pasv = in["cant_find_pasv"].(int)
		ret.Pasv_addr_ne_server = in["pasv_addr_ne_server"].(int)
		ret.Smp_create_fail = in["smp_create_fail"].(int)
		ret.Data_server_conn_fail = in["data_server_conn_fail"].(int)
		ret.Data_send_fail = in["data_send_fail"].(int)
		ret.Epsv = in["epsv"].(int)
		ret.Cant_find_epsv = in["cant_find_epsv"].(int)
		ret.Data_curr = in["data_curr"].(int)
		ret.Data_total = in["data_total"].(int)
		ret.Auth_unsupported = in["auth_unsupported"].(int)
		ret.Adat = in["adat"].(int)
		ret.Unsupported_pbsz_value = in["unsupported_pbsz_value"].(int)
		ret.Unsupported_prot_value = in["unsupported_prot_value"].(int)
		ret.Unsupported_command = in["unsupported_command"].(int)
		ret.Control_to_clear = in["control_to_clear"].(int)
		ret.Control_to_ssl = in["control_to_ssl"].(int)
		ret.Bad_sequence = in["bad_sequence"].(int)
		ret.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		ret.Smp_v6_fail = in["smp_v6_fail"].(int)
		ret.Smp_v4_fail = in["smp_v4_fail"].(int)
		ret.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		ret.Cl_est_err = in["cl_est_err"].(int)
		ret.Ser_connecting_err = in["ser_connecting_err"].(int)
		ret.Server_response_err = in["server_response_err"].(int)
		ret.Cl_request_err = in["cl_request_err"].(int)
		ret.Data_conn_start_err = in["data_conn_start_err"].(int)
		ret.Data_serv_connecting_err = in["data_serv_connecting_err"].(int)
		ret.Data_serv_connected_err = in["data_serv_connected_err"].(int)
		ret.Request = in["request"].(int)
		ret.Capability = in["capability"].(int)
		ret.Start_tls = in["start_tls"].(int)
		ret.Login = in["login"].(int)
		ret.Realloc_error = in["realloc_error"].(int)
		ret.Alloc_error = in["alloc_error"].(int)
		ret.Boundary_error = in["boundary_error"].(int)
		ret.Negative_error = in["negative_error"].(int)
	}
	return ret
}

func dataToEndpointSlbVirtualServerPortStats51(d *schema.ResourceData) edpt.SlbVirtualServerPortStats51 {
	var ret edpt.SlbVirtualServerPortStats51
	ret.Inst.PortNumber = d.Get("port_number").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Stats = getObjectSlbVirtualServerPortStats51Stats(d.Get("stats").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
