package reconciler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
)

var (
	errorNodeNotManaged  = errors.New("Node not managed")
	nodeRolesToConfigMap = map[string]string{
		"node-role.kubernetes.io/master":  "node-config-master",
		"node-role.kubernetes.io/infra":   "node-config-infra",
		"node-role.kubernetes.io/compute": "node-config-compute",
	}
)

const (
	customNodeLabel  = "node-role.kubernetes.io/custom"
	defaultNamespace = "openshift-node"
)

func getConfigMapName(node *corev1.Node) (types.NamespacedName, error) {
	if configMap, ok := node.Labels[customNodeLabel]; ok {
		namespace := defaultNamespace
		name := configMap
		if namespaceName := strings.Split(configMap, "/"); len(namespaceName) == 2 {
			namespace = namespaceName[0]
			name = namespaceName[1]
		}
		return types.NamespacedName{Namespace: namespace, Name: name}, nil
	}
	for nodeRole, configMap := range nodeRolesToConfigMap {
		if _, ok := node.Labels[nodeRole]; ok {
			return types.NamespacedName{Namespace: defaultNamespace, Name: configMap}, nil
		}
	}
	return types.NamespacedName{}, errorNodeNotManaged
}

type Reconciler struct {
}

func New() *Reconciler {
	return &Reconciler{}
}

func (r *Reconciler) Reconcile(ctx *context.Context, node *corev1.Node) error {
	configMapName, err := getConfigMapName(node)
	if err != nil {
		return err
	}
	if node.Spec.ConfigSource != nil && node.Spec.ConfigSource.ConfigMap.Name == configMapName.Name && node.Spec.ConfigSource.ConfigMap.Namespace == configMapName.Namespace {
		return fmt.Errorf("Already Added ConfigMap %v/%v to Node %v", configMapName.Namespace, configMapName.Name, node.Name)
	}
	configMap := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      configMapName.Name,
			Namespace: configMapName.Namespace,
		},
	}
	if err := sdk.Get(configMap); err != nil {
		return err
	}
	node.Spec.ConfigSource = &corev1.NodeConfigSource{
		ConfigMap: &corev1.ConfigMapNodeConfigSource{
			Namespace:        cm.Namespace,
			Name:             cm.Name,
			KubeletConfigKey: "kubelet",
		},
	}
	if err = sdk.Update(node); err != nil {
		return err
	}
	log.Printf("Added ConfigMap %v/%v to Node %v", configMap.Namespace, configMap.Name, node.Name)
	return nil
}
