package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLifIpOspfOspfGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_lif_ip_ospf_ospf_global`: Global setting for Open Shortest Path First for IPv4 (OSPF)\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLifIpOspfOspfGlobalCreate,
		UpdateContext: resourceInterfaceLifIpOspfOspfGlobalUpdate,
		ReadContext:   resourceInterfaceLifIpOspfOspfGlobalRead,
		DeleteContext: resourceInterfaceLifIpOspfOspfGlobalDelete,

		Schema: map[string]*schema.Schema{
			"authentication_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable authentication",
						},
						"value": {
							Type: schema.TypeString, Optional: true, Description: "'message-digest': Use message-digest authentication; 'null': Use no authentication;",
						},
					},
				},
			},
			"authentication_key": {
				Type: schema.TypeString, Optional: true, Description: "Authentication password (key) (The OSPF password (key))",
			},
			"bfd_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bfd": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable BFD",
						},
					},
				},
			},
			"cost": {
				Type: schema.TypeInt, Optional: true, Description: "Interface cost",
			},
			"database_filter_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"database_filter": {
							Type: schema.TypeString, Optional: true, Description: "'all': Filter all LSA;",
						},
						"out": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Outgoing LSA",
						},
					},
				},
			},
			"dead_interval": {
				Type: schema.TypeInt, Optional: true, Default: 40, Description: "Interval after which a neighbor is declared dead (Seconds)",
			},
			"disable": {
				Type: schema.TypeString, Optional: true, Description: "'all': All functionality;",
			},
			"hello_interval": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Time between HELLO packets (Seconds)",
			},
			"message_digest_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message_digest_key": {
							Type: schema.TypeInt, Optional: true, Description: "Message digest authentication password (key) (Key id)",
						},
						"md5": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"md5_value": {
										Type: schema.TypeString, Optional: true, Description: "The OSPF password (1-16)",
									},
								},
							},
						},
					},
				},
			},
			"mtu": {
				Type: schema.TypeInt, Optional: true, Description: "OSPF interface MTU (MTU size)",
			},
			"mtu_ignore": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignores the MTU in DBD packets",
			},
			"network": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"broadcast": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF broadcast multi-access network",
						},
						"non_broadcast": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF NBMA network",
						},
						"point_to_point": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF point-to-point network",
						},
						"point_to_multipoint": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF point-to-multipoint network",
						},
						"p2mp_nbma": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify non-broadcast point-to-multipoint network",
						},
					},
				},
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
			"ifname": {
				Type: schema.TypeString, Required: true, Description: "Ifname",
			},
		},
	}
}
func resourceInterfaceLifIpOspfOspfGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpOspfOspfGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpOspfOspfGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIpOspfOspfGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLifIpOspfOspfGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpOspfOspfGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpOspfOspfGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIpOspfOspfGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLifIpOspfOspfGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpOspfOspfGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpOspfOspfGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLifIpOspfOspfGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpOspfOspfGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpOspfOspfGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceLifIpOspfOspfGlobalAuthenticationCfg(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalAuthenticationCfg {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobalAuthenticationCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = in["authentication"].(int)
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalBfdCfg(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalBfdCfg {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobalBfdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalDatabaseFilterCfg(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DatabaseFilter = in["database_filter"].(string)
		ret.Out = in["out"].(int)
	}
	return ret
}

func getSliceInterfaceLifIpOspfOspfGlobalMessageDigestCfg(d []interface{}) []edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfg
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = getObjectInterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5(in["md5"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Md5Value = in["md5_value"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalNetwork(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalNetwork {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobalNetwork
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Broadcast = in["broadcast"].(int)
		ret.NonBroadcast = in["non_broadcast"].(int)
		ret.PointToPoint = in["point_to_point"].(int)
		ret.PointToMultipoint = in["point_to_multipoint"].(int)
		ret.P2mpNbma = in["p2mp_nbma"].(int)
	}
	return ret
}

func dataToEndpointInterfaceLifIpOspfOspfGlobal(d *schema.ResourceData) edpt.InterfaceLifIpOspfOspfGlobal {
	var ret edpt.InterfaceLifIpOspfOspfGlobal
	ret.Inst.AuthenticationCfg = getObjectInterfaceLifIpOspfOspfGlobalAuthenticationCfg(d.Get("authentication_cfg").([]interface{}))
	ret.Inst.AuthenticationKey = d.Get("authentication_key").(string)
	ret.Inst.BfdCfg = getObjectInterfaceLifIpOspfOspfGlobalBfdCfg(d.Get("bfd_cfg").([]interface{}))
	ret.Inst.Cost = d.Get("cost").(int)
	ret.Inst.DatabaseFilterCfg = getObjectInterfaceLifIpOspfOspfGlobalDatabaseFilterCfg(d.Get("database_filter_cfg").([]interface{}))
	ret.Inst.DeadInterval = d.Get("dead_interval").(int)
	ret.Inst.Disable = d.Get("disable").(string)
	ret.Inst.HelloInterval = d.Get("hello_interval").(int)
	ret.Inst.MessageDigestCfg = getSliceInterfaceLifIpOspfOspfGlobalMessageDigestCfg(d.Get("message_digest_cfg").([]interface{}))
	ret.Inst.Mtu = d.Get("mtu").(int)
	ret.Inst.MtuIgnore = d.Get("mtu_ignore").(int)
	ret.Inst.Network = getObjectInterfaceLifIpOspfOspfGlobalNetwork(d.Get("network").([]interface{}))
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.RetransmitInterval = d.Get("retransmit_interval").(int)
	ret.Inst.TransmitDelay = d.Get("transmit_delay").(int)
	//omit uuid
	ret.Inst.Ifname = d.Get("ifname").(string)
	return ret
}
