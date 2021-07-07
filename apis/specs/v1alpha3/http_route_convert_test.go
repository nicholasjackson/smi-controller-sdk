package v1alpha3

import (
	"testing"

	"github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha4"
	assert "github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestConvertHTTPRouteGroupToConvertsFromAlpha4ToAlpha3(t *testing.T) {
	v3Test := &HTTPRouteGroup{}

	err := v3Test.ConvertFrom(v4HTTPRoute)
	assert.NoError(t, err)

	assert.Equal(t, v4HTTPRoute.ObjectMeta, v3Test.ObjectMeta)
	assert.Equal(t, v4HTTPRoute.TypeMeta.Kind, v3Test.TypeMeta.Kind)
	assert.Equal(t, GroupVersion.Identifier(), v3Test.TypeMeta.APIVersion)

	for i, m := range v4HTTPRoute.Spec.Matches {
		v3 := v3Test.Spec.Matches[i]

		assert.Equal(t, m.Name, v3.Name)

		// check the methods
		assert.Equal(t, m.Methods[0], v3.Methods[0])
		assert.Equal(t, m.Methods[1], v3.Methods[1])

		assert.Equal(t, m.PathRegex, v3.PathRegex)

		// check the headers
		for k, v := range m.Headers {
			assert.Equal(t, v3.Headers[k], v)
		}
	}
}

func TestConvertHTTPRouteGroupToConvertsFromAlpha3ToAlpha4(t *testing.T) {
	v4Test := &v1alpha4.HTTPRouteGroup{}

	err := v3HTTPRoute.ConvertTo(v4Test)
	assert.NoError(t, err)

	assert.Equal(t, v3HTTPRoute.ObjectMeta, v4Test.ObjectMeta)
	assert.Equal(t, v3HTTPRoute.TypeMeta.Kind, v4Test.TypeMeta.Kind)
	assert.Equal(t, v1alpha4.GroupVersion.Identifier(), v4Test.TypeMeta.APIVersion)

	for i, m := range v3HTTPRoute.Spec.Matches {
		v4 := v4Test.Spec.Matches[i]

		assert.Equal(t, m.Name, v4.Name)

		// check the methods
		assert.Equal(t, m.Methods[0], v4.Methods[0])
		assert.Equal(t, m.Methods[1], v4.Methods[1])

		assert.Equal(t, m.PathRegex, v4.PathRegex)

		// check the headers
		for k, v := range m.Headers {
			assert.Equal(t, v4.Headers[k], v)
		}
	}
}

var v4HTTPRoute = &v1alpha4.HTTPRouteGroup{
	TypeMeta: v1.TypeMeta{
		Kind:       "HTTPRouteGroup",
		APIVersion: "v1alpha4",
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v4HTTPRouteGroup",
		Namespace: "default",
	},
	Spec: v1alpha4.HTTPRouteGroupSpec{
		Matches: []v1alpha4.HTTPMatch{
			v1alpha4.HTTPMatch{
				Name:      "testing1",
				Methods:   []string{"GET", "POST"},
				PathRegex: ".*",
				Headers:   map[string]string{"Foo": "Bar", "One": "Two"},
			},
			v1alpha4.HTTPMatch{
				Name:      "testing2",
				Methods:   []string{"DELETE", "POST"},
				PathRegex: "/post",
				Headers:   map[string]string{"abc": "123", "Mario": "Luigi"},
			},
		},
	},
}

var v3HTTPRoute = &HTTPRouteGroup{
	TypeMeta: v1.TypeMeta{
		Kind:       "HTTPRouteGroup",
		APIVersion: "v1alpha3",
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v3HTTPRouteGroup",
		Namespace: "default",
	},
	Spec: HTTPRouteGroupSpec{
		Matches: []HTTPMatch{
			HTTPMatch{
				Name:      "testing1",
				Methods:   []string{"GET", "POST"},
				PathRegex: ".*",
				Headers:   map[string]string{"Foo": "Bar", "One": "Two"},
			},
			HTTPMatch{
				Name:      "testing2",
				Methods:   []string{"DELETE", "POST"},
				PathRegex: "/post",
				Headers:   map[string]string{"abc": "123", "Mario": "Luigi"},
			},
		},
	},
}
