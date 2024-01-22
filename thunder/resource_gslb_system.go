package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSystem() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_system`: GSLB system options\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbSystemCreate,
		UpdateContext: resourceGslbSystemUpdate,
		ReadContext:   resourceGslbSystemRead,
		DeleteContext: resourceGslbSystemDelete,

		Schema: map[string]*schema.Schema{
			"age_interval": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Interval to age runtime statistics. 0: never age, default is 10 (Time, unit: sec, default is 10)",
			},
			"geo_location_iana": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Load built-in IANA table",
			},
			"gslb_group": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "GSLB Group",
			},
			"gslb_load_file_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_location_load_filename": {
							Type: schema.TypeString, Optional: true, Description: "Specify file to be loaded",
						},
						"template_name": {
							Type: schema.TypeString, Optional: true, Description: "CSV template to load this file",
						},
					},
				},
			},
			"gslb_service_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "GSLB Service-IP",
			},
			"gslb_site": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "GSLB Site",
			},
			"hostname": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "System's Network Name",
			},
			"ip_ttl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "TTL of IP packets, default is 0 (IP TTL value, default is 0)",
			},
			"module": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify Auto Map Module",
			},
			"slb_device": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SLB Device",
			},
			"slb_server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SLB Server",
			},
			"slb_virtual_server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SLB Virtual Server",
			},
			"ttl": {
				Type: schema.TypeInt, Optional: true, Default: 300, Description: "Specify Auto Map TTL (TTL, default is 300)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"wait": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable GSLB until timeout if system is not ready (Time, unit: sec, default is 0)",
			},
		},
	}
}
func resourceGslbSystemCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSystemCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSystem(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSystemRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbSystemUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSystemUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSystem(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSystemRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbSystemDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSystemDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSystem(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbSystemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSystemRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSystem(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbSystemGslbLoadFileList(d []interface{}) []edpt.GslbSystemGslbLoadFileList {

	count1 := len(d)
	ret := make([]edpt.GslbSystemGslbLoadFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSystemGslbLoadFileList
		oi.GeoLocationLoadFilename = in["geo_location_load_filename"].(string)
		oi.TemplateName = in["template_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbSystem(d *schema.ResourceData) edpt.GslbSystem {
	var ret edpt.GslbSystem
	ret.Inst.AgeInterval = d.Get("age_interval").(int)
	ret.Inst.GeoLocationIana = d.Get("geo_location_iana").(int)
	ret.Inst.GslbGroup = d.Get("gslb_group").(int)
	ret.Inst.GslbLoadFileList = getSliceGslbSystemGslbLoadFileList(d.Get("gslb_load_file_list").([]interface{}))
	ret.Inst.GslbServiceIp = d.Get("gslb_service_ip").(int)
	ret.Inst.GslbSite = d.Get("gslb_site").(int)
	ret.Inst.Hostname = d.Get("hostname").(int)
	ret.Inst.IpTtl = d.Get("ip_ttl").(int)
	ret.Inst.Module = d.Get("module").(int)
	ret.Inst.SlbDevice = d.Get("slb_device").(int)
	ret.Inst.SlbServer = d.Get("slb_server").(int)
	ret.Inst.SlbVirtualServer = d.Get("slb_virtual_server").(int)
	ret.Inst.Ttl = d.Get("ttl").(int)
	//omit uuid
	ret.Inst.Wait = d.Get("wait").(int)
	return ret
}
