package thread

import (
	"clean_arch_api/backend/util/errors"
	"context"
	"time"
)

const (
	defaultChanelTask = 5
	DefaultTimeOut    = 10 * time.Second
)

type AsyncFunc func(ctx context.Context) errors.Error

func RunTask(f func() errors.Error) AsyncFunc {
	return func(ctx context.Context) errors.Error {
		ch := make(chan errors.Error, defaultChanelTask)
		go func() {
			ch <- f()
		}()
		select {
		case <-ctx.Done():
			<-ch
			if ctx.Err() != nil {
				return errors.NewInternalServerError()
			} else {
				return nil
			}
		case err := <-ch:
			return err
		}
	}
}

/*
平行に処理を行う
*/
func RunParallel(funcs ...AsyncFunc) AsyncFunc {
	return func(ctx context.Context) errors.Error {
		childCtx, cancelAll := context.WithCancel(ctx)
		defer cancelAll()

		doneCh := make(chan struct{}, len(funcs))
		errCh := make(chan errors.Error, len(funcs))
		recoverCh := make(chan interface{}, len(funcs))

		for _, f := range funcs {
			go func(_f AsyncFunc) {
				defer func() {
					r := recover()
					if r != nil {
						recoverCh <- r
					}
				}()

				if _f == nil {
					doneCh <- struct{}{}
					return
				}

				if err := _f(childCtx); err != nil {
					errCh <- err
					return
				}
				doneCh <- struct{}{}
			}(f)
		}

		for i := 0; i < len(funcs); i++ {
			select {
			case <-ctx.Done():
				if ctx.Err() != nil {
					return errors.NewInternalServerError()
				} else {
					return nil
				}
			case <-doneCh:
			case err := <-errCh:
				return err
			case r := <-recoverCh:
				panic(r)
			}
		}
		return nil
	}
}

/*
直列に処理を行う
*/
func RunSerial(fs ...AsyncFunc) AsyncFunc {
	return func(ctx context.Context) errors.Error {
		for _, f := range fs {
			if f == nil {
				continue
			}
			if err := f(ctx); err != nil {
				return err
			}
		}
		return nil
	}
}
