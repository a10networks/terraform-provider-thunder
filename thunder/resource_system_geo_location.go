package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGeoLocation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_geo_location`: Configure system global geo-location\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemGeoLocationCreate,
		UpdateContext: resourceSystemGeoLocationUpdate,
		ReadContext:   resourceSystemGeoLocationRead,
		DeleteContext: resourceSystemGeoLocationDelete,

		Schema: map[string]*schema.Schema{
			"entry_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_locn_obj_name": {
							Type: schema.TypeString, Required: true, Description: "Specify geo-location name, section range is (1-15)",
						},
						"geo_locn_multiple_addresses": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"first_ip_address": {
										Type: schema.TypeString, Optional: true, Description: "Specify IP information (Specify IP address)",
									},
									"geol_ipv4_mask": {
										Type: schema.TypeString, Optional: true, Description: "Specify IPv4 mask",
									},
									"ip_addr2": {
										Type: schema.TypeString, Optional: true, Description: "Specify IP address range",
									},
									"first_ipv6_address": {
										Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address",
									},
									"geol_ipv6_mask": {
										Type: schema.TypeInt, Optional: true, Description: "Specify IPv6 mask",
									},
									"ipv6_addr2": {
										Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address range",
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
			"geo_location_geolite2_asn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load built-in Maxmind GeoLite2-ASN database. Database available from http://www.maxmind.com",
			},
			"geo_location_geolite2_city": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load built-in Maxmind GeoLite2-City database. Database available from http://www.maxmind.com",
			},
			"geo_location_geolite2_country": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load built-in Maxmind GeoLite2-Country database. Database available from http://www.maxmind.com",
			},
			"geo_location_iana": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Load built-in IANA Database",
			},
			"geo_location_iana_system": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load built-in IANA Database",
			},
			"geolite2_asn_include_ipv6": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include IPv6 address",
			},
			"geolite2_city_include_ipv6": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include IPv6 address",
			},
			"geolite2_country_include_ipv6": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include IPv6 address",
			},
			"geoloc_load_file_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_location_load_filename": {
							Type: schema.TypeString, Optional: true, Description: "Specify file to be loaded",
						},
						"geo_location_load_file_include_ipv6": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include IPv6 address",
						},
						"template_name": {
							Type: schema.TypeString, Optional: true, Description: "CSV template to load this file",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemGeoLocationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeoLocationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeoLocation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemGeoLocationRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemGeoLocationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeoLocationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeoLocation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemGeoLocationRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemGeoLocationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeoLocationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeoLocation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemGeoLocationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeoLocationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeoLocation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemGeoLocationEntryList(d []interface{}) []edpt.SystemGeoLocationEntryList {

	count1 := len(d)
	ret := make([]edpt.SystemGeoLocationEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeoLocationEntryList
		oi.GeoLocnObjName = in["geo_locn_obj_name"].(string)
		oi.GeoLocnMultipleAddresses = getSliceSystemGeoLocationEntryListGeoLocnMultipleAddresses(in["geo_locn_multiple_addresses"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemGeoLocationEntryListGeoLocnMultipleAddresses(d []interface{}) []edpt.SystemGeoLocationEntryListGeoLocnMultipleAddresses {

	count1 := len(d)
	ret := make([]edpt.SystemGeoLocationEntryListGeoLocnMultipleAddresses, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeoLocationEntryListGeoLocnMultipleAddresses
		oi.FirstIpAddress = in["first_ip_address"].(string)
		oi.GeolIpv4Mask = in["geol_ipv4_mask"].(string)
		oi.IpAddr2 = in["ip_addr2"].(string)
		oi.FirstIpv6Address = in["first_ipv6_address"].(string)
		oi.GeolIpv6Mask = in["geol_ipv6_mask"].(int)
		oi.Ipv6Addr2 = in["ipv6_addr2"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemGeoLocationGeolocLoadFileList(d []interface{}) []edpt.SystemGeoLocationGeolocLoadFileList {

	count1 := len(d)
	ret := make([]edpt.SystemGeoLocationGeolocLoadFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeoLocationGeolocLoadFileList
		oi.GeoLocationLoadFilename = in["geo_location_load_filename"].(string)
		oi.GeoLocationLoadFileIncludeIpv6 = in["geo_location_load_file_include_ipv6"].(int)
		oi.TemplateName = in["template_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemGeoLocation(d *schema.ResourceData) edpt.SystemGeoLocation {
	var ret edpt.SystemGeoLocation
	ret.Inst.EntryList = getSliceSystemGeoLocationEntryList(d.Get("entry_list").([]interface{}))
	ret.Inst.GeoLocationGeolite2Asn = d.Get("geo_location_geolite2_asn").(int)
	ret.Inst.GeoLocationGeolite2City = d.Get("geo_location_geolite2_city").(int)
	ret.Inst.GeoLocationGeolite2Country = d.Get("geo_location_geolite2_country").(int)
	ret.Inst.GeoLocationIana = d.Get("geo_location_iana").(int)
	ret.Inst.GeoLocationIanaSystem = d.Get("geo_location_iana_system").(int)
	ret.Inst.Geolite2AsnIncludeIpv6 = d.Get("geolite2_asn_include_ipv6").(int)
	ret.Inst.Geolite2CityIncludeIpv6 = d.Get("geolite2_city_include_ipv6").(int)
	ret.Inst.Geolite2CountryIncludeIpv6 = d.Get("geolite2_country_include_ipv6").(int)
	ret.Inst.GeolocLoadFileList = getSliceSystemGeoLocationGeolocLoadFileList(d.Get("geoloc_load_file_list").([]interface{}))
	//omit uuid
	return ret
}
