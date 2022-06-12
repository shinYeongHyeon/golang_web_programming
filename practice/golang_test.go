package practice

import (
	"context"
	"github.com/stretchr/testify/assert"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestGolang(t *testing.T) {
	t.Run("string test", func(t *testing.T) {
		str := "Ann,Jenny,Tom,Zico"
		actual := strings.Split(str, ",")
		expected := []string{"Ann", "Jenny", "Tom", "Zico"}
		assert.Equal(t, expected, actual)
	})

	t.Run("goroutine에서 slice에 값 추가해보기", func(t *testing.T) {
		var numbers []int
		var waitGroup sync.WaitGroup
		waitGroup.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				numbers = append(numbers, i)
			}
			waitGroup.Done()
		}()
		waitGroup.Wait()

		var expected []int // actual : [0 1 2 ... 100]
		for i := 0; i < 100; i++ {
			expected = append(expected, i)
		}

		assert.ElementsMatch(t, expected, numbers)
	})

	t.Run("fan out, fan in", func(t *testing.T) {
		inputCh := generate()
		outputCh := make(chan int, 3)
		var waitGroup sync.WaitGroup
		waitGroup.Add(3)

		go func(ch <-chan int) {
			for {
				select {
				case value, ok := <-ch:
					if !ok {
						return
					}
					outputCh <- value * 10
					waitGroup.Done()
				}
			}
		}(inputCh)
		waitGroup.Wait()
		close(outputCh)

		var actual []int
		for value := range outputCh {
			actual = append(actual, value)
		}
		expected := []int{10, 20, 30}
		assert.Equal(t, expected, actual)
	})

	t.Run("context timeout", func(t *testing.T) {
		startTime := time.Now()
		add := time.Second * 3
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, add)
		defer cancel()

		var endTime time.Time
		select {
		case <-ctx.Done():
			endTime = time.Now()
			break
		}

		assert.True(t, endTime.After(startTime.Add(add)))
	})

	t.Run("context deadline", func(t *testing.T) {
		add := time.Second * 3
		ctx := context.Background()
		startTime := time.Now()
		ctx, cancel := context.WithDeadline(ctx, startTime.Add(add))
		defer cancel()

		var endTime time.Time
		select {
		case <-ctx.Done():
			endTime = time.Now()
			break
		}

		assert.True(t, endTime.After(startTime.Add(add)))
	})

	t.Run("context value", func(t *testing.T) {
		const existKey1 = "exist_key1"
		const key1Value = "value1"
		const existKey2 = "exist_key2"
		const key2Value = "value2"

		ctx := context.Background()
		ctx = context.WithValue(ctx, existKey1, key1Value)
		ctx = context.WithValue(ctx, existKey2, key2Value)

		assert.Equal(t, key1Value, ctx.Value(existKey1))
		assert.Equal(t, key2Value, ctx.Value(existKey2))
		assert.Nil(t, ctx.Value("none_key"))
	})
}

func generate() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 3; i++ {
			ch <- i
		}
	}()
	return ch
}
