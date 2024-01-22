package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemPath() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_path`: Layer-2 path\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemPathCreate,
		UpdateContext: resourceSystemPathUpdate,
		ReadContext:   resourceSystemPathRead,
		DeleteContext: resourceSystemPathDelete,

		Schema: map[string]*schema.Schema{
			"ifpair_eth_end": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet port",
			},
			"ifpair_eth_start": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Interface number)",
			},
			"ifpair_trunk_end": {
				Type: schema.TypeInt, Optional: true, Description: "Trunk Group",
			},
			"ifpair_trunk_start": {
				Type: schema.TypeInt, Optional: true, Description: "Trunk groups",
			},
			"l2hm_attach": {
				Type: schema.TypeString, Optional: true, Description: "Monitor Name",
			},
			"l2hm_path_name": {
				Type: schema.TypeString, Required: true, Description: "Monitor Name",
			},
			"l2hm_setup_test_api": {
				Type: schema.TypeInt, Optional: true, Description: "Test-API Interface (Ethernet Interface)",
			},
			"l2hm_vlan": {
				Type: schema.TypeInt, Optional: true, Description: "VLAN id",
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
func resourceSystemPathCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPathCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPath(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemPathRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemPathUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPathUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPath(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemPathRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemPathDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPathDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPath(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemPathRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPathRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPath(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemPath(d *schema.ResourceData) edpt.SystemPath {
	var ret edpt.SystemPath
	ret.Inst.IfpairEthEnd = d.Get("ifpair_eth_end").(int)
	ret.Inst.IfpairEthStart = d.Get("ifpair_eth_start").(int)
	ret.Inst.IfpairTrunkEnd = d.Get("ifpair_trunk_end").(int)
	ret.Inst.IfpairTrunkStart = d.Get("ifpair_trunk_start").(int)
	ret.Inst.L2hmAttach = d.Get("l2hm_attach").(string)
	ret.Inst.L2hmPathName = d.Get("l2hm_path_name").(string)
	ret.Inst.L2hmSetupTestApi = d.Get("l2hm_setup_test_api").(int)
	ret.Inst.L2hmVlan = d.Get("l2hm_vlan").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
