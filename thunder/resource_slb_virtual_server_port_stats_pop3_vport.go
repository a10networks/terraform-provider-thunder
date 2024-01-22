package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbVirtualServerPortStats52() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_virtual_server_port_stats_pop3_vport`: Statistics for the object port\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbVirtualServerPortStats52Create,
		UpdateContext: resourceSlbVirtualServerPortStats52Update,
		ReadContext:   resourceSlbVirtualServerPortStats52Read,
		DeleteContext: resourceSlbVirtualServerPortStats52Delete,

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
						"pop3_vport": {
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
									"line_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "line too long",
									},
									"line_mem_freed": {
										Type: schema.TypeInt, Optional: true, Description: "request line freed",
									},
									"invalid_start_line": {
										Type: schema.TypeInt, Optional: true, Description: "invalid start line",
									},
									"stls": {
										Type: schema.TypeInt, Optional: true, Description: "stls cmd",
									},
									"request_dont_care": {
										Type: schema.TypeInt, Optional: true, Description: "other cmd",
									},
									"unsupported_command": {
										Type: schema.TypeInt, Optional: true, Description: "Unsupported cmd",
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
									"request": {
										Type: schema.TypeInt, Optional: true, Description: "Total POP3 Request",
									},
									"control_to_ssl": {
										Type: schema.TypeInt, Optional: true, Description: "Control chn ssl",
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
func resourceSlbVirtualServerPortStats52Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats52Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats52(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortStats52Read(ctx, d, meta)
	}
	return diags
}

func resourceSlbVirtualServerPortStats52Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats52Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats52(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortStats52Read(ctx, d, meta)
	}
	return diags
}
func resourceSlbVirtualServerPortStats52Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats52Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats52(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbVirtualServerPortStats52Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats52Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats52(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbVirtualServerPortStats52Stats(d []interface{}) edpt.SlbVirtualServerPortStats52Stats {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats52Stats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Pop3_vport = getObjectSlbVirtualServerPortStats52StatsPop3_vport(in["pop3_vport"].([]interface{}))
	}
	return ret
}

func getObjectSlbVirtualServerPortStats52StatsPop3_vport(d []interface{}) edpt.SlbVirtualServerPortStats52StatsPop3_vport {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats52StatsPop3_vport
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Num = in["num"].(int)
		ret.Curr = in["curr"].(int)
		ret.Total = in["total"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Line_mem_freed = in["line_mem_freed"].(int)
		ret.Invalid_start_line = in["invalid_start_line"].(int)
		ret.Stls = in["stls"].(int)
		ret.Request_dont_care = in["request_dont_care"].(int)
		ret.Unsupported_command = in["unsupported_command"].(int)
		ret.Bad_sequence = in["bad_sequence"].(int)
		ret.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		ret.Smp_v6_fail = in["smp_v6_fail"].(int)
		ret.Smp_v4_fail = in["smp_v4_fail"].(int)
		ret.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		ret.Cl_est_err = in["cl_est_err"].(int)
		ret.Ser_connecting_err = in["ser_connecting_err"].(int)
		ret.Server_response_err = in["server_response_err"].(int)
		ret.Cl_request_err = in["cl_request_err"].(int)
		ret.Request = in["request"].(int)
		ret.Control_to_ssl = in["control_to_ssl"].(int)
	}
	return ret
}

func dataToEndpointSlbVirtualServerPortStats52(d *schema.ResourceData) edpt.SlbVirtualServerPortStats52 {
	var ret edpt.SlbVirtualServerPortStats52
	ret.Inst.PortNumber = d.Get("port_number").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Stats = getObjectSlbVirtualServerPortStats52Stats(d.Get("stats").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
