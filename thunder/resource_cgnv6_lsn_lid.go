package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnLid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_lid`: Create an LSN Lid\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnLidCreate,
		UpdateContext: resourceCgnv6LsnLidUpdate,
		ReadContext:   resourceCgnv6LsnLidRead,
		DeleteContext: resourceCgnv6LsnLidDelete,

		Schema: map[string]*schema.Schema{
			"conn_rate_limit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"conn_rate_limit_val": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum connections per second (Default: No limit)",
						},
					},
				},
			},
			"ds_lite": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inside_src_permit_list": {
							Type: schema.TypeString, Optional: true, Description: "Class-List of IPv4 addresses permitted (Class-list to match for DS-Lite)",
						},
					},
				},
			},
			"extended_user_quota": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tcp_service_port": {
										Type: schema.TypeInt, Optional: true, Description: "Port (Port Value)",
									},
									"tcp_sessions": {
										Type: schema.TypeInt, Optional: true, Description: "Number of Extended Quota sessions allowed for this service",
									},
								},
							},
						},
						"udp": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"udp_service_port": {
										Type: schema.TypeInt, Optional: true, Description: "Port (Port Value)",
									},
									"udp_sessions": {
										Type: schema.TypeInt, Optional: true, Description: "Number of Extended Quota sessions allowed for this service",
									},
								},
							},
						},
					},
				},
			},
			"lid_number": {
				Type: schema.TypeInt, Required: true, Description: "LSN Lid",
			},
			"lsn_rule_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": {
							Type: schema.TypeString, Optional: true, Description: "Apply LSN Rule-List on Destination (LSN Rule-List Name)",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Optional: true, Description: "LSN Lid Name",
			},
			"override": {
				Type: schema.TypeString, Optional: true, Default: "none", Description: "'none': Apply source NAT if configured (default); 'drop': Drop packets that match this LSN lid; 'pass-through': Layer-3 route packets that match this LSN lid and do not apply source NAT;",
			},
			"respond_to_user_mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use the user's source MAC for the next hop rather than the routing table (default: off)",
			},
			"source_nat_pool": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pool_name": {
							Type: schema.TypeString, Optional: true, Description: "Source NAT Pool or Pool-Group",
						},
						"shared": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use a shared source NAT pool or pool-group",
						},
					},
				},
			},
			"user_quota": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp": {
							Type: schema.TypeInt, Optional: true, Description: "User Quota for ICMP identifiers (NAT port quota per user (default: not configured))",
						},
						"quota_udp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"udp_quota": {
										Type: schema.TypeInt, Optional: true, Description: "NAT port quota per user (default: not configured)",
									},
									"udp_reserve": {
										Type: schema.TypeInt, Optional: true, Description: "Number of ports to reserve per user (default: same as user-quota value) (Reserved quota per user (default: same as user-quota value))",
									},
								},
							},
						},
						"quota_tcp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tcp_quota": {
										Type: schema.TypeInt, Optional: true, Description: "NAT port quota per user (default: not configured)",
									},
									"tcp_reserve": {
										Type: schema.TypeInt, Optional: true, Description: "Number of ports to reserve per user (default: same as user-quota value) (Reserved quota per user (default: same as user-quota value))",
									},
								},
							},
						},
						"session": {
							Type: schema.TypeInt, Optional: true, Description: "User Quota for number of data sessions",
						},
					},
				},
			},
			"user_quota_prefix_length": {
				Type: schema.TypeInt, Optional: true, Description: "NAT64/DS-Lite user quota prefix length (Prefix Length (Default: Uses the global NAT64/DS-Lite configured value))",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6LsnLidCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnLidCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnLid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnLidRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnLidUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnLidUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnLid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnLidRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnLidDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnLidDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnLid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnLidRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnLidRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnLid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6LsnLidConnRateLimit(d []interface{}) edpt.Cgnv6LsnLidConnRateLimit {

	count1 := len(d)
	var ret edpt.Cgnv6LsnLidConnRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConnRateLimitVal = in["conn_rate_limit_val"].(int)
	}
	return ret
}

func getObjectCgnv6LsnLidDsLite(d []interface{}) edpt.Cgnv6LsnLidDsLite {

	count1 := len(d)
	var ret edpt.Cgnv6LsnLidDsLite
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.InsideSrcPermitList = in["inside_src_permit_list"].(string)
	}
	return ret
}

func getObjectCgnv6LsnLidExtendedUserQuota(d []interface{}) edpt.Cgnv6LsnLidExtendedUserQuota {

	count1 := len(d)
	var ret edpt.Cgnv6LsnLidExtendedUserQuota
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp = getSliceCgnv6LsnLidExtendedUserQuotaTcp(in["tcp"].([]interface{}))
		ret.Udp = getSliceCgnv6LsnLidExtendedUserQuotaUdp(in["udp"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6LsnLidExtendedUserQuotaTcp(d []interface{}) []edpt.Cgnv6LsnLidExtendedUserQuotaTcp {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnLidExtendedUserQuotaTcp, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnLidExtendedUserQuotaTcp
		oi.TcpServicePort = in["tcp_service_port"].(int)
		oi.TcpSessions = in["tcp_sessions"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6LsnLidExtendedUserQuotaUdp(d []interface{}) []edpt.Cgnv6LsnLidExtendedUserQuotaUdp {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnLidExtendedUserQuotaUdp, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnLidExtendedUserQuotaUdp
		oi.UdpServicePort = in["udp_service_port"].(int)
		oi.UdpSessions = in["udp_sessions"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6LsnLidLsnRuleList(d []interface{}) edpt.Cgnv6LsnLidLsnRuleList {

	count1 := len(d)
	var ret edpt.Cgnv6LsnLidLsnRuleList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Destination = in["destination"].(string)
	}
	return ret
}

func getObjectCgnv6LsnLidSourceNatPool(d []interface{}) edpt.Cgnv6LsnLidSourceNatPool {

	count1 := len(d)
	var ret edpt.Cgnv6LsnLidSourceNatPool
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PoolName = in["pool_name"].(string)
		ret.Shared = in["shared"].(int)
	}
	return ret
}

func getObjectCgnv6LsnLidUserQuota(d []interface{}) edpt.Cgnv6LsnLidUserQuota {

	count1 := len(d)
	var ret edpt.Cgnv6LsnLidUserQuota
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Icmp = in["icmp"].(int)
		ret.QuotaUdp = getObjectCgnv6LsnLidUserQuotaQuotaUdp(in["quota_udp"].([]interface{}))
		ret.QuotaTcp = getObjectCgnv6LsnLidUserQuotaQuotaTcp(in["quota_tcp"].([]interface{}))
		ret.Session = in["session"].(int)
	}
	return ret
}

func getObjectCgnv6LsnLidUserQuotaQuotaUdp(d []interface{}) edpt.Cgnv6LsnLidUserQuotaQuotaUdp {

	count1 := len(d)
	var ret edpt.Cgnv6LsnLidUserQuotaQuotaUdp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UdpQuota = in["udp_quota"].(int)
		ret.UdpReserve = in["udp_reserve"].(int)
	}
	return ret
}

func getObjectCgnv6LsnLidUserQuotaQuotaTcp(d []interface{}) edpt.Cgnv6LsnLidUserQuotaQuotaTcp {

	count1 := len(d)
	var ret edpt.Cgnv6LsnLidUserQuotaQuotaTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TcpQuota = in["tcp_quota"].(int)
		ret.TcpReserve = in["tcp_reserve"].(int)
	}
	return ret
}

func dataToEndpointCgnv6LsnLid(d *schema.ResourceData) edpt.Cgnv6LsnLid {
	var ret edpt.Cgnv6LsnLid
	ret.Inst.ConnRateLimit = getObjectCgnv6LsnLidConnRateLimit(d.Get("conn_rate_limit").([]interface{}))
	ret.Inst.DsLite = getObjectCgnv6LsnLidDsLite(d.Get("ds_lite").([]interface{}))
	ret.Inst.ExtendedUserQuota = getObjectCgnv6LsnLidExtendedUserQuota(d.Get("extended_user_quota").([]interface{}))
	ret.Inst.LidNumber = d.Get("lid_number").(int)
	ret.Inst.LsnRuleList = getObjectCgnv6LsnLidLsnRuleList(d.Get("lsn_rule_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Override = d.Get("override").(string)
	ret.Inst.RespondToUserMac = d.Get("respond_to_user_mac").(int)
	ret.Inst.SourceNatPool = getObjectCgnv6LsnLidSourceNatPool(d.Get("source_nat_pool").([]interface{}))
	ret.Inst.UserQuota = getObjectCgnv6LsnLidUserQuota(d.Get("user_quota").([]interface{}))
	ret.Inst.UserQuotaPrefixLength = d.Get("user_quota_prefix_length").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
