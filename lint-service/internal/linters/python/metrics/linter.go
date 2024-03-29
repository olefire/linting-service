package metrics

import (
	"encoding/json"
	"fmt"
	"lint-service/internal/models"
	"lint-service/pkg"
)

type Metrics struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Complexity int    `json:"complexity"`
	Rank       string `json:"rank"`
	Lineno     int    `json:"lineno"`
}

var (
	Radon models.Linter = "Radon"
)

type Linter struct{}

func (l *Linter) LintFile(file models.SourceFile) (models.LintResult, error) {

	content, stderr, err := pkg.Execute(file.Code, "radon", "cc", "-j", "-")

	if stderr != "" {
		fmt.Printf("lintFile stderr: %s\n", stderr)
	}

	if err != nil {
		return models.LintResult{Linter: Radon}, err
	}

	var data map[string][]Metrics

	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return models.LintResult{Linter: Radon}, err
	}
	lintRes := data["-"]
	issues := make([]models.LintCodeIssue, len(lintRes))

	for i, metric := range lintRes {
		issues[i].Message = fmt.Sprintf("Cyclomatic complexity of %s %s is : %d. Сode quality assessment: %s ", metric.Type, metric.Name, metric.Complexity, metric.Rank)
		issues[i].Line = metric.Lineno
	}

	return models.LintResult{Issues: issues, Linter: Radon}, nil
}
