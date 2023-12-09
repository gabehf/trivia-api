package trivia

import (
	"encoding/json"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"sync"
)

// represents the structure of trivia questions stored
// in the trivia.json file
type Question struct {
	Question string   `json:"question"`
	Category string   `json:"category"`
	Format   string   `json:"format"`
	Choices  []string `json:"choices"`
	Answer   string   `json:"answer"`
}

type Questions struct {
	Categories []string
	M          map[string][]Question
	lock       *sync.RWMutex
}

func (q *Questions) Init() {
	q.lock = &sync.RWMutex{}
}

func (q *Questions) Load(r io.Reader) error {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.M == nil {
		q.M = make(map[string][]Question, 0)
	}
	err := json.NewDecoder(r).Decode(&q.M)
	if err != nil {
		return err
	}
	if q.Categories == nil {
		q.Categories = make([]string, 0)
	}
	for key := range q.M {
		q.Categories = append(q.Categories, key)
	}
	return nil
}

func (q *Questions) categoryExists(cat string) bool {
	q.lock.RLock()
	defer q.lock.RUnlock()
	cat = strings.ToLower(cat)
	if q.M[cat] == nil || len(q.M[cat]) < 1 {
		return false
	}
	return true
}

func (q *Questions) getRandomCategory() string {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.Categories[rand.Int()%len(q.Categories)]
}

// Gets a random question from the category, if specified.
func (q *Questions) GetRandomQuestion(category string) (*Question, int) {
	q.lock.RLock()
	defer q.lock.RUnlock()
	// NOTE: it is okay to call another function that locks the RWMutex here,
	// as it will not cause a deadlock since CategoryExists only locks the Read
	if category == "" {
		category = q.getRandomCategory()
	} else if !q.categoryExists(category) {
		return nil, 0
	}
	category = strings.ToLower(category)
	qIndex := rand.Int() % len(q.M[category])
	return &q.M[category][qIndex], qIndex
}

func (q *Questions) GetQuestionById(id string) *Question {
	// get values from question_id
	questionSlice := strings.Split(id, "|")
	if len(questionSlice) != 2 {
		return nil
	}
	category, indexS := questionSlice[0], questionSlice[1]
	category = strings.ToLower(category)
	index, err := strconv.Atoi(indexS)
	if err != nil {
		return nil
	}

	q.lock.RLock()
	defer q.lock.RUnlock()
	// ensure category exists
	if !q.categoryExists(category) {
		return nil
	}
	// ensure question index is valid
	if len(q.M[category]) <= index {
		return nil
	}

	// retrieve question
	return &q.M[category][index]
}
