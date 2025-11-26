package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionTrustlist() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_detection_trustlist`: Configure trusted IPv4/IPv6 classlists\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDetectionTrustlistCreate,
		UpdateContext: resourceDdosDetectionTrustlistUpdate,
		ReadContext:   resourceDdosDetectionTrustlistRead,
		DeleteContext: resourceDdosDetectionTrustlistDelete,

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
		},
	}
}
func resourceDdosDetectionTrustlistCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionTrustlistCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionTrustlist(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionTrustlistRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDetectionTrustlistUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionTrustlistUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionTrustlist(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionTrustlistRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDetectionTrustlistDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionTrustlistDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionTrustlist(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDetectionTrustlistRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionTrustlistRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionTrustlist(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDetectionTrustlist(d *schema.ResourceData) edpt.DdosDetectionTrustlist {
	var ret edpt.DdosDetectionTrustlist
	//omit uuid
	ret.Inst.V4ClassList = d.Get("v4_class_list").(string)
	ret.Inst.V6ClassList = d.Get("v6_class_list").(string)
	return ret
}
