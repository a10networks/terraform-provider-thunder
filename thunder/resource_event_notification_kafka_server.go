package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEventNotificationKafkaServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_event_notification_kafka_server`: Set remote kafka server ip address\n\n__PLACEHOLDER__",
		CreateContext: resourceEventNotificationKafkaServerCreate,
		UpdateContext: resourceEventNotificationKafkaServerUpdate,
		ReadContext:   resourceEventNotificationKafkaServerRead,
		DeleteContext: resourceEventNotificationKafkaServerDelete,

		Schema: map[string]*schema.Schema{
			"host_ipv4": {
				Type: schema.TypeString, Optional: true, Description: "Set kafka Broker ip address or hostname",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Default: 9092, Description: "Set remote kafka port number (Remote kafka port number 1-32767, default is 9092)",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'pr-acos-harmony-topic': L7 PR logs sent; 'avro-device-status-topic': Device Status Metrics sent; 'avro-partition-metrics-topic': Partition Metrics sent; 'avro-generic-sent': Generic Metrics sent; 'pr-acos-harmony-topic-enqueue-err': L7 PR dropped,enq error on acos queues; 'pr-acos-harmony-topic-dequeue-err': L7 PR dropped,enq error analytics queues; 'avro-generic-failed-encoding': Generic Metrics dropped,encoding error; 'avro-generic-failed-sending': Generic Metrics dropped,sending failure; 'avro-device-status-topic-enqueue-err': Device Status dropped,enq error on acos queues; 'avro-device-status-topic-dequeue-err': Device Status dropped,enq error analytics queues; 'avro-partition-metrics-topic-enqueue-err': Part metrics dropped,enq error on acos queues; 'avro-partition-metrics-topic-dequeue-err': Part metrics dropped,enq error analytics queues; 'kafka-unknown-topic-dequeue-err': Unknown type dropped,enq error analytics queues; 'kafka-broker-down': Messages dropped,analytics down; 'kafka-queue-full-err': Messages dropped,acos analytics queue full; 'pr-throttle-drop': L7 PR dropped,log throttling; 'pr-not-allowed-drop': L7 PR dropped, not allowed to be sent; 'pr-be-ttfb-anomaly': L7 PR back-end ttfb is negative; 'pr-be-ttlb-anomaly': L7 PR back-end ttlb is negative; 'pr-in-latency-threshold-exceed': L7 PR on latency threshold exceeded; 'pr-out-latency-threshold-exceed': L7 PR out latency threshold exceeded; 'pr-out-latency-anomaly': L7 PR out latency negative; 'pr-in-latency-anomaly': L7 PR on latency negative; 'kafka-topic-error': Module not supported by analytics; 'pc-encoding-failed': L4 PC logs dropped,encoding error; 'pc-acos-harmony-topic': L4 PC logs sent; 'pc-acos-harmony-topic-dequeue-err': L4 PC logs dropped,enq error analytics queues; 'cgn-pc-acos-harmony-topic': CGN PC logs sent; 'cgn-pc-acos-harmony-topic-dequeue-err': CGN PC logs dropped,enq error analytics queues; 'cgn-pe-acos-harmony-topic': CGN PE logs sent; 'cgn-pe-acos-harmony-topic-dequeue-err': CGN PE logs dropped,enq error analytics queues; 'fw-pc-acos-harmony-topic': FW PC logs sent; 'fw-pc-acos-harmony-topic-dequeue-err': FW PC logs dropped,enq error analytics queues; 'fw-deny-pc-acos-harmony-topic': FW DENY PC logs sent; 'fw-deny-pc-acos-harmony-topic-dequeue-err': FW DENY PC logs dropped,enq error analytics queues; 'fw-rst-pc-acos-harmony-topic': FW RST PC logs sent; 'fw-rst-pc-acos-harmony-topic-dequeue-err': FW RST PC logs dropped,enq error analytics queues; 'cgn-summary-error-acos-harmony-topic': CGN PE logs sent; 'cgn-summary-error-acos-harmony-topic-dequeue-err': CGN PE logs dropped,enq error analytics queues; 'rule-set-application-metrics-topic': AppFW metrics sent; 'rule-set-application-metrics-topic-dequeue-err': AppFW metrics dropped,enq error analytics queues; 'slb-ssl-stats-metrics-topic': SSL metrics sent; 'slb-ssl-stats-metrics-topic-dequeue-err': SSL metrics dropped,enq error analytics queues; 'slb-client-ssl-counters-metrics-topic': Client SSL metrics sent; 'slb-client-ssl-counters-metrics-topic-dequeue-err': Cilent SSL metrics dropped,enq error analytics qs; 'slb-server-ssl-counters-metrics-topic': Server SSL metrics sent; 'slb-server-ssl-counters-metrics-topic-dequeue-err': Server SSL metrics dropped,enq error analytics qs; 'pc-throttle-drop': L4 PC logs dropped,throttling; 'metrics-dropped-pt-missing': Metrics dropped,missing partition tenant mapping; 'ssli-pc-acos-harmony-topic': SSLi PC topic counter from acos to harmony; 'ssli-pc-acos-harmony-topic-dequeue-err': SSLi PC topic to harmony dequeue error; 'ssli-pe-acos-harmony-topic': SSLi PE topic counter from acos to harmony; 'ssli-pe-acos-harmony-topic-dequeue-err': SSLi PE topic to harmony dequeue error; 'analytics-bus-restart': Analytics bus restart count; 'waf-learn-pr-topic': WAF learn topic counter; 'waf-learn-pr-topic-dequeue-err': WAF learn metrics dropped,enq error analytics qs; 'waf-events-topic': WAF events topic counter; 'waf-events-topic-dequeue-err': WAF events metrics dropped,enq error analytics qs; 'visibility-topn-harmony-topic': Visibility TopN sent; 'visibility-topn-harmony-topic-dequeue-err': Visibility TopN metrics dropped,enq error analytics qs; 'hc-logs-sent-to-master': HC logs sent to master; 'hc-logs-received-from-blade': HC logs received from blade; 'hc-oper-sent-to-master': HC oper to master; 'hc-oper-received-from-blade': HC oper received from blade; 'hc-counters-sent-to-master': HC counters sent to master; 'hc-counters-received-from-blade': HC counters received from blade; 'hc-counters-dropped-from-blade': HC counters dropped from blade (uuid or size mismatch); 'pe-acos-harmony-topic': L7 PE logs sent; 'pe-acos-harmony-topic-enqueue-err': L7 PE dropped,enq error on acos queues; 'pe-acos-harmony-topic-dequeue-err': L7 PE dropped,enq error analytics queues; 'vpn-ipsec-sa-metrics-topic': IPSec SA metrics sent; 'vpn-ipsec-sa-metrics-topic-dequeue-err': IPSec SA metrics dropped,enq error analytics qs; 'vpn-ike-gateway-metrics-topic': IKE gateway metrics sent; 'vpn-ike-gateway-metrics-topic-dequeue-err': IKE gateway metrics dropped,enq error analytics qs; 'vpn-stats-metrics-topic': VPN STATS metrics sent; 'vpn-stats-metrics-topic-dequeue-err': VPN STATS metrics dropped,enq error analytics qs; 'cgn-port-usage-hstgrm-acos-harmony-topic': CGN Port Usage Histogram HC Export; 'cgn-port-usage-hstgrm-acos-harmony-topic-dequeue-err': CGN Port Usage Histogram HC Export Failed; 'avro-system-env-topic': System environment sent; 'avro-system-env-dequeue-err': System Environmet dropped,enq error analytics queues; 'cert-pinning-list-topic': Cert-pinning candidate list sent; 'cert-pinning-list-topic-dequeue-err': Cert-pinning candidate list dropped,enq error analytics queues; 'ngwaf-hc-ep-topic': NGWAF HC PE export; 'ngwaf-hc-ep-topic-dequeue-err': NGWAF HC PE export failed; 'ngwaf-hc-metrics-topic': NGWAF HC metrics export; 'ngwaf-hc-metrics-topic-dequeue-err': NGWAF HC metrics export failed;",
						},
					},
				},
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port for connections",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceEventNotificationKafkaServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEventNotificationKafkaServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEventNotificationKafkaServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEventNotificationKafkaServerRead(ctx, d, meta)
	}
	return diags
}

func resourceEventNotificationKafkaServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEventNotificationKafkaServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEventNotificationKafkaServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEventNotificationKafkaServerRead(ctx, d, meta)
	}
	return diags
}
func resourceEventNotificationKafkaServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEventNotificationKafkaServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEventNotificationKafkaServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceEventNotificationKafkaServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEventNotificationKafkaServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEventNotificationKafkaServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceEventNotificationKafkaServerSamplingEnable(d []interface{}) []edpt.EventNotificationKafkaServerSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.EventNotificationKafkaServerSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EventNotificationKafkaServerSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointEventNotificationKafkaServer(d *schema.ResourceData) edpt.EventNotificationKafkaServer {
	var ret edpt.EventNotificationKafkaServer
	ret.Inst.HostIpv4 = d.Get("host_ipv4").(string)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.SamplingEnable = getSliceEventNotificationKafkaServerSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
