package lower

import (
	"github.com/mochilang/mochi-jvm/transpiler/aotir"
	"github.com/mochilang/mochi-jvm/transpiler/javasrc"
)

// lowerHttpGetExpr lowers HttpGetExpr to dev.mochi.runtime.io.Fetch.get(url).
func (l *lowerer) lowerHttpGetExpr(e *aotir.HttpGetExpr) (javasrc.Expr, error) {
	url, err := l.lowerExpr(e.URL)
	if err != nil {
		return nil, err
	}
	return &javasrc.StaticCallExpr{
		Class:  "dev.mochi.runtime.io.Fetch",
		Method: "get",
		Args:   []javasrc.Expr{url},
	}, nil
}

// lowerJsonDecodeExpr lowers JsonDecodeExpr to dev.mochi.runtime.io.JSON.decode(input).
func (l *lowerer) lowerJsonDecodeExpr(e *aotir.JsonDecodeExpr) (javasrc.Expr, error) {
	input, err := l.lowerExpr(e.Input)
	if err != nil {
		return nil, err
	}
	return &javasrc.StaticCallExpr{
		Class:  "dev.mochi.runtime.io.JSON",
		Method: "decode",
		Args:   []javasrc.Expr{input},
	}, nil
}
