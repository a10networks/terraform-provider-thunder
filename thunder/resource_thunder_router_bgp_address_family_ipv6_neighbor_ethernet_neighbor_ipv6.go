package thunder

//Thunder resource RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6

import (
	"context"
	"strconv"
	"strings"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Create,
		UpdateContext: resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Update,
		ReadContext:   resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read,
		DeleteContext: resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Delete,
		Schema: map[string]*schema.Schema{
			"ethernet": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"peer_group_name": {
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

func resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6 (Inside resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Create) ")
		name1 := d.Get("as_number").(int)
		name := strconv.Itoa(name1)
		name2 := d.Get("ethernet").(int)
		data := dataToRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6 --")
		d.SetId(strconv.Itoa(name1) + "," + strconv.Itoa(name2))
		err := go_thunder.PostRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6 (Inside resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(client.Token, name1, name2, client.Host)
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

func resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6   (Inside resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Update) ")
		data := dataToRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6 ")
		err := go_thunder.PutRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(client.Token, name1, name2, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read(ctx, d, meta)
}

func dataToRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(d *schema.ResourceData) go_thunder.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6 {
	var vc go_thunder.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6
	var c go_thunder.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Instance
	c.Ethernet = d.Get("ethernet").(int)
	c.PeerGroupName = d.Get("peer_group_name").(string)

	vc.Ethernet = c
	return vc
}
