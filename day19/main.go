package main

import (
	"fmt"
	"strings"

	"github.com/radoslawbiesek/aoc2023/utils"
)

const LT = "<"
const GT = ">"
const ACCEPTED = "A"
const REJECTED = "R"
const WORKFLOW_START = "in"

type rule struct {
	parameterName string
	condition     string
	value         int
	next          string
}
type workflow struct {
	name  string
	rules []rule
}
type workflows map[string]workflow
type part = map[string]int

func parseWorkflow(line string) (workflow workflow) {
	splitted := strings.Split(line, "{")
	workflow.name = splitted[0]
	rulesStr := splitted[1][:len(splitted[1])-1]
	for _, ruleStr := range strings.Split(rulesStr, ",") {
		rule := rule{}
		if strings.Contains(ruleStr, LT) {
			rule.condition = LT
		} else if strings.Contains(ruleStr, GT) {
			rule.condition = GT
		}

		if rule.condition == "" {
			rule.next = ruleStr
		} else {
			rule.parameterName = string(ruleStr[0])
			ruleStr = ruleStr[2:]
			splitted := strings.Split(ruleStr, ":")
			rule.value = utils.ParseInt(splitted[0])
			rule.next = splitted[1]
		}

		workflow.rules = append(workflow.rules, rule)
	}
	return
}

func parsePart(line string) part {
	line = line[1 : len(line)-1]
	part := part{}
	for _, ratingStr := range strings.Split(line, ",") {
		splitted := strings.Split(ratingStr, "=")
		name := splitted[0]
		value := utils.ParseInt(splitted[1])
		part[name] = value
	}
	return part
}

func getInput(path string) (workflows, []part) {
	content := utils.GetLines(path, "\n\n")

	workflowStr := content[0]
	workflows := workflows{}
	for _, line := range strings.Split(workflowStr, "\n") {
		workflow := parseWorkflow(line)
		workflows[workflow.name] = workflow
	}

	partsStr := content[1]
	parts := []part{}
	for _, line := range strings.Split(partsStr, "\n") {
		parts = append(parts, parsePart(line))
	}

	return workflows, parts
}

func checkPart1(part *part, workflows *workflows, currentWorkflowName string) (total int) {
	if currentWorkflowName == REJECTED {
		return
	}
	if currentWorkflowName == ACCEPTED {
		for _, value := range *part {
			total += value
		}
		return
	}
	currWorkflow := (*workflows)[currentWorkflowName]
	for _, rule := range currWorkflow.rules {
		if rule.condition == "" {
			return checkPart1(part, workflows, rule.next)
		}
		value := (*part)[rule.parameterName]
		if (rule.condition == LT && value < rule.value) || (rule.condition == GT && value > rule.value) {
			return checkPart1(part, workflows, rule.next)
		}
	}
	return
}

func part1(path string) (total int) {
	workflows, parts := getInput(path)
	for _, part := range parts {
		total += checkPart1(&part, &workflows, WORKFLOW_START)
	}
	return
}

func checkPart2(rules []rule, workflows *workflows, acceptedRules *[][]rule, currentWorkflowName string) {
	if currentWorkflowName == REJECTED {
		return
	}
	if currentWorkflowName == ACCEPTED {
		(*acceptedRules) = append((*acceptedRules), rules)
		return
	}
	currWorkflow := (*workflows)[currentWorkflowName]
	for _, rule := range currWorkflow.rules {
		if rule.condition == "" {
			checkPart2(rules, workflows, acceptedRules, rule.next)
		} else {
			checkPart2(append(rules, rule), workflows, acceptedRules, rule.next)
		}
	}
	return
}

func part2(path string) (total int) {
	workflows, _ := getInput(path)
	acceptedRules := [][]rule{}
	checkPart2([]rule{}, &workflows, &acceptedRules, WORKFLOW_START)
	return
}

func main() {
	fmt.Println("Test input: ")
	fmt.Printf("Part 1: %d\n", part1("./test-input.txt"))
	fmt.Printf("Part 2: %d\n", part2("./test-input.txt"))
	fmt.Println("")
	fmt.Println("Input: ")
	fmt.Printf("Part 1: %d\n", part1("./input.txt"))
	fmt.Printf("Part 2: %d\n", part2("./input.txt"))
}
