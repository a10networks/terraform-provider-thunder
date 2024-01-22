package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerHostHostName() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_host_host_name`: Specify IPV6 hosts to receive SNMP traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerHostHostNameCreate,
		UpdateContext: resourceSnmpServerHostHostNameUpdate,
		ReadContext:   resourceSnmpServerHostHostNameRead,
		DeleteContext: resourceSnmpServerHostHostNameDelete,

		Schema: map[string]*schema.Schema{
			"hostname": {
				Type: schema.TypeString, Required: true, Description: "Hostname of SNMP trap host",
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
func resourceSnmpServerHostHostNameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostHostNameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerHostHostName(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerHostHostNameRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerHostHostNameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostHostNameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerHostHostName(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerHostHostNameRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerHostHostNameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostHostNameDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerHostHostName(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerHostHostNameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostHostNameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerHostHostName(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerHostHostName(d *schema.ResourceData) edpt.SnmpServerHostHostName {
	var ret edpt.SnmpServerHostHostName
	ret.Inst.Hostname = d.Get("hostname").(string)
	ret.Inst.UdpPort = d.Get("udp_port").(int)
	ret.Inst.User = d.Get("user").(string)
	//omit uuid
	ret.Inst.V1V2cComm = d.Get("v1_v2c_comm").(string)
	ret.Inst.Version = d.Get("version").(string)
	return ret
}
