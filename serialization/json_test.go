package serialization

import "testing"

// mock struct
type MockPerson struct {
	Name string `json:"name"`
	Age  int
}

func TestToJSON(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "string test",
			args: args{
				in: "my test string",
			},
			want:    "my test string",
			wantErr: false,
		}, {
			name: "struct test",
			args: args{
				in: MockPerson{
					Name: "paul",
					Age:  21,
				},
			},
			want:    "{\"name\":\"paul\",\"Age\":21}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSON(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
