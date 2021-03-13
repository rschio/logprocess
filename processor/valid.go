package processor

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

func validRequest(r *Request) error {
	if r == nil {
		return errors.New("nil Request")
	}
	if r.Size < 0 {
		return errors.New("invalid size")
	}
	return nil
}

func validLatencies(l *Latencies) error {
	if l == nil {
		return errors.New("nil Latencies")
	}
	switch {
	case l.Proxy < 0:
		return errors.New("invalid latency proxy")
	case l.Gateway < 0:
		return errors.New("invalid gateway proxy")
	case l.Request < 0:
		return errors.New("invalid request proxy")
	}
	return nil
}

func validAuthenticatedEntity(a *AuthenticatedEntity) error {
	if a == nil {
		return errors.New("nil AuthenticatedEntity")
	}
	_, err := uuid.Parse(a.ConsumerID.UUID)
	return err
}

func validResponse(r *Response) error {
	if r == nil {
		return errors.New("nil Response")
	}
	switch {
	case r.Status < 0:
		return errors.New("invalid status")
	case r.Size < 0:
		return errors.New("invalid size")
	}
	if _, err := strconv.Atoi(r.Headers.ContentLength); err != nil {
		return errors.New("headers content length is not a int")
	}
	return nil
}

func validRoute(r *Route) error {
	if r == nil {
		return errors.New("nil Route")
	}
	if _, err := uuid.Parse(r.Service.ID); err != nil {
		return err
	}
	_, err := uuid.Parse(r.ID)
	return err
}

func validService(s *Service) error {
	if s == nil {
		return errors.New("nil Service")
	}
	switch {
	case s.Port < 0:
		return errors.New("invalid port")
	case s.ConnectTimeout < 0:
		return errors.New("invalid connect timeout")
	case s.ReadTimeout < 0:
		return errors.New("invalid read timeout")
	case s.WriteTimeout < 0:
		return errors.New("invalid write timeout")
	case s.Retries < 0:
		return errors.New("invalid retries")
	}
	return nil
}

func ValidRecord(r *Record) error {
	if r == nil {
		return fmt.Errorf("nil Record")
	}
	if err := validRequest(&r.Request); err != nil {
		return fmt.Errorf("invalid request: %v", err)
	}
	if err := validResponse(&r.Response); err != nil {
		return fmt.Errorf("invalid response: %v", err)
	}
	if err := validAuthenticatedEntity(&r.AuthenticatedEntity); err != nil {
		return err
	}
	if err := validRoute(&r.Route); err != nil {
		return fmt.Errorf("invalid route: %v", err)
	}
	if err := validService(&r.Service); err != nil {
		return fmt.Errorf("invalid service: %v", err)
	}
	if err := validLatencies(&r.Latencies); err != nil {
		return fmt.Errorf("invalid latencies: %v", err)
	}
	return nil
}
