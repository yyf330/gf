package main

import (
	"context"
	"github.com/yyf330/gf/os/glog"
)

func main() {
	glog.SetCtxKeys("Trace-Id", "Span-Id", "Test")
	ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
	ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

	glog.Ctx(ctx).Print(1, 2, 3)
}
