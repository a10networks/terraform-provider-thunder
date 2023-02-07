module github.com/a10networks/terraform-provider-thunder

go 1.16

replace github.com/go_thunder/thunder => ./go_thunder/thunder

replace util => ./util/UtilLog

require (
	github.com/aws/aws-sdk-go v1.37.0 // indirect
	github.com/clarketm/json v1.15.7
	github.com/go_thunder/thunder v0.0.0-00010101000000-000000000000
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/hashicorp/hcl/v2 v2.8.2 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/pkg/errors v0.9.1
	golang.org/x/tools v0.0.0-20201028111035-eafbe7b904eb // indirect
	google.golang.org/api v0.34.0 // indirect
	util v0.0.0-00010101000000-000000000000

)
