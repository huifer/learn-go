package function

import "testing"

func Test_add4(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// 添加单元测试数据
		{name: "1+1", args: args{1, 1}, want: 2},
		{name: "1+2", args: args{1, 2}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add4(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("add4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAdd(t *testing.B) {
	for i := 0; i < t.N; i++ {
		add4(1, 2)
	}
}
