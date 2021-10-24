package examples

import (
	"context"
	"fmt"
	cnf "github.com/niouzjh/cloud-native-flow/pkg"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// 状态机，判断状态转移，返回动作Action
	//Action ac = StateMachine.process(req)

	// 流程中心，执行 动作Action
	//CloudNativeFlow.execute(createAction)

	if err := cnf.CNF.Run(createFlow); err != nil {
		fmt.Println("error occurs")
	}

	return ctrl.Result{}, nil
}
