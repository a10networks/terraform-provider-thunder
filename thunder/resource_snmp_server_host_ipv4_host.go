package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerHostIpv4Host() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_host_ipv4_host`: Specify IPV4 hosts to receive SNMP traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerHostIpv4HostCreate,
		UpdateContext: resourceSnmpServerHostIpv4HostUpdate,
		ReadContext:   resourceSnmpServerHostIpv4HostRead,
		DeleteContext: resourceSnmpServerHostIpv4HostDelete,

		Schema: map[string]*schema.Schema{
			"ipv4_addr": {
				Type: schema.TypeString, Required: true, Description: "IPV4 address of SNMP trap host",
			},
			"udp_port": {
				Type: schema.TypeInt, Optional: true, Default: 162, Description: "The trap host's UDP port number(default: 162)",
			},
			"user": {
				Type: schema.TypeString, Optional: true, Description: "SNMPv3 user to send traps (User Name)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"v1_v2c_comm": {
				Type: schema.TypeString, Optional: true, Description: "SNMPv1/v2c community string",
			},
			"version": {
				Type: schema.TypeString, Required: true, Description: "'v1': Use SNMPv1; 'v2c': Use SNMPv2c; 'v3': User SNMPv3;",
			},
		},
	}
}
func resourceSnmpServerHostIpv4HostCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostIpv4HostCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerHostIpv4Host(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerHostIpv4HostRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerHostIpv4HostUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostIpv4HostUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerHostIpv4Host(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerHostIpv4HostRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerHostIpv4HostDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostIpv4HostDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerHostIpv4Host(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerHostIpv4HostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostIpv4HostRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerHostIpv4Host(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerHostIpv4Host(d *schema.ResourceData) edpt.SnmpServerHostIpv4Host {
	var ret edpt.SnmpServerHostIpv4Host
	ret.Inst.Ipv4Addr = d.Get("ipv4_addr").(string)
	ret.Inst.UdpPort = d.Get("udp_port").(int)
	ret.Inst.User = d.Get("user").(string)
	//omit uuid
	ret.Inst.V1V2cComm = d.Get("v1_v2c_comm").(string)
	ret.Inst.Version = d.Get("version").(string)
	return ret
}
