# linear-stats
A simple program that reads in a data file and calculates the Simple Linear Regression and Pearson Correlation Coefficient.

## Installation
1. Run `git clone https://01.gritlab.ax/git/ylee/linear-stats.git` in the terminal at the desired destination.

## Usage
Run `go run . <filename>` in the terminal to run the program.

The data file should be lines of integers.
The line number will be the x value and the integer will be the y value. Note: the x value starts counting at 0.

Example:
````
123
124
134
131
111
145
189
...
````

Expected output:
```
Linear Regression Line: y = <value>x + <value>
Pearson Correlation Coefficient: <value>
```

The program will display the statistic in the format above. The values for Linear Regression Line is to 6 decimal places. The Pearson Correlation Coeffiecnt is to 10 decimal places.

## Implementation
The functions `calcSums()`, `GetSimpLinearModel()` and `GetPearsonCoef()` are implemented to produce the results.

`calcSums()` calculate the:
- sum of x
- sum of y
- sum of x*y
- sum of x^2
- sum of y^2

These sums are used to carry the calculation for the Simple Linear Regression Line and the Pearson Correlation Coefficient.