package db

import (
	"context"
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/mgiks/typo-typer/internal/storage"
)

func Seed(store storage.Store) error {
	texts := generateTexts(50)
	for _, text := range texts {
		if err := store.Text().Create(context.Background(), text); err != nil {
			return fmt.Errorf("failed to create text: %w", err)
		}
	}
	return nil
}

var words = []string{
	"apple", "bridge", "candle", "drift", "ember", "forest",
	"glove", "harbor", "island", "jungle", "kitten", "lantern",
	"meadow", "napkin", "orange", "pencil", "quartz", "river",
	"shadow", "thunder", "umbrella", "velvet", "window", "xylophone",
	"yogurt", "zephyr", "anchor", "button", "castle", "desert",
	"engine", "feather", "garden", "hammer", "icicle", "jacket",
	"kernel", "ladder", "mirror", "needle", "ocean", "planet",
	"quiver", "rocket", "silver", "tunnel", "unicorn", "violet",
	"whistle", "xenon", "yonder", "zipper", "avalanche", "basket",
	"compass", "dolphin", "echo", "falcon", "galaxy", "helmet",
	"inkwell", "jigsaw", "kayak", "library", "mountain", "nectar",
	"orchard", "parrot", "quicksand", "rainbow", "saddle", "temple",
	"utensil", "valley", "wallet", "xylem", "yearbook", "zodiac",
	"artist", "blossom", "crystal", "dragon", "emerald", "fountain",
	"guitar", "horizon", "insect", "jasmine", "kingdom", "legend",
	"marble", "notebook", "opal", "pirate", "quinoa", "radar",
	"sunrise", "traveler",
}

func generateTexts(count int) []*storage.Text {
	texts := make([]*storage.Text, count)
	for i := range count {
		wordCount := rand.IntN(len(words)) + 1
		wrds := make([]string, wordCount)
		for i := range wordCount {
			wrds[i] = words[rand.IntN(len(words))]
		}
		texts[i] = &storage.Text{
			Content: strings.Join(wrds, " "),
		}
	}
	return texts
}
