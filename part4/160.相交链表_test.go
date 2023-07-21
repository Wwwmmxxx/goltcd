package goltcd

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	intersect := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{Val: 5},
		},
	}

	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case01",
			args: args{
				headA: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  1,
						Next: intersect,
					},
				},
				headB: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  1,
							Next: intersect,
						},
					},
				},
			},
			want: intersect,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
