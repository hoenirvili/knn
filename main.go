package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x     float64 // x coordonate
	y     float64 // y coordonate
	label string  // the classifier
}

type Labels struct {
	name  string
	count int
}

func (p Point) String() string {
	return fmt.Sprintf("[*] X = %f, Y = %f Label = %s\n", p.x, p.y, p.label)
}

type Data struct {
	point    Point   // point that has coordonates x, y and a label
	distance float64 // the distance from X To point
}

func (d Data) String() string {
	return fmt.Sprintf(
		"X = %f Y = %f, Distance = %f Label = %s\n",
		d.point.x, d.point.y, d.distance, d.point.label,
	)
}

// type Block implements the sort interface
type Block []Data

func (b Block) Len() int           { return len(b) }
func (b Block) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b Block) Less(i, j int) bool { return b[i].distance < b[j].distance }

func EuclidianDistance(A Point, X Point) (distance float64, err error) {
	distance = math.Sqrt(math.Pow((X.x-A.x), 2) + math.Pow((X.y-A.y), 2))
	if distance < 0 {
		return 0, fmt.Errorf("Invalid euclidian distance, the result is negative")
	}

	return distance, nil
}

func NewLoadedData(csvPath string) (data []Data, err error) {
	// open file descriptor
	fd, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	// make a new reader out of the file descriptor
	reader := csv.NewReader(fd)
	// read all records and parse them as strings
	records, err := reader.ReadAll()
	fmt.Println("[*] Loading records")
	fmt.Println()
	// output them into a debugging friendly manour
	lines := len(records)
	columns := len(records[0])
	if columns < 3 {
		return nil, fmt.Errorf("Cannot not load this data")
	}
	for i := 0; i < lines; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("%s\t  ", records[i][j])
		}
		if i == 0 {
			fmt.Println()
		}
		fmt.Println()
	}
	fmt.Println()
	// parse the raw string data into a []Data slice
	var value float64
	data = make([]Data, lines-1, lines-1)
	// we need now to parse every record and make it all floats exept the
	// label column , we will leave it as string
	for i := 1; i < lines; i++ {
		value, err = strconv.ParseFloat(records[i][0], 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse X value: %v", err)
		}
		data[i-1].point.x = value
		value, err = strconv.ParseFloat(records[i][1], 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse Y value: %v", err)
		}
		data[i-1].point.y = value
		data[i-1].point.label = records[i][2]
	}
	return data, nil
}

// logOutError checks if we have a valid
// error and if we have print it and exit the program
func logOutError(err error) {
	if err != nil {
		fmt.Printf("[!] %s\n", err.Error())
		os.Exit(1)
	}
}

// In pattern recognition, the k-Nearest Neighbors algorithm (or k-NN for short)
// is a non-parametric method used for classification and regression.
// [1] In both cases, the input consists of the k closest training examples
// in the feature space. The output depends on whether k-NN is used for classification or regression:
//
// In k-NN classification, the output is a class membership.
// An object is classified by a majority vote of its neighbors,
// with the object being assigned to the class most common among
// its k nearest neighbors (k is a positive integer, typically small).
// If k = 1, then the object is simply assigned to the class of that single nearest neighbor.
//
// In k-NN regression, the output is the property value for the object.
// This value is the average of the values of its k nearest neighbors.
// k-NN is a type of instance-based learning, or lazy learning,
// where the function is only approximated locally and all computation
// is deferred until classification. The k-NN algorithm is among the simplest of all machine learning algorithms.
func Knn(data []Data, k byte, X *Point) (err error) {
	n := len(data)
	// compute every point distance with X
	for i := 0; i < n; i++ {
		if data[i].distance, err = EuclidianDistance(data[i].point, *X); err != nil {
			return err
		}
	}

	var blk Block
	blk = data
	// sort the data in ascending order order
	sort.Sort(blk)
	var save []Labels
	for i := byte(0); i < k; i++ {
		save = foundLabelAndIncrement(data[i].point.label, save)
	}

	fmt.Printf("[*] Using k as %d\n", k)
	fmt.Println()
	fmt.Printf("[*] %+v\n", save)
	fmt.Println()

	max := 0
	var maxLabel string
	m := len(save)
	for i := 0; i < m; i++ {
		if max < save[i].count {
			max = save[i].count
			maxLabel = save[i].name
		}
	}

	X.label = maxLabel
	return nil
}

// foundLabelAndINcrement takes a one label and it searches the labels we
// already saved and if we found one increment the count of that label
// if we didn't find any label that means that the label is new and we must
// add it to the slice
// if the labels is empty then just add and return it
func foundLabelAndIncrement(label string, labels []Labels) []Labels {
	if labels == nil {
		labels = append(labels, Labels{
			name:  label,
			count: 1,
		})
		return labels
	}

	count := len(labels)
	for i := 0; i < count; i++ {
		if strings.Compare(labels[i].name, label) == 0 {
			labels[i].count++
			return labels
		}
	}

	return append(labels, Labels{
		name:  label,
		count: 1,
	})
}

func main() {
	fmt.Println()
	fmt.Println("[*] KNN algorithm")
	// set up default k we wish to compute
	k := []byte{1, 3, 5, 7, 9, 11}

	data, err := NewLoadedData("data.csv")
	logOutError(err)

	// read from keyboard
	var X Point
	fmt.Println("[i] Give the X point to look for")
	fmt.Printf("x = ")
	fmt.Scanf("%f", &X.x)
	fmt.Printf("\ny = ")
	fmt.Scanf("%f", &X.y)
	fmt.Println()

	n := len(k)
	for i := 0; i < n; i++ {
		err = Knn(data, k[i], &X)
		if i == 0 {
			fmt.Println(data)
		}
		logOutError(err)
		fmt.Printf("[*] Result for X is ")
		fmt.Println(X)
	}
}
