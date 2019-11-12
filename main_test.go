package main

import "testing"

const (
	output = `interface dialer 1
interface dialer 0 overload
interface FastEthernet2
interface FastEthernet3
interface FastEthernet4
interface FastEthernet5
interface FastEthernet6
interface FastEthernet7
interface FastEthernet8
interface FastEthernet9
interface FastEthernet0
interface FastEthernet1
interface Dot11Radio0
interface Dot11Radio0.1
interface Dot11Radio0.2
interface Dot11Radio0.3
interface Vlan1
interface Vlan2
interface Vlan3
interface BVI1
interface BVI2
interface BVI3`
)

func Test_testInterpreter(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "normal output",
			want: output,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testInterpreter(); got != tt.want {
				t.Errorf("testInterpreter() = %v, want %v", got, tt.want)
			}
		})
	}
}
