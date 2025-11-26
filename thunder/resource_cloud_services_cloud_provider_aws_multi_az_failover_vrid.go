package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudServicesCloudProviderAwsMultiAzFailoverVrid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cloud_services_cloud_provider_aws_multi_az_failover_vrid`: VRID list\n\n__PLACEHOLDER__",
		CreateContext: resourceCloudServicesCloudProviderAwsMultiAzFailoverVridCreate,
		UpdateContext: resourceCloudServicesCloudProviderAwsMultiAzFailoverVridUpdate,
		ReadContext:   resourceCloudServicesCloudProviderAwsMultiAzFailoverVridRead,
		DeleteContext: resourceCloudServicesCloudProviderAwsMultiAzFailoverVridDelete,

		Schema: map[string]*schema.Schema{
			"fip_dest": {
				Type: schema.TypeString, Optional: true, Description: "Alien FIP Destination CIDR Block",
			},
			"fip_interface_id": {
				Type: schema.TypeString, Optional: true, Description: "Data-Out Interface ID",
			},
			"route_table_id": {
				Type: schema.TypeString, Optional: true, Description: "Specifies the route table id of the interface",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vip_dest": {
				Type: schema.TypeString, Optional: true, Description: "Alien VIP Destination CIDR Block",
			},
			"vip_interface_id": {
				Type: schema.TypeString, Optional: true, Description: "Data-In Interface ID",
			},
			"vip_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vip_number": {
							Type: schema.TypeInt, Required: true, Description: "Specify ha vip Number",
						},
						"private_ip": {
							Type: schema.TypeString, Optional: true, Description: "Private IP address of VIP",
						},
						"elastic_ip": {
							Type: schema.TypeString, Optional: true, Description: "Elastic IP address of VIP",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"vrid_number": {
				Type: schema.TypeInt, Required: true, Description: "Specify ha VRRP-A vrid Number",
			},
		},
	}
}
func resourceCloudServicesCloudProviderAwsMultiAzFailoverVridCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAwsMultiAzFailoverVridCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAwsMultiAzFailoverVrid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesCloudProviderAwsMultiAzFailoverVridRead(ctx, d, meta)
	}
	return diags
}

func resourceCloudServicesCloudProviderAwsMultiAzFailoverVridUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAwsMultiAzFailoverVridUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAwsMultiAzFailoverVrid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesCloudProviderAwsMultiAzFailoverVridRead(ctx, d, meta)
	}
	return diags
}
func resourceCloudServicesCloudProviderAwsMultiAzFailoverVridDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAwsMultiAzFailoverVridDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAwsMultiAzFailoverVrid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCloudServicesCloudProviderAwsMultiAzFailoverVridRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAwsMultiAzFailoverVridRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAwsMultiAzFailoverVrid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCloudServicesCloudProviderAwsMultiAzFailoverVridVipList(d []interface{}) []edpt.CloudServicesCloudProviderAwsMultiAzFailoverVridVipList {

	count1 := len(d)
	ret := make([]edpt.CloudServicesCloudProviderAwsMultiAzFailoverVridVipList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.CloudServicesCloudProviderAwsMultiAzFailoverVridVipList
		oi.VipNumber = in["vip_number"].(int)
		oi.PrivateIp = in["private_ip"].(string)
		oi.ElasticIp = in["elastic_ip"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCloudServicesCloudProviderAwsMultiAzFailoverVrid(d *schema.ResourceData) edpt.CloudServicesCloudProviderAwsMultiAzFailoverVrid {
	var ret edpt.CloudServicesCloudProviderAwsMultiAzFailoverVrid
	ret.Inst.FipDest = d.Get("fip_dest").(string)
	ret.Inst.FipInterfaceId = d.Get("fip_interface_id").(string)
	ret.Inst.RouteTableId = d.Get("route_table_id").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VipDest = d.Get("vip_dest").(string)
	ret.Inst.VipInterfaceId = d.Get("vip_interface_id").(string)
	ret.Inst.VipList = getSliceCloudServicesCloudProviderAwsMultiAzFailoverVridVipList(d.Get("vip_list").([]interface{}))
	ret.Inst.VridNumber = d.Get("vrid_number").(int)
	return ret
}
