package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.Parse()

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot find the specified file: %s\n  %s\n", filePath, err)
	}
	defer f.Close()

	morphs, err := newMorpheme(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	morphs.sortByAppearance().rankByCounts().drawLogLogGraph()
}

// morpheme represents the mapping list of MeCab format.
type morpheme map[string]string

// morphemes represents lists of Morpheme.
type morphemes []*morpheme

// newMorpheme implements the constructor of Morpheme.
// MeCab outputs the following data structure through morphological analysis.
// <Surface>\t<POS>,<POS subtyping1>,<POS subtyping2>,<POS subtyping3>,<Conjugation Form>,<Conjugation>,<Base>,<Furigana>,<Pronunciation>
func newMorpheme(r io.Reader) (morphemes, error) {
	morphs := morphemes{}
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		if sc.Text() == "EOS" || sc.Text() == "" {
			continue
		}
		// separation of the surface field and the rest of the fields.
		surf := strings.Split(sc.Text(), "\t")
		// separation of the rest of the fields.
		other := strings.Split(surf[1], ",")
		morphs = append(morphs, &morpheme{
			"surface": surf[0],
			"base":    other[6],
			"pos":     other[0],
			"pos1":    other[1],
		})
	}
	return morphs, nil
}

// sortCounter is used for sorting the key by counts
type sortCounter struct {
	Key   string
	Count int
}

// countSorsortCountersters represents the slice of Countsorter
type sortCounters []sortCounter

func (counters sortCounters) String() string {
	var buf bytes.Buffer
	for _, counter := range counters {
		// remove the first character "{"
		ans := fmt.Sprintf("%+v\n", counter)[1:]
		buf.WriteString(strings.Replace(ans, "}", "", -1))
	}
	return strings.TrimRight(buf.String(), "\n")
}

// sortByAppearance returns the surface verb stably sorted by their appearance
func (morphes *morphemes) sortByAppearance() sortCounters {
	// counter counts frequency of word
	counter := map[string]int{}
	var sortedWords sortCounters
	var word string

	// Use separate slice container to memorize stable iteration order
	// See Iteration order section in the following URL:
	// https://blog.golang.org/go-maps-in-action
	var keys []string

	for _, morph := range *morphes {
		word = (*morph)["surface"]
		if _, ok := counter[word]; ok {
			counter[word]++
			continue
		}
		// Memorize original order of words
		keys = append(keys, word)
		counter[word] = 1
	}

	for _, key := range keys {
		// Assign sorter in fixed order using memorized keys
		sortedWords = append(sortedWords, sortCounter{
			Key:   key,
			Count: counter[key],
		})
	}

	sort.SliceStable(sortedWords, func(i, j int) bool {
		return sortedWords[i].Count > sortedWords[j].Count
	})

	return sortedWords
}

// CountRank represents the ranking information
type CountRank struct {
	Count int
	Rank  int
}

// CountRanking represents the slice of CountRank
type CountRanking []CountRank

// rankByCounts aggregates the counts of appearance and returns the ranking
func (counters sortCounters) rankByCounts() CountRanking {
	var (
		ranking        CountRanking
		prevCount      int
		currentRank    int
		sameRankCounts int
	)

	for _, sortedWord := range counters {
		if prevCount != sortedWord.Count {
			currentRank += 1 + sameRankCounts
			sameRankCounts = 0
		} else {
			sameRankCounts++
		}
		ranking = append(ranking, CountRank{
			Count: sortedWord.Count,
			Rank:  currentRank,
		})
		prevCount = sortedWord.Count
	}

	return ranking
}

// drawLogLogGraph outputs the graph log log graph
func (ranking CountRanking) drawLogLogGraph() {
	points := make(plotter.XYs, len(ranking))

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	for i, result := range ranking {
		points[i].X = math.Log10(float64(result.Rank))
		points[i].Y = math.Log10(float64(result.Count))
	}

	p.Title.Text = "Ranking of word counts"
	p.X.Label.Text = "Logarithm of ranking"
	p.Y.Label.Text = "Logarithm of word counts"

	err = plotutil.AddLinePoints(p, "Ranking", points)
	if err != nil {
		panic(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "zipfs-distribution.png"); err != nil {
		panic(err)
	}
}
