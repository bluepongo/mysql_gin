package install

import "testing"

func TestUnTarGz(t *testing.T) {
	type args struct {
		srcFilePath string
		destDirPath string
	}
	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		{
			name: "Untargz",
			args: args{
				srcFilePath: "",
				destDirPath: "",
			},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnTarGz(tt.args.srcFilePath, tt.args.destDirPath); (err != nil) != tt.expected {
				t.Errorf("ExecCommand() error = %v, wantErr %v", err, tt.expected)
			}
		})
	}
}
