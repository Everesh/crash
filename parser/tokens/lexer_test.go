package tokens

import "testing"

func TestLexerNext(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []rune
	}{
		{
			name:  "ascii string consumed in order",
			input: "abc",
			want:  []rune{'a', 'b', 'c'},
		},
		{
			name:  "unicode runes consumed correctly",
			input: "héllo",
			want:  []rune{'h', 'é', 'l', 'l', 'o'},
		},
		{
			name:  "empty string yields nothing",
			input: "",
			want:  []rune{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLexer(tt.input)
			var got []rune
			for {
				r, ok := l.next()
				if !ok {
					break
				}
				got = append(got, r)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("rune[%d]: got %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestLexerNextExhausted(t *testing.T) {
	l := newLexer("a")
	l.next()
	_, ok := l.next()
	if ok {
		t.Error("expected ok=false after exhausting lexer")
	}
}

func TestLexerPeek(t *testing.T) {
	l := newLexer("ab")

	r, ok := l.peek()
	if !ok || r != 'a' {
		t.Fatalf("peek: got (%q, %v), want ('a', true)", r, ok)
	}

	r, ok = l.peek()
	if !ok || r != 'a' {
		t.Fatalf("second peek: got (%q, %v), want ('a', true)", r, ok)
	}

	l.next()
	r, ok = l.peek()
	if !ok || r != 'b' {
		t.Fatalf("peek after next: got (%q, %v), want ('b', true)", r, ok)
	}
}

func TestLexerPeekExhausted(t *testing.T) {
	l := newLexer("a")
	l.next()
	_, ok := l.peek()
	if ok {
		t.Error("expected ok=false peeking exhausted lexer")
	}
}
