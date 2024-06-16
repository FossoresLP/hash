package hash_test

import (
	"testing"
	"time"

	"github.com/fossoreslp/hash"
)

type Example struct {
	Name   string
	Values []int
	Detail map[string]interface{}
	Nested struct {
		SubName   string
		Recursive *Example
	}
}

type Realistic struct {
	ID          uint64
	Active      bool
	Title       string
	Paragraph   string
	Text        string
	ValueString string
	ValueInt    int
	ValueFloat  float64
	ValueBool   bool
	ValueSlice  []string
	AccessID    uint64
	Ordering    uint64
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	Recursive   *Realistic
}

func BenchmarkHashStruct(b *testing.B) {
	example := Example{
		Name:   "example",
		Values: []int{1, 2, 3},
		Detail: map[string]interface{}{
			"age":    30,
			"tags":   []string{"go", "programming"},
			"info":   "example struct for hashing",
			"floats": []float64{1.2, 3.4, 5.6, 7.8, 9.0},
		},
		Nested: struct {
			SubName   string
			Recursive *Example
		}{
			SubName: "sub",
			Recursive: &Example{
				Name:   "recursive",
				Values: []int{4, 5, 6},
				Detail: nil,
				Nested: struct {
					SubName   string
					Recursive *Example
				}{
					SubName:   "sub",
					Recursive: nil,
				},
			},
		},
	}
	for i := 0; i < b.N; i++ {
		_ = hash.Hash(example)
	}
}

func BenchmarkHashStructRealistic(b *testing.B) {
	s := Realistic{
		ID:          1,
		Active:      true,
		Title:       "Lorem ipsum dolor sit amet",
		Paragraph:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Id consectetur purus ut faucibus pulvinar. Ornare quam viverra orci sagittis eu volutpat odio facilisis. Convallis a cras semper auctor. Nec dui nunc mattis enim ut tellus elementum.",
		Text:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Id consectetur purus ut faucibus pulvinar. Ornare quam viverra orci sagittis eu volutpat odio facilisis. Convallis a cras semper auctor. Nec dui nunc mattis enim ut tellus elementum. Ut tellus elementum sagittis vitae et leo duis. Luctus accumsan tortor posuere ac ut consequat semper viverra. Nibh cras pulvinar mattis nunc sed blandit libero volutpat sed. Viverra justo nec ultrices dui sapien eget. Eros donec ac odio tempor. Dapibus ultrices in iaculis nunc sed. Suscipit adipiscing bibendum est ultricies. Vulputate odio ut enim blandit volutpat maecenas volutpat. Vestibulum morbi blandit cursus risus at ultrices mi tempus imperdiet. Neque vitae tempus quam pellentesque nec. Hendrerit dolor magna eget est lorem ipsum dolor sit. Id diam maecenas ultricies mi eget mauris pharetra. Eleifend quam adipiscing vitae proin sagittis nisl rhoncus mattis rhoncus.\n\nProin libero nunc consequat interdum varius sit amet mattis vulputate. Viverra ipsum nunc aliquet bibendum. Tempor orci dapibus ultrices in. In iaculis nunc sed augue. Nam libero justo laoreet sit amet cursus sit amet dictum. Dui nunc mattis enim ut tellus. Sed arcu non odio euismod lacinia at quis. Dolor sit amet consectetur adipiscing elit duis tristique sollicitudin nibh. Gravida rutrum quisque non tellus orci ac. Velit dignissim sodales ut eu sem integer vitae justo eget. Auctor eu augue ut lectus arcu bibendum at. Turpis massa tincidunt dui ut ornare lectus sit amet. Dui sapien eget mi proin sed libero. Proin libero nunc consequat interdum varius. Lectus mauris ultrices eros in cursus turpis. Enim nulla aliquet porttitor lacus luctus accumsan tortor posuere ac. Aliquet enim tortor at auctor urna nunc.",
		ValueString: "value",
		ValueInt:    1,
		ValueFloat:  1.1,
		ValueBool:   true,
		ValueSlice:  []string{"a", "b", "c"},
		AccessID:    1,
		Ordering:    1,
		CreatedAt:   &time.Time{},
		UpdatedAt:   &time.Time{},
		DeletedAt:   &time.Time{},
		Recursive: &Realistic{
			ID:          2,
			Active:      false,
			Title:       "Lorem ipsum dolor sit amet",
			Paragraph:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Id consectetur purus ut faucibus pulvinar. Ornare quam viverra orci sagittis eu volutpat odio facilisis. Convallis a cras semper auctor. Nec dui nunc mattis enim ut tellus elementum.",
			Text:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Id consectetur purus ut faucibus pulvinar. Ornare quam viverra orci sagittis eu volutpat odio facilisis. Convallis a cras semper auctor. Nec dui nunc mattis enim ut tellus elementum. Ut tellus elementum sagittis vitae et leo duis. Luctus accumsan tortor posuere ac ut consequat semper viverra. Nibh cras pulvinar mattis nunc sed blandit libero volutpat sed. Viverra justo nec ultrices dui sapien eget. Eros donec ac odio tempor. Dapibus ultrices in iaculis nunc sed. Suscipit adipiscing bibendum est ultricies. Vulputate odio ut enim blandit volutpat maecenas volutpat. Vestibulum morbi blandit cursus risus at ultrices mi tempus imperdiet. Neque vitae tempus quam pellentesque nec. Hendrerit dolor magna eget est lorem ipsum dolor sit. Id diam maecenas ultricies mi eget mauris pharetra. Eleifend quam adipiscing vitae proin sagittis nisl rhoncus mattis rhoncus.\n\nProin libero nunc consequat interdum varius sit amet mattis vulputate. Viverra ipsum nunc aliquet bibendum. Tempor orci dapibus ultrices in. In iaculis nunc sed augue. Nam libero justo laoreet sit amet cursus sit amet dictum. Dui nunc mattis enim ut tellus. Sed arcu non odio euismod lacinia at quis. Dolor sit amet consectetur adipiscing elit duis tristique sollicitudin nibh. Gravida rutrum quisque non tellus orci ac. Velit dignissim sodales ut eu sem integer vitae justo eget. Auctor eu augue ut lectus arcu bibendum at. Turpis massa tincidunt dui ut ornare lectus sit amet. Dui sapien eget mi proin sed libero. Proin libero nunc consequat interdum varius. Lectus mauris ultrices eros in cursus turpis. Enim nulla aliquet porttitor lacus luctus accumsan tortor posuere ac. Aliquet enim tortor at auctor urna nunc.",
			ValueString: "value",
			ValueInt:    1,
			ValueFloat:  1.1,
			ValueBool:   true,
			ValueSlice:  []string{"a", "b", "c"},
			AccessID:    1,
			Ordering:    1,
			CreatedAt:   &time.Time{},
			UpdatedAt:   &time.Time{},
			DeletedAt:   &time.Time{},
		},
	}
	for i := 0; i < b.N; i++ {
		_ = hash.Hash(s)
	}
}
