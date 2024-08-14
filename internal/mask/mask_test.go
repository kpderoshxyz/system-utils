package mask

import (
    "testing"
    "strings"
)

func TestMaskMethod(t *testing.T) {
    // Define test table go.dev/wiki/TableDrivenTests
    var tests = []struct{
        input string
        expected string
    }{
        {"testingmylongstring", "tes****************[19]"},
        {"test", "tes*[4]"},
        {"te", "te[2]"},
    }
    for _, test := range tests {
        t.Logf("input: %s, expected:%s", test.input, test.expected)
        actual := Mask(test.input)
        t.Logf("%s != %s, %v", actual, test.expected, actual != test.expected)
        if strings.Compare(actual, test.expected) != 0 {
            t.Fatalf("Test failed, input: %s, expected: %s, actual: %s", test.input, test.expected, actual)
        }
    }
}

