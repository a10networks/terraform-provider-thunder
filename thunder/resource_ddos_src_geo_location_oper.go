package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcGeoLocationOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_src_geo_location_oper`: Operational Status for the object geo-location\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosSrcGeoLocationOperRead,

		Schema: map[string]*schema.Schema{
			"geolocation_name": {
				Type: schema.TypeString, Required: true, Description: "Geolocation Name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ddos_entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_address_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"src_address_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"state_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"level_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_connections": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"connection_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_connection_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"connection_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_packet_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"packet_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_kbit_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"kbit_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_frag_packet_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"frag_packet_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat1": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat1_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat2": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat2_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat3": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat3_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat4": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat4_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat5": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat5_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat6": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat6_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat7": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat7_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat8": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat8_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"age_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"lockup_time_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dynamic_entry_count": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dynamic_entry_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sflow_source_id": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"debug_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"entry_displayed_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"service_displayed_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all_entries": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"l4_type_str": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"app_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"black_listed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"white_listed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"authenticated": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all_l4_types": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"l4_ext_rate": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hw_blacklisted": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceDdosSrcGeoLocationOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcGeoLocationOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcGeoLocationOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosSrcGeoLocationOperOper := setObjectDdosSrcGeoLocationOperOper(res)
		d.Set("oper", DdosSrcGeoLocationOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosSrcGeoLocationOperOper(ret edpt.DataDdosSrcGeoLocationOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_entry_list":         setSliceDdosSrcGeoLocationOperOperDdos_entry_list(ret.DtDdosSrcGeoLocationOper.Oper.Ddos_entry_list),
			"entry_displayed_count":   ret.DtDdosSrcGeoLocationOper.Oper.EntryDisplayedCount,
			"service_displayed_count": ret.DtDdosSrcGeoLocationOper.Oper.ServiceDisplayedCount,
			"all_entries":             ret.DtDdosSrcGeoLocationOper.Oper.AllEntries,
			"l4_type_str":             ret.DtDdosSrcGeoLocationOper.Oper.L4TypeStr,
			"app_type":                ret.DtDdosSrcGeoLocationOper.Oper.AppType,
			"exceeded":                ret.DtDdosSrcGeoLocationOper.Oper.Exceeded,
			"black_listed":            ret.DtDdosSrcGeoLocationOper.Oper.BlackListed,
			"white_listed":            ret.DtDdosSrcGeoLocationOper.Oper.WhiteListed,
			"authenticated":           ret.DtDdosSrcGeoLocationOper.Oper.Authenticated,
			"all_l4_types":            ret.DtDdosSrcGeoLocationOper.Oper.AllL4Types,
			"l4_ext_rate":             ret.DtDdosSrcGeoLocationOper.Oper.L4ExtRate,
			"hw_blacklisted":          ret.DtDdosSrcGeoLocationOper.Oper.HwBlacklisted,
		},
	}
}

func setSliceDdosSrcGeoLocationOperOperDdos_entry_list(d []edpt.DdosSrcGeoLocationOperOperDdos_entry_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dst_address_str"] = item.DstAddressStr
		in["src_address_str"] = item.SrcAddressStr
		in["port_str"] = item.PortStr
		in["state_str"] = item.StateStr
		in["level_str"] = item.LevelStr
		in["current_connections"] = item.CurrentConnections
		in["connection_limit"] = item.ConnectionLimit
		in["current_connection_rate"] = item.CurrentConnectionRate
		in["connection_rate_limit"] = item.ConnectionRateLimit
		in["current_packet_rate"] = item.CurrentPacketRate
		in["packet_rate_limit"] = item.PacketRateLimit
		in["current_kbit_rate"] = item.CurrentKbitRate
		in["kbit_rate_limit"] = item.KbitRateLimit
		in["current_frag_packet_rate"] = item.CurrentFragPacketRate
		in["frag_packet_rate_limit"] = item.FragPacketRateLimit
		in["current_app_stat1"] = item.CurrentAppStat1
		in["app_stat1_limit"] = item.AppStat1Limit
		in["current_app_stat2"] = item.CurrentAppStat2
		in["app_stat2_limit"] = item.AppStat2Limit
		in["current_app_stat3"] = item.CurrentAppStat3
		in["app_stat3_limit"] = item.AppStat3Limit
		in["current_app_stat4"] = item.CurrentAppStat4
		in["app_stat4_limit"] = item.AppStat4Limit
		in["current_app_stat5"] = item.CurrentAppStat5
		in["app_stat5_limit"] = item.AppStat5Limit
		in["current_app_stat6"] = item.CurrentAppStat6
		in["app_stat6_limit"] = item.AppStat6Limit
		in["current_app_stat7"] = item.CurrentAppStat7
		in["app_stat7_limit"] = item.AppStat7Limit
		in["current_app_stat8"] = item.CurrentAppStat8
		in["app_stat8_limit"] = item.AppStat8Limit
		in["age_str"] = item.AgeStr
		in["lockup_time_str"] = item.LockupTimeStr
		in["dynamic_entry_count"] = item.DynamicEntryCount
		in["dynamic_entry_limit"] = item.DynamicEntryLimit
		in["sflow_source_id"] = item.SflowSourceId
		in["debug_str"] = item.DebugStr
		result = append(result, in)
	}
	return result
}

func getObjectDdosSrcGeoLocationOperOper(d []interface{}) edpt.DdosSrcGeoLocationOperOper {

	count1 := len(d)
	var ret edpt.DdosSrcGeoLocationOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosSrcGeoLocationOperOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
		ret.EntryDisplayedCount = in["entry_displayed_count"].(int)
		ret.ServiceDisplayedCount = in["service_displayed_count"].(int)
		ret.AllEntries = in["all_entries"].(int)
		ret.L4TypeStr = in["l4_type_str"].(string)
		ret.AppType = in["app_type"].(string)
		ret.Exceeded = in["exceeded"].(int)
		ret.BlackListed = in["black_listed"].(int)
		ret.WhiteListed = in["white_listed"].(int)
		ret.Authenticated = in["authenticated"].(int)
		ret.AllL4Types = in["all_l4_types"].(int)
		ret.L4ExtRate = in["l4_ext_rate"].(string)
		ret.HwBlacklisted = in["hw_blacklisted"].(string)
	}
	return ret
}

func getSliceDdosSrcGeoLocationOperOperDdos_entry_list(d []interface{}) []edpt.DdosSrcGeoLocationOperOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosSrcGeoLocationOperOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSrcGeoLocationOperOperDdos_entry_list
		oi.DstAddressStr = in["dst_address_str"].(string)
		oi.SrcAddressStr = in["src_address_str"].(string)
		oi.PortStr = in["port_str"].(string)
		oi.StateStr = in["state_str"].(string)
		oi.LevelStr = in["level_str"].(string)
		oi.CurrentConnections = in["current_connections"].(string)
		oi.ConnectionLimit = in["connection_limit"].(string)
		oi.CurrentConnectionRate = in["current_connection_rate"].(string)
		oi.ConnectionRateLimit = in["connection_rate_limit"].(string)
		oi.CurrentPacketRate = in["current_packet_rate"].(string)
		oi.PacketRateLimit = in["packet_rate_limit"].(string)
		oi.CurrentKbitRate = in["current_kbit_rate"].(string)
		oi.KbitRateLimit = in["kbit_rate_limit"].(string)
		oi.CurrentFragPacketRate = in["current_frag_packet_rate"].(string)
		oi.FragPacketRateLimit = in["frag_packet_rate_limit"].(string)
		oi.CurrentAppStat1 = in["current_app_stat1"].(string)
		oi.AppStat1Limit = in["app_stat1_limit"].(string)
		oi.CurrentAppStat2 = in["current_app_stat2"].(string)
		oi.AppStat2Limit = in["app_stat2_limit"].(string)
		oi.CurrentAppStat3 = in["current_app_stat3"].(string)
		oi.AppStat3Limit = in["app_stat3_limit"].(string)
		oi.CurrentAppStat4 = in["current_app_stat4"].(string)
		oi.AppStat4Limit = in["app_stat4_limit"].(string)
		oi.CurrentAppStat5 = in["current_app_stat5"].(string)
		oi.AppStat5Limit = in["app_stat5_limit"].(string)
		oi.CurrentAppStat6 = in["current_app_stat6"].(string)
		oi.AppStat6Limit = in["app_stat6_limit"].(string)
		oi.CurrentAppStat7 = in["current_app_stat7"].(string)
		oi.AppStat7Limit = in["app_stat7_limit"].(string)
		oi.CurrentAppStat8 = in["current_app_stat8"].(string)
		oi.AppStat8Limit = in["app_stat8_limit"].(string)
		oi.AgeStr = in["age_str"].(string)
		oi.LockupTimeStr = in["lockup_time_str"].(string)
		oi.DynamicEntryCount = in["dynamic_entry_count"].(string)
		oi.DynamicEntryLimit = in["dynamic_entry_limit"].(string)
		oi.SflowSourceId = in["sflow_source_id"].(string)
		oi.DebugStr = in["debug_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosSrcGeoLocationOper(d *schema.ResourceData) edpt.DdosSrcGeoLocationOper {
	var ret edpt.DdosSrcGeoLocationOper

	ret.GeolocationName = d.Get("geolocation_name").(string)

	ret.Oper = getObjectDdosSrcGeoLocationOperOper(d.Get("oper").([]interface{}))
	return ret
}
