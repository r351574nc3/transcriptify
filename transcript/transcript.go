package transcript

var (
	transcript *Transcript
)

func Setup() *Transcript {
	transcript = transcript.New()

	return transcript
}