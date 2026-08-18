package reflection

import (
	"reflect"
	"testing"
)

type (
	Profile struct {
		Age  int
		City string
	}

	Person struct {
		Name    string
		Profile Profile
	}
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Umbra"},
			[]string{"Umbra"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Umbra", "Brooklyn"},
			[]string{"Umbra", "Brooklyn"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Umbra", 45},
			[]string{"Umbra"},
		},
		{
			"nested fields",
			Person{
				"Umbra",
				Profile{45, "Brooklyn"},
			},
			[]string{"Umbra", "Brooklyn"},
		},
		{
			"pointers to things",
			&Person{
				"Umbra",
				Profile{45, "Brooklyn"},
			},
			[]string{"Umbra", "Brooklyn"},
		},
		{
			"slices",
			[]Profile{
				{45, "Brooklyn"},
				{23, "Orlando"},
			},
			[]string{"Brooklyn", "Orlando"},
		},
		{
			"arrays",
			[2]Profile{
				{23, "Orlando"},
				{45, "Brooklyn"},
			},
			[]string{"Orlando", "Brooklyn"},
		},
		{
			"maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertLength(t, got, len(aMap))
		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{45, "Brooklyn"}
			aChannel <- Profile{23, "Orlando"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Brooklyn", "Orlando"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{45, "Brooklyn"}, Profile{23, "Orlando"}
		}

		var got []string
		want := []string{"Brooklyn", "Orlando"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}

func assertLength(t testing.TB, got []string, want int) {
	t.Helper()
	if len(got) != want {
		t.Errorf("got %d values but expected %d", len(got), want)
	}
}
