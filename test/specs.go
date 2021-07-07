package main

import (
	specsv1alpha1 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha1"
	specsv1alpha2 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha2"
	specsv1alpha3 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha3"
	specsv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha4"
	"k8s.io/kubectl/pkg/scheme"
)

func setupSpecs() error {
	err := specsv1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	err = specsv1alpha2.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	err = specsv1alpha3.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	err = specsv1alpha4.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	return nil
}
