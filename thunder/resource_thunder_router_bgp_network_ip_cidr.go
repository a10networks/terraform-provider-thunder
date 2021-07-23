package thunder

//Thunder resource RouterBgpNetworkIpCidr

import (
	"context"
	"strconv"
	"strings"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpNetworkIpCidr() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterBgpNetworkIpCidrCreate,
		UpdateContext: resourceRouterBgpNetworkIpCidrUpdate,
		ReadContext:   resourceRouterBgpNetworkIpCidrRead,
		DeleteContext: resourceRouterBgpNetworkIpCidrDelete,
		Schema: map[string]*schema.Schema{
			"network_ipv4_cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"route_map": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"backdoor": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"comm_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"as_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"process_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sequence": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceRouterBgpNetworkIpCidrCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpNetworkIpCidr (Inside resourceRouterBgpNetworkIpCidrCreate) ")
		name2 := d.Get("network_ipv4_cidr").(string)
		name1 := d.Get("as_number").(int)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpNetworkIpCidr(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNetworkIpCidr --")
		d.SetId(strconv.Itoa(name1) + "," + name2)
		err := go_thunder.PostRouterBgpNetworkIpCidr(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpNetworkIpCidrRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpNetworkIpCidrRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading RouterBgpNetworkIpCidr (Inside resourceRouterBgpNetworkIpCidrRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpNetworkIpCidr(client.Token, name1, name2, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return nil
}

func resourceRouterBgpNetworkIpCidrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying RouterBgpNetworkIpCidr   (Inside resourceRouterBgpNetworkIpCidrUpdate) ")
		data := dataToRouterBgpNetworkIpCidr(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNetworkIpCidr ")
		err := go_thunder.PutRouterBgpNetworkIpCidr(client.Token, name1, name2, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpNetworkIpCidrRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpNetworkIpCidrDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceRouterBgpNetworkIpCidrDelete) " + name1)
		err := go_thunder.DeleteRouterBgpNetworkIpCidr(client.Token, name1, name2, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToRouterBgpNetworkIpCidr(d *schema.ResourceData) go_thunder.RouterBgpNetworkIpCidr {
	var vc go_thunder.RouterBgpNetworkIpCidr
	var c go_thunder.RouterBgpNetworkIPCidrInstance
	c.NetworkIpv4Cidr = d.Get("network_ipv4_cidr").(string)
	c.RouteMap = d.Get("route_map").(string)
	c.Backdoor = d.Get("backdoor").(int)
	c.Description = d.Get("description").(string)
	c.CommValue = d.Get("comm_value").(string)

	vc.NetworkIpv4Cidr = c
	return vc
}
