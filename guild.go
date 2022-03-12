package guild

type guild struct {
	Guild               string
	LogChannel          string
	VerifyChannel       string
	UnverifiedRole      string
	Automodnword        string
	Automodlinks        string
	Automodads          bool
	Auotmodiploggers    bool
	Automodmassmention  bool
	Massmentionrate     int
	Masscapsrate        int
	Automodmasscaps     bool
	Automodipv4         bool
	Automodipv6         bool
	Automodcc           bool
	Automodphonenumbers bool
	Automodssn          bool
	Automoddehoist      bool
	Prefix              bool
	Ticket              bool
	TicketRole          string
	MuteRole            string
	Autorole            bool
	Autoroles           []string
}

func token(token string) {
}
