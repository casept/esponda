package main

import (
	_ "github.com/casept/esponda/statik"

	"github.com/bwmarrin/discordgo"
)

func createAndSendSchillerQuoteMaymay(s *discordgo.Session, maymayText string, channelID string) {
	createAndSendMaymay(s, maymayText, channelID, createSchillerQuoteMaymay(maymayText))
}

func createSchillerQuoteMaymay(maymayText string) ImageAnnotationRecipe {
	return ImageAnnotationRecipe{
		InputImage:     "SchillerQuoteBlank.png",
		CaptionFont:    "RalewayUnicode.ttf",
		CaptionColor:   "#042B66",
		CaptionText:    maymayText,
		TextDimensions: "590x260",
		TextGeometry:   "+420+120",
	}
}

func createAndSendSchillerSpeechMaymay(s *discordgo.Session, maymayText string, channelID string) {
	createAndSendMaymay(s, maymayText, channelID, createSchillerSpeechMaymay(maymayText))
}

func createSchillerSpeechMaymay(maymayText string) ImageAnnotationRecipe {
	return ImageAnnotationRecipe{
		InputImage:     "SchillerSpeechBlank.png",
		CaptionFont:    "Helvetica.ttf",
		CaptionColor:   "#FFFFFF",
		CaptionText:    maymayText,
		TextDimensions: "1100x1000",
		TextGeometry:   "+100+100",
	}
}

func createAndSendEspondaMaymay(s *discordgo.Session, maymayText string, channelID string) {
	createAndSendMaymay(s, maymayText, channelID, createEspondaMaymay(maymayText))
}

func createEspondaMaymay(maymayText string) ImageAnnotationRecipe {
	return ImageAnnotationRecipe{
		InputImage:     "EspondaBlank.png",
		CaptionFont:    "Helvetica.ttf",
		CaptionColor:   "#042B66",
		CaptionText:    maymayText,
		TextDimensions: "420x250",
		TextGeometry:   "+460+125",
	}
}
