package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Logging() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_logging`: CGNV6 Logging Statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LoggingCreate,
		UpdateContext: resourceCgnv6LoggingUpdate,
		ReadContext:   resourceCgnv6LoggingRead,
		DeleteContext: resourceCgnv6LoggingDelete,

		Schema: map[string]*schema.Schema{
			"nat_quota_exceeded": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": {
							Type: schema.TypeString, Optional: true, Default: "warning", Description: "'warning': Log level Warning (Default); 'critical': Log level Critical; 'notice': Log level Notice;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"nat_resource_exhausted": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": {
							Type: schema.TypeString, Optional: true, Default: "critical", Description: "'warning': Log level Warning; 'critical': Log level Critical (Default); 'notice': Log level Notice;",
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
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'tcp-session-created': TCP Session Created; 'tcp-session-deleted': TCP Session Deleted; 'tcp-port-allocated': TCP Port Allocated; 'tcp-port-freed': TCP Port Freed; 'tcp-port-batch-allocated': TCP Port Batch Allocated; 'tcp-port-batch-freed': TCP Port Batch Freed; 'udp-session-created': UDP Session Created; 'udp-session-deleted': UDP Session Deleted; 'udp-port-allocated': UDP Port Allocated; 'udp-port-freed': UDP Port Freed; 'udp-port-batch-allocated': UDP Port Batch Allocated; 'udp-port-batch-freed': UDP Port Batch Freed; 'icmp-session-created': ICMP Session Created; 'icmp-session-deleted': ICMP Session Deleted; 'icmp-resource-allocated': ICMP Resource Allocated; 'icmp-resource-freed': ICMP Resource Freed; 'icmpv6-session-created': ICMPV6 Session Created; 'icmpv6-session-deleted': ICMPV6 Session Deleted; 'icmpv6-resource-allocated': ICMPV6 Resource Allocated; 'icmpv6-resource-freed': ICMPV6 Resource Freed; 'gre-session-created': GRE Session Created; 'gre-session-deleted': GRE Session Deleted; 'gre-resource-allocated': GRE Resource Allocated; 'gre-resource-freed': GRE Resource Freed; 'esp-session-created': ESP Session Created; 'esp-session-deleted': ESP Session Deleted; 'esp-resource-allocated': ESP Resource Allocated; 'esp-resource-freed': ESP Resource Freed; 'fixed-nat-user-ports': Fixed NAT Inside User Port Mapping; 'fixed-nat-disable-config-logged': Fixed NAT Periodic Configs Logged; 'fixed-nat-disable-config-logs-sent': Fixed NAT Periodic Config Logs Sent; 'fixed-nat-periodic-config-logs-sent': Fixed NAT Disabled Configs Logged; 'fixed-nat-periodic-config-logged': Fixed NAT Disabled Config Logs Sent; 'fixed-nat-interim-updated': Fixed NAT Interim Updated; 'enhanced-user-log': Enhanced User Log; 'log-sent': Log Packets Sent; 'log-dropped': Log Packets Dropped; 'conn-tcp-established': TCP Connection Established; 'conn-tcp-dropped': TCP Connection Lost; 'tcp-port-overloading-allocated': TCP Port Overloading Allocated; 'tcp-port-overloading-freed': TCP Port Overloading Freed; 'udp-port-overloading-allocated': UDP Port Overloading Allocated; 'udp-port-overloading-freed': UDP Port Overloading Freed; 'http-request-logged': HTTP Request Logged; 'reduced-logs-by-destination': Reduced Logs by Destination Protocol and Port; 'out-of-buffers': Out of Buffers; 'add-msg-failed': Add Message to Buffer Failed; 'rtsp-port-allocated': RTSP UDP Port Allocated; 'rtsp-port-freed': RTSP UDP Port Freed; 'conn-tcp-create-failed': TCP Connection Failed; 'ipv4-frag-applied': IPv4 Fragmentation Applied; 'ipv4-frag-failed': IPv4 Fragmentation Failed; 'ipv6-frag-applied': IPv6 Fragmentation Applied; 'ipv6-frag-failed': IPv6 Fragmentation Failed; 'interim-update-scheduled': Port Allocation Interim Update Scheduled; 'interim-update-schedule-failed': Port Allocation Interim Update Failed; 'interim-update-terminated': Port Allocation Interim Update Terminated; 'interim-update-memory-freed': Port Allocation Interim Update Memory Freed; 'interim-update-no-buff-retried': Port Allocation Interim Update Memory Freed; 'tcp-port-batch-interim-updated': TCP Port Batch Interim Updated; 'udp-port-batch-interim-updated': UDP Port Batch Interim Updated; 'port-block-accounting-freed': Port Allocation Accounting Freed; 'port-block-accounting-allocated': Port Allocation Accounting Allocated; 'log-message-too-long': Log message too long; 'http-out-of-order-dropped': HTTP out-of-order dropped; 'http-alloc-failed': HTTP Request Info Allocation Failed; 'http-frag-merge-failed-dropped': HTTP frag merge failed dropped; 'http-malloc': HTTP mem allocate; 'http-mfree': HTTP mem free; 'http-spm-alloc-type0': HTTP Conn SPM Type 0 allocate; 'http-spm-alloc-type1': HTTP Conn SPM Type 1 allocate; 'http-spm-alloc-type2': HTTP Conn SPM Type 2 allocate; 'http-spm-alloc-type3': HTTP Conn SPM Type 3 allocate; 'http-spm-alloc-type4': HTTP Conn SPM Type 4 allocate; 'http-spm-free-type0': HTTP Conn SPM Type 0 free; 'http-spm-free-type1': HTTP Conn SPM Type 1 free; 'http-spm-free-type2': HTTP Conn SPM Type 2 free; 'http-spm-free-type3': HTTP Conn SPM Type 3 free; 'http-spm-free-type4': HTTP Conn SPM Type 4 free; 'iddos-l3-entry-create': iDDoS L3 Entry Created; 'iddos-l3-entry-delete': iDDoS L3 Entry Deleted; 'iddos-l4-entry-create': iDDoS L4 Entry Created; 'iddos-l4-entry-delete': iDDoS L4 Entry Deleted;",
						},
					},
				},
			},
			"source_address": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"tcp_svr_status": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceCgnv6LoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Logging(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Logging(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LoggingRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Logging(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Logging(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6LoggingNatQuotaExceeded84(d []interface{}) edpt.Cgnv6LoggingNatQuotaExceeded84 {

	count1 := len(d)
	var ret edpt.Cgnv6LoggingNatQuotaExceeded84
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Level = in["level"].(string)
		//omit uuid
	}
	return ret
}

func getObjectCgnv6LoggingNatResourceExhausted85(d []interface{}) edpt.Cgnv6LoggingNatResourceExhausted85 {

	count1 := len(d)
	var ret edpt.Cgnv6LoggingNatResourceExhausted85
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Level = in["level"].(string)
		//omit uuid
	}
	return ret
}

func getSliceCgnv6LoggingSamplingEnable(d []interface{}) []edpt.Cgnv6LoggingSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LoggingSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LoggingSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6LoggingSourceAddress86(d []interface{}) edpt.Cgnv6LoggingSourceAddress86 {

	var ret edpt.Cgnv6LoggingSourceAddress86
	return ret
}

func getObjectCgnv6LoggingTcpSvrStatus87(d []interface{}) edpt.Cgnv6LoggingTcpSvrStatus87 {

	var ret edpt.Cgnv6LoggingTcpSvrStatus87
	return ret
}

func dataToEndpointCgnv6Logging(d *schema.ResourceData) edpt.Cgnv6Logging {
	var ret edpt.Cgnv6Logging
	ret.Inst.NatQuotaExceeded = getObjectCgnv6LoggingNatQuotaExceeded84(d.Get("nat_quota_exceeded").([]interface{}))
	ret.Inst.NatResourceExhausted = getObjectCgnv6LoggingNatResourceExhausted85(d.Get("nat_resource_exhausted").([]interface{}))
	ret.Inst.SamplingEnable = getSliceCgnv6LoggingSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SourceAddress = getObjectCgnv6LoggingSourceAddress86(d.Get("source_address").([]interface{}))
	ret.Inst.TcpSvrStatus = getObjectCgnv6LoggingTcpSvrStatus87(d.Get("tcp_svr_status").([]interface{}))
	//omit uuid
	return ret
}
