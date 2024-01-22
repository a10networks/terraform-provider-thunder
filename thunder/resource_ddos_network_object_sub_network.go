package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectSubNetwork() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_sub_network`: Configure sub-network in a DDos Network Object\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectSubNetworkCreate,
		UpdateContext: resourceDdosNetworkObjectSubNetworkUpdate,
		ReadContext:   resourceDdosNetworkObjectSubNetworkRead,
		DeleteContext: resourceDdosNetworkObjectSubNetworkDelete,

		Schema: map[string]*schema.Schema{
			"host_anomaly_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static_pkt_rate_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Packet rate of per host",
						},
						"static_byte_rate_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Byte rate of per host",
						},
					},
				},
			},
			"sub_network_anomaly_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static_sub_network_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Packet rate of the sub-network",
						},
						"static_sub_network_byte_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Byte rate of the sub-network",
						},
					},
				},
			},
			"subnet_ip_addr": {
				Type: schema.TypeString, Required: true, Description: "IPv4 Subnet/host, supported prefix range is from 24 to 32",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "ObjectName",
			},
		},
	}
}
func resourceDdosNetworkObjectSubNetworkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSubNetworkCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSubNetwork(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSubNetworkRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectSubNetworkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSubNetworkUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSubNetwork(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSubNetworkRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectSubNetworkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSubNetworkDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSubNetwork(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectSubNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSubNetworkRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSubNetwork(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosNetworkObjectSubNetworkHostAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectSubNetworkHostAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetworkHostAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticPktRateThreshold = in["static_pkt_rate_threshold"].(int)
		ret.StaticByteRateThreshold = in["static_byte_rate_threshold"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectSubNetworkSubNetworkAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectSubNetworkSubNetworkAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticSubNetworkPktRate = in["static_sub_network_pkt_rate"].(int)
		ret.StaticSubNetworkByteRate = in["static_sub_network_byte_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosNetworkObjectSubNetwork(d *schema.ResourceData) edpt.DdosNetworkObjectSubNetwork {
	var ret edpt.DdosNetworkObjectSubNetwork
	ret.Inst.HostAnomalyThreshold = getObjectDdosNetworkObjectSubNetworkHostAnomalyThreshold(d.Get("host_anomaly_threshold").([]interface{}))
	ret.Inst.SubNetworkAnomalyThreshold = getObjectDdosNetworkObjectSubNetworkSubNetworkAnomalyThreshold(d.Get("sub_network_anomaly_threshold").([]interface{}))
	ret.Inst.SubnetIpAddr = d.Get("subnet_ip_addr").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ObjectName = d.Get("object_name").(string)
	return ret
}
