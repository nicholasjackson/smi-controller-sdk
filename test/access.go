package main

import (
	"k8s.io/kubectl/pkg/scheme"

	accessv1alpha1 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha1"
	accessv1alpha2 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha2"
	accessv1alpha3 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha3"
)

func setupAccess() error {
	err := accessv1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	err = accessv1alpha2.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	err = accessv1alpha3.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	return nil
}
