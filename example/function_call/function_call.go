package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const apiURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent"

type ContentPart struct {
	Text string `json:"text"`
}

type Content struct {
	Role  string        `json:"role"`
	Parts []ContentPart `json:"parts"`
}

type ToolParamProperty struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Items       *struct {
		Type string `json:"type"`
	} `json:"items,omitempty"`
}

type ToolParameters struct {
	Type       string                       `json:"type"`
	Properties map[string]ToolParamProperty `json:"properties"`
	Required   []string                     `json:"required"`
}

type FunctionDeclaration struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Parameters  ToolParameters `json:"parameters"`
}

type Tool struct {
	FunctionDeclarations []FunctionDeclaration `json:"functionDeclarations"`
}

type RequestBody struct {
	Contents []Content `json:"contents"`
	Tools    []Tool    `json:"tools"`
}

func main() {
	ctx := context.Background()

	apiKey := "AIzaSyCKURVV8jEX3CsRu_4pysxmJm3IH4mr8VU"
	if apiKey == "" {
		fmt.Println("Please set GEMINI_API_KEY environment variable")
		return
	}

	reqBody := RequestBody{
		Contents: []Content{
			{
				Role: "user",
				Parts: []ContentPart{
					{Text: "Schedule a meeting with Bob and Alice for 03/27/2025 at 10:00 AM about the Q3 planning."},
				},
			},
		},
		Tools: []Tool{
			{
				FunctionDeclarations: []FunctionDeclaration{
					{
						Name:        "schedule_meeting",
						Description: "Schedules a meeting with specified attendees at a given time and date.",
						Parameters: ToolParameters{
							Type: "object",
							Properties: map[string]ToolParamProperty{
								"attendees": {
									Type:        "array",
									Description: "List of people attending the meeting.",
									Items: &struct {
										Type string `json:"type"`
									}{Type: "string"},
								},
								"date":  {Type: "string", Description: "Date of the meeting (e.g., '2024-07-29')"},
								"time":  {Type: "string", Description: "Time of the meeting (e.g., '15:00')"},
								"topic": {Type: "string", Description: "The subject or topic of the meeting."},
							},
							Required: []string{"attendees", "date", "time", "topic"},
						},
					},
				},
			},
		},
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", apiURL+"?key="+apiKey, bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-goog-api-key", apiKey)

	client := http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
