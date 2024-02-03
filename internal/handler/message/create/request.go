package create

import (
	"fmt"
	"net/http"
	"time"

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

	if val, ex := v["email"]; ex {
		switch val := val.(type) {
		case string:
			d.Email = val
		default:
			return data{}, fmt.Errorf("email error: %w", pkg.ErrInvalidType)
		}
	} else {
		return data{}, fmt.Errorf("email error: %w", pkg.ErrMissingParam)
	}

	if val, ex := v["title"]; ex {
		switch val := val.(type) {
		case string:
			d.Title = val
		default:
			return data{}, fmt.Errorf("title error: %w", pkg.ErrInvalidType)
		}
	} else {
		return data{}, fmt.Errorf("title error: %w", pkg.ErrMissingParam)
	}

	if val, ex := v["content"]; ex {
		switch val := val.(type) {
		case string:
			d.Content = val
		default:
			return data{}, fmt.Errorf("content error: %w", pkg.ErrInvalidType)
		}
	} else {
		return data{}, fmt.Errorf("content error: %w", pkg.ErrMissingParam)
	}

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

	if val, ex := v["insert_time"]; ex {
		switch val := val.(type) {
		case string:
			if tm, err := time.Parse(time.RFC3339Nano, val); err != nil {
				return data{}, fmt.Errorf("insert_time error: %w", pkg.ErrInvalidType)
			} else {
				d.InsertTime = tm
			}
		default:
			return data{}, fmt.Errorf("insert_time error: %w", pkg.ErrInvalidType)
		}
	} else {
		return data{}, fmt.Errorf("insert_time error: %w", pkg.ErrMissingParam)
	}

	return d, nil
}
