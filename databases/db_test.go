package databases

import (
	"sync"
	"testing"
)

func TestConnectDB(t *testing.T) {
	type args struct {
		wg *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConnectDB(tt.args.wg)
		})
	}
}
