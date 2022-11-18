package stax

import (
	"testing"

	"github.com/matryer/is"
)

func TestStax(t *testing.T) {
	t.Run("test pop", func(t *testing.T) {
		is := is.New(t)
		items := []Item{
			Item{
				Instruction: "foo",
			},
			Item{
				Instruction: "bar",
			},
			Item{
				Instruction: "biz",
			},
		}
		s := NewStack(items)
		got := s.Pop()
		is.Equal(got.Instruction, "biz")
	})

	t.Run("test push", func(t *testing.T) {
		is := is.New(t)
		items := []Item{
			Item{
				Instruction: "foo",
			},
			Item{
				Instruction: "bar",
			},
			Item{
				Instruction: "biz",
			},
		}
		s := NewStack(items)
		s.Push(Item{Instruction: "baz"})
		got := s.Pop()
		is.Equal(got.Instruction, "baz")
	})
}
