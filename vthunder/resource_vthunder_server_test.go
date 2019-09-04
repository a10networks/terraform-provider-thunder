package vthunder

import (
	"fmt"
	"reflect"
	"testing"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var NAME = "rs9"

var TEST_SR_RESOURCE = `
resource "vthunder_server" "rs9" {
health_check_disable=1
name="` + NAME + `"
host="10.0.9.4"
port_list {
health_check_disable=1
port_number=80
protocol="tcp"
} 
}
`

//Acceptance test
func TestAccVthunderServer_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testCheckServerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: TEST_SR_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					testCheckServerExists(NAME, true),
					resource.TestCheckResourceAttr("vthunder_server.rs9", "name", NAME),
					resource.TestCheckResourceAttr("vthunder_server.rs9", "host", "10.0.9.4"),
					resource.TestCheckResourceAttr("vthunder_server.rs9", "health_check_disable", "1"),
					resource.TestCheckResourceAttr("vthunder_server.rs9", "port_list.0.health_check_disable", "1"),
					resource.TestCheckResourceAttr("vthunder_server.rs9", "port_list.0.port_number", "80"),
					resource.TestCheckResourceAttr("vthunder_server.rs9", "port_list.0.protocol", "tcp"),
				),
			},
		},
	})
}

func TestAccVthunderServer_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testCheckServerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: TEST_SR_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					testCheckServerExists(NAME, true),
				),
				ResourceName:      NAME,
				ImportState:       false,
				ImportStateVerify: true,
			},
		},
	})
}


