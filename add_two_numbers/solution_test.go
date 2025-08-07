package add_two_numbers

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1+2=3",
			args: args{
				l1: &ListNode{
					Val: 1,
				},
				l2: &ListNode{
					Val: 2,
				},
			},
			want: &ListNode{
				Val: 3,
			},
		},
		{
			name: "10+2=12",
			args: args{
				l1: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
					},
				},
				l2: &ListNode{
					Val: 2,
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
		{
			name: "432+654=1086",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
				l2: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
			want: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 8,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
		},
		{
			name: "243+564=708",
			args: args{
				l1: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
				l2: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 8,
					},
				},
			},
		},
		{
			name: "43+654=697",
			args: args{
				l1: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
				l2: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 6,
					},
				},
			},
		},
		{
			name: "0+0=0",
			args: args{
				l1: &ListNode{
					Val: 0,
				},
				l2: &ListNode{
					Val: 0,
				},
			},
			want: &ListNode{
				Val: 0,
			},
		},
		{
			name: "nil+nil=nil",
			args: args{
				l1: nil,
				l2: nil,
			},
			want: &ListNode{Val: 0},
		},
		{
			name: "100+10=110",
			args: args{
				l1: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
		{
			name: "9999999+9999=10009998",
			args: args{
				l1: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 9,
										Next: &ListNode{
											Val:  9,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
				l2: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val:  1,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
