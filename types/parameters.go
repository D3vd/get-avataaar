package types

// Parameters structure used for URL query generation
type Parameters struct {
	AccessoriesType string `url:"accessoriesType"`
	AvatarStyle     string `url:"avatarStyle"`
	ClotheColor     string `url:"clotheColor"`
	ClotheType      string `url:"clotheType"`
	EyebrowType     string `url:"eyebrowType"`
	EyeType         string `url:"eyeType"`
	FacialHairColor string `url:"facialHairColor"`
	FacialHairType  string `url:"facialHairType"`
	GraphicType     string `url:"graphicType"`
	HairColor       string `url:"hairColor"`
	HatColor        string `url:"hatColor"`
	MouthType       string `url:"mouthType"`
	SkinColor       string `url:"skinColor"`
	TopType         string `url:"topType"`
}
