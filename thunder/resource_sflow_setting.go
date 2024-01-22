package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowSetting() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_setting`: Configure sFlow\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowSettingCreate,
		UpdateContext: resourceSflowSettingUpdate,
		ReadContext:   resourceSflowSettingRead,
		DeleteContext: resourceSflowSettingDelete,

		Schema: map[string]*schema.Schema{
			"append_mapping_info": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow TPS to always send mapping ctr block (260, 271, and 272)",
			},
			"counter_polling_interval": {
				Type: schema.TypeInt, Optional: true, Default: 20, Description: "sFlow counter polling interval, default is 20",
			},
			"default_counter_polling_mtu": {
				Type: schema.TypeInt, Optional: true, Default: 1500, Description: "Default MTU for counter-polling packets - DDoS 3.2 format only (Default: 1500)",
			},
			"local_collection": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable local sflow collection; 'disable': Disable local sflow collection;",
			},
			"local_t1_polling_interval": {
				Type: schema.TypeInt, Optional: true, Description: "Set sFlow local counter polling interval for T1 stats",
			},
			"local_t2_polling_interval": {
				Type: schema.TypeInt, Optional: true, Description: "Set sFlow local counter polling interval for T2 stats",
			},
			"management_link_utilization": {
				Type: schema.TypeInt, Optional: true, Description: "limit management link speed in (Mbps)",
			},
			"management_link_utilization_percentage": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "percentage limit of the management link speed (Default is 30%)",
			},
			"max_header": {
				Type: schema.TypeInt, Optional: true, Description: "Configure maximum number of bytes that should be copied from a sampled packet (default: 128) (The maximum number of bytes (Default: 128))",
			},
			"packet_sampling_rate": {
				Type: schema.TypeInt, Optional: true, Default: 1000, Description: "sFlow packet sampling rate, default is 1000",
			},
			"port_range_end": {
				Type: schema.TypeInt, Optional: true, Description: "Source port-range end",
			},
			"port_range_start": {
				Type: schema.TypeInt, Optional: true, Description: "Source port-range",
			},
			"randomize_source_port": {
				Type: schema.TypeString, Optional: true, Default: "packet-sampling-only", Description: "'enable': Randomize source port; 'disable': Fix source port 6343; 'packet-sampling-only': Only randomized source port for packet-sampling (Default);",
			},
			"source_ip_use_mgmt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management interface's IP address for source IP of sFlow packets",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSflowSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowSettingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowSetting(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowSettingRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowSettingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowSetting(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowSettingRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowSettingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowSetting(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowSettingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowSetting(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSflowSetting(d *schema.ResourceData) edpt.SflowSetting {
	var ret edpt.SflowSetting
	ret.Inst.AppendMappingInfo = d.Get("append_mapping_info").(int)
	ret.Inst.CounterPollingInterval = d.Get("counter_polling_interval").(int)
	ret.Inst.DefaultCounterPollingMtu = d.Get("default_counter_polling_mtu").(int)
	ret.Inst.LocalCollection = d.Get("local_collection").(string)
	ret.Inst.LocalT1PollingInterval = d.Get("local_t1_polling_interval").(int)
	ret.Inst.LocalT2PollingInterval = d.Get("local_t2_polling_interval").(int)
	ret.Inst.ManagementLinkUtilization = d.Get("management_link_utilization").(int)
	ret.Inst.ManagementLinkUtilizationPercentage = d.Get("management_link_utilization_percentage").(int)
	ret.Inst.MaxHeader = d.Get("max_header").(int)
	ret.Inst.PacketSamplingRate = d.Get("packet_sampling_rate").(int)
	ret.Inst.PortRangeEnd = d.Get("port_range_end").(int)
	ret.Inst.PortRangeStart = d.Get("port_range_start").(int)
	ret.Inst.RandomizeSourcePort = d.Get("randomize_source_port").(string)
	ret.Inst.SourceIpUseMgmt = d.Get("source_ip_use_mgmt").(int)
	//omit uuid
	return ret
}
