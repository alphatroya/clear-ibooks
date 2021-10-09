package main

import (
	"strings"
	"testing"
)

func Test_read(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name: "case 1",
			input: `
			«Глава 12. Атака на внешние сущности XML (XXE)»

			Отрывок из книги: Хоффман Э. «Безопасность веб-приложений». Apple Books.
			`,
			want: "Глава 12. Атака на внешние сущности XML (XXE)",
		},
		{
			name: "case 2",
			input: `
			«Глава 12. Атака на внешние сущности XML (XXE)
			»

			Отрывок из книги: Хоффман Э. «Безопасность веб-приложений». Apple Books.
			`,
			want: "Глава 12. Атака на внешние сущности XML (XXE)\n",
		},
		{
			name: "case 3",
			input: `
			«
			Глава 12. Атака на внешние сущности XML (XXE)
			»

			Отрывок из книги: Хоффман Э. «Безопасность веб-приложений». Apple Books.
			`,
			want: "\nГлава 12. Атака на внешние сущности XML (XXE)\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			if got := read(reader); got != tt.want {
				t.Errorf("read(%s) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
