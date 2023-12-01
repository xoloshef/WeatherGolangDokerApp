package service_test

import (
	"testing"

	"WeatherGoDokerApp/internal/app/service"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	// Определение тестовых случаев.
	tests := []struct {
		input    string
		expected string
	}{
		{"abcabcbb", "abc"},
		{"bbbb", "b"},
		{"pwwkew", "wke"},в
		{"", ""},
		{"a", "a"},в
	}
	// Создание нового экземпляра сервиса.
	s := service.New()
	// Итерация по тестовым случаям.
	for _, test := range tests {
		// Вызов метода LongestSubstring для каждого тестового случая.
		result := s.LongestSubstring(test.input)
		// Проверка результата с ожидаемым значением.
		if result != test.expected {
			t.Errorf("Input string- %s, wait- %s, have- %s", test.input, test.expected, result)
		}
	}
}
