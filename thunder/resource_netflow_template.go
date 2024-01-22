package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_netflow_template`: IPFIX Custom Template\n\n__PLACEHOLDER__",
		CreateContext: resourceNetflowTemplateCreate,
		UpdateContext: resourceNetflowTemplateUpdate,
		ReadContext:   resourceNetflowTemplateRead,
		DeleteContext: resourceNetflowTemplateDelete,

		Schema: map[string]*schema.Schema{
			"information_element_blk": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"information_element": {
							Type: schema.TypeString, Optional: true, Description: "'fwd-tuple-vnp-id': Session forward tuple partition id (ID: 33028); 'rev-tuple-vnp-id': Session reverse tuple partition id (ID: 33029); 'source-ipv4-address': IPv4 source address in the IP packet header (ID: 8); 'source-ipv4-prefix-len': Prefix length for IPv4 source address(ID: 9); 'dest-ipv4-address': IPv4 destination address in the IP packet header (ID: 12); 'dest-ipv4-prefix-len': Prefix length for IPv4 dest address(ID: 13); 'source-ipv6-address': IPv6 source address in the IP packet header (ID: 27); 'source-ipv6-prefix-len': Prefix length for IPv6 source address(ID:29); 'dest-ipv6-address': IPv6 destination address in the IP packet header (ID:28); 'dest-ipv6-prefix-len': Prefix length for IPv6 dest address (ID:30); 'post-nat-source-ipv4-address': IPv4 natted source address (ID: 225); 'post-nat-dest-ipv4-address': IPv4 natted destination address(ID: 226); 'post-nat-source-ipv6-address': IPv6 natted source address (ID: 281); 'post-nat-dest-ipv6-address': IPv6 natted destination address (ID: 282); 'source-port': Source port identifier in the transport header (ID: 7); 'dest-port': Destination port identifier in the transport header (ID: 11); 'post-nat-source-port': L4 natted source port(ID: 227); 'post-nat-dest-port': L4 natted destination port (ID: 228); 'fwd-tuple-type': Session forward tuple type (ID: 33024); 'rev-tuple-type': Session reverse tuple type (ID: 33025); 'ip-proto': Value of the protocol number in the IP packet header (ID: 4); 'flow-direction': Flow direction: 0:inbound(To an outside interface)/1:outbound(To an inside interface) (ID: 61); 'tcp-control-bits': Cumulative of all the TCP flags seen for this flow (ID: 6); 'fwd-bytes': Incoming bytes associated with an IP Flow (ID: 1); 'fwd-packets': Incoming packets associated with an IP Flow (ID: 2); 'rev-bytes': Delta bytes in reverse direction of bidirectional flow record (ID: 32769); 'rev-packets': Delta packets in reverse direction of bidirectional flow record (ID: 32770); 'in-port': Incoming interface port (ID: 10); 'out-port': Outcoming interface port (ID: 14); 'in-interface': Incoming interface name e.g. ethernet 0 (ID: 82); 'out-interface': Outcoming interface name e.g. ethernet 0 (ID: 32850); 'port-range-start': Port number identifying the start of a range of ports (ID: 361); 'port-range-end': Port number identifying the end of a range of ports (ID: 362); 'port-range-step-size': Step size in a port range (ID: 363); 'port-range-num-ports': Number of ports in a port range (ID: 364); 'rule-name': Rule Name (ID: 33034); 'rule-set-name': Rule-Set Name (ID: 33035); 'fw-source-zone': Firewall Source Zone Name (ID: 33036); 'fw-dest-zone': Firewall Dest Zone Name (ID: 33037); 'application-id': Application ID (ID: 95); 'application-name': Application Name (ID: 96); 'imsi': Subscriber Attribute IMSI (ID: 455); 'msisdn': Subscriber Attribute MSISDN (ID: 456); 'imei': Subscriber Attribute IMEI (ID: 33030); 'radius-custom1': Radius Attribute Custom 1 (ID: 33031); 'radius-custom2': Radius Attribute Custom 2(ID: 33032); 'radius-custom3': Radius Attribute Custom 3 (ID:33033); 'radius-custom4': Radius Attribute Custom 4 (ID: 33067); 'radius-custom5': Radius Attribute Custom 5(ID: 33068); 'radius-custom6': Radius Attribute Custom 6 (ID:33069); 'flow-start-msec': The absolute timestamp of the first packet of the flow (ID: 152); 'flow-duration-msec': Difference in time between the first observed packet of this flow and the last observed packet of this flow (4 bytes) (ID: 161); 'flow-duration-msec-64': Difference in time between the first observed packet of this flow and the last observed packet of this flow (8 bytes) (ID: 33039); 'flow-end-msec': The absolute timestamp of the last packet of the flow (ID: 153); 'nat-event': Indicates a NAT event (ID: 230); 'fw-event': Indicates a FW session event(ID: 233); 'fw-deny-reset-event': Indicates a FW deny/reset event (ID: 33038); 'cgn-flow-direction': Flow direction: 0:inbound(To an outside interface)/1:outbound(To an inside interface)/2:hairpin(From an inside interface to an inside interface) (ID: 33040); 'fw-dest-fqdn': Firewall matched fqdn(ID: 33041); 'flow-end-reason': A10 flow end reason(ID: 33042); 'gtp-deny-reason': Indicates a GTP deny event (ID: 33043); 'gtp-apn': Indicates GTP APN (ID: 33044); 'gtp-steid': Indicates GTP Source TEID (ID: 33045); 'gtp-dteid': Indicates GTP Destination TEID (ID: 33046); 'gtp-selection-mode': Indicates GTP Selection Mode (ID: 33047); 'gtp-mcc': Indicates the MCC of the Serving Network (ID: 33048); 'gtp-mnc': Indicates the MNC of the serving network (ID: 33049); 'gtp-rat-type': Indicates the RAT Type received in the GTP Control packet (ID: 33050); 'gtp-pdn-pdp-type': Indicates the PDN/PDP Type in the GTP Control Packet (ID: 33051); 'gtp-uli': Indicates the User Location Information (ID: 33052); 'gtp-enduser-v4-addr': Indicates the Subscriber IPv4 Address (ID: 33053); 'gtp-enduser-v6-addr': Indicates the Subscriber IPv6 Address (ID: 33054); 'gtp-bearer-id-or-nsapi': Indicates the EPS Bearer ID or NSAPI of the Subscriber (ID: 33055); 'gtp-qci': Indicates the QoS Profile or Traffic Class of the subscriber (ID: 33056); 'gtp-info-event-ind': Indicates a GTP Info event(ID: 33057); 'gtp-restarted-node-ipv4': Restarted S5 Node IPV4 Address(ID: 33058); 'gtp-restarted-node-ipv6': Restarted S5 Node IPV6 Address(ID: 33059); 'gtp-c-tunnels-removed-with-node-restart': Indicates GTP-C tunnels removed by Node restart (ID: 33060); 'radius-imsi': Subscriber Attribute IMSI (Deprecated Field) (ID: 455); 'radius-msisdn': Subscriber Attribute MSISDN (Deprecated Field) (ID: 456); 'radius-imei': Subscriber Attribute IMEI (Deprecated Field) (ID: 33030); 'event-time-msec': The absolute time in milliseconds of an event observation(ID: 323); 'security-event-type': Type of security event(ID: 33063); 'limit-exceeded-count': Limit exceeded count for FW concurrent session(ID: 33062); 'rate-limit-key': Rate Limit Key(ID: 33064); 'rate-limit-type': Rate Limit Type(ID: 33065); 'rate-limit-drop-count': Rate Limit Drop Count(ID: 33066);",
						},
					},
				},
			},
			"ipfix_template_id": {
				Type: schema.TypeInt, Optional: true, Description: "Custom IPFIX Template ID",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "IPFIX CUSTOM Template Name",
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
func resourceNetflowTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceNetflowTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceNetflowTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetflowTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceNetflowTemplateInformationElementBlk(d []interface{}) []edpt.NetflowTemplateInformationElementBlk {

	count1 := len(d)
	ret := make([]edpt.NetflowTemplateInformationElementBlk, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowTemplateInformationElementBlk
		oi.InformationElement = in["information_element"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetflowTemplate(d *schema.ResourceData) edpt.NetflowTemplate {
	var ret edpt.NetflowTemplate
	ret.Inst.InformationElementBlk = getSliceNetflowTemplateInformationElementBlk(d.Get("information_element_blk").([]interface{}))
	ret.Inst.IpfixTemplateId = d.Get("ipfix_template_id").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
