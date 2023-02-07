package thunder

import (
	"errors"
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSnmpServerHostIpv4Host() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_host_ipv4_host`: Specify hosts to receive SNMP traps\n\n_",
		CreateContext: resourceSnmpServerHostIpv4HostCreate,
		UpdateContext: resourceSnmpServerHostIpv4HostUpdate,
		ReadContext:   resourceSnmpServerHostIpv4HostRead,
		DeleteContext: resourceSnmpServerHostIpv4HostDelete,
		Schema: map[string]*schema.Schema{
			"ipv4_addr": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "IPv4 address of SNMP trap host",
				ValidateFunc: validation.IsIPv4Address,
			},
			"udp_port": {
				Type: schema.TypeInt, Optional: true, Default: 162, Description: "The trap host's UDP port number(default: 162)",
				ValidateFunc: validation.IntBetween(1, 65535),
			},
			"user": {
				Type: schema.TypeString, Optional: true, Description: "SNMPv3 user to send traps (User Name)",
				ValidateFunc: validation.StringLenBetween(1, 31),
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"v1_v2c_comm": {
				Type: schema.TypeString, Optional: true, Description: "SNMPv1/v2c community string",
				ValidateFunc: validation.StringLenBetween(1, 31),
			},
			"version": {
				Type: schema.TypeString, Optional: true, ForceNew: true, Description: "'v1': Use SNMPv1; 'v2c': Use SNMPv2c; 'v3': User SNMPv3;",
				ValidateFunc: validation.StringInSlice([]string{"v1", "v2c", "v3"}, false),
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
		if client.Partition == "shared" {
			obj := dataToEndpointSnmpServerHostIpv4Host(d)
			err := runtimeCheckSnmpServerHostIpv4Host(d)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(obj.GetId())
			err = obj.Post(client.Token, client.Host, logger)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			obj := dataToEndpointSnmpServerHostIpv4HostL3v(d)
			d.SetId(obj.GetId())
			err := obj.Post(client.Token, client.Host, logger)
			if err != nil {
				return diag.FromErr(err)
			}
		}
		return resourceSnmpServerHostIpv4HostRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerHostIpv4HostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostIpv4HostRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		if client.Partition == "shared" {
			obj := edpt.SnmpServerHostIpv4Host{}
			err := obj.Get(client.Token, client.Host, d.Id(), logger)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			obj := edpt.SnmpServerHostIpv4HostL3v{}
			err := obj.Get(client.Token, client.Host, d.Id(), logger)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}
	return diags
}

func resourceSnmpServerHostIpv4HostUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostIpv4HostUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		if client.Partition == "shared" {
			obj := dataToEndpointSnmpServerHostIpv4Host(d)
			err := runtimeCheckSnmpServerHostIpv4Host(d)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(obj.GetId())
			err = obj.Put(client.Token, client.Host, logger)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			obj := dataToEndpointSnmpServerHostIpv4HostL3v(d)
			d.SetId(obj.GetId())
			err := obj.Put(client.Token, client.Host, logger)
			if err != nil {
				return diag.FromErr(err)
			}
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
		if client.Partition == "shared" {
			obj := edpt.SnmpServerHostIpv4Host{}
			err := obj.Delete(client.Token, client.Host, d.Id(), logger)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			obj := edpt.SnmpServerHostIpv4HostL3v{}
			err := obj.Delete(client.Token, client.Host, d.Id(), logger)
			if err != nil {
				return diag.FromErr(err)
			}
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

func dataToEndpointSnmpServerHostIpv4HostL3v(d *schema.ResourceData) edpt.SnmpServerHostIpv4HostL3v {
	var ret edpt.SnmpServerHostIpv4HostL3v
	ret.Inst.Ipv4Addr = d.Get("ipv4_addr").(string)
	//omit uuid
	return ret
}

func runtimeCheckSnmpServerHostIpv4Host(d *schema.ResourceData) error {
	snmp_version,flag := d.GetOk("version")
	if false == flag {
		return errors.New("Need 'version' for SNMP server in shared partition")
	}
	if snmp_version == "v3" {
		_,flag = d.GetOk("user")
		if false == flag {
			return errors.New("Need 'user' for SNMPv3")
		}
	} else {
		_,flag = d.GetOk("v1_v2c_comm")
		if false == flag {
			return errors.New("Need 'v1_v2c_comm' for SNMPv1 or SNMPv2c")
		}
	}
	return nil
}
