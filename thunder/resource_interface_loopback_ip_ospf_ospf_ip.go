package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLoopbackIpOspfOspfIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_loopback_ip_ospf_ospf_ip`: IP address configuration for Open Shortest Path First for IPv4 (OSPF)\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLoopbackIpOspfOspfIpCreate,
		UpdateContext: resourceInterfaceLoopbackIpOspfOspfIpUpdate,
		ReadContext:   resourceInterfaceLoopbackIpOspfOspfIpRead,
		DeleteContext: resourceInterfaceLoopbackIpOspfOspfIpDelete,

		Schema: map[string]*schema.Schema{
			"authentication": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable authentication",
			},
			"authentication_key": {
				Type: schema.TypeString, Optional: true, Description: "Authentication password (key) (The OSPF password (key))",
			},
			"cost": {
				Type: schema.TypeInt, Optional: true, Description: "Interface cost",
			},
			"database_filter": {
				Type: schema.TypeString, Optional: true, Description: "'all': Filter all LSA;",
			},
			"dead_interval": {
				Type: schema.TypeInt, Optional: true, Default: 40, Description: "Interval after which a neighbor is declared dead (Seconds)",
			},
			"hello_interval": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Time between HELLO packets (Seconds)",
			},
			"ip_addr": {
				Type: schema.TypeString, Required: true, Description: "Address of interface",
			},
			"message_digest_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message_digest_key": {
							Type: schema.TypeInt, Optional: true, Description: "Message digest authentication password (key) (Key id)",
						},
						"md5_value": {
							Type: schema.TypeString, Optional: true, Description: "The OSPF password (1-16)",
						},
					},
				},
			},
			"mtu_ignore": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignores the MTU in DBD packets",
			},
			"out": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Outgoing LSA",
			},
			"priority": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Router priority",
			},
			"retransmit_interval": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Time between retransmitting lost link state advertisements (Seconds)",
			},
			"transmit_delay": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Link state transmit delay (Seconds)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeString, Optional: true, Description: "'message-digest': Use message-digest authentication; 'null': Use no authentication;",
			},
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceLoopbackIpOspfOspfIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpOspfOspfIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpOspfOspfIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLoopbackIpOspfOspfIpRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLoopbackIpOspfOspfIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpOspfOspfIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpOspfOspfIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLoopbackIpOspfOspfIpRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLoopbackIpOspfOspfIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpOspfOspfIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpOspfOspfIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLoopbackIpOspfOspfIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpOspfOspfIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpOspfOspfIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceInterfaceLoopbackIpOspfOspfIpMessageDigestCfg(d []interface{}) []edpt.InterfaceLoopbackIpOspfOspfIpMessageDigestCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpOspfOspfIpMessageDigestCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpOspfOspfIpMessageDigestCfg
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5Value = in["md5_value"].(string)
		//omit encrypted
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceLoopbackIpOspfOspfIp(d *schema.ResourceData) edpt.InterfaceLoopbackIpOspfOspfIp {
	var ret edpt.InterfaceLoopbackIpOspfOspfIp
	ret.Inst.Authentication = d.Get("authentication").(int)
	ret.Inst.AuthenticationKey = d.Get("authentication_key").(string)
	ret.Inst.Cost = d.Get("cost").(int)
	ret.Inst.DatabaseFilter = d.Get("database_filter").(string)
	ret.Inst.DeadInterval = d.Get("dead_interval").(int)
	ret.Inst.HelloInterval = d.Get("hello_interval").(int)
	ret.Inst.IpAddr = d.Get("ip_addr").(string)
	ret.Inst.MessageDigestCfg = getSliceInterfaceLoopbackIpOspfOspfIpMessageDigestCfg(d.Get("message_digest_cfg").([]interface{}))
	ret.Inst.MtuIgnore = d.Get("mtu_ignore").(int)
	ret.Inst.Out = d.Get("out").(int)
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.RetransmitInterval = d.Get("retransmit_interval").(int)
	ret.Inst.TransmitDelay = d.Get("transmit_delay").(int)
	//omit uuid
	ret.Inst.Value = d.Get("value").(string)
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
