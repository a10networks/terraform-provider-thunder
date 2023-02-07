provider "thunder" {
  address = ""
  username = ""
  password = ""
}

 
resource "thunder_slb_template_http" "test_http" {
      
      name = "Template_HTTP_XFF"
      insert_client_ip = 1
      insert_client_ip_header_name = "X-Forwarded-For"             
}


resource "thunder_slb_template_connection_reuse" "test_connection_reuse" {
      
      name = "Template_ConRe-Use_5000"
      limit_per_server = 5000
      timeout = 300            
}

resource "thunder_slb_template_persist_cookie" "test_persist_cookie" {
      
       name = "Template_Persist_Cookie_SSL"
       encrypt_level = 0
        secure = 1
        httponly = 1         
}

resource "thunder_slb_template_persist_source_ip" "test_persist_source_ip" {
      
       name = "Template_Persist_Source_IP"
       timeout = 3       
}

resource "thunder_slb_template_tcp_proxy" "test_tcp_proxy" {
      
        name = "Template_TCP-Proxy_300"
        idle_timeout = 300
        reset_fwd = 1
        reset_rev = 1      
}

resource "thunder_import" "test_import_aflex" {
     
		# depends_on = [thunder_service_group.test_group]
    remote_file = "scp://root@10.10.10.10:/home/visibility/vip-switch-ar2"
    use_mgmt_port =  1
    aflex =  "vip-switch-ar2"
    password = "password"
    overwrite =  1
}


resource "thunder_ip_nat_pool" "test_pool" {
        
        pool_name = "SNAT_VRID1"
        start_address = "1.1.1.1"
        end_address = "1.1.1.10"
        netmask = "/24"
        vrid = 4
        ip_rr = 1
        port_overload = 1        
}

resource "thunder_server" "rs10" {
  health_check_disable=1
  name="rs10"
  host="10.0.3.2"
  port_list {
    health_check_disable=1
    port_number=8080
    protocol="tcp"
  }
}

resource "thunder_service_group" "test_group" {

  name = "test_pool_group_dev"
  protocol="TCP"
  member_list {
    name=thunder_server.rs10.name
    port=8080
  }
}

resource "thunder_virtual_server" "test_virtual" {
      
      # depends_on = [thunder_service_group.test_group, thunder_import.test_import_aflex, thunder_slb_template_http.test_http]
      name = "VS_10.10.10.101"
      ip_address = "10.10.10.101"
      vrid = 4
      port_list {
          port_number = 80
          protocol = "http"
          aflex_scripts {
              aflex = thunder_import.test_import_aflex.aflex
          }
          no_auto_up_on_aflex = 1
          pool = thunder_ip_nat_pool.test_pool.pool_name
          service_group = thunder_service_group.test_group.name
          use_rcv_hop_for_resp = 1
          template_connection_reuse = thunder_slb_template_connection_reuse.test_connection_reuse.name
          template_persist_cookie = thunder_slb_template_persist_cookie.test_persist_cookie.name
          template_http = thunder_slb_template_http.test_http.name
          template_tcp_proxy_client = thunder_slb_template_tcp_proxy.test_tcp_proxy.name
          template_tcp_proxy_server = thunder_slb_template_tcp_proxy.test_tcp_proxy.name
      }        
            
}
