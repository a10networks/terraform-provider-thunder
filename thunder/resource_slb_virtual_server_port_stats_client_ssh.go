package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbVirtualServerPortStats49() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_virtual_server_port_stats_client_ssh`: Statistics for the object port\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbVirtualServerPortStats49Create,
		UpdateContext: resourceSlbVirtualServerPortStats49Update,
		ReadContext:   resourceSlbVirtualServerPortStats49Read,
		DeleteContext: resourceSlbVirtualServerPortStats49Delete,

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
						"client_ssh": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"successful_handshakes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"failed_handshakes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"forwarding_errors": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"virtual_server_name": {
				Type: schema.TypeString, Required: true, Description: "Virtual_server_name",
			},
		},
	}
}
func resourceSlbVirtualServerPortStats49Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats49Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats49(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortStats49Read(ctx, d, meta)
	}
	return diags
}

func resourceSlbVirtualServerPortStats49Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats49Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats49(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortStats49Read(ctx, d, meta)
	}
	return diags
}
func resourceSlbVirtualServerPortStats49Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats49Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats49(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbVirtualServerPortStats49Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats49Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats49(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbVirtualServerPortStats49Stats(d []interface{}) edpt.SlbVirtualServerPortStats49Stats {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats49Stats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClientSsh = getObjectSlbVirtualServerPortStats49StatsClientSsh(in["client_ssh"].([]interface{}))
	}
	return ret
}

func getObjectSlbVirtualServerPortStats49StatsClientSsh(d []interface{}) edpt.SlbVirtualServerPortStats49StatsClientSsh {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats49StatsClientSsh
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Successful_handshakes = in["successful_handshakes"].(int)
		ret.Failed_handshakes = in["failed_handshakes"].(int)
		ret.Forwarding_errors = in["forwarding_errors"].(int)
	}
	return ret
}

func dataToEndpointSlbVirtualServerPortStats49(d *schema.ResourceData) edpt.SlbVirtualServerPortStats49 {
	var ret edpt.SlbVirtualServerPortStats49
	ret.Inst.PortNumber = d.Get("port_number").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Stats = getObjectSlbVirtualServerPortStats49Stats(d.Get("stats").([]interface{}))
	ret.Inst.Virtual_server_name = d.Get("virtual_server_name").(string)
	return ret
}
