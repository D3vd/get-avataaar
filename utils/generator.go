package utils

import (
	"fmt"
	"get-avataaar/data"
	"get-avataaar/types"
	"log"

	"github.com/google/go-querystring/query"
)

// GenerateURL ...
func GenerateURL(avatarStyle string) (url string) {

	baseURL := "https://avataaars.io/"

	p := types.Parameters{
		AccessoriesType: GetRandom(data.AccessoriesType),
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

	if avatarStyle == "Circle" {
		p.AvatarStyle = "Circle"
	} else if avatarStyle == "Transparent" {
		p.AvatarStyle = "Transparent"
	} else {
		p.AvatarStyle = GetRandom(data.AvatarStyle)
	}

	v, err := query.Values(p)

	if err != nil {
		log.Println(err)
	}

	url = fmt.Sprintf("%s?%s", baseURL, v.Encode())

	return
}
