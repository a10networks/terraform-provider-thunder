package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcGeoLocation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_geo_location`: Configure Static Geolocation Entry\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcGeoLocationCreate,
		UpdateContext: resourceDdosSrcGeoLocationUpdate,
		ReadContext:   resourceDdosSrcGeoLocationRead,
		DeleteContext: resourceDdosSrcGeoLocationDelete,

		Schema: map[string]*schema.Schema{
			"app_type_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'dns': dns; 'http': http; 'ssl-l4': ssl-l4; 'sip': sip;",
						},
						"template": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ssl_l4": {
										Type: schema.TypeString, Optional: true, Description: "DDOS SSL-L4 template",
									},
									"dns": {
										Type: schema.TypeString, Optional: true, Description: "DDOS DNS template",
									},
									"http": {
										Type: schema.TypeString, Optional: true, Description: "DDOS HTTP template",
									},
									"sip": {
										Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
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
			"bypass": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always permit for the Source to bypass all feature & limit checks",
			},
			"description": {
				Type: schema.TypeString, Optional: true, Description: "Description for this Geolocation Entry",
			},
			"geolocation_name": {
				Type: schema.TypeString, Required: true, Description: "Geolocation Name",
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"l4_type_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Description: "'permit': Whitelist incoming packets for protocol; 'deny': Blacklist incoming packets for protocol;",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"template": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tcp": {
										Type: schema.TypeString, Optional: true, Description: "DDOS TCP template",
									},
									"udp": {
										Type: schema.TypeString, Optional: true, Description: "DDOS UDP template",
									},
									"other": {
										Type: schema.TypeString, Optional: true, Description: "DDOS OTHER template",
									},
									"template_icmp_v4": {
										Type: schema.TypeString, Optional: true, Description: "DDOS ICMP template",
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
			"log_periodic": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable periodic log while event is continuing",
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logging": {
							Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
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
		},
	}
}
func resourceDdosSrcGeoLocationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcGeoLocationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcGeoLocation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcGeoLocationRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcGeoLocationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcGeoLocationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcGeoLocation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcGeoLocationRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcGeoLocationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcGeoLocationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcGeoLocation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcGeoLocationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcGeoLocationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcGeoLocation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosSrcGeoLocationAppTypeList(d []interface{}) []edpt.DdosSrcGeoLocationAppTypeList {

	count1 := len(d)
	ret := make([]edpt.DdosSrcGeoLocationAppTypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSrcGeoLocationAppTypeList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosSrcGeoLocationAppTypeListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosSrcGeoLocationAppTypeListTemplate(d []interface{}) edpt.DdosSrcGeoLocationAppTypeListTemplate {

	count1 := len(d)
	var ret edpt.DdosSrcGeoLocationAppTypeListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getSliceDdosSrcGeoLocationL4TypeList(d []interface{}) []edpt.DdosSrcGeoLocationL4TypeList {

	count1 := len(d)
	ret := make([]edpt.DdosSrcGeoLocationL4TypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSrcGeoLocationL4TypeList
		oi.Protocol = in["protocol"].(string)
		oi.Action = in["action"].(string)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosSrcGeoLocationL4TypeListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosSrcGeoLocationL4TypeListTemplate(d []interface{}) edpt.DdosSrcGeoLocationL4TypeListTemplate {

	count1 := len(d)
	var ret edpt.DdosSrcGeoLocationL4TypeListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Other = in["other"].(string)
		ret.TemplateIcmpV4 = in["template_icmp_v4"].(string)
	}
	return ret
}

func getObjectDdosSrcGeoLocationTemplate(d []interface{}) edpt.DdosSrcGeoLocationTemplate {

	count1 := len(d)
	var ret edpt.DdosSrcGeoLocationTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func dataToEndpointDdosSrcGeoLocation(d *schema.ResourceData) edpt.DdosSrcGeoLocation {
	var ret edpt.DdosSrcGeoLocation
	ret.Inst.AppTypeList = getSliceDdosSrcGeoLocationAppTypeList(d.Get("app_type_list").([]interface{}))
	ret.Inst.Bypass = d.Get("bypass").(int)
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.GeolocationName = d.Get("geolocation_name").(string)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.L4TypeList = getSliceDdosSrcGeoLocationL4TypeList(d.Get("l4_type_list").([]interface{}))
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.Template = getObjectDdosSrcGeoLocationTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
