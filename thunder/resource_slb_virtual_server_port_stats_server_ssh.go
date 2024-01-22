package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbVirtualServerPortStats53() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_virtual_server_port_stats_server_ssh`: Statistics for the object port\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbVirtualServerPortStats53Create,
		UpdateContext: resourceSlbVirtualServerPortStats53Update,
		ReadContext:   resourceSlbVirtualServerPortStats53Read,
		DeleteContext: resourceSlbVirtualServerPortStats53Delete,

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
						"server_ssh": {
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbVirtualServerPortStats53Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats53Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats53(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortStats53Read(ctx, d, meta)
	}
	return diags
}

func resourceSlbVirtualServerPortStats53Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats53Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats53(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortStats53Read(ctx, d, meta)
	}
	return diags
}
func resourceSlbVirtualServerPortStats53Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats53Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats53(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbVirtualServerPortStats53Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats53Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats53(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbVirtualServerPortStats53Stats(d []interface{}) edpt.SlbVirtualServerPortStats53Stats {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats53Stats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServerSsh = getObjectSlbVirtualServerPortStats53StatsServerSsh(in["server_ssh"].([]interface{}))
	}
	return ret
}

func getObjectSlbVirtualServerPortStats53StatsServerSsh(d []interface{}) edpt.SlbVirtualServerPortStats53StatsServerSsh {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats53StatsServerSsh
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Successful_handshakes = in["successful_handshakes"].(int)
		ret.Failed_handshakes = in["failed_handshakes"].(int)
		ret.Forwarding_errors = in["forwarding_errors"].(int)
	}
	return ret
}

func dataToEndpointSlbVirtualServerPortStats53(d *schema.ResourceData) edpt.SlbVirtualServerPortStats53 {
	var ret edpt.SlbVirtualServerPortStats53
	ret.Inst.PortNumber = d.Get("port_number").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Stats = getObjectSlbVirtualServerPortStats53Stats(d.Get("stats").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
