package cronjob

import (
	"context"

	"github.com/hillfolk/goutils/ctxdb"
	"github.com/hillfolk/goutils/echomiddleware"
	"github.com/hillfolk/goutils/kafka"

	"github.com/go-xorm/xorm"
)

 echomiddleware.ContextDBName

func ContextDB(service string, xormEngine *xorm.Engine, kafkaConfig kafka.Config) Middleware {
	ctxdb := ctxdb.New(xormEngine, service, kafkaConfig)

	return func(next HandlerFunc) HandlerFunc {
		return func(ctx context.Context) error {
			session := ctxdb.NewSession(ctx)
			defer session.Close()

			ctx = context.WithValue(ctx, echomiddleware.ContextDBName, session)

			return next(ctx)
		}
	}
}

func ContextMDB(service string, xormEngine *xorm.Engine, kafkaConfig kafka.Config) Middleware {
	ctxdb := ctxdb.New(xormEngine, service, kafkaConfig)

	return func(next HandlerFunc) HandlerFunc {
		return func(ctx context.Context) error {
			session := ctxdb.NewSession(ctx)
			defer session.Close()

			ctx = context.WithValue(ctx, echomiddleware.ContextMYNUMBERDBName, session)

			return next(ctx)
		}
	}
}
