//
// THIS FILE IS GENERATED BY fastapi
// DO NOT MODIFY BY MANUAL
//
package module1

import "github.com/funny/fastapi"
import link "github.com/funny/link"

func (_ *Service) ServiceID() byte {
	return 1
}
func (_ *Service) NewRequest(id byte) fastapi.Message {
	switch id {
	case 1:
		return &AddReq{}
	}
	return nil
}
func (_ *Service) NewResponse(id byte) fastapi.Message {
	switch id {
	case 1:
		return &AddRsp{}
	}
	return nil
}
func (s *Service) HandleRequest(session *link.Session, req fastapi.Message) {
	switch req.MessageID() {
	case 1:
		session.Send(s.Add(session, req.(*AddReq)))
	default:
		panic("Unhandled Message Type")
	}
}
func (this *AddReq) ServiceID() byte {
	return 1
}
func (this *AddReq) MessageID() byte {
	return 1
}
func (this *AddReq) Identity() string {
	return "Service.AddReq"
}
func (this *AddRsp) ServiceID() byte {
	return 1
}
func (this *AddRsp) MessageID() byte {
	return 1
}
func (this *AddRsp) Identity() string {
	return "Service.AddRsp"
}
