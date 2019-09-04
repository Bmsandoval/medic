package bash

import "fmt"

var (
	KubeListPodsByName  = "kubectl get pods --all-namespaces -o=jsonpath='{range .items[*]}{\"\\n\"}{.metadata.name}{end}'"
	KubeListCtxtsByName    = "kubectl config get-contexts -o name"
	KubeListReleasesByName = "helm ls -q"
	KubeGetCurrentContext  = "kubectl config current-context"
	KubeStart = ". ~/projects/k8s-cluster-management/scripts/minikube.sh"
	KubeStop = "minikube stop"
)

// Kubernetes
func KubePortForward(pod string, ports string) string { return fmt.Sprintf("kubectl port-forward %s %s", pod, ports) }
func KubeTailLogs(pod string) string { return fmt.Sprintf("kubectl logs %s -f", pod) }
func KubeDeletePod(pod string) string { return fmt.Sprintf("kubectl delete pod %s", pod) }
func KubeSshPod(pod string) string { return fmt.Sprintf("kubectl exec %s -it /bin/sh", pod) }
func KubeKillPod(pod string) string { return fmt.Sprintf("kubectl exec -it %s -- killall main", pod) }
func KubeSetContext(context string) string { return fmt.Sprintf("kubectl config use-context %s", context) }
func KubeDescribePod(pod string) string { return fmt.Sprintf("kubectl describe pod %s", pod) }
func KubeEventsOfPod(pod string) string { return fmt.Sprintf("kubectl get events | grep %s", pod)}
func KubeStatusOfPod(pod *string) string {
	command := "kubectl get pods -w"
	if pod == nil {
		return command
	}
	return fmt.Sprintf("%s | grep %s", command, *pod)
}

// Helm
func HelmDeletePurge(release string) string { return fmt.Sprintf("helm delete --purge %s", release) }

