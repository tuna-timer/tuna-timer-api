package utils

import (
	"context"

	"gopkg.in/mgo.v2"
)

const (
	// ContextKeyMongoSession is a key under which a request-scoped mongodb session is stored on Context
	ContextKeyMongoSession = "mongoSession"
)

// GetMongoSessionFromContext returns request-scoped DB transaction from the context
func GetMongoSessionFromContext(ctx context.Context) *mgo.Session {
	return ctx.Value(ContextKeyMongoSession).(*mgo.Session)
}

// PutMongoSessionInContext saves request-scoped mongo session in context
func PutMongoSessionInContext(parentContext context.Context, session *mgo.Session) context.Context {
	return context.WithValue(parentContext, ContextKeyMongoSession, session)
}
