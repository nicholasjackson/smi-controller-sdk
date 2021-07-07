package main

import (
	"k8s.io/kubectl/pkg/scheme"

	splitv1alpha1 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha1"
	splitv1alpha2 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha2"
	splitv1alpha3 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha3"
	splitv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"
)

func setupSplits() error {
	err := splitv1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	err = splitv1alpha2.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	err = splitv1alpha3.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	err = splitv1alpha4.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	return nil
}
