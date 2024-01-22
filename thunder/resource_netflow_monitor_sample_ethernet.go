package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowMonitorSampleEthernet() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_netflow_monitor_sample_ethernet`: select ethernet interface to monitor\n\n__PLACEHOLDER__",
		CreateContext: resourceNetflowMonitorSampleEthernetCreate,
		UpdateContext: resourceNetflowMonitorSampleEthernetUpdate,
		ReadContext:   resourceNetflowMonitorSampleEthernetRead,
		DeleteContext: resourceNetflowMonitorSampleEthernetDelete,

		Schema: map[string]*schema.Schema{
			"ifindex": {
				Type: schema.TypeInt, Required: true, Description: "Ethernet interface number",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceNetflowMonitorSampleEthernetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorSampleEthernetCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorSampleEthernet(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorSampleEthernetRead(ctx, d, meta)
	}
	return diags
}

func resourceNetflowMonitorSampleEthernetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorSampleEthernetUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorSampleEthernet(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorSampleEthernetRead(ctx, d, meta)
	}
	return diags
}
func resourceNetflowMonitorSampleEthernetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorSampleEthernetDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorSampleEthernet(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetflowMonitorSampleEthernetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorSampleEthernetRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorSampleEthernet(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetflowMonitorSampleEthernet(d *schema.ResourceData) edpt.NetflowMonitorSampleEthernet {
	var ret edpt.NetflowMonitorSampleEthernet
	ret.Inst.Ifindex = d.Get("ifindex").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
