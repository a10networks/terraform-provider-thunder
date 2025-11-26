package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcGeoLocationAppType() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_geo_location_app_type`: DDOS APP type\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcGeoLocationAppTypeCreate,
		UpdateContext: resourceDdosSrcGeoLocationAppTypeUpdate,
		ReadContext:   resourceDdosSrcGeoLocationAppTypeRead,
		DeleteContext: resourceDdosSrcGeoLocationAppTypeDelete,

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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"geolocation_name": {
				Type: schema.TypeString, Required: true, Description: "GeolocationName",
			},
		},
	}
}
func resourceDdosSrcGeoLocationAppTypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcGeoLocationAppTypeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcGeoLocationAppType(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcGeoLocationAppTypeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcGeoLocationAppTypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcGeoLocationAppTypeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcGeoLocationAppType(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcGeoLocationAppTypeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcGeoLocationAppTypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcGeoLocationAppTypeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcGeoLocationAppType(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcGeoLocationAppTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcGeoLocationAppTypeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcGeoLocationAppType(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosSrcGeoLocationAppTypeTemplate(d []interface{}) edpt.DdosSrcGeoLocationAppTypeTemplate {

	count1 := len(d)
	var ret edpt.DdosSrcGeoLocationAppTypeTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func dataToEndpointDdosSrcGeoLocationAppType(d *schema.ResourceData) edpt.DdosSrcGeoLocationAppType {
	var ret edpt.DdosSrcGeoLocationAppType
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Template = getObjectDdosSrcGeoLocationAppTypeTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.GeolocationName = d.Get("geolocation_name").(string)
	return ret
}
