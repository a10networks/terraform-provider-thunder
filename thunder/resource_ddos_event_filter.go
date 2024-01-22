package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosEventFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_event_filter`: Configure DDOS debug event filter\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosEventFilterCreate,
		UpdateContext: resourceDdosEventFilterUpdate,
		ReadContext:   resourceDdosEventFilterRead,
		DeleteContext: resourceDdosEventFilterDelete,

		Schema: map[string]*schema.Schema{
			"black_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"black_list_dst": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dst entry/port is black-listed",
						},
						"black_list_src": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Src entry/port is black-listed",
						},
					},
				},
			},
			"drop": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"drop_src": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet is dropped because of src",
						},
						"drop_dst": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet is dropped because of dst",
						},
						"drop_black_list": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet is dropped because of black-list",
						},
					},
				},
			},
			"filter_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"l4_type_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp;",
						},
						"tcp_auth": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tcp_auth_init": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet that inits syn-auth/action-on-ack",
									},
									"tcp_auth_pass": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet that passes syn-auth/action-on-ack",
									},
									"tcp_auth_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet that fails syn-auth/action-on-ack",
									},
								},
							},
						},
						"retrans_syn_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"retrans_syn": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "TCP SYN retransmission",
									},
									"retrans_syn_exceed": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "TCP SYN retransmission exceed",
									},
								},
							},
						},
						"out_of_seq": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "TCP out-of-seq pkts",
						},
						"zero_window": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "TCP zero window pkts",
						},
						"udp_auth": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"udp_auth_init": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet that inits spoof-detect",
									},
									"udp_auth_pass": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet that passes spoof-detect",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"white_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"white_list_dst": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dst entry/port is white-listed",
						},
						"white_list_src": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Src entry/port is white-listed",
						},
					},
				},
			},
		},
	}
}
func resourceDdosEventFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosEventFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosEventFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosEventFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosEventFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosEventFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosEventFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosEventFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosEventFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosEventFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosEventFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosEventFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosEventFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosEventFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosEventFilterBlackList(d []interface{}) edpt.DdosEventFilterBlackList {

	count1 := len(d)
	var ret edpt.DdosEventFilterBlackList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BlackListDst = in["black_list_dst"].(int)
		ret.BlackListSrc = in["black_list_src"].(int)
	}
	return ret
}

func getObjectDdosEventFilterDrop(d []interface{}) edpt.DdosEventFilterDrop {

	count1 := len(d)
	var ret edpt.DdosEventFilterDrop
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DropSrc = in["drop_src"].(int)
		ret.DropDst = in["drop_dst"].(int)
		ret.DropBlackList = in["drop_black_list"].(int)
	}
	return ret
}

func getSliceDdosEventFilterL4TypeList(d []interface{}) []edpt.DdosEventFilterL4TypeList {

	count1 := len(d)
	ret := make([]edpt.DdosEventFilterL4TypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosEventFilterL4TypeList
		oi.Protocol = in["protocol"].(string)
		oi.TcpAuth = getObjectDdosEventFilterL4TypeListTcpAuth(in["tcp_auth"].([]interface{}))
		oi.RetransSynCfg = getObjectDdosEventFilterL4TypeListRetransSynCfg(in["retrans_syn_cfg"].([]interface{}))
		oi.OutOfSeq = in["out_of_seq"].(int)
		oi.ZeroWindow = in["zero_window"].(int)
		oi.UdpAuth = getObjectDdosEventFilterL4TypeListUdpAuth(in["udp_auth"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosEventFilterL4TypeListTcpAuth(d []interface{}) edpt.DdosEventFilterL4TypeListTcpAuth {

	count1 := len(d)
	var ret edpt.DdosEventFilterL4TypeListTcpAuth
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TcpAuthInit = in["tcp_auth_init"].(int)
		ret.TcpAuthPass = in["tcp_auth_pass"].(int)
		ret.TcpAuthFail = in["tcp_auth_fail"].(int)
	}
	return ret
}

func getObjectDdosEventFilterL4TypeListRetransSynCfg(d []interface{}) edpt.DdosEventFilterL4TypeListRetransSynCfg {

	count1 := len(d)
	var ret edpt.DdosEventFilterL4TypeListRetransSynCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RetransSyn = in["retrans_syn"].(int)
		ret.RetransSynExceed = in["retrans_syn_exceed"].(int)
	}
	return ret
}

func getObjectDdosEventFilterL4TypeListUdpAuth(d []interface{}) edpt.DdosEventFilterL4TypeListUdpAuth {

	count1 := len(d)
	var ret edpt.DdosEventFilterL4TypeListUdpAuth
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UdpAuthInit = in["udp_auth_init"].(int)
		ret.UdpAuthPass = in["udp_auth_pass"].(int)
	}
	return ret
}

func getObjectDdosEventFilterWhiteList(d []interface{}) edpt.DdosEventFilterWhiteList {

	count1 := len(d)
	var ret edpt.DdosEventFilterWhiteList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.WhiteListDst = in["white_list_dst"].(int)
		ret.WhiteListSrc = in["white_list_src"].(int)
	}
	return ret
}

func dataToEndpointDdosEventFilter(d *schema.ResourceData) edpt.DdosEventFilter {
	var ret edpt.DdosEventFilter
	ret.Inst.BlackList = getObjectDdosEventFilterBlackList(d.Get("black_list").([]interface{}))
	ret.Inst.Drop = getObjectDdosEventFilterDrop(d.Get("drop").([]interface{}))
	ret.Inst.FilterName = d.Get("filter_name").(string)
	ret.Inst.L4TypeList = getSliceDdosEventFilterL4TypeList(d.Get("l4_type_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.WhiteList = getObjectDdosEventFilterWhiteList(d.Get("white_list").([]interface{}))
	return ret
}
