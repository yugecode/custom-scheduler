package plugins

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	"strings"
)

// 插件名称
const Name = "custom-plugin"

type Args struct {
	FavoriteColor  string `json:"favorite_color,omitempty"`
	FavoriteNumber int    `json:"favorite_number,omitempty"`
	ThanksTo       string `json:"thanks_to,omitempty"`
}

type Sample struct {
	args   *Args
	handle framework.FrameworkHandle
}

func (s *Sample) Name() string {
	return Name
}

func (sr Sample) PreBind(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	fmt.Println("进入PreBind........")
	fmt.Printf("pod:%s,namespace:%s\n", pod.Name, pod.Namespace)
	fmt.Printf("nodeName:%v\n", nodeName)
	fmt.Println("-----------------------------------------")
	if pod == nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("pod cannot be nil"))
	}
	if !strings.Contains(pod.Name, "c9521cd4") {
		return framework.NewStatus(framework.Error, fmt.Sprintf("pod name need contain c9521cd4"))
		//binding := &v1.Binding{
		//	ObjectMeta: v12.ObjectMeta{Namespace: pod.Namespace, Name: pod.Name, UID: pod.UID},
		//	Target:     v1.ObjectReference{Kind: "Node", Name: "master1"},
		//}
		//sr.handle.ClientSet().CoreV1().Pods(pod.Namespace).Bind(ctx,binding,v12.CreateOptions{})
	}
	return nil
}


func (sr Sample) Bind(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	fmt.Println("进入Bind........")
	fmt.Printf("pod:%s,namespace:%s\n", pod.Name, pod.Namespace)
	fmt.Printf("nodeName:%v\n", nodeName)
	fmt.Println("-----------------------------------------")
	if pod == nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("pod cannot be nil"))
	}
	// 包含c9521cd4的pod调度到master1节点
	if strings.Contains(pod.Name, "c9521cd4") {
		binding := &v1.Binding{
			ObjectMeta: v12.ObjectMeta{Namespace: pod.Namespace, Name: pod.Name, UID: pod.UID},
			Target:     v1.ObjectReference{Kind: "Node", Name: "master1"},
		}
		err := sr.handle.ClientSet().CoreV1().Pods(pod.Namespace).Bind(ctx, binding, v12.CreateOptions{})
		if err != nil {
			return framework.NewStatus(framework.Error, err.Error())
		}
		return nil
	}
	return nil
}

// type PluginFactory = func(configuration *runtime.Unknown, f FrameworkHandle) (Plugin, error)
func New(configuration *runtime.Unknown, f framework.FrameworkHandle) (framework.Plugin, error) {
	args := &Args{}
	if err := framework.DecodeInto(configuration, args); err != nil {
		return nil, err
	}
	klog.V(3).Infof("get plugin config args: %+v", args)
	return &Sample{
		args:   args,
		handle: f,
	}, nil
}
