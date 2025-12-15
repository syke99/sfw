package web

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/syke99/sfw/app/spinner"
	"github.com/syke99/sfw/pkg/models"
)

func (s *wb) startSpinnerSource(ctx context.Context, baseHandler http.HandlerFunc, source string, spinnerType spinner.Type, lines chan []byte, errs chan<- error) {
	if spinnerType == spinner.Web {
		invertChain(s.mux, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			msg, err := io.ReadAll(r.Body)
			if err != nil {
				errs <- err
				return
			}

			lines <- msg

			// TODO: send status back to webhook
			return
		}), source)
	}

	spinnerTicker := time.NewTicker(time.Second)
	defer spinnerTicker.Stop()

	started := false

	for {
		select {
		case <-ctx.Done():
			errs <- ctx.Err()
			return
		case <-spinnerTicker.C:
			if !started {
				started = true

				switch spinnerType {
				case spinner.File:
					file, err := os.Open(source)
					if err != nil {
						errs <- err
						return
					}

					scanner := bufio.NewReader(file)

					ticker := time.NewTicker(time.Second)
					defer ticker.Stop()

					go func() {
						for {
							select {
							case <-ctx.Done():
								errs <- ctx.Err()
								return
							case <-ticker.C:
								line, err := scanner.ReadBytes('\n')
								if err != nil {
									if err != io.EOF {
										errs <- err
									}
								}

								lines <- line
							}
						}
					}()
				default:
					errs <- fmt.Errorf("unsupported spinner type: %T", spinnerType)
					return
				}
			}
		}
	}
}

func (s *wb) startSpinner(ctx context.Context, id int, sp spinner.Spinner, in <-chan []byte, errs chan<- error) {
	for {
		select {
		case <-ctx.Done():
			errs <- ctx.Err()
			return
		case msg, ok := <-in:
			if !ok {
				continue
			}
			message := models.Message{
				ID:   id,
				Data: msg,
			}
			sp.Cast(ctx, message, errs)
		}
	}
}
