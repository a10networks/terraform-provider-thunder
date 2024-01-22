package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDiameter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_diameter`: diameter template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDiameterCreate,
		UpdateContext: resourceSlbTemplateDiameterUpdate,
		ReadContext:   resourceSlbTemplateDiameterRead,
		DeleteContext: resourceSlbTemplateDiameterDelete,

		Schema: map[string]*schema.Schema{
			"avp_code": {
				Type: schema.TypeInt, Optional: true, Description: "avp code",
			},
			"avp_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"avp": {
							Type: schema.TypeInt, Optional: true, Description: "customize avps for cer to the server (avp number)",
						},
						"int32": {
							Type: schema.TypeInt, Optional: true, Description: "32 bits integer",
						},
						"int64": {
							Type: schema.TypeInt, Optional: true, Description: "64 bits integer",
						},
						"string": {
							Type: schema.TypeString, Optional: true, Description: "String (string name, max length 127 bytes)",
						},
						"mandatory": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "mandatory avp",
						},
					},
				},
			},
			"avp_string": {
				Type: schema.TypeString, Optional: true, Description: "pattern to be matched in the avp string name, max length 127 bytes",
			},
			"customize_cea": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "customizing cea response",
			},
			"dwr_time": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "dwr health-check timer interval (in 100 milli second unit, default is 100, 0 means unset this option)",
			},
			"dwr_up_retry": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "number of successful dwr health-check before declaring target up",
			},
			"forward_to_latest_server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward client message to the latest server that sends message with the same session id",
			},
			"forward_unknown_session_id": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward server message even it has unknown session id",
			},
			"idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "user sesison idle timeout (in minutes, default is 5)",
			},
			"load_balance_on_session_id": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load balance based on the session id",
			},
			"message_code_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message_code": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"multiple_origin_host": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "allowing multiple origin-host to a single server",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "diameter template Name",
			},
			"origin_host": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"origin_host_name": {
							Type: schema.TypeString, Optional: true, Description: "origin-host name avp",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"origin_realm": {
				Type: schema.TypeString, Optional: true, Description: "origin-realm name avp",
			},
			"product_name": {
				Type: schema.TypeString, Optional: true, Description: "product name avp",
			},
			"relaxed_origin_host": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Relaxed Origin-Host Format",
			},
			"service_group_name": {
				Type: schema.TypeString, Optional: true, Description: "service group name, this is the service group that the message needs to be copied to",
			},
			"session_age": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "user session age allowed (default 10), this is not idle-time (in minutes)",
			},
			"terminate_on_cca_t": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "remove diameter session when receiving CCA-T message",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vendor_id": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "vendor-id avp (Vendor Id)",
			},
		},
	}
}
func resourceSlbTemplateDiameterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDiameterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDiameter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDiameterRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDiameterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDiameterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDiameter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDiameterRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDiameterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDiameterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDiameter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDiameterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDiameterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDiameter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateDiameterAvpList(d []interface{}) []edpt.SlbTemplateDiameterAvpList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDiameterAvpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDiameterAvpList
		oi.Avp = in["avp"].(int)
		oi.Int32 = in["int32"].(int)
		oi.Int64 = in["int64"].(int)
		oi.String = in["string"].(string)
		oi.Mandatory = in["mandatory"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateDiameterMessageCodeList(d []interface{}) []edpt.SlbTemplateDiameterMessageCodeList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDiameterMessageCodeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDiameterMessageCodeList
		oi.MessageCode = in["message_code"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDiameterOriginHost1416(d []interface{}) edpt.SlbTemplateDiameterOriginHost1416 {

	count1 := len(d)
	var ret edpt.SlbTemplateDiameterOriginHost1416
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OriginHostName = in["origin_host_name"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointSlbTemplateDiameter(d *schema.ResourceData) edpt.SlbTemplateDiameter {
	var ret edpt.SlbTemplateDiameter
	ret.Inst.AvpCode = d.Get("avp_code").(int)
	ret.Inst.AvpList = getSliceSlbTemplateDiameterAvpList(d.Get("avp_list").([]interface{}))
	ret.Inst.AvpString = d.Get("avp_string").(string)
	ret.Inst.CustomizeCea = d.Get("customize_cea").(int)
	ret.Inst.DwrTime = d.Get("dwr_time").(int)
	ret.Inst.DwrUpRetry = d.Get("dwr_up_retry").(int)
	ret.Inst.ForwardToLatestServer = d.Get("forward_to_latest_server").(int)
	ret.Inst.ForwardUnknownSessionId = d.Get("forward_unknown_session_id").(int)
	ret.Inst.IdleTimeout = d.Get("idle_timeout").(int)
	ret.Inst.LoadBalanceOnSessionId = d.Get("load_balance_on_session_id").(int)
	ret.Inst.MessageCodeList = getSliceSlbTemplateDiameterMessageCodeList(d.Get("message_code_list").([]interface{}))
	ret.Inst.MultipleOriginHost = d.Get("multiple_origin_host").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.OriginHost = getObjectSlbTemplateDiameterOriginHost1416(d.Get("origin_host").([]interface{}))
	ret.Inst.OriginRealm = d.Get("origin_realm").(string)
	ret.Inst.ProductName = d.Get("product_name").(string)
	ret.Inst.RelaxedOriginHost = d.Get("relaxed_origin_host").(int)
	ret.Inst.ServiceGroupName = d.Get("service_group_name").(string)
	ret.Inst.SessionAge = d.Get("session_age").(int)
	ret.Inst.TerminateOnCcaT = d.Get("terminate_on_cca_t").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VendorId = d.Get("vendor_id").(int)
	return ret
}
