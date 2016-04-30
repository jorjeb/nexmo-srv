package handler

import (
	"net/url"

	"github.com/jorjeb/nexmo-srv/nexmo"
	sms "github.com/jorjeb/nexmo-srv/proto/sms"
	"github.com/micro/go-micro/errors"
	"github.com/ttacon/libphonenumber"
	"golang.org/x/net/context"
)

type SMS struct{}

func (s *SMS) Send(ctx context.Context, req *sms.SendRequest, rsp *sms.SendResponse) error {
	if len(req.To) == 0 {
		return errors.BadRequest("go.micro.srv.nexmo.SMS.Send", "Recipient cannot be blank")
	}

	if len(req.From) == 0 {
		return errors.BadRequest("go.micro.srv.nexmo.SMS.Send", "Sender ID cannot be blank")
	}

	num, err := libphonenumber.Parse(req.To, req.RegionCode)
	if err != nil {
		return errors.BadRequest("go.micro.srv.nexmo.SMS.Send", err.Error())
	}

	to := libphonenumber.Format(num, libphonenumber.E164)

	u := url.Values{}

	u.Set("to", to)
	u.Set("from", req.To)
	u.Set("text", req.Text)

	if err = nexmo.SendSMS(u, rsp); err != nil {
		return errors.InternalServerError("go.micro.srv.nexmo.SMS.Send", err.Error())
	}

	return nil
}
