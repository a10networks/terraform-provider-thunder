package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectTrustlist() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_trustlist`: Configure trusted source IP list in a DDoS Network Object\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectTrustlistCreate,
		UpdateContext: resourceDdosNetworkObjectTrustlistUpdate,
		ReadContext:   resourceDdosNetworkObjectTrustlistRead,
		DeleteContext: resourceDdosNetworkObjectTrustlistDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"v4_class_list": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 Class-list name",
			},
			"v6_class_list": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 Class-list name",
			},
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "ObjectName",
			},
		},
	}
}
func resourceDdosNetworkObjectTrustlistCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTrustlistCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTrustlist(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectTrustlistRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectTrustlistUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTrustlistUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTrustlist(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectTrustlistRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectTrustlistDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTrustlistDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTrustlist(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectTrustlistRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTrustlistRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTrustlist(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosNetworkObjectTrustlist(d *schema.ResourceData) edpt.DdosNetworkObjectTrustlist {
	var ret edpt.DdosNetworkObjectTrustlist
	//omit uuid
	ret.Inst.V4ClassList = d.Get("v4_class_list").(string)
	ret.Inst.V6ClassList = d.Get("v6_class_list").(string)
	ret.Inst.ObjectName = d.Get("object_name").(string)
	return ret
}
