package bulb

import (
	"fmt"
	"testing"
)

func TestLifetime(t *testing.T) {

	type namedLifetime struct {
		name  string
		value Lifetime
	}

	expectedDefinedLifetimes := []namedLifetime{
		{
			name:  "Transient",
			value: Transient,
		},
		{
			name:  "Scoped",
			value: Scoped,
		},
		{
			name:  "Singleton",
			value: Singleton,
		},
	}

	expectedUndefinedLifetimes := []namedLifetime{
		{
			name:  "undefined",
			value: Undefined,
		},
		{
			name:  "Lifetime(127)",
			value: Lifetime(127),
		},
	}

	t.Run(".Defined", func(t *testing.T) {

		t.Run("Lifetime is defined", func(t *testing.T) {
			for _, underTest := range expectedDefinedLifetimes {
				t.Run(fmt.Sprintf("%s/returns true", underTest.name), func(t *testing.T) {
					expected := true
					actual := underTest.value.Defined()
					if expected != actual {
						t.Fatalf("expected .Defined() to return %t; got %t", expected, actual)
					}
				})
			}
		})

		t.Run("Lifetime is not defined", func(t *testing.T) {
			for _, underTest := range expectedUndefinedLifetimes {
				t.Run(fmt.Sprintf("%s/returns true", underTest.name), func(t *testing.T) {
					expected := false
					actual := underTest.value.Defined()
					if expected != actual {
						t.Fatalf("expected .Defined() to return %t; got %t", expected, actual)
					}
				})
			}
		})
	})

	t.Run(".String", func(t *testing.T) {

		t.Run("Lifetime is defined", func(t *testing.T) {
			for _, underTest := range expectedDefinedLifetimes {
				t.Run(fmt.Sprintf("%[1]s/returns %[1]q", underTest.name), func(t *testing.T) {
					expected := underTest.name
					actual := underTest.value.String()
					if expected != actual {
						t.Fatalf("expected .String() to return %q; got %q", expected, actual)
					}
				})
			}
		})

		t.Run("Lifetime is not defined", func(t *testing.T) {
			for _, underTest := range expectedUndefinedLifetimes {
				t.Run(fmt.Sprintf(`%s/returns "Undefined"`, underTest.name), func(t *testing.T) {
					expected := "Undefined"
					actual := underTest.value.String()
					if expected != actual {
						t.Fatalf("expected .String() to return %q; got %q", expected, actual)
					}
				})
			}
		})
	})
}
