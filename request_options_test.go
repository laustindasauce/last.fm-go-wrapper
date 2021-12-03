package lastfm

import (
	"testing"
)

func TestOptions(t *testing.T) {
	t.Parallel()

	resultSet := processOptions(
		LimitOpt(2),
		PageOpt(1),
		ArtistOpt("Lauv"),
		AlbumOpt("Sour"),
		TrackOpt("I like me better"),
		MbidOpt("03c91c40-49a6-44a7-90e7-a700edf97a62"),
		UsernameOpt("RJ"),
		LangOpt("eng"),
		CountryOpt("United States"),
		LocationOpt("West"),
		TagOpt("rock"),
		FromOpt(1638310000),
		ToOpt(1638316805),
		AutocorrectOpt(One),
		ExtendedOpt(Zero),
		TaggingTypeOpt(AlbumTag),
	)

	expected := "album=Sour&artist=Lauv&autocorrect=1&country=United+States&extended=0&from=1638310000&lang=eng&limit=2&location=West&mbid=03c91c40-49a6-44a7-90e7-a700edf97a62&page=1&tag=rock&taggingtype=album&to=1638316805&track=I+like+me+better&username=RJ"
	actual := resultSet.urlParams.Encode()
	if actual != expected {
		t.Errorf("Expected '%v', got '%v'", expected, actual)
	}
}
