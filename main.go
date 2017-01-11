package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"k8s.io/kubernetes/pkg/api"
	_ "k8s.io/kubernetes/pkg/api/install"
	"k8s.io/kubernetes/pkg/api/validation"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/yaml"
)

func main() {
	//	format := flag.String("format", "", "format of the file -json or yaml - if unset, attempt to detect from file extension")
	flag.Parse()

	if len(os.Args) == 0 {
		fmt.Println("need a file")
		os.Exit(-2)
	}

	//	v1.E

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	json, err := yaml.ToJSON(data)
	if err != nil {
		log.Fatal("unable to coerce into json ", err)
	}
	obj, err := runtime.Decode(api.Codecs.UniversalDecoder(), json)
	if err != nil {
		log.Fatal("unable to decode object ", err)
	}

	pod, ok := obj.(*api.Pod)
	if !ok {
		log.Fatalf("invalid pod: %#v", obj)
	}

	if errs := validation.ValidatePod(pod); len(errs) > 0 {
		log.Fatalf("invalid pod: %v\n", errs)
	}

	if len(pod.Spec.Containers) == 0 {
		log.Fatal("pod has no containers")
	}

}
