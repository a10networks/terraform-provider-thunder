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

func getObjectSystemIpThreatListIpv4DestList1576(d []interface{}) edpt.SystemIpThreatListIpv4DestList1576 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv4DestList1576
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv4DestListClassListCfg1577(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv4DestListClassListCfg1577(d []interface{}) []edpt.SystemIpThreatListIpv4DestListClassListCfg1577 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv4DestListClassListCfg1577, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv4DestListClassListCfg1577
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv4InternetHostList1578(d []interface{}) edpt.SystemIpThreatListIpv4InternetHostList1578 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv4InternetHostList1578
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv4InternetHostListClassListCfg1579(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv4InternetHostListClassListCfg1579(d []interface{}) []edpt.SystemIpThreatListIpv4InternetHostListClassListCfg1579 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv4InternetHostListClassListCfg1579, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv4InternetHostListClassListCfg1579
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv4SourceList1580(d []interface{}) edpt.SystemIpThreatListIpv4SourceList1580 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv4SourceList1580
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv4SourceListClassListCfg1581(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv4SourceListClassListCfg1581(d []interface{}) []edpt.SystemIpThreatListIpv4SourceListClassListCfg1581 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv4SourceListClassListCfg1581, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv4SourceListClassListCfg1581
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv6DestList1582(d []interface{}) edpt.SystemIpThreatListIpv6DestList1582 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv6DestList1582
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv6DestListClassListCfg1583(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv6DestListClassListCfg1583(d []interface{}) []edpt.SystemIpThreatListIpv6DestListClassListCfg1583 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv6DestListClassListCfg1583, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv6DestListClassListCfg1583
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv6InternetHostList1584(d []interface{}) edpt.SystemIpThreatListIpv6InternetHostList1584 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv6InternetHostList1584
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv6InternetHostListClassListCfg1585(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv6InternetHostListClassListCfg1585(d []interface{}) []edpt.SystemIpThreatListIpv6InternetHostListClassListCfg1585 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv6InternetHostListClassListCfg1585, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv6InternetHostListClassListCfg1585
		oi.ClassList = in["class_list"].(string)
		oi.IpThreatActionTmpl = in["ip_threat_action_tmpl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemIpThreatListIpv6SourceList1586(d []interface{}) edpt.SystemIpThreatListIpv6SourceList1586 {

	count1 := len(d)
	var ret edpt.SystemIpThreatListIpv6SourceList1586
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListCfg = getSliceSystemIpThreatListIpv6SourceListClassListCfg1587(in["class_list_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemIpThreatListIpv6SourceListClassListCfg1587(d []interface{}) []edpt.SystemIpThreatListIpv6SourceListClassListCfg1587 {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListIpv6SourceListClassListCfg1587, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListIpv6SourceListClassListCfg1587
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
	ret.Inst.Ipv4DestList = getObjectSystemIpThreatListIpv4DestList1576(d.Get("ipv4_dest_list").([]interface{}))
	ret.Inst.Ipv4InternetHostList = getObjectSystemIpThreatListIpv4InternetHostList1578(d.Get("ipv4_internet_host_list").([]interface{}))
	ret.Inst.Ipv4SourceList = getObjectSystemIpThreatListIpv4SourceList1580(d.Get("ipv4_source_list").([]interface{}))
	ret.Inst.Ipv6DestList = getObjectSystemIpThreatListIpv6DestList1582(d.Get("ipv6_dest_list").([]interface{}))
	ret.Inst.Ipv6InternetHostList = getObjectSystemIpThreatListIpv6InternetHostList1584(d.Get("ipv6_internet_host_list").([]interface{}))
	ret.Inst.Ipv6SourceList = getObjectSystemIpThreatListIpv6SourceList1586(d.Get("ipv6_source_list").([]interface{}))
	ret.Inst.SamplingEnable = getSliceSystemIpThreatListSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
