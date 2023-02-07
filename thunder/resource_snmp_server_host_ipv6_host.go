package thunder

import (
	"errors"
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSnmpServerHostIpv6Host() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_host_ipv6_host`: Specify hosts to receive SNMP traps\n\n_",
		CreateContext: resourceSnmpServerHostIpv6HostCreate,
		UpdateContext: resourceSnmpServerHostIpv6HostUpdate,
		ReadContext:   resourceSnmpServerHostIpv6HostRead,
		DeleteContext: resourceSnmpServerHostIpv6HostDelete,
		Schema: map[string]*schema.Schema{
			"ipv6_addr": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "IPv6 address of SNMP trap host",
				ValidateFunc: validation.IsIPv6Address,
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
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "'v1': Use SNMPv1; 'v2c': Use SNMPv2c; 'v3': User SNMPv3;",
				ValidateFunc: validation.StringInSlice([]string{"v1", "v2c", "v3"}, false),
			},
		},
	}
}

func resourceSnmpServerHostIpv6HostCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostIpv6HostCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		if client.Partition == "shared" {
			obj := dataToEndpointSnmpServerHostIpv6Host(d)
			err := runtimeCheckSnmpServerHostIpv6Host(d)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(obj.GetId())
			err = obj.Post(client.Token, client.Host, logger)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			err := errors.New("this resource is unavailable for private partitions")
			return diag.FromErr(err)
		}
		return resourceSnmpServerHostIpv6HostRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerHostIpv6HostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostIpv6HostRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		if client.Partition == "shared" {
			obj := edpt.SnmpServerHostIpv6Host{}
			err := obj.Get(client.Token, client.Host, d.Id(), logger)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			err := errors.New("this resource is unavailable for private partitions")
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerHostIpv6HostUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostIpv6HostUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		if client.Partition == "shared" {
			obj := dataToEndpointSnmpServerHostIpv6Host(d)
			err := runtimeCheckSnmpServerHostIpv6Host(d)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(obj.GetId())
			err = obj.Put(client.Token, client.Host, logger)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			err := errors.New("this resource is unavailable for private partitions")
			return diag.FromErr(err)
		}
		return resourceSnmpServerHostIpv6HostRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerHostIpv6HostDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostIpv6HostDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		if client.Partition == "shared" {
			obj := edpt.SnmpServerHostIpv6Host{}
			err := obj.Delete(client.Token, client.Host, d.Id(), logger)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			err := errors.New("this resource is unavailable for private partitions")
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerHostIpv6Host(d *schema.ResourceData) edpt.SnmpServerHostIpv6Host {
	var ret edpt.SnmpServerHostIpv6Host
	ret.Inst.Ipv6Addr = d.Get("ipv6_addr").(string)
	ret.Inst.UdpPort = d.Get("udp_port").(int)
	ret.Inst.User = d.Get("user").(string)
	//omit uuid
	ret.Inst.V1V2cComm = d.Get("v1_v2c_comm").(string)
	ret.Inst.Version = d.Get("version").(string)
	return ret
}

func runtimeCheckSnmpServerHostIpv6Host(d *schema.ResourceData) error {
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
