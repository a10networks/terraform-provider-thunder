package thunder

import (
	"errors"
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSnmpServerHostHostName() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_host_host_name`: Specify hosts to receive SNMP traps\n\n_",
		CreateContext: resourceSnmpServerHostHostNameCreate,
		UpdateContext: resourceSnmpServerHostHostNameUpdate,
		ReadContext:   resourceSnmpServerHostHostNameRead,
		DeleteContext: resourceSnmpServerHostHostNameDelete,
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Hostname of SNMP trap host",
				ValidateFunc: validation.StringLenBetween(1, 31),
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

func resourceSnmpServerHostHostNameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostHostNameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		if client.Partition == "shared" {
			obj := dataToEndpointSnmpServerHostHostName(d)
			err := runtimeCheckData(d)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(obj.GetId())
			err = obj.Post(client.Token, client.Host, logger)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			obj := dataToEndpointSnmpServerHostHostNameL3v(d)
			d.SetId(obj.GetId())
			err := obj.Post(client.Token, client.Host, logger)
			if err != nil {
				return diag.FromErr(err)
			}
		}
		return resourceSnmpServerHostHostNameRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerHostHostNameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostHostNameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		if client.Partition == "shared" {
			obj := edpt.SnmpServerHostHostName{}
			err := obj.Get(client.Token, client.Host, d.Id(), logger)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			obj := edpt.SnmpServerHostHostNameL3v{}
			err := obj.Get(client.Token, client.Host, d.Id(), logger)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}
	return diags
}

func resourceSnmpServerHostHostNameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerHostHostNameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		if client.Partition == "shared" {
			obj := dataToEndpointSnmpServerHostHostName(d)
			err := runtimeCheckData(d)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(obj.GetId())
			err = obj.Put(client.Token, client.Host, logger)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			obj := dataToEndpointSnmpServerHostHostNameL3v(d)
			d.SetId(obj.GetId())
			err := obj.Put(client.Token, client.Host, logger)
			if err != nil {
				return diag.FromErr(err)
			}
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
		if client.Partition == "shared" {
			obj := edpt.SnmpServerHostHostName{}
			err := obj.Delete(client.Token, client.Host, d.Id(), logger)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			obj := edpt.SnmpServerHostHostNameL3v{}
			err := obj.Delete(client.Token, client.Host, d.Id(), logger)
			if err != nil {
				return diag.FromErr(err)
			}
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

func dataToEndpointSnmpServerHostHostNameL3v(d *schema.ResourceData) edpt.SnmpServerHostHostNameL3v {
	var ret edpt.SnmpServerHostHostNameL3v
	ret.Inst.Hostname = d.Get("hostname").(string)
	//omit uuid
	return ret
}

func runtimeCheckData(d *schema.ResourceData) error {
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
