package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-client-go/internal/config"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"strings"
)

var introspect = &cobra.Command{
	Use:   "introspect",
	Short: "introspect - Introspect the schema",
	Long:  `Introspect the schema. Query is dependent on the permissions on your API key and the version of the API.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Introspect")
		config.Load("./.env")

		c := config.Get()

		// Define the URL
		url := fmt.Sprintf("https://app.myshopify.com/services/graphql/introspection/admin.graphql?api_client_api_key=%s&api_version=%s",
			strings.TrimSpace(c.ShopifyApiKey),
			strings.TrimSpace(c.ApiVersion))

		// Make HTTP GET request
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error making HTTP request:", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		// Check if the response status code is OK (200)
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error: received status code %d\n", resp.StatusCode)
			os.Exit(1)
		}

		// Read the response body
		jsonBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			os.Exit(1)
		}

		var schema schema
		if err := json.Unmarshal(jsonBytes, &schema); err != nil {
			fmt.Println("Error parsing JSON:", err)
			os.Exit(1)
		}

		graphql := processSchema(schema)

		if err := os.WriteFile("schema.graphql", []byte(graphql), 0644); err != nil {
			fmt.Println("Error writing GraphQL schema:", err)
			os.Exit(1)
		}

		fmt.Println("Successfully converted JSON to GraphQL schema")
	},
}

type SchemaType struct {
	Name string `json:"name"`
}

type schema struct {
	Data struct {
		Schema struct {
			QueryType    SchemaType  `json:"queryType"`
			MutationType SchemaType  `json:"mutationType"`
			Types        []Type      `json:"types"`
			Directives   []Directive `json:"directives"`
		} `json:"__schema"`
	} `json:"data"`
}

type Type struct {
	Name          string      `json:"name"`
	Kind          string      `json:"kind"`
	Description   string      `json:"description"`
	Fields        []Field     `json:"fields"`
	InputFields   []Field     `json:"inputFields"`
	Interfaces    []TypeRef   `json:"interfaces"`
	EnumValues    []EnumValue `json:"enumValues"`
	PossibleTypes []TypeRef   `json:"possibleTypes"`
}

type TypeRef struct {
	Kind   string   `json:"kind"`
	Name   string   `json:"name"`
	OfType *TypeRef `json:"ofType"`
}

type Field struct {
	Name              string     `json:"name"`
	Description       string     `json:"description"`
	Type              TypeRef    `json:"type"`
	Args              []Argument `json:"args"`
	IsDeprecated      bool       `json:"isDeprecated"`
	DeprecationReason string     `json:"deprecationReason"`
}

type Argument struct {
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Type              TypeRef `json:"type"`
	DefaultValue      string  `json:"defaultValue"`
	IsDeprecated      bool    `json:"isDeprecated"`
	DeprecationReason string  `json:"deprecationReason"`
}

type EnumValue struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	IsDeprecated      bool   `json:"isDeprecated"`
	DeprecationReason string `json:"deprecationReason"`
}

type Directive struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Locations   []string   `json:"locations"`
	Args        []Argument `json:"args"`
}

func sanitizeString(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return strings.TrimSpace(s)
}

func getFullTypeName(typeRef TypeRef) string {
	if typeRef.Kind == "NON_NULL" {
		return getFullTypeName(*typeRef.OfType) + "!"
	}
	if typeRef.Kind == "LIST" {
		return "[" + getFullTypeName(*typeRef.OfType) + "]"
	}
	return typeRef.Name
}

func processSchema(schema schema) string {
	var builder strings.Builder

	// Write schema definition
	builder.WriteString("schema {\n")
	builder.WriteString(fmt.Sprintf("  query: %s\n", schema.Data.Schema.QueryType.Name))
	builder.WriteString(fmt.Sprintf("  mutation: %s\n", schema.Data.Schema.MutationType.Name))
	builder.WriteString("}\n\n")

	// Write directives
	for _, d := range schema.Data.Schema.Directives {
		if d.Description != "" {
			builder.WriteString(fmt.Sprintf("\"\"\"%s\"\"\"\n", sanitizeString(d.Description)))
		}

		directiveStr := fmt.Sprintf("directive @%s", d.Name)

		if len(d.Args) > 0 {
			args := make([]string, 0, len(d.Args))
			for _, arg := range d.Args {
				argStr := fmt.Sprintf("%s: %s", arg.Name, getFullTypeName(arg.Type))
				if arg.DefaultValue != "" {
					argStr += fmt.Sprintf(" = %s", arg.DefaultValue)
				}
				args = append(args, argStr)
			}
			directiveStr += fmt.Sprintf("(%s)", strings.Join(args, ", "))
		}

		directiveStr += fmt.Sprintf(" on %s\n", strings.Join(d.Locations, " | "))
		builder.WriteString(directiveStr)
	}
	builder.WriteString("\n")

	// Process types
	for _, t := range schema.Data.Schema.Types {
		if strings.HasPrefix(t.Name, "__") {
			continue
		}

		if t.Description != "" {
			builder.WriteString(fmt.Sprintf("\"\"\"%s\"\"\"\n", sanitizeString(t.Description)))
		}

		switch t.Kind {
		case "SCALAR":
			builder.WriteString(fmt.Sprintf("scalar %s\n\n", t.Name))

		case "OBJECT":
			if t.Fields != nil {
				interfaceStr := ""
				if len(t.Interfaces) > 0 {
					interfaces := make([]string, 0, len(t.Interfaces))
					for _, i := range t.Interfaces {
						interfaces = append(interfaces, i.Name)
					}
					interfaceStr = fmt.Sprintf(" implements %s", strings.Join(interfaces, " & "))
				}

				builder.WriteString(fmt.Sprintf("type %s%s {\n", t.Name, interfaceStr))

				for _, f := range t.Fields {
					if strings.HasPrefix(f.Name, "__") {
						continue
					}

					if f.Description != "" {
						builder.WriteString(fmt.Sprintf("  \"\"\"%s\"\"\"\n", sanitizeString(f.Description)))
					}

					typeName := getFullTypeName(f.Type)
					fieldStr := fmt.Sprintf("  %s", f.Name)

					if f.Args != nil && len(f.Args) > 0 {
						args := make([]string, 0, len(f.Args))
						for _, arg := range f.Args {
							argType := getFullTypeName(arg.Type)
							argStr := fmt.Sprintf("%s: %s", arg.Name, argType)
							if arg.DefaultValue != "" {
								argStr += fmt.Sprintf(" = %s", arg.DefaultValue)
							}
							args = append(args, argStr)
						}
						fieldStr += fmt.Sprintf("(%s)", strings.Join(args, ", "))
					}

					fieldStr += fmt.Sprintf(": %s", typeName)

					if f.IsDeprecated {
						reason := f.DeprecationReason
						if reason == "" {
							reason = "No longer supported"
						}
						fieldStr += fmt.Sprintf(" @deprecated(reason: \"%s\")", sanitizeString(reason))
					}

					fieldStr += "\n"
					builder.WriteString(fieldStr)
				}
				builder.WriteString("}\n\n")
			}

		case "INTERFACE":
			if t.Fields != nil {
				builder.WriteString(fmt.Sprintf("interface %s {\n", t.Name))
				for _, f := range t.Fields {
					if strings.HasPrefix(f.Name, "__") {
						continue
					}

					if f.Description != "" {
						builder.WriteString(fmt.Sprintf("  \"\"\"%s\"\"\"\n", sanitizeString(f.Description)))
					}

					typeName := getFullTypeName(f.Type)
					fieldStr := fmt.Sprintf("  %s", f.Name)

					if f.Args != nil && len(f.Args) > 0 {
						args := make([]string, 0, len(f.Args))
						for _, arg := range f.Args {
							argType := getFullTypeName(arg.Type)
							argStr := fmt.Sprintf("%s: %s", arg.Name, argType)
							if arg.DefaultValue != "" {
								argStr += fmt.Sprintf(" = %s", arg.DefaultValue)
							}
							args = append(args, argStr)
						}
						fieldStr += fmt.Sprintf("(%s)", strings.Join(args, ", "))
					}

					fieldStr += fmt.Sprintf(": %s", typeName)

					if f.IsDeprecated {
						reason := f.DeprecationReason
						if reason == "" {
							reason = "No longer supported"
						}
						fieldStr += fmt.Sprintf(" @deprecated(reason: \"%s\")", sanitizeString(reason))
					}

					fieldStr += "\n"
					builder.WriteString(fieldStr)
				}
				builder.WriteString("}\n\n")
			}

		case "UNION":
			if t.PossibleTypes != nil {
				possibleTypes := make([]string, 0, len(t.PossibleTypes))
				for _, pt := range t.PossibleTypes {
					possibleTypes = append(possibleTypes, pt.Name)
				}
				builder.WriteString(fmt.Sprintf("union %s = %s\n\n", t.Name, strings.Join(possibleTypes, " | ")))
			}

		case "ENUM":
			if t.EnumValues != nil {
				builder.WriteString(fmt.Sprintf("enum %s {\n", t.Name))
				for _, ev := range t.EnumValues {
					if strings.HasPrefix(ev.Name, "__") {
						continue
					}

					if ev.Description != "" {
						builder.WriteString(fmt.Sprintf("  \"\"\"%s\"\"\"\n", sanitizeString(ev.Description)))
					}

					enumStr := fmt.Sprintf("  %s", ev.Name)
					if ev.IsDeprecated {
						reason := ev.DeprecationReason
						if reason == "" {
							reason = "No longer supported"
						}
						enumStr += fmt.Sprintf(" @deprecated(reason: \"%s\")", sanitizeString(reason))
					}
					builder.WriteString(enumStr + "\n")
				}
				builder.WriteString("}\n\n")
			}

		case "INPUT_OBJECT":
			if t.InputFields != nil {
				builder.WriteString(fmt.Sprintf("input %s {\n", t.Name))
				for _, f := range t.InputFields {
					if strings.HasPrefix(f.Name, "__") {
						continue
					}

					if f.Description != "" {
						builder.WriteString(fmt.Sprintf("  \"\"\"%s\"\"\"\n", sanitizeString(f.Description)))
					}

					typeName := getFullTypeName(f.Type)
					inputStr := fmt.Sprintf("  %s: %s", f.Name, typeName)
					if f.IsDeprecated {
						reason := f.DeprecationReason
						if reason == "" {
							reason = "No longer supported"
						}
						inputStr += fmt.Sprintf(" @deprecated(reason: \"%s\")", sanitizeString(reason))
					}
					builder.WriteString(inputStr + "\n")
				}
				builder.WriteString("}\n\n")
			}
		}
	}

	return builder.String()
}
