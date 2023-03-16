package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	storage_pct_metric_name = "storage_pct" // "scaling.thresholds.storage_pct"

	tenant_id_label_name    = "tenant_id"
	cluster_name_label_name = "cluster_name"
	template_label_name     = "template"
)

var storagePct = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: storage_pct_metric_name,
		Help: "Percentage of storage that should be held available for failure tolerance and defrag.",
	},
	[]string{cluster_name_label_name, tenant_id_label_name},
)

var clusterTemplateInfo = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "cluster_template_info",
	},
	[]string{template_label_name, cluster_name_label_name, tenant_id_label_name},
)

func setClusterMetrics(labels map[string]string) {
	clusterTemplateInfo.With(prometheus.Labels{
		template_label_name:     labels[template_label_name],
		tenant_id_label_name:    labels[tenant_id_label_name],
		cluster_name_label_name: labels[cluster_name_label_name],
	}).Set(1)

	template := templates[labels[template_label_name]]

	if template[storage_pct_metric_name] == 0 {
		// don't create a storage pct metric if the value is 0
		return
	}
	storagePct.With(prometheus.Labels{
		tenant_id_label_name:    labels[tenant_id_label_name],
		cluster_name_label_name: labels[cluster_name_label_name],
	},
	).Set(template[storage_pct_metric_name])
}
