package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"sigs.k8s.io/kind/pkg/apis/config/v1alpha4"
)

func main() {
	cluster := &v1alpha4.Cluster{
		TypeMeta: v1alpha4.TypeMeta{
			Kind:       "Cluster",
			APIVersion: "kind.sigs.k8s.io/v1alpha4",
		},
	}
	v1alpha4.SetDefaultsCluster(cluster)

	tmpl, err := yaml.Marshal(cluster)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create("kind_tmpl.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	file.Write(tmpl)
}
