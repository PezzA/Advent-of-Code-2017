# advent-of-go

Contains go solutions for https://adventofcode.com/

To get the code run

    go get github.com\pezza\advent-of-go


from the root of the package, execute the following commands - 

: to run all tests

    go test ./... -v

to run the test for a particular day run

    go test ./day<day number padded with a 0>

to build the project run

    go install github.com\pezza\advent-of-go


and then assuming you have the %gohome% bin folder in your %path% 

    advent-of-go <day number>


