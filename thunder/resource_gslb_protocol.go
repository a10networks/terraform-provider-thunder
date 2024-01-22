package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbProtocol() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_protocol`: Specify GSLB Message Protocol parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbProtocolCreate,
		UpdateContext: resourceGslbProtocolUpdate,
		ReadContext:   resourceGslbProtocolRead,
		DeleteContext: resourceGslbProtocolDelete,

		Schema: map[string]*schema.Schema{
			"auto_detect": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Automatically detect SLB Config",
			},
			"disable_new_gslb_sync": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable new gslb config sync",
			},
			"enable_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Required: true, Description: "'controller': Enable/Disable GSLB protocol as GSLB controller; 'device': Enable/Disable GSLB protocol as site device;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"limit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ardt_query": {
							Type: schema.TypeInt, Optional: true, Default: 200, Description: "Query Messages of Active RDT, default is 200 (Number)",
						},
						"ardt_response": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "Response Messages of Active RDT, default is 1000 (Number)",
						},
						"ardt_session": {
							Type: schema.TypeInt, Optional: true, Default: 32768, Description: "Sessions of Active RDT, default is 32768 (Number)",
						},
						"conn_response": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Response Messages of Connection Load, default is no limit (Number)",
						},
						"response": {
							Type: schema.TypeInt, Optional: true, Default: 3600, Description: "Amount of Response Messages, default is 3600 (Number)",
						},
						"message": {
							Type: schema.TypeInt, Optional: true, Default: 10000, Description: "Amount of Messages, default is 10000 (Number)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"msg_format_acos_2x": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Run GSLB Protocol in compatible mode with a ACOS 2.x GSLB peer",
			},
			"ping_site": {
				Type: schema.TypeString, Optional: true, Description: "name of site or ip address to ping",
			},
			"secure": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Secure; 'disable': Disable Secure (default); 'enable-fallback': Fall back to non-secure if fail;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"status_interval": {
				Type: schema.TypeInt, Optional: true, Description: "Specify GSLB Message Protocol update period (The GSLB Protocol update interval (seconds), default is 30)",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port for connections in Shared Partition",
			},
			"use_mgmt_port_for_all_partitions": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port for connections in all L3v Partitions",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceGslbProtocolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocol(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbProtocolRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbProtocolUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocol(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbProtocolRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbProtocolDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocol(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbProtocolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbProtocolRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbProtocol(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbProtocolEnableList(d []interface{}) []edpt.GslbProtocolEnableList {

	count1 := len(d)
	ret := make([]edpt.GslbProtocolEnableList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbProtocolEnableList
		oi.Type = in["type"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbProtocolLimit394(d []interface{}) edpt.GslbProtocolLimit394 {

	count1 := len(d)
	var ret edpt.GslbProtocolLimit394
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ArdtQuery = in["ardt_query"].(int)
		ret.ArdtResponse = in["ardt_response"].(int)
		ret.ArdtSession = in["ardt_session"].(int)
		ret.ConnResponse = in["conn_response"].(int)
		ret.Response = in["response"].(int)
		ret.Message = in["message"].(int)
		//omit uuid
	}
	return ret
}

func getObjectGslbProtocolSecure395(d []interface{}) edpt.GslbProtocolSecure395 {

	count1 := len(d)
	var ret edpt.GslbProtocolSecure395
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointGslbProtocol(d *schema.ResourceData) edpt.GslbProtocol {
	var ret edpt.GslbProtocol
	ret.Inst.AutoDetect = d.Get("auto_detect").(int)
	ret.Inst.DisableNewGslbSync = d.Get("disable_new_gslb_sync").(int)
	ret.Inst.EnableList = getSliceGslbProtocolEnableList(d.Get("enable_list").([]interface{}))
	ret.Inst.Limit = getObjectGslbProtocolLimit394(d.Get("limit").([]interface{}))
	ret.Inst.MsgFormatAcos2x = d.Get("msg_format_acos_2x").(int)
	ret.Inst.PingSite = d.Get("ping_site").(string)
	ret.Inst.Secure = getObjectGslbProtocolSecure395(d.Get("secure").([]interface{}))
	ret.Inst.StatusInterval = d.Get("status_interval").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.UseMgmtPortForAllPartitions = d.Get("use_mgmt_port_for_all_partitions").(int)
	//omit uuid
	return ret
}
