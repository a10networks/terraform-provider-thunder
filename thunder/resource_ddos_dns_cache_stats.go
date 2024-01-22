package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsCacheStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dns_cache_stats`: Statistics for the object dns-cache\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDnsCacheStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "DNS Cache Instance Name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_cached_fqdn": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_cached_records": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fqdn_a": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fqdn_aaaa": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fqdn_cname": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fqdn_ns": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fqdn_mx": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fqdn_soa": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fqdn_srv": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fqdn_txt": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fqdn_ptr": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fqdn_other": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fqdn_wildcard": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fqdn_delegation": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"shard_size": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resp_ext_size": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"a_record": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aaaa_record": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cname_record": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ns_record": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mx_record": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"soa_record": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"srv_record": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"txt_record": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ptr_record": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"other_record": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fqdn_in_shard_filter": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceDdosDnsCacheStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDnsCacheStatsStats := setObjectDdosDnsCacheStatsStats(res)
		d.Set("stats", DdosDnsCacheStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDnsCacheStatsStats(ret edpt.DataDdosDnsCacheStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_cached_fqdn":    ret.DtDdosDnsCacheStats.Stats.TotalCachedFqdn,
			"total_cached_records": ret.DtDdosDnsCacheStats.Stats.TotalCachedRecords,
			"fqdn_a":               ret.DtDdosDnsCacheStats.Stats.FqdnA,
			"fqdn_aaaa":            ret.DtDdosDnsCacheStats.Stats.FqdnAaaa,
			"fqdn_cname":           ret.DtDdosDnsCacheStats.Stats.FqdnCname,
			"fqdn_ns":              ret.DtDdosDnsCacheStats.Stats.FqdnNs,
			"fqdn_mx":              ret.DtDdosDnsCacheStats.Stats.FqdnMx,
			"fqdn_soa":             ret.DtDdosDnsCacheStats.Stats.FqdnSoa,
			"fqdn_srv":             ret.DtDdosDnsCacheStats.Stats.FqdnSrv,
			"fqdn_txt":             ret.DtDdosDnsCacheStats.Stats.FqdnTxt,
			"fqdn_ptr":             ret.DtDdosDnsCacheStats.Stats.FqdnPtr,
			"fqdn_other":           ret.DtDdosDnsCacheStats.Stats.FqdnOther,
			"fqdn_wildcard":        ret.DtDdosDnsCacheStats.Stats.FqdnWildcard,
			"fqdn_delegation":      ret.DtDdosDnsCacheStats.Stats.FqdnDelegation,
			"shard_size":           ret.DtDdosDnsCacheStats.Stats.ShardSize,
			"resp_ext_size":        ret.DtDdosDnsCacheStats.Stats.RespExtSize,
			"a_record":             ret.DtDdosDnsCacheStats.Stats.ARecord,
			"aaaa_record":          ret.DtDdosDnsCacheStats.Stats.AaaaRecord,
			"cname_record":         ret.DtDdosDnsCacheStats.Stats.CnameRecord,
			"ns_record":            ret.DtDdosDnsCacheStats.Stats.NsRecord,
			"mx_record":            ret.DtDdosDnsCacheStats.Stats.MxRecord,
			"soa_record":           ret.DtDdosDnsCacheStats.Stats.SoaRecord,
			"srv_record":           ret.DtDdosDnsCacheStats.Stats.SrvRecord,
			"txt_record":           ret.DtDdosDnsCacheStats.Stats.TxtRecord,
			"ptr_record":           ret.DtDdosDnsCacheStats.Stats.PtrRecord,
			"other_record":         ret.DtDdosDnsCacheStats.Stats.OtherRecord,
			"fqdn_in_shard_filter": ret.DtDdosDnsCacheStats.Stats.FqdnInShardFilter,
		},
	}
}

func getObjectDdosDnsCacheStatsStats(d []interface{}) edpt.DdosDnsCacheStatsStats {

	count1 := len(d)
	var ret edpt.DdosDnsCacheStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TotalCachedFqdn = in["total_cached_fqdn"].(int)
		ret.TotalCachedRecords = in["total_cached_records"].(int)
		ret.FqdnA = in["fqdn_a"].(int)
		ret.FqdnAaaa = in["fqdn_aaaa"].(int)
		ret.FqdnCname = in["fqdn_cname"].(int)
		ret.FqdnNs = in["fqdn_ns"].(int)
		ret.FqdnMx = in["fqdn_mx"].(int)
		ret.FqdnSoa = in["fqdn_soa"].(int)
		ret.FqdnSrv = in["fqdn_srv"].(int)
		ret.FqdnTxt = in["fqdn_txt"].(int)
		ret.FqdnPtr = in["fqdn_ptr"].(int)
		ret.FqdnOther = in["fqdn_other"].(int)
		ret.FqdnWildcard = in["fqdn_wildcard"].(int)
		ret.FqdnDelegation = in["fqdn_delegation"].(int)
		ret.ShardSize = in["shard_size"].(int)
		ret.RespExtSize = in["resp_ext_size"].(int)
		ret.ARecord = in["a_record"].(int)
		ret.AaaaRecord = in["aaaa_record"].(int)
		ret.CnameRecord = in["cname_record"].(int)
		ret.NsRecord = in["ns_record"].(int)
		ret.MxRecord = in["mx_record"].(int)
		ret.SoaRecord = in["soa_record"].(int)
		ret.SrvRecord = in["srv_record"].(int)
		ret.TxtRecord = in["txt_record"].(int)
		ret.PtrRecord = in["ptr_record"].(int)
		ret.OtherRecord = in["other_record"].(int)
		ret.FqdnInShardFilter = in["fqdn_in_shard_filter"].(int)
	}
	return ret
}

func dataToEndpointDdosDnsCacheStats(d *schema.ResourceData) edpt.DdosDnsCacheStats {
	var ret edpt.DdosDnsCacheStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectDdosDnsCacheStatsStats(d.Get("stats").([]interface{}))
	return ret
}
