package openai

import (
	"bufio"
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type AudioRequest struct {
	InputFile  string
	OutPutFile string
} // struct for audio request

func SetFile(InputFile, OutputFile string) *AudioRequest {
	A := AudioRequest{
		InputFile:  InputFile,
		OutPutFile: OutputFile,
	}
	return &A
}
func Connect(apikey string) *openai.Client {
	client := openai.NewClient(apikey)
	return client
}

func Transcribe(client *openai.Client, Name string) {
	ctx := context.Background()
	fmt.Println("Transcribing: ", Name)
	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: "./transcribe.mp3",
	}
	resp, err := client.CreateTranscription(ctx, req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return
	}

	fmt.Println("Transcription saved to: ", Name+".txt")
	f, err := os.Create(Name + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString(resp.Text)
	w.Flush()

}

func SubmitToGPT4(client *openai.Client, inputFile string, query string) {
	// Read the input file
	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		return
	}

	// Prepare the GPT-4 request
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query + string(inputData),
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	fmt.Println(inputFile)
	filename := inputFile + ".query.txt"
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString(resp.Choices[0].Message.Content)
	w.Flush()
}
