module github.com/directxman12/k8s-prometheus-adapter

go 1.13

require (
	github.com/NYTimes/gziphandler v1.0.1 // indirect
	github.com/evanphx/json-patch v4.5.0+incompatible // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/kubernetes-incubator/custom-metrics-apiserver v0.0.0-20200618121405-54026617ec44
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.7.0
	github.com/pkg/errors v0.9.0 // indirect
	github.com/prometheus/client_golang v1.0.0
	github.com/prometheus/common v0.10.0
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.4.0
	gopkg.in/yaml.v2 v2.2.8
	k8s.io/api v0.18.5
	k8s.io/apimachinery v0.18.5
	k8s.io/client-go v0.18.5
	k8s.io/component-base v0.18.5
	k8s.io/klog v1.0.0
	k8s.io/metrics v0.18.5
	sigs.k8s.io/metrics-server v0.3.7-0.20200925134111-c39853110962
)

// forced by the inclusion of sigs.k8s.io/metrics-server's use of this in their go.mod
replace k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1 => ./localvendor/k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1
