package main

// import whatever packages you will require here
import (
    "bufio"
    "encoding/csv"
    "os"
    "fmt"
    "io"
)

// use the following sample dataset
/*
| x     | y     |
|-------|-------|
| 4512  | 1530  |
| 3738  | 1297  |
| 4261  | 1335  |
| 3777  | 1282  |
| 4177  | 1590  |
*/

// expect the dataset to be in csv form

// FUNCTION DEFINITION
// add parameters and change the return type as necessary

// function to read csv file (dataset)

func read_csv(filename string)  {
    f, _ := os.Open(filename)
    r := csv.NewReader(f)
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }

        if err != nil {
            panic(err)
        }

        fmt.Println(record)
        fmt.Println(len(record))
        for value := range record {
            fmt.Printf("  %v\n", record[value])
        }
    }
}


// function to calculate weights

// // FORMULA
// W = Inverse(X' . X) . X' . y
// where X' is the transpose of X and . denotes dot product

func calc_weights() {

}

// function to predict y values

// // FORMULA
// y = X . W
// Where W is the calculated weights and . denotes dot product

func predict_y() {

}

// function to calculate mean
func mean(data []float64) float64 {
    if len(data) == 0 {
        return 0.0
    }

    var sum float64 = 0

    for _, value := range data {
        sum += value
    }

    return sum / float64(len(data))
}

// function to plot regression line

func plot_regression_line() {
	// insert code here
}

// function to plot the data points

func plot_data_points() {
	// insert code here
}

// function to calculate R-squared value to measure accuracy of model

// // FORMULA
// R^2 = (Σ(y_pred - y')^2)/(Σ(y-y')^2)
// where y_pred is the predicted value for y, y' is the mean

func calculate_r_squared() float32 {
	// insert code here
}

// MAIN FUNCTION

func main() {

}
