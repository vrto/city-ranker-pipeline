package parallel_mapper

import (
	"reflect"
	"testing"
)

func TestCity_String(t *testing.T) {
	type fields struct {
		name       string
		population int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "should print Prague",
			fields: fields{
				name:       "Prague",
				population: 1350000,
			},
			want: "Prague (1,350,000)",
		},
		{
			name: "should print Brno",
			fields: fields{
				name:       "Brno",
				population: 383000,
			},
			want: "Brno (383,000)",
		},
		{
			name: "should print České Budějovice",
			fields: fields{
				name:       "České Budějovice",
				population: 94229,
			},
			want: "České Budějovice (94,229)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &City{
				name:       tt.fields.name,
				population: tt.fields.population,
			}
			if got := this.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCity(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  *City
	}{
		{
			name:  "should parse Prague",
			input: "Prague (1,350,000)",
			want: &City{
				name:       "Prague",
				population: 1350000,
			},
		},
		{
			name:  "should parse Brno",
			input: "Brno (383,000)",
			want: &City{
				name:       "Brno",
				population: 383000,
			},
		},
		{
			name:  "should parse a city with accented character",
			input: "České Budějovice (94,229)",
			want: &City{
				name:       "České Budějovice",
				population: 94229,
			},
		},
		{
			name:  "should return nil on invalid format",
			input: "invalid format",
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCity(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
