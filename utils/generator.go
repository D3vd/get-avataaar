package utils

import (
	"fmt"
	"get-avataaar/data"
	"get-avataaar/types"
	"log"

	"github.com/google/go-querystring/query"
)

// GenerateURL ...
func GenerateURL() (url string) {

	baseURL := "https://avataaars.io/"

	p := types.Parameters{
		AccessoriesType: GetRandom(data.AccessoriesType),
		AvatarStyle:     GetRandom(data.AvatarStyle),
		ClotheColor:     GetRandom(data.ClotheColor),
		ClotheType:      GetRandom(data.ClotheType),
		EyebrowType:     GetRandom(data.EyebrowType),
		EyeType:         GetRandom(data.EyeType),
		FacialHairColor: GetRandom(data.FacialHairColor),
		FacialHairType:  GetRandom(data.FacialHairType),
		GraphicType:     GetRandom(data.GraphicType),
		HairColor:       GetRandom(data.HairColor),
		HatColor:        GetRandom(data.HatColor),
		MouthType:       GetRandom(data.MouthType),
		SkinColor:       GetRandom(data.SkinColor),
		TopType:         GetRandom(data.TopType),
	}

	v, err := query.Values(p)

	if err != nil {
		log.Println(err)
	}

	url = fmt.Sprintf("%s?%s", baseURL, v.Encode())

	return
}
