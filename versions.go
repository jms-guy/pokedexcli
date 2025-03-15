package main

//Version enum "structure" used to validate the input by the user, making sure it's a valid, working game version.\\

type Version int

const (
	Red Version = iota
	Blue
	Yellow
	Gold
	Silver
	Crystal
	Ruby
	Sapphire
	Emerald
	FireRed
	LeafGreen
	Diamond
	Pearl
	Platinum
	HeartGold
	SoulSilver
	Black
	White
	Black2
	White2
	X
	Y
	Sun
	Moon
	Sword
	Shield
	LegendsArceus
	Scarlet
	Violet
)

var versionName = map[Version]string{
	Blue:	"blue",
	Red:	"red",
	Yellow:	"yellow",
	Gold:	"gold",
	Silver:	"silver",
	Crystal:"crystal",
	Ruby:	"ruby",
	Sapphire:"sapphire",
	Emerald:"emerald",
	FireRed:"firered",
	LeafGreen:"leafgreen",
	Diamond:"diamond",
	Pearl:"pearl",
	Platinum:"platinum",
	HeartGold:"heartgold",
	SoulSilver:"soulsilver",
	Black:	"black",
	White:	"white",
	Black2:	"black-2",
	White2:	"white-2",
	X:	"x",
	Y:	"y",
	Sun:	"sun",
	Moon:	"moon",
	Sword:	"sword",
	Shield:	"shield",
	LegendsArceus:	"legends-arceus",
	Scarlet:	"scarlet",
	Violet:	"violet",
}
