package scheduler

import (
	"fmt"
	"math/rand"
	"sort"
)

const (
	populationSize = 1000
	maxGenerations = 10
)

type ByFitness []*Chromosome

func (a ByFitness) Len() int           { return len(a) }
func (a ByFitness) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFitness) Less(i, j int) bool { return a[i].GetFitness() < a[j].GetFitness() }
// SchedulerMain handles the scheduling process
type SchedulerMain struct {
	firstList     []*Chromosome
	newList        []*Chromosome
	firstListFitness float64
	newListFitness  float64
}

// finalSon holds the final suitable Chromosome
var finalSon *Chromosome

func NewSchedulerMain() *SchedulerMain {
	sm := &SchedulerMain{}

	// Print input data
	newUtility := Utility{}
    newUtility.PrintInputData()

	// Generate slots
	another:=NewTimeTable()

	// Print slots
	newUtility.PrintSlots(another)
	// Initialize the first generation of chromosomes
	sm.initializePopulation()

	// Generate new generations of chromosomes using crossovers and mutation
	sm.createNewGenerations()

	return sm
}

// createNewGenerations creates new generations using crossovers and mutations
func (sm *SchedulerMain) createNewGenerations() {
	var father, mother, son *Chromosome

	nogenerations := 0

	for nogenerations < maxGenerations {
		sm.newList = []*Chromosome{}
		sm.newListFitness = 0

		// Elitism: Add the top 10% chromosomes as they are
		for i := 0; i < populationSize/10; i++ {
			sm.newList = append(sm.newList, sm.firstList[i].DeepClone())
			sm.newListFitness += sm.firstList[i].GetFitness()
		}

		// Add other members after performing crossover and mutation
		i := populationSize / 10
		for i < populationSize {
			father = sm.selectParentRoulette()
			mother = sm.selectParentRoulette()

			// Crossover
			if rand.Float64() < inputData.CrossoverRate {
				son = sm.crossover(father, mother)
			} else {
				son = father
			}

			// Mutation
			sm.customMutation(son)

			if son.Fitness == 1 {
				fmt.Println("Selected Chromosome is:")
				son.PrintChromosome()
				break
			}

			sm.newList = append(sm.newList, son)
			sm.newListFitness += son.GetFitness()
			i++
		}

		// If a suitable chromosome with fitness 1 is found
		if len(sm.newList) < populationSize {
			fmt.Println("****************************************************************************************")
			fmt.Printf("\n\nSuitable Timetable has been generated in the %dth Chromosome of %d generation with fitness 1.\n", len(sm.newList), nogenerations+2)
			fmt.Println("\nGenerated Timetable is:")
			if finalSon != nil {
				finalSon.PrintTimeTable()
			} else {
				fmt.Println("finalSon is nil, cannot print timetable.")
			}
			finalSon = son
			break
		}
		finalSon = son
		// If chromosome with required fitness not found in this generation
		sm.firstList = sm.newList
		sort.Sort(ByFitness(sm.newList))
		sort.Sort(ByFitness(sm.firstList))
		fmt.Printf("**************************     Generation %d     ********************************************\n", nogenerations+2)
		sm.printGeneration(sm.newList)
		nogenerations++
	}
}
// selectParentRoulette selects a parent using Roulette Wheel Selection from the top 10% chromosomes
func (sm *SchedulerMain) selectParentRoulette() *Chromosome {
	firstListFitness := sm.firstListFitness / 10
	randomDouble := rand.Float64() * firstListFitness
	currentsum := 0.0
	i := 0

	for currentsum <= randomDouble {
		currentsum += sm.firstList[i].GetFitness()
		i++
	}
	return sm.firstList[i-1].DeepClone()
}

// customMutation performs a custom mutation
func (sm *SchedulerMain) customMutation(c *Chromosome) {
	oldFitness := c.GetFitness()
	geneno := rand.Intn(inputData.NoStudentGroup)

	i := 0
	newFitness := 0.0
	for newFitness < oldFitness {
		c.Gene[geneno] = *NewGene(geneno)
		newFitness = c.GetFitness()
		i++
		if i >= 500000 {
			break
		}
	}
}

// crossover performs a two-point crossover
func (sm *SchedulerMain) crossover(father, mother *Chromosome) *Chromosome {
	randomInt := rand.Intn(inputData.NoStudentGroup)
	temp := father.Gene[randomInt].DeepClone()
	father.Gene[randomInt] = *mother.Gene[randomInt].DeepClone()
	mother.Gene[randomInt] = *temp

	if father.GetFitness() > mother.GetFitness() {
		return father
	}
	return mother
}

// initializePopulation creates the initial population of chromosomes
func (sm *SchedulerMain) initializePopulation() {
	sm.firstList = []*Chromosome{}
	sm.firstListFitness = 0

	for i := 0; i < populationSize; i++ {
		c := NewChromosome()
		sm.firstList = append(sm.firstList, c)
		sm.firstListFitness += c.Fitness
	}

	sort.Sort(ByFitness(sm.firstList))
	fmt.Println("----------Initial Generation-----------")
	sm.printGeneration(sm.firstList)
}

// printGeneration prints important details of a generation
func (sm *SchedulerMain) printGeneration(list []*Chromosome) {
	fmt.Println("Fetching details from this generation...")

	for i := 0; i < 4; i++ {
		fmt.Printf("Chromosome no.%d: %f\n", i, list[i].GetFitness())
		list[i].PrintChromosome()
		fmt.Println()
	}

	fmt.Printf("Chromosome no. %d: %f\n", populationSize/10+1, list[populationSize/10+1].GetFitness())
	fmt.Printf("Chromosome no. %d: %f\n", populationSize/5+1, list[populationSize/5+1].GetFitness())
	fmt.Printf("Most fit chromosome from this generation has fitness = %f\n", list[0].GetFitness())
}

// selectParentBest selects a parent from the best chromosomes
func (sm *SchedulerMain) selectParentBest(list []*Chromosome) *Chromosome {
	randomInt := rand.Intn(100)
	return list[randomInt].DeepClone()
}

// mutation performs a simple mutation
func (sm *SchedulerMain) mutation(c *Chromosome) {
	geneno := rand.Intn(inputData.NoStudentGroup)
	temp := c.Gene[geneno].SlotNo[0]
	for i := 0; i < inputData.DaysPerWeek*inputData.HoursPerDay-1; i++ {
		c.Gene[geneno].SlotNo[i] = c.Gene[geneno].SlotNo[i+1]
	}
	c.Gene[geneno].SlotNo[inputData.DaysPerWeek*inputData.HoursPerDay-1] = temp
}

// swapMutation performs a swap mutation
func (sm *SchedulerMain) swapMutation(c *Chromosome) {
	geneno := rand.Intn(inputData.NoStudentGroup)
	slotno1 := rand.Intn(inputData.HoursPerDay * inputData.DaysPerWeek)
	slotno2 := rand.Intn(inputData.HoursPerDay * inputData.DaysPerWeek)

	temp := c.Gene[geneno].SlotNo[slotno1]
	c.Gene[geneno].SlotNo[slotno1] = c.Gene[geneno].SlotNo[slotno2]
	c.Gene[geneno].SlotNo[slotno2] = temp
}
