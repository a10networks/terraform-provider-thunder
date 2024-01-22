package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernetIpOspfOspfGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ethernet_ip_ospf_ospf_global`: Global setting for Open Shortest Path First for IPv4 (OSPF)\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceEthernetIpOspfOspfGlobalCreate,
		UpdateContext: resourceInterfaceEthernetIpOspfOspfGlobalUpdate,
		ReadContext:   resourceInterfaceEthernetIpOspfOspfGlobalRead,
		DeleteContext: resourceInterfaceEthernetIpOspfOspfGlobalDelete,

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
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceEthernetIpOspfOspfGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpOspfOspfGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIpOspfOspfGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetIpOspfOspfGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceEthernetIpOspfOspfGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpOspfOspfGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIpOspfOspfGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetIpOspfOspfGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceEthernetIpOspfOspfGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpOspfOspfGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIpOspfOspfGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceEthernetIpOspfOspfGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpOspfOspfGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIpOspfOspfGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceEthernetIpOspfOspfGlobalAuthenticationCfg(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobalAuthenticationCfg {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobalAuthenticationCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = in["authentication"].(int)
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpOspfOspfGlobalBfdCfg(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobalBfdCfg {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobalBfdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectInterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DatabaseFilter = in["database_filter"].(string)
		ret.Out = in["out"].(int)
	}
	return ret
}

func getSliceInterfaceEthernetIpOspfOspfGlobalMessageDigestCfg(d []interface{}) []edpt.InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = getObjectInterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5(in["md5"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Md5Value = in["md5_value"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceEthernetIpOspfOspfGlobalNetwork(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobalNetwork {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobalNetwork
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

func dataToEndpointInterfaceEthernetIpOspfOspfGlobal(d *schema.ResourceData) edpt.InterfaceEthernetIpOspfOspfGlobal {
	var ret edpt.InterfaceEthernetIpOspfOspfGlobal
	ret.Inst.AuthenticationCfg = getObjectInterfaceEthernetIpOspfOspfGlobalAuthenticationCfg(d.Get("authentication_cfg").([]interface{}))
	ret.Inst.AuthenticationKey = d.Get("authentication_key").(string)
	ret.Inst.BfdCfg = getObjectInterfaceEthernetIpOspfOspfGlobalBfdCfg(d.Get("bfd_cfg").([]interface{}))
	ret.Inst.Cost = d.Get("cost").(int)
	ret.Inst.DatabaseFilterCfg = getObjectInterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg(d.Get("database_filter_cfg").([]interface{}))
	ret.Inst.DeadInterval = d.Get("dead_interval").(int)
	ret.Inst.Disable = d.Get("disable").(string)
	ret.Inst.HelloInterval = d.Get("hello_interval").(int)
	ret.Inst.MessageDigestCfg = getSliceInterfaceEthernetIpOspfOspfGlobalMessageDigestCfg(d.Get("message_digest_cfg").([]interface{}))
	ret.Inst.Mtu = d.Get("mtu").(int)
	ret.Inst.MtuIgnore = d.Get("mtu_ignore").(int)
	ret.Inst.Network = getObjectInterfaceEthernetIpOspfOspfGlobalNetwork(d.Get("network").([]interface{}))
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.RetransmitInterval = d.Get("retransmit_interval").(int)
	ret.Inst.TransmitDelay = d.Get("transmit_delay").(int)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
