package thunder

//Thunder resource RouterOspfArea

import (
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"strings"
	"util"
)

func resourceRouterOspfArea() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspfAreaCreate,
		Update: resourceRouterOspfAreaUpdate,
		Read:   resourceRouterOspfAreaRead,
		Delete: resourceRouterOspfAreaDelete,
		Schema: map[string]*schema.Schema{
			"area_ipv4": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"area_num": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"auth_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"message_digest": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"filter_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_list": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"acl_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"acl_direction": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"plist_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"plist_direction": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"nssa_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nssa": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"no_redistribution": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"no_summary": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"translator_role": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"default_information_originate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"metric": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"metric_type": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"default_cost": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"range_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"area_range_prefix": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"option": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"shortcut": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"stub_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stub": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"no_summary": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"virtual_link_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"virtual_link_ip_addr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"bfd": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"hello_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dead_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"retransmit_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"transmit_delay": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"virtual_link_authentication": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"virtual_link_auth_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"authentication_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"message_digest_key": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"md5": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"as_number": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"process_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceRouterOspfAreaCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterOspfArea (Inside resourceRouterOspfAreaCreate) ")
		name2 := d.Get("area_ipv4").(string)
		name3 := d.Get("area_num").(int)
		name1 := d.Get("process_id").(int)
		name := strconv.Itoa(name1)
		data := dataToRouterOspfArea(d)
		logger.Println("[INFO] received formatted data from method data to RouterOspfArea --")
		d.SetId(strconv.Itoa(name1) + "," + name2 + "," + strconv.Itoa(name3))
		go_thunder.PostRouterOspfArea(client.Token, name, data, client.Host)

		return resourceRouterOspfAreaRead(d, meta)

	}
	return nil
}

func resourceRouterOspfAreaRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterOspfArea (Inside resourceRouterOspfAreaRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		name3 := id[2]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterOspfArea(client.Token, name1, name2, name3, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceRouterOspfAreaUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		name3 := id[2]
		logger.Println("[INFO] Modifying RouterOspfArea   (Inside resourceRouterOspfAreaUpdate) ")
		data := dataToRouterOspfArea(d)
		logger.Println("[INFO] received formatted data from method data to RouterOspfArea ")
		go_thunder.PutRouterOspfArea(client.Token, name1, name2, name3, data, client.Host)

		return resourceRouterOspfAreaRead(d, meta)

	}
	return nil
}

func resourceRouterOspfAreaDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		name3 := id[2]
		logger.Println("[INFO] Deleting instance (Inside resourceRouterOspfAreaDelete) " + name1)
		err := go_thunder.DeleteRouterOspfArea(client.Token, name1, name2, name3, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToRouterOspfArea(d *schema.ResourceData) go_thunder.RouterOspfArea {
	var vc go_thunder.RouterOspfArea
	var c go_thunder.RouterOspfAreaInstance
	c.AreaIpv4 = d.Get("area_ipv4").(string)
	c.AreaNum = d.Get("area_num").(int)

	var obj1 go_thunder.RouterOspfAreaAuthCfg
	prefix1 := "auth_cfg.0."
	obj1.Authentication = d.Get(prefix1 + "authentication").(int)
	obj1.MessageDigest = d.Get(prefix1 + "message_digest").(int)

	c.Authentication = obj1

	FilterListsCount := d.Get("filter_lists.#").(int)
	c.FilterList = make([]go_thunder.RouterOspfAreaFilterLists, 0, FilterListsCount)

	for i := 0; i < FilterListsCount; i++ {
		var obj2 go_thunder.RouterOspfAreaFilterLists
		prefix2 := fmt.Sprintf("filter_lists.%d.", i)
		obj2.FilterList = d.Get(prefix2 + "filter_list").(int)
		obj2.AclName = d.Get(prefix2 + "acl_name").(string)
		obj2.AclDirection = d.Get(prefix2 + "acl_direction").(string)
		obj2.PlistName = d.Get(prefix2 + "plist_name").(string)
		obj2.PlistDirection = d.Get(prefix2 + "plist_direction").(string)
		c.FilterList = append(c.FilterList, obj2)
	}

	var obj3 go_thunder.RouterOspfAreaNssaCfg
	prefix3 := "nssa_cfg.0."
	obj3.Nssa = d.Get(prefix3 + "nssa").(int)
	obj3.NoRedistribution = d.Get(prefix3 + "no_redistribution").(int)
	obj3.NoSummary = d.Get(prefix3 + "no_summary").(int)
	obj3.TranslatorRole = d.Get(prefix3 + "translator_role").(string)
	obj3.DefaultInformationOriginate = d.Get(prefix3 + "default_information_originate").(int)
	obj3.Metric = d.Get(prefix3 + "metric").(int)
	obj3.MetricType = d.Get(prefix3 + "metric_type").(int)

	c.Nssa = obj3

	c.DefaultCost = d.Get("default_cost").(int)

	RangeListCount := d.Get("range_list.#").(int)
	c.AreaRangePrefix = make([]go_thunder.RouterOspfAreaRangeList, 0, RangeListCount)

	for i := 0; i < RangeListCount; i++ {
		var obj4 go_thunder.RouterOspfAreaRangeList
		prefix4 := fmt.Sprintf("range_list.%d.", i)
		obj4.AreaRangePrefix = d.Get(prefix4 + "area_range_prefix").(string)
		obj4.Option = d.Get(prefix4 + "option").(string)
		c.AreaRangePrefix = append(c.AreaRangePrefix, obj4)
	}

	c.Shortcut = d.Get("shortcut").(string)

	var obj5 go_thunder.RouterOspfAreaStubCfg
	prefix5 := "stub_cfg.0."
	obj5.Stub = d.Get(prefix5 + "stub").(int)
	obj5.NoSummary = d.Get(prefix5 + "no_summary").(int)

	c.Stub = obj5

	VirtualLinkListCount := d.Get("virtual_link_list.#").(int)
	c.VirtualLinkIPAddr = make([]go_thunder.RouterOspfAreaVirtualLinkList, 0, VirtualLinkListCount)

	for i := 0; i < VirtualLinkListCount; i++ {
		var obj6 go_thunder.RouterOspfAreaVirtualLinkList
		prefix6 := fmt.Sprintf("virtual_link_list.%d.", i)
		obj6.VirtualLinkIPAddr = d.Get(prefix6 + "virtual_link_ip_addr").(string)
		obj6.Bfd = d.Get(prefix6 + "bfd").(int)
		obj6.HelloInterval = d.Get(prefix6 + "hello_interval").(int)
		obj6.DeadInterval = d.Get(prefix6 + "dead_interval").(int)
		obj6.RetransmitInterval = d.Get(prefix6 + "retransmit_interval").(int)
		obj6.TransmitDelay = d.Get(prefix6 + "transmit_delay").(int)
		obj6.VirtualLinkAuthentication = d.Get(prefix6 + "virtual_link_authentication").(int)
		obj6.VirtualLinkAuthType = d.Get(prefix6 + "virtual_link_auth_type").(string)
		obj6.AuthenticationKey = d.Get(prefix6 + "authentication_key").(string)
		obj6.MessageDigestKey = d.Get(prefix6 + "message_digest_key").(int)
		obj6.Md5 = d.Get(prefix6 + "md5").(string)
		c.VirtualLinkIPAddr = append(c.VirtualLinkIPAddr, obj6)
	}

	vc.AreaIpv4 = c
	return vc
}
