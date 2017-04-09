package models

type Response struct {
	Items []ResponseItem `json:"items"`
}

type ResponseItem struct {
	Uid          string    `json:"uid"`
	Valid        bool      `json:"valid"`
	Title        string    `json:"title"`
	Subtitle     string    `json:"subtitle"`
	Arg          string    `json:"arg"`
	Autocomplete string    `json:"autocomplete"`
	Icon         IconModel `json:"icon"`
	Text         TextModel `json:"text"`
	Mod          ModModel  `json:"mods"`
}

type IconModel struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type TextModel struct {
	Copy      string `json:"copy"`
	Largetype string `json:"largetype"`
}

type ModModel struct {
	Shift ModItems `json:"shift"`
	Cmd   ModItems `json:"cmd"`
}

type ModItems struct {
	Valid    bool   `json:"valid"`
	Arg      string `json:"arg"`
	Subtitle string `json:"subtitle"`
}
