package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	setClusterMetrics(map[string]string{
		template_label_name:     templateName1,
		tenant_id_label_name:    "tenant_id_1",
		cluster_name_label_name: "cluster_name_1",
	})
	setClusterMetrics(map[string]string{
		template_label_name:     templateName2,
		tenant_id_label_name:    "tenant_id_2",
		cluster_name_label_name: "cluster_name_2",
	})
	setClusterMetrics(map[string]string{
		template_label_name:     templateName3,
		tenant_id_label_name:    "tenant_id_1",
		cluster_name_label_name: "cluster_name_3",
	})
	setClusterMetrics(map[string]string{
		template_label_name:     templateName4,
		tenant_id_label_name:    "tenant_id_1",
		cluster_name_label_name: "cluster_name_4",
	})

	http.Handle("/metrics", promhttp.Handler())
	log.Println("listening on", ":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
