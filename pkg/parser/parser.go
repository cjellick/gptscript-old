package parser

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/acorn-io/gptscript/pkg/openai"
	"github.com/acorn-io/gptscript/pkg/types"
	"golang.org/x/exp/maps"
)

var defaultToolSchema = types.JSONSchema{
	Property: types.Property{
		Type: "object",
	},
	Properties: map[string]types.Property{
		openai.DefaultPromptParameter: {
			Description: "Prompt to send to the tool or assistant. This may be instructions or question",
			Type:        "string",
		},
	},
	Required: []string{openai.DefaultPromptParameter},
}

func normalize(key string) string {
	return strings.ToLower(strings.ReplaceAll(key, " ", ""))
}

func toBool(line string) (bool, error) {
	if line == "true" {
		return true, nil
	} else if line != "false" {
		return false, fmt.Errorf("invalid boolean parameter, must be \"true\" or \"false\", got [%s]", line)
	}
	return false, nil
}

func csv(line string) (result []string) {
	for _, part := range strings.Split(line, ",") {
		result = append(result, strings.TrimSpace(part))
	}
	return
}

func addArg(line string, tool *types.Tool) error {
	if tool.Arguments == nil {
		tool.Arguments = &types.JSONSchema{
			Property: types.Property{
				Type: "object",
			},
			Properties: map[string]types.Property{},
		}
	}

	key, value, ok := strings.Cut(line, ":")
	if !ok {
		return fmt.Errorf("invalid arg format: %s", line)
	}

	tool.Arguments.Properties[key] = types.Property{
		Description: value,
		Type:        "string",
	}

	return nil
}

func isParam(line string, tool *types.Tool) (_ bool, err error) {
	key, value, ok := strings.Cut(line, ":")
	if !ok {
		return false, nil
	}
	value = strings.TrimSpace(value)
	switch normalize(key) {
	case "tool":
		fallthrough
	case "name":
		tool.Name = strings.ToLower(value)
	case "modelname":
		tool.ModelName = value
	case "description":
		tool.Description = value
	case "tools":
		tool.Tools = csv(strings.ToLower(value))
	case "args":
		fallthrough
	case "arg":
		if err := addArg(value, tool); err != nil {
			return false, err
		}
	case "vision":
		tool.Vision, err = toBool(value)
		if err != nil {
			return false, err
		}
	case "maxtokens":
		tool.MaxTokens, err = strconv.Atoi(value)
		if err != nil {
			return false, err
		}
	case "cache":
		b, err := toBool(value)
		if err != nil {
			return false, err
		}
		tool.Cache = &b
	case "jsonresponse":
		tool.JSONResponse, err = toBool(value)
		if err != nil {
			return false, err
		}
	default:
		return false, nil
	}

	return true, nil
}

type ErrLine struct {
	Line int
	Err  error
}

func (e *ErrLine) Unwrap() error {
	return e.Err
}

func (e *ErrLine) Error() string {
	return fmt.Sprintf("line %d: %v", e.Line, e.Err)
}

func errLine(lineNo int, err error) error {
	return &ErrLine{
		Line: lineNo,
		Err:  err,
	}
}

type context struct {
	tool         types.Tool
	instructions []string
	inBody       bool
}

func (c *context) finish(tools *[]types.Tool) {
	c.tool.Instructions = strings.TrimSpace(strings.Join(c.instructions, ""))
	if c.tool.Instructions != "" || c.tool.Name != "" {
		if !c.tool.IsCommand() && c.tool.Arguments == nil {
			c.tool.Arguments = &defaultToolSchema
		}
		*tools = append(*tools, c.tool)
	}
	*c = context{}
}

func Parse(input io.Reader) (types.Tool, types.ToolSet, error) {
	scan := bufio.NewScanner(input)

	var (
		tools   []types.Tool
		context context
		lineNo  int
	)

	for scan.Scan() {
		lineNo++
		line := scan.Text() + "\n"

		if line == "---\n" {
			context.finish(&tools)
			continue
		}

		if !context.inBody {
			if strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "#!") {
				continue
			}
			if strings.TrimSpace(line) == "" {
				continue
			}
			if isParam, err := isParam(line, &context.tool); err != nil {
				return types.Tool{}, nil, errLine(lineNo, err)
			} else if isParam {
				continue
			}
		}

		context.inBody = true
		context.instructions = append(context.instructions, line)
	}

	context.finish(&tools)

	var (
		mainTool types.Tool
		toolSet  = types.ToolSet{}
	)

	for i, tool := range tools {
		if i == 0 {
			mainTool = tool
		}
		if tool.Name != "" {
			if _, ok := toolSet[tool.Name]; ok {
				return mainTool, nil, fmt.Errorf("duplicate tool named %s", tool.Name)
			}
			toolSet[tool.Name] = tool
		}
	}

	for _, tool := range append(maps.Values(toolSet), mainTool) {
		for _, tool := range tool.Tools {
			if _, ok := toolSet[tool]; !ok {
				return mainTool, nil, fmt.Errorf("missing reference to tool named: %s", tool)
			}
		}
	}

	return mainTool, toolSet, nil
}
