package examples

import cnf "github.com/niouzjh/cloud-native-flow/pkg"

var (
	// 创建实例任务流
	createFlow = cnf.Flow{
		Name:    "createFlow",
		Comment: "create a instance",
		Steps: []cnf.Step{
			{Name: "CreatePod", Func: CreatePod, Comment: "create pods"},
			{Name: "ListPod", Func: ListPod, Comment: "list all pod"},
			{Name: "GenerateConfig", Func: GenerateConfig, Comment: "generate configs"},
			{Name: "FlushParam", SubFlow: flushParamFlow, Comment: "flush param in pod"},
		},
	}

	// 刷新参数任务流
	flushParamFlow = cnf.Flow{
		Name:    "flushParamFlow",
		Comment: "flush param of instance",
		Steps: []cnf.Step{
			{Name: "ListPod", Func: ListPod, Comment: "list all pod"},
		},
	}

	// 升配任务流
	updateFlow = cnf.Flow{
		Name:    "upgradeFlow",
		Comment: "upgrade the instance",
		Steps: []cnf.Step{
			{Name: "...", Func: ListPod, Comment: "upgrade steps..."},
		},
	}
)
