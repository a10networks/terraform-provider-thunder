package file

//pseudo endpoint for upload and delete SSL CRL files

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	endpnt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi/oper"
)

type SslCrl struct {
	Name        string
	Protocol    string
	Host        string
	Path        string
	Username    string
	Password    string
	UseMgmtPort int
	Overwrite   int
}

func (p *SslCrl) GetId() string {
	return p.Name
}

func (p *SslCrl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("file.SslCrl::Post")

	realInst := endpnt.Import{}
	realInst.Inst.SslCrl = p.Name
	realInst.Inst.UseMgmtPort = p.UseMgmtPort
	realInst.Inst.Overwrite = p.Overwrite
	if len(p.Password) > 0 {
		realInst.Inst.Password = p.Password
	}
	realInst.Inst.RemoteFile = axapi.GenUrlForDownload(p.Protocol, p.Host, p.Path, p.Username)

	return realInst.Post(authToken, host, logger)
}

func (p *SslCrl) Get(authToken string, host, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.SslCrl::Get")
	err := oper.FileSslCrlGet(authToken, host, instId, logger)
	return err
}

func (p *SslCrl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.SslCrl::Delet")
	obj := endpnt.PkiDelete{}
	obj.Inst.Crl = instId
	return obj.Post(authToken, host, logger)
}
