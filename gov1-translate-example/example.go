package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sourceLanguageCode := "auto"
	targetLanguageCode := "zh"
	text := "Hello world"

	svc := translate.New(sess)
	resp, err := svc.Text(&translate.TextInput{
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
