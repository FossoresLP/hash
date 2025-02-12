package hash_test

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"
	"time"

	"github.com/fossoreslp/hash"
)

type Example struct {
	Name   string
	Values []int
	Detail map[string]interface{}
	Typed  map[string]string
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
		Typed: map[string]string{
			"key": "value",
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
				Typed:  nil,
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
	b.Run("xxhash64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hash.Hash(example)
		}
	})
	b.Run("xxhash128", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hash.Hash128(example)
		}
	})
	b.Run("sha256", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hash.HashWithHash(example, sha256.New())
		}
	})
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
	b.Run("xxhash64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hash.Hash(s)
		}
	})
	b.Run("xxhash128", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hash.Hash128(s)
		}
	})
	b.Run("sha256", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hash.HashWithHash(s, sha256.New())
		}
	})
}

func BenchmarkHashString(b *testing.B) {
	s := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Id consectetur purus ut faucibus pulvinar. Ornare quam viverra orci sagittis eu volutpat odio facilisis. Convallis a cras semper auctor. Nec dui nunc mattis enim ut tellus elementum. Ut tellus elementum sagittis vitae et leo duis. Luctus accumsan tortor posuere ac ut consequat semper viverra. Nibh cras pulvinar mattis nunc sed blandit libero volutpat sed. Viverra justo nec ultrices dui sapien eget. Eros donec ac odio tempor. Dapibus ultrices in iaculis nunc sed. Suscipit adipiscing bibendum est ultricies. Vulputate odio ut enim blandit volutpat maecenas volutpat. Vestibulum morbi blandit cursus risus at ultrices mi tempus imperdiet. Neque vitae tempus quam pellentesque nec. Hendrerit dolor magna eget est lorem ipsum dolor sit. Id diam maecenas ultricies mi eget mauris pharetra. Eleifend quam adipiscing vitae proin sagittis nisl rhoncus mattis rhoncus.\n\nProin libero nunc consequat interdum varius sit amet mattis vulputate. Viverra ipsum nunc aliquet bibendum. Tempor orci dapibus ultrices in. In iaculis nunc sed augue. Nam libero justo laoreet sit amet cursus sit amet dictum. Dui nunc mattis enim ut tellus. Sed arcu non odio euismod lacinia at quis. Dolor sit amet consectetur adipiscing elit duis tristique sollicitudin nibh. Gravida rutrum quisque non tellus orci ac. Velit dignissim sodales ut eu sem integer vitae justo eget. Auctor eu augue ut lectus arcu bibendum at. Turpis massa tincidunt dui ut ornare lectus sit amet. Dui sapien eget mi proin sed libero. Proin libero nunc consequat interdum varius. Lectus mauris ultrices eros in cursus turpis. Enim nulla aliquet porttitor lacus luctus accumsan tortor posuere ac. Aliquet enim tortor at auctor urna nunc."
	b.Run("xxhash64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hash.Hash(s)
		}
	})
	b.Run("xxhash128", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hash.Hash128(s)
		}
	})
	b.Run("sha256", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hash.HashWithHash(s, sha256.New())
		}
	})
}

func BenchmarkHashBytes(b *testing.B) {
	bts := make([]byte, 16384)
	_, err := rand.Read(bts)
	if err != nil {
		b.Fatal(err)
	}
	b.Run("xxhash64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hash.Hash(bts)
		}
		b.SetBytes(int64(len(bts)))
	})
	b.Run("xxhash128", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hash.Hash128(bts)
		}
		b.SetBytes(int64(len(bts)))
	})
	b.Run("sha256", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hash.HashWithHash(bts, sha256.New())
		}
		b.SetBytes(int64(len(bts)))
	})
}
