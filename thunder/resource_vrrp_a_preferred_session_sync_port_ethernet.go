package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAPreferredSessionSyncPortEthernet() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_preferred_session_sync_port_ethernet`: preferred-session-sync-port ethernet\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAPreferredSessionSyncPortEthernetCreate,
		UpdateContext: resourceVrrpAPreferredSessionSyncPortEthernetUpdate,
		ReadContext:   resourceVrrpAPreferredSessionSyncPortEthernetRead,
		DeleteContext: resourceVrrpAPreferredSessionSyncPortEthernetDelete,

		Schema: map[string]*schema.Schema{
			"pre_eth": {
				Type: schema.TypeInt, Required: true, Description: "Ethernet interface number",
			},
			"pre_vlan": {
				Type: schema.TypeInt, Optional: true, Description: "Interface VLAN (VLAN ID)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVrrpAPreferredSessionSyncPortEthernetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAPreferredSessionSyncPortEthernetCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAPreferredSessionSyncPortEthernet(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAPreferredSessionSyncPortEthernetRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAPreferredSessionSyncPortEthernetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAPreferredSessionSyncPortEthernetUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAPreferredSessionSyncPortEthernet(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAPreferredSessionSyncPortEthernetRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAPreferredSessionSyncPortEthernetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAPreferredSessionSyncPortEthernetDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAPreferredSessionSyncPortEthernet(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAPreferredSessionSyncPortEthernetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAPreferredSessionSyncPortEthernetRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAPreferredSessionSyncPortEthernet(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVrrpAPreferredSessionSyncPortEthernet(d *schema.ResourceData) edpt.VrrpAPreferredSessionSyncPortEthernet {
	var ret edpt.VrrpAPreferredSessionSyncPortEthernet
	ret.Inst.PreEth = d.Get("pre_eth").(int)
	ret.Inst.PreVlan = d.Get("pre_vlan").(int)
	//omit uuid
	return ret
}
