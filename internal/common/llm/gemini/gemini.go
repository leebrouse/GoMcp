package gemini

import (
	"context"
	"fmt"
	"log"

	"github.com/leebrouse/GoMcp/internal/common/llm"
	"github.com/leebrouse/GoMcp/internal/common/model"
	"github.com/robermar23/langchaingo/tools/duckduckgo"
	"google.golang.org/genai"
)

type GeminiLLM struct {
	model.LlmModel
}

// NewGeminiLLM creates a new Gemini LLM
func NewGeminiLLM(apiKey string, modelName string, embedder string) llm.LLM {
	return &GeminiLLM{
		LlmModel: model.LlmModel{
			ApiKey:   apiKey,
			Model:    modelName,
			Embedder: embedder,
		},
	}
}

// GenerateText generates text using the Gemini model
func (llm *GeminiLLM) GenerateText(ctx context.Context, prompt string) (string, error) {

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  llm.ApiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return fmt.Sprintf("Error initializing client: %v", err), err
	}

	result, err := client.Models.GenerateContent(
		ctx,
		llm.Model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return fmt.Sprintf("Error generating content: %v", err), err
	}

	return result.Text(), nil

}

// Read local PDF and do operation using the Gemini model
func (llm *GeminiLLM) ReadDocument(ctx context.Context, path string, prompt string) (string, error) {

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  llm.ApiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return fmt.Sprintf("Error initializing client: %v", err), err
	}

	// localPdfPath := "./Kubernetes in Action (Marko Luksa) (Z-Library).pdf"

	uploadConfig := &genai.UploadFileConfig{MIMEType: "application/pdf"}
	uploadedFile, _ := client.Files.UploadFromPath(ctx, path, uploadConfig)

	promptParts := []*genai.Part{
		genai.NewPartFromURI(uploadedFile.URI, uploadedFile.MIMEType),
		genai.NewPartFromText(prompt),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(promptParts, genai.RoleUser),
	}

	result, err := client.Models.GenerateContent(
		ctx,
		llm.Model,
		contents,
		nil,
	)
	if err != nil {
		return fmt.Sprintf("Error Generate Content: %v", err), err
	}

	return result.Text(), nil
}

// Search info from the duckduckgo
func (llm *GeminiLLM) DuckDuckGoSearch(ctx context.Context, keyword string) (string, error) {
	// create DuckDuckGo search tool
	searchTool, err := duckduckgo.New(5, "LangChainGo/1.0")
	if err != nil {
		return fmt.Sprintf("Error init a new DuckDuckGo Search tool : %v", err), err
	}

	// execute searching
	result, err := searchTool.Call(ctx, keyword)
	if err != nil {
		return fmt.Sprintf("Fail to Search info by using DuckDuckGo: %v", err), err
	}

	// return result
	return result, nil
}

// Embeding prompt form string to byte for the RAG system
func (llm *GeminiLLM) Embeding(ctx context.Context, prompt string, role genai.Role) ([]float32, error) {
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Tips (EN): Role denotes the speaker/identity of this content, e.g., RoleUser (user), RoleModel (model).
	// In generation tasks, it affects instruction priority and safety policies; in embedding tasks, it's mostly metadata and typically does not affect vector results.
	contents := []*genai.Content{
		genai.NewContentFromText(prompt, role),
	}

	result, err := client.Models.EmbedContent(ctx,
		llm.Embedder,
		contents,
		nil,
	)
	if err != nil {
		return nil, err
	}

	// Convert embeddings to []float32
	if len(result.Embeddings) == 0 {
		return nil, fmt.Errorf("no embeddings returned")
	}

	embedding := result.Embeddings[0]
	values := make([]float32, len(embedding.Values))
	for i, v := range embedding.Values {
		values[i] = float32(v)
	}

	return values, nil
}
