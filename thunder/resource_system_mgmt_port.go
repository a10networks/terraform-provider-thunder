package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMgmtPort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_mgmt_port`: set management port\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemMgmtPortCreate,
		UpdateContext: resourceSystemMgmtPortUpdate,
		ReadContext:   resourceSystemMgmtPortRead,
		DeleteContext: resourceSystemMgmtPortDelete,

		Schema: map[string]*schema.Schema{
			"mac_address": {
				Type: schema.TypeString, Optional: true, Description: "mac-address to be configured as mgmt port",
			},
			"pci_address": {
				Type: schema.TypeString, Optional: true, Description: "pci-address to be configured as mgmt port",
			},
			"port_index": {
				Type: schema.TypeInt, Optional: true, Description: "port index to be configured (Specify port index)",
			},
		},
	}
}
func resourceSystemMgmtPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMgmtPortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMgmtPort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMgmtPortRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemMgmtPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMgmtPortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMgmtPort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMgmtPortRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemMgmtPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMgmtPortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMgmtPort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemMgmtPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMgmtPortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMgmtPort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemMgmtPort(d *schema.ResourceData) edpt.SystemMgmtPort {
	var ret edpt.SystemMgmtPort
	ret.Inst.MacAddress = d.Get("mac_address").(string)
	ret.Inst.PciAddress = d.Get("pci_address").(string)
	ret.Inst.PortIndex = d.Get("port_index").(int)
	return ret
}
