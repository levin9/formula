package fs

import (
	"fmt"
	"github.com/levin9/formula/opt"
	"reflect"
)

type AndFunction struct {
}

func (*AndFunction) Name() string {
	return "&&"
}

func (f *AndFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("function %s required two arguments", f.Name())
	}

	arg0, err := (*args[0]).Evaluate(context)
	if err != nil {
		return nil, err
	}

	if arg0.Type != reflect.Bool {
		return nil, fmt.Errorf("the first argument of function %s should be bool", f.Name())
	}

	arg1, err1 := (*args[1]).Evaluate(context)
	if err1 != nil {
		return nil, err1
	}

	if arg1.Type != reflect.Bool {
		return nil, fmt.Errorf("the first argument of function %s should be bool", f.Name())
	}
	if (arg0.Value == true) && (arg1.Value == true) {
		return opt.NewArgumentWithType(true, reflect.Bool), nil
	}
	return opt.NewArgumentWithType(false, reflect.Bool), nil
}

func NewAndFunction() *AndFunction {
	return &AndFunction{}
}
