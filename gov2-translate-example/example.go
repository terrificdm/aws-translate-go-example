package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/translate"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		fmt.Println(err)
		return
	}

	sourceLanguageCode := "auto"
	targetLanguageCode := "zh"
	text := "hello world"

	svc := translate.NewFromConfig(cfg)
	resp, err := svc.TranslateText(context.TODO(), &translate.TranslateTextInput{
		SourceLanguageCode: aws.String(sourceLanguageCode),
		TargetLanguageCode: aws.String(targetLanguageCode),
		Text:               aws.String(text),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	resultStr := *resp.TranslatedText
	fmt.Println(resultStr)
}
