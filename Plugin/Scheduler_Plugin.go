package main

import (
	"context"
	"fmt"

	"k8s.io/kubernetes/pkg/scheduler/framework"
	corev1 "k8s.io/api/core/v1"
)

type ExamplePlugin struct{}

// 新插件实现接口
func (p *ExamplePlugin) Name() string {
	return "ExamplePlugin"
}

func (p *ExamplePlugin) Filter(ctx context.Context, state *framework.CycleState, pod *corev1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	// 调度过滤逻辑
	fmt.Printf("Filtering pod %s on node %s\n", pod.Name, nodeInfo.Node().Name)
	return nil
}

func NewExamplePlugin() framework.Plugin {
	return &ExamplePlugin{}
}
