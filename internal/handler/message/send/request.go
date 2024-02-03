package send

import (
	"fmt"
	"net/http"

	"github.com/flashlabs/mailer-micro-service/internal/handler"
	"github.com/flashlabs/mailer-micro-service/pkg"
)

func requestData(r *http.Request) (data, error) {
	p, err := handler.Payload(r)
	if err != nil {
		return data{}, fmt.Errorf("error while executing handler.Payload: %w", pkg.ErrInvalidPayload)
	}

	v := p
	d := data{}

	if val, ex := v["mailing_id"]; ex {
		switch val := val.(type) {
		case float64:
			d.MailingID = int(val)
		default:
			return data{}, fmt.Errorf("mailing_id error: %w", pkg.ErrInvalidType)
		}
	} else {
		return data{}, fmt.Errorf("mailing_id error: %w", pkg.ErrMissingParam)
	}

	return d, nil
}
