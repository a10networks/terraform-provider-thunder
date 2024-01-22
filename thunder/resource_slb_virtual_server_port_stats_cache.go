package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbVirtualServerPortStats() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_virtual_server_port_stats_cache`: Statistics for the object port\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbVirtualServerPortStatsCreate,
		UpdateContext: resourceSlbVirtualServerPortStatsUpdate,
		ReadContext:   resourceSlbVirtualServerPortStatsRead,
		DeleteContext: resourceSlbVirtualServerPortStatsDelete,

		Schema: map[string]*schema.Schema{
			"port_number": {
				Type: schema.TypeInt, Required: true, Description: "Port",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': TCP LB service; 'udp': UDP Port; 'others': for no tcp/udp protocol, do IP load balancing; 'diameter': diameter port; 'dns-tcp': DNS service over TCP; 'dns-udp': DNS service over UDP; 'fast-http': Fast HTTP Port; 'fix': FIX Port; 'ftp': File Transfer Protocol Port; 'ftp-proxy': ftp proxy port; 'http': HTTP Port; 'https': HTTPS port; 'imap': imap proxy port; 'mlb': Message based load balancing; 'mms': Microsoft Multimedia Service Port; 'mysql': mssql port; 'mssql': mssql; 'pop3': pop3 proxy port; 'radius': RADIUS Port; 'rtsp': Real Time Streaming Protocol Port; 'sip': Session initiation protocol over UDP; 'sip-tcp': Session initiation protocol over TCP; 'sips': Session initiation protocol over TLS; 'smpp-tcp': SMPP service over TCP; 'spdy': spdy port; 'spdys': spdys port; 'smtp': SMTP Port; 'mqtt': MQTT Port; 'mqtts': MQTTS Port; 'ssl-proxy': Generic SSL proxy; 'ssli': SSL insight; 'ssh': SSH Port; 'tcp-proxy': Generic TCP proxy; 'tftp': TFTP Port; 'fast-fix': Fast FIX port; 'http-over-quic': HTTP3-over-quic port;",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cache": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hits": {
										Type: schema.TypeInt, Optional: true, Description: "Cache hits",
									},
									"miss": {
										Type: schema.TypeInt, Optional: true, Description: "Cache misses",
									},
									"bytes_served": {
										Type: schema.TypeInt, Optional: true, Description: "Bytes served from cache",
									},
									"total_req": {
										Type: schema.TypeInt, Optional: true, Description: "Total requests received",
									},
									"caching_req": {
										Type: schema.TypeInt, Optional: true, Description: "Total requests to cache",
									},
									"nc_req_header": {
										Type: schema.TypeInt, Optional: true, Description: "slbTemplateCacheNcReqHeader, help nc_req_header",
									},
									"nc_res_header": {
										Type: schema.TypeInt, Optional: true, Description: "slbTemplateCacheNcResHeader, help nc_res_header",
									},
									"rv_success": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rv_failure": {
										Type: schema.TypeInt, Optional: true, Description: "slbTemplateCacheRvFailure, help rv_failure",
									},
									"ims_request": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nm_response": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_type_cl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_type_ce": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_type_304": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_type_other": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_no_compress": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_gzip": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_deflate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_other": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nocache_match": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"match": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalidate_match": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"content_toobig": {
										Type: schema.TypeInt, Optional: true, Description: "slbTemplateCacheContentToobig, help content_toobig",
									},
									"content_toosmall": {
										Type: schema.TypeInt, Optional: true, Description: "slbTemplateCacheContentToosmall, help content_toosmall",
									},
									"entry_create_failures": {
										Type: schema.TypeInt, Optional: true, Description: "slbTemplateCacheEntryCreateFailures, help entry_create_failures",
									},
									"mem_size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"entry_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"replaced_entry": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aging_entry": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cleaned_entry": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_type_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_save_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_br": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbVirtualServerPortStatsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStatsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortStatsRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbVirtualServerPortStatsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStatsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortStatsRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbVirtualServerPortStatsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStatsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbVirtualServerPortStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbVirtualServerPortStatsStats1475(d []interface{}) edpt.SlbVirtualServerPortStatsStats1475 {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStatsStats1475
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Cache = getObjectSlbVirtualServerPortStatsStatsCache(in["cache"].([]interface{}))
	}
	return ret
}

func getObjectSlbVirtualServerPortStatsStatsCache(d []interface{}) edpt.SlbVirtualServerPortStatsStatsCache {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStatsStatsCache
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
		ret.Miss = in["miss"].(int)
		ret.Bytes_served = in["bytes_served"].(int)
		ret.Total_req = in["total_req"].(int)
		ret.Caching_req = in["caching_req"].(int)
		ret.Nc_req_header = in["nc_req_header"].(int)
		ret.Nc_res_header = in["nc_res_header"].(int)
		ret.Rv_success = in["rv_success"].(int)
		ret.Rv_failure = in["rv_failure"].(int)
		ret.Ims_request = in["ims_request"].(int)
		ret.Nm_response = in["nm_response"].(int)
		ret.Rsp_type_cl = in["rsp_type_cl"].(int)
		ret.Rsp_type_ce = in["rsp_type_ce"].(int)
		ret.Rsp_type_304 = in["rsp_type_304"].(int)
		ret.Rsp_type_other = in["rsp_type_other"].(int)
		ret.Rsp_no_compress = in["rsp_no_compress"].(int)
		ret.Rsp_gzip = in["rsp_gzip"].(int)
		ret.Rsp_deflate = in["rsp_deflate"].(int)
		ret.Rsp_other = in["rsp_other"].(int)
		ret.Nocache_match = in["nocache_match"].(int)
		ret.Match = in["match"].(int)
		ret.Invalidate_match = in["invalidate_match"].(int)
		ret.Content_toobig = in["content_toobig"].(int)
		ret.Content_toosmall = in["content_toosmall"].(int)
		ret.Entry_create_failures = in["entry_create_failures"].(int)
		ret.Mem_size = in["mem_size"].(int)
		ret.Entry_num = in["entry_num"].(int)
		ret.Replaced_entry = in["replaced_entry"].(int)
		ret.Aging_entry = in["aging_entry"].(int)
		ret.Cleaned_entry = in["cleaned_entry"].(int)
		ret.Rsp_type_stream = in["rsp_type_stream"].(int)
		ret.Header_save_error = in["header_save_error"].(int)
		ret.Rsp_br = in["rsp_br"].(int)
	}
	return ret
}

func dataToEndpointSlbVirtualServerPortStats(d *schema.ResourceData) edpt.SlbVirtualServerPortStats {
	var ret edpt.SlbVirtualServerPortStats
	ret.Inst.PortNumber = d.Get("port_number").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Stats = getObjectSlbVirtualServerPortStatsStats1475(d.Get("stats").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
