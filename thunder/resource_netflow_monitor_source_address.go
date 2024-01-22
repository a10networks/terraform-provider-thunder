package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowMonitorSourceAddress() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_netflow_monitor_source_address`: Specify source address of Netflow packet\n\n__PLACEHOLDER__",
		CreateContext: resourceNetflowMonitorSourceAddressCreate,
		UpdateContext: resourceNetflowMonitorSourceAddressUpdate,
		ReadContext:   resourceNetflowMonitorSourceAddressRead,
		DeleteContext: resourceNetflowMonitorSourceAddressDelete,

		Schema: map[string]*schema.Schema{
			"ip": {
				Type: schema.TypeString, Optional: true, Description: "Specify source IP address",
			},
			"ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Specify source IPv6 address",
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
func resourceNetflowMonitorSourceAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorSourceAddressCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorSourceAddress(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorSourceAddressRead(ctx, d, meta)
	}
	return diags
}

func resourceNetflowMonitorSourceAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorSourceAddressUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorSourceAddress(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorSourceAddressRead(ctx, d, meta)
	}
	return diags
}
func resourceNetflowMonitorSourceAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorSourceAddressDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorSourceAddress(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetflowMonitorSourceAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorSourceAddressRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorSourceAddress(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetflowMonitorSourceAddress(d *schema.ResourceData) edpt.NetflowMonitorSourceAddress {
	var ret edpt.NetflowMonitorSourceAddress
	ret.Inst.Ip = d.Get("ip").(string)
	ret.Inst.Ipv6 = d.Get("ipv6").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
