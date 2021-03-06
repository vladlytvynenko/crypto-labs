package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

const UPPER_ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type TrigramGenetic struct {
	secret          []byte
	trigrams        map[string]float64
	trigramStandard float64
}

type Population struct {
	key        string
	estimation float64
}

func substitutionWithoutDest(secret []byte, key []byte) []byte {
	dest := make([]byte, len(secret))
	substitution(dest, secret, key)
	return dest
}

func substitution(dest []byte, secret []byte, key []byte) {
	for i, c := range secret {
		dest[i] = key[c-65]
	}
}

func NewTrigramGenetic(trigramPath string, englishTextPath string, secret []byte) *TrigramGenetic {
	t := TrigramGenetic{
		secret: secret,
	}
	t.LoadTrigrams(trigramPath)
	t.LoadEnglishStandard(englishTextPath)
	return &t
}

func (t *TrigramGenetic) LoadEnglishStandard(path string) {
	t.trigramStandard = t.CalculateTrigramIndex(GetTextFromFile(path))
}

func (t *TrigramGenetic) LoadTrigrams(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	trigramsCounts := make(map[string]int64)
	t.trigrams = make(map[string]float64)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var sum int64

	for scanner.Scan() {
		splited := strings.Split(scanner.Text(), "\t")
		key := splited[0]
		value, err := strconv.ParseInt(splited[1], 10, 64)
		if err != nil {
			panic(err)
		}
		trigramsCounts[strings.ToUpper(key)] = value
		sum += value
	}
	for key, value := range trigramsCounts {
		newVal := float64(value) / float64(sum)
		t.trigrams[key] = math.Log10(newVal)
	}
}

func (t *TrigramGenetic) TrigramEstimation(text string) float64 {
	index := t.CalculateTrigramIndex(text)
	return math.Abs(index - t.trigramStandard)
}

func (t *TrigramGenetic) CalculateTrigramIndex(text string) float64 {
	index := 0.
	for i := 0; i < len(text)-2; i++ {
		index += t.trigrams[text[i:i+3]]
	}
	return index / float64(len(text)-2)
}

func (t *TrigramGenetic) GetInitialPopulation(populationCount int) []Population {
	randomAlphabet := []byte(UPPER_ALPHABET)
	out := make([]Population, populationCount)
	for i := 0; i < populationCount; i++ {
		rand.Shuffle(len(randomAlphabet), func(i, j int) {
			randomAlphabet[i], randomAlphabet[j] = randomAlphabet[j], randomAlphabet[i]
		})
		out[i] = Population{
			key:        string(randomAlphabet),
			estimation: -1,
		}
	}
	return out
}

func (t *TrigramGenetic) GetBestFromPopulation(population []Population, aliveCount int) []Population {
	allPopulation := make([]Population, len(population))
	for i, p := range population {
		allPopulation[i] = Population{
			key:        p.key,
			estimation: t.TrigramEstimation(p.key),
		}
	}

	sort.SliceStable(allPopulation, func(i, j int) bool {
		return allPopulation[i].estimation < allPopulation[j].estimation
	})
	seen := make(map[string]struct{}, 0)
	out := make([]Population, 0)
	for _, p := range allPopulation {
		if len(out) == aliveCount {
			break
		}
		if _, ok := seen[p.key]; !ok {
			out = append(out, p)
			seen[p.key] = struct{}{}
		}
	}

	return out
}

func (t *TrigramGenetic) ChangeDecision(child []byte, firstParent []byte, secondParent []byte, i int) []byte {
	if !bytes.Contains(child, firstParent[i:i+1]) {
		return child
	}
	newIdx := bytes.Index(child, firstParent[i:i+1])
	child[newIdx] = 0
	child = t.ChangeDecision(child, firstParent, secondParent, newIdx)
	child[newIdx] = firstParent[newIdx]
	return child
}

func (t *TrigramGenetic) Cross(first Population, second Population) Population {

	child := make([]byte, 0)
	firstParent := []byte(first.key)
	secondParent := []byte(second.key)

	for i := 0; i < len(firstParent); i++ {

		//we need to change decision if child has already 2 of parents values
		if bytes.Contains(child, firstParent[i:i+1]) && bytes.Contains(child, secondParent[i:i+1]) {
			child = t.ChangeDecision(child, firstParent, secondParent, i)
		}

		if bytes.Contains(child, firstParent[i:i+1]) {
			child = append(child, secondParent[i])
		} else if bytes.Contains(child, secondParent[i:i+1]) {
			child = append(child, firstParent[i])
		} else {
			if rand.Int()%2 == 0 {
				child = append(child, firstParent[i])
			} else {
				child = append(child, secondParent[i])
			}
		}
	}

	return Population{
		key:        string(child),
		estimation: -1,
	}

}
func (t *TrigramGenetic) Crossing(population []Population) []Population {
	children := make([]Population, 0)
	for i := 1; i < len(population)*2; i++ {
		idx1 := rand.Int63n(int64(len(population)))
		idx2 := rand.Int63n(int64(len(population)))
		// we need different rand numbers
		for idx1 == idx2 {
			idx2 = rand.Int63n(int64(len(population)))
		}
		child := t.Cross(population[idx1], population[idx2])
		children = append(children, child)
	}

	return append(children, population...)
}

func (t *TrigramGenetic) Mutate(population Population) Population {
	key := []byte(population.key)
	firstRand := rand.Int63n(int64(len(key)))
	secondRand := rand.Int63n(int64(len(key)))
	key[firstRand], key[secondRand] = key[secondRand], key[firstRand]
	return Population{
		key:        string(key),
		estimation: -1,
	}

}

func (t *TrigramGenetic) MutatePopulation(population []Population) {
	for i := 0; i < len(population); i++ {
		if rand.Int63n(100) <= 20 {
			population[i] = t.Mutate(population[i])
		}
	}
}

func UniqLen(p []Population) int {
	res := make(map[string]struct{}, 0)
	for _, pp := range p {
		res[pp.key] = struct{}{}
	}
	return len(res)
}

func (t *TrigramGenetic) SubstitutionWithGeneticAlgorithm() string {
	population := t.GetInitialPopulation(3000)
	best := t.GetBestFromPopulation(population, 1)[0]
	for generation := 0; t.TrigramEstimation(best.key) >= 0.021; generation++ {
		if generation%100 == 0 {
			fmt.Println(
				"Generation:", generation,
			)
			fmt.Println(
				"Estimation:", t.TrigramEstimation(best.key),
			)
			fmt.Println(
				"Key:", best.key,
			)
			fmt.Println(
				"Result:", string(substitutionWithoutDest(t.secret, []byte(best.key))),
			)
		}
		bestFromPopulation := t.GetBestFromPopulation(population, 1000)
		children := t.Crossing(bestFromPopulation)
		t.MutatePopulation(children)
		population = t.GetBestFromPopulation(children, len(children))
		best = population[0]
	}
	fmt.Println(population)
	return string(substitutionWithoutDest(t.secret, []byte(t.GetBestFromPopulation(population, 1)[0].key)))
}
