package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudServicesCloudProviderAwsMultiAzFailoverVridVip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cloud_services_cloud_provider_aws_multi_az_failover_vrid_vip`: list of VIP IPs\n\n__PLACEHOLDER__",
		CreateContext: resourceCloudServicesCloudProviderAwsMultiAzFailoverVridVipCreate,
		UpdateContext: resourceCloudServicesCloudProviderAwsMultiAzFailoverVridVipUpdate,
		ReadContext:   resourceCloudServicesCloudProviderAwsMultiAzFailoverVridVipRead,
		DeleteContext: resourceCloudServicesCloudProviderAwsMultiAzFailoverVridVipDelete,

		Schema: map[string]*schema.Schema{
			"elastic_ip": {
				Type: schema.TypeString, Optional: true, Description: "Elastic IP address of VIP",
			},
			"private_ip": {
				Type: schema.TypeString, Optional: true, Description: "Private IP address of VIP",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vip_number": {
				Type: schema.TypeInt, Required: true, Description: "Specify ha vip Number",
			},
			"vrid_number": {
				Type: schema.TypeString, Required: true, Description: "VridNumber",
			},
		},
	}
}
func resourceCloudServicesCloudProviderAwsMultiAzFailoverVridVipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAwsMultiAzFailoverVridVipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAwsMultiAzFailoverVridVip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesCloudProviderAwsMultiAzFailoverVridVipRead(ctx, d, meta)
	}
	return diags
}

func resourceCloudServicesCloudProviderAwsMultiAzFailoverVridVipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAwsMultiAzFailoverVridVipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAwsMultiAzFailoverVridVip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesCloudProviderAwsMultiAzFailoverVridVipRead(ctx, d, meta)
	}
	return diags
}
func resourceCloudServicesCloudProviderAwsMultiAzFailoverVridVipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAwsMultiAzFailoverVridVipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAwsMultiAzFailoverVridVip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCloudServicesCloudProviderAwsMultiAzFailoverVridVipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAwsMultiAzFailoverVridVipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAwsMultiAzFailoverVridVip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCloudServicesCloudProviderAwsMultiAzFailoverVridVip(d *schema.ResourceData) edpt.CloudServicesCloudProviderAwsMultiAzFailoverVridVip {
	var ret edpt.CloudServicesCloudProviderAwsMultiAzFailoverVridVip
	ret.Inst.ElasticIp = d.Get("elastic_ip").(string)
	ret.Inst.PrivateIp = d.Get("private_ip").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VipNumber = d.Get("vip_number").(int)
	ret.Inst.VridNumber = d.Get("vrid_number").(string)
	return ret
}
