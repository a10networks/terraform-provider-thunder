package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpThreatList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ip_threat_list`: Configure System IP Threat List\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIpThreatListCreate,
		UpdateContext: resourceSystemIpThreatListUpdate,
		ReadContext:   resourceSystemIpThreatListRead,
		DeleteContext: resourceSystemIpThreatListDelete,

		Schema: map[string]*schema.Schema{
			"ipv4_dest_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_list_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list": {
										Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name)",
									},
									"ip_threat_action_tmpl": {
										Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number)",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ipv4_internet_host_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"white_list": {
							Type: schema.TypeString, Optional: true, Description: "Bind exception-list (class-list name)",
						},
						"class_list_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list": {
										Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name)",
									},
									"ip_threat_action_tmpl": {
										Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number)",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ipv4_source_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_list_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list": {
										Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name)",
									},
									"ip_threat_action_tmpl": {
										Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number)",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ipv6_dest_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_list_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list": {
										Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name)",
									},
									"ip_threat_action_tmpl": {
										Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number)",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ipv6_internet_host_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"white_list": {
							Type: schema.TypeString, Optional: true, Description: "Bind exception-list (class-list name)",
						},
						"class_list_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list": {
										Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name)",
									},
									"ip_threat_action_tmpl": {
										Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number)",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ipv6_source_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_list_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list": {
										Type: schema.TypeString, Optional: true, Description: "Bind class-list (class-list name)",
									},
									"ip_threat_action_tmpl": {
										Type: schema.TypeInt, Optional: true, Description: "Bind ip-threat-action Template (ip-threat-action Template number)",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'packet_hit_count_in_sw': Packet Hit Count in SW; 'packet_hit_count_in_spe': Packet Hit Count in SPE; 'entries_added_in_sw': Entries Added in SW; 'entries_removed_from_sw': Entries Removed from SW; 'entries_added_in_spe': Entries Added in SPE; 'entries_removed_from_spe': Entries Removed from SPE; 'error_out_of_memory': Out of memory Error; 'error_out_of_spe_entries': Out of SPE Entries Error;",
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
func resourceSystemIpThreatListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpThreatListRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIpThreatListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpThreatListRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIpThreatListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIpThreatListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSystemIpThreatListIpv4DestList1682(d []interface{}) edpt.SystemIpThreatListIpv4DestList1682 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv4DestList1682
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv4DestListClassListCfg1683(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv4DestListClassListCfg1683(d []interface{}) []edpt.SystemIpThreatListIpv4DestListClassListCfg1683 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv4DestListClassListCfg1683, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv4DestListClassListCfg1683
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv4InternetHostList1684(d []interface{}) edpt.SystemIpThreatListIpv4InternetHostList1684 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv4InternetHostList1684
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.WhiteList = in["white_list"].(string)
		ret.ClassListCfg = getSliceSystemIpThreatListIpv4InternetHostListClassListCfg1685(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv4InternetHostListClassListCfg1685(d []interface{}) []edpt.SystemIpThreatListIpv4InternetHostListClassListCfg1685 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv4InternetHostListClassListCfg1685, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv4InternetHostListClassListCfg1685
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv4SourceList1686(d []interface{}) edpt.SystemIpThreatListIpv4SourceList1686 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv4SourceList1686
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv4SourceListClassListCfg1687(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv4SourceListClassListCfg1687(d []interface{}) []edpt.SystemIpThreatListIpv4SourceListClassListCfg1687 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv4SourceListClassListCfg1687, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv4SourceListClassListCfg1687
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv6DestList1688(d []interface{}) edpt.SystemIpThreatListIpv6DestList1688 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv6DestList1688
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv6DestListClassListCfg1689(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv6DestListClassListCfg1689(d []interface{}) []edpt.SystemIpThreatListIpv6DestListClassListCfg1689 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv6DestListClassListCfg1689, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv6DestListClassListCfg1689
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv6InternetHostList1690(d []interface{}) edpt.SystemIpThreatListIpv6InternetHostList1690 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv6InternetHostList1690
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.WhiteList = in["white_list"].(string)
		ret.ClassListCfg = getSliceSystemIpThreatListIpv6InternetHostListClassListCfg1691(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv6InternetHostListClassListCfg1691(d []interface{}) []edpt.SystemIpThreatListIpv6InternetHostListClassListCfg1691 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv6InternetHostListClassListCfg1691, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv6InternetHostListClassListCfg1691
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv6SourceList1692(d []interface{}) edpt.SystemIpThreatListIpv6SourceList1692 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv6SourceList1692
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv6SourceListClassListCfg1693(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv6SourceListClassListCfg1693(d []interface{}) []edpt.SystemIpThreatListIpv6SourceListClassListCfg1693 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv6SourceListClassListCfg1693, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv6SourceListClassListCfg1693
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemIpThreatListSamplingEnable(d []interface{}) []edpt.SystemIpThreatListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemIpThreatList(d *schema.ResourceData) edpt.SystemIpThreatList {
	var ret edpt.SystemIpThreatList
	ret.Inst.Ipv4DestList = getObjectSystemIpThreatListIpv4DestList1682(d.Get("ipv4_dest_list").([]interface{}))
	ret.Inst.Ipv4InternetHostList = getObjectSystemIpThreatListIpv4InternetHostList1684(d.Get("ipv4_internet_host_list").([]interface{}))
	ret.Inst.Ipv4SourceList = getObjectSystemIpThreatListIpv4SourceList1686(d.Get("ipv4_source_list").([]interface{}))
	ret.Inst.Ipv6DestList = getObjectSystemIpThreatListIpv6DestList1688(d.Get("ipv6_dest_list").([]interface{}))
	ret.Inst.Ipv6InternetHostList = getObjectSystemIpThreatListIpv6InternetHostList1690(d.Get("ipv6_internet_host_list").([]interface{}))
	ret.Inst.Ipv6SourceList = getObjectSystemIpThreatListIpv6SourceList1692(d.Get("ipv6_source_list").([]interface{}))
	ret.Inst.SamplingEnable = getSliceSystemIpThreatListSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