//Unit test for utility method to create Server structure
func TestDataToServer(t *testing.T) {
	HealthCheckDisable := 1
	StatsDataAction := "enable"
	SlowStart := 1
	Weight := 1
	SpoofingCache := 1
	ResolveAs := "r"
	ConnLimit := 1
	UUID := "id"
	FqdnName := "name"
	ExternalIP := "1.1.1.1"
	HealthCheckShared := "enable"
	Ipv6 := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	TemplateServer := "s9"
	ServerIpv6Addr := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	SharedPartitionHealthCheck := 1
	Host := "2.3.4.5"
	ExtendedStats := 1
	ConnResume := 1
	Name := "name"
	UserTag := "tag"
	Action := "enable"
	HealthCheck := "enable"
	NoLogging := 1

	resourceDataMap := map[string]interface{}{
		"health_check_disable":          HealthCheckDisable,
		"stats_data_action":             StatsDataAction,
		"slow_start":                    SlowStart,
		"weight":                        Weight,
		"spoofing_cache":                SpoofingCache,
		"resolve_as":                    ResolveAs,
		"conn_limit":                    ConnLimit,
		"uuid":                          UUID,
		"fqdn_name":                     FqdnName,
		"external_ip":                   ExternalIP,
		"health_check_shared":           HealthCheckShared,
		"ipv6":                          Ipv6,
		"template_server":               TemplateServer,
		"server_ipv6_addr":              ServerIpv6Addr,
		"shared_partition_health_check": SharedPartitionHealthCheck,
		"host":                          Host,
		"extended_stats":                ExtendedStats,
		"conn_resume":                   ConnResume,
		"name":                          Name,
		"user_tag":                      UserTag,
		"action":                        Action,
		"health_check":                  HealthCheck,
		"no_logging":                    NoLogging,
		"port_list":                     map[string]interface{}{},
		"alternate_server":              map[string]interface{}{},
		"sampling_enable":               map[string]interface{}{},
	}

	resourceSchema := map[string]*schema.Schema{
		"health_check_disable": &schema.Schema{
			Type: schema.TypeInt,
		},
		"stats_data_action": &schema.Schema{
			Type: schema.TypeString,
		},
		"slow_start": &schema.Schema{
			Type: schema.TypeInt,
		},
		"weight": &schema.Schema{
			Type: schema.TypeInt,
		},
		"spoofing_cache": &schema.Schema{
			Type: schema.TypeInt,
		},
		"resolve_as": &schema.Schema{
			Type: schema.TypeString,
		},
		"conn_limit": &schema.Schema{
			Type: schema.TypeInt,
		},
		"uuid": &schema.Schema{
			Type: schema.TypeString,
		},
		"fqdn_name": &schema.Schema{
			Type: schema.TypeString,
		},
		"external_ip": &schema.Schema{
			Type: schema.TypeString,
		},
		"health_check_shared": &schema.Schema{
			Type: schema.TypeString,
		},
		"ipv6": &schema.Schema{
			Type: schema.TypeString,
		},
		"template_server": &schema.Schema{
			Type: schema.TypeString,
		},
		"server_ipv6_addr": &schema.Schema{
			Type: schema.TypeString,
		},
		"shared_partition_health_check": &schema.Schema{
			Type: schema.TypeInt,
		},
		"host": &schema.Schema{
			Type: schema.TypeString,
		},
		"extended_stats": &schema.Schema{
			Type: schema.TypeInt,
		},
		"conn_resume": &schema.Schema{
			Type: schema.TypeInt,
		},
		"name": &schema.Schema{
			Type: schema.TypeString,
		},
		"user_tag": &schema.Schema{
			Type: schema.TypeString,
		},
		"action": &schema.Schema{
			Type: schema.TypeString,
		},
		"health_check": &schema.Schema{
			Type: schema.TypeString,
		},
		"no_logging": &schema.Schema{
			Type: schema.TypeInt,
		},
		"port_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_check": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"extended_stats": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"rport_health_check_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"range": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"stats_data_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_rport_health_check": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"conn_resume": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"health_check_disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"support_http2": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"follow_port_protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"health_check_follow_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"no_logging": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"alternate_port": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alternate": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"alternate_server_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"alternate_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"no_ssl": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"weight": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"sampling_enable": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"port_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"conn_limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_port": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_server_ssl": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		"alternate_server": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alternate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"alternate_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		"sampling_enable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
	}

	resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, resourceDataMap)

	var s go_vthunder.Server
	var sInstance go_vthunder.ServerInstance

	sInstance.HealthCheckDisable = HealthCheckDisable
	sInstance.StatsDataAction = StatsDataAction
	sInstance.SlowStart = SlowStart
	sInstance.Weight = Weight
	sInstance.SpoofingCache = SpoofingCache
	sInstance.ResolveAs = ResolveAs
	sInstance.ConnLimit = ConnLimit
	sInstance.UUID = UUID
	sInstance.FqdnName = FqdnName
	sInstance.ExternalIP = ExternalIP
	sInstance.HealthCheckShared = HealthCheckShared
	sInstance.Ipv6 = Ipv6
	sInstance.TemplateServer = TemplateServer
	sInstance.ServerIpv6Addr = ServerIpv6Addr
	sInstance.SharedPartitionHealthCheck = SharedPartitionHealthCheck
	sInstance.Host = Host
	sInstance.ExtendedStats = ExtendedStats
	sInstance.ConnResume = ConnResume
	sInstance.Name = Name
	sInstance.UserTag = UserTag
	sInstance.Action = Action
	sInstance.HealthCheck = HealthCheck
	sInstance.NoLogging = NoLogging
	sInstance.PortNumber = make([]go_vthunder.PortLists, 0, 1)
	sInstance.Counters1 = make([]go_vthunder.SamplingEnable, 0, 1)
	sInstance.AlternateName = make([]go_vthunder.AlternateServer, 0, 1)
	s.Name = sInstance
	
	cases := []struct {
		input  *schema.ResourceData
		output go_vthunder.Server
	}{
		{
			input:  resourceLocalData,
			output: s,
		},
	}

	for _, tc := range cases {
		output := dataToServer("name",resourceLocalData)
		if !reflect.DeepEqual(output, tc.output) {
			t.Fatalf("Got:\n\n%#v\n\nExpected:\n\n%#v", output, tc.output)
		}
	}

}

func testCheckServerExists(name string, exists bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		client := testAccProvider.Meta().(vThunder)
		vs, err := go_vthunder.GetServer(client.Token, name, client.Host)

		if err != nil {
			return err
		}
		if exists && vs == nil {
			return fmt.Errorf(" server %s was not created.", name)
		}
		if !exists && vs != nil {
			return fmt.Errorf(" server %s still exists.", name)
		}
		return nil
	}
}

func testCheckServerDestroyed(s *terraform.State) error {
	client := testAccProvider.Meta().(vThunder)
	for _, rs := range s.RootModule().Resources {

		name := rs.Primary.ID
		sr, err := go_vthunder.GetServer(client.Token, name, client.Host)
		if err != nil {
			return err
		}
		if sr == nil {
			return fmt.Errorf(" server %s not destroyed.", name)
		}
	}
	return nil
}
