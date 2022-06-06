package data

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)
var tracer  = otel.Tracer("gorm")

func RegisterCallbacks(ctx context.Context, db *gorm.DB) {
	registerCallbacks(ctx, db)
}

func registerCallbacks(ctx context.Context, db *gorm.DB) {
	prefix := db.Dialector.Name() + ":"

	db.Callback().Create().Before("gorm:begin_transaction").Register("aotgorm_before_create", newBefore(prefix+"create"))
	db.Callback().Create().After("gorm:commit_or_rollback_transaction").Register("otgorm_after_create", newAfter())

	db.Callback().Update().Before("gorm:begin_transaction").Register("otgorm_before_update", newBefore(prefix+"update"))
	db.Callback().Update().After("gorm:commit_or_rollback_transaction").Register("otgorm_after_update", newAfter())

	db.Callback().Query().Before("gorm:query").Register("otgorm_before_query", newBefore(prefix+"query"))
	db.Callback().Query().After("gorm:after_query").Register("otgorm_after_query", newAfter())

	db.Callback().Delete().Before("gorm:begin_transaction").Register("otgorm_before_delete", newBefore( prefix+"delete"))
	db.Callback().Delete().After("gorm:commit_or_rollback_transactio").Register("otgorm_after_delete", newAfter())

	db.Callback().Row().Before("gorm:row").Register("otgorm_before_row", newBefore( prefix+"row"))
	db.Callback().Row().After("gorm:row").Register("otgorm_after_row", newAfter())

	db.Callback().Raw().Before("gorm:raw").Register("otgorm_before_raw", newBefore(prefix+"raw"))
	db.Callback().Raw().After("gorm:raw").Register("otgorm_after_raw", newAfter())
}

func newBefore(name string) func(*gorm.DB) {
	return func(db *gorm.DB) {
		if !trace.SpanFromContext(db.Statement.Context).IsRecording() {
			return
		}

		ctx, span := tracer.Start(db.Statement.Context, name)
		span.SetAttributes(
			attribute.String("db.system", "mysql"),
			attribute.String("db.table", db.Statement.Table),
			attribute.String("mdl", name),
		)
		db.Statement.Context = ctx
		// span.End()
	}
}

func newAfter() func(*gorm.DB) {
	return func(db *gorm.DB) {
		span := trace.SpanFromContext(db.Statement.Context)
		if err := db.Error; err != nil {
			recordError(db.Statement.Context, span, err)
		} else {
			span.SetStatus(codes.Ok, "OK")
		}
		span.End()
	}
}

func recordError(ctx context.Context, span trace.Span, err error) {
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
}
