package mask

import (
    "fmt"
    "strings"
)

// Calculate the masked string when length is greater than 3
func Mask(text string) string {
    if len(text) >= 3 {
        return fmt.Sprintf("%s%s[%d]", text[:3], strings.Repeat("*", len(text) - 3), len(text))
    } else {
        return fmt.Sprintf("%s[%d]", text, len(text))
    }
}
