package generator

import (
	"fmt"
)

// GenerateFalcoRule genera una regla de Falco básica basada en el nombre de la regla.
func GenerateFalcoRule(ruleName string) (string, error) {
	if ruleName == "" {
		return "", fmt.Errorf("rule name cannot be empty")
	}

	rule := fmt.Sprintf(`
- rule: %s
  desc: Detects suspicious activity
  condition: evt.type=open and fd.name=/etc/shadow
  output: "Suspicious activity detected on %s"
  priority: WARNING
  tags: [user, rules]
`, ruleName, ruleName)

	return rule, nil
}
