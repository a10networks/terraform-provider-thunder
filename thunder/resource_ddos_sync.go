package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSync() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_sync`: Sync White list to peer devices\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSyncCreate,
		UpdateContext: resourceDdosSyncUpdate,
		ReadContext:   resourceDdosSyncRead,
		DeleteContext: resourceDdosSyncDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable",
			},
			"local_ip": {
				Type: schema.TypeString, Optional: true, Description: "Local IP address for White list sync",
			},
			"peer_ip_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"peer_ip": {
							Type: schema.TypeString, Optional: true, Description: "IP Address",
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
func resourceDdosSyncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSyncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSync(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSyncRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSyncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSyncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSync(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSyncRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSyncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSyncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSync(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSyncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSyncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSync(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosSyncPeerIpCfg(d []interface{}) []edpt.DdosSyncPeerIpCfg {

	count1 := len(d)
	ret := make([]edpt.DdosSyncPeerIpCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSyncPeerIpCfg
		oi.PeerIp = in["peer_ip"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosSync(d *schema.ResourceData) edpt.DdosSync {
	var ret edpt.DdosSync
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.LocalIp = d.Get("local_ip").(string)
	ret.Inst.PeerIpCfg = getSliceDdosSyncPeerIpCfg(d.Get("peer_ip_cfg").([]interface{}))
	//omit uuid
	return ret
}
