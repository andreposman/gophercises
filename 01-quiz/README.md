# Quiz Game

This Go program is designed to create a quiz game based on questions and answers provided in a CSV file.

## Part 1
Create a program that reads in a quiz from a CSV file and presents the quiz to the user, keeping track of the number of correct and incorrect answers. The CSV file should contain questions and their corresponding answers in a format like:

```
5+5,10
7+3,10
1+1,2
...
```

The user should be able to customize the filename of the CSV file via a flag. At the end of the quiz, the program should output the total number of correct answers and the total number of questions asked.

## Part 2
Adapt the program from Part 1 to include a timer. The default time limit should be 30 seconds, but the user should be able to customize it via a flag.

Before starting the quiz, the user should be prompted to press a key to start the timer. Questions should be presented one by one, and regardless of whether the answer is correct or wrong, the next question should be asked immediately. If the timer expires before the user completes the quiz, the program should end the quiz early. Unanswered questions should be considered incorrect.

At the end of the quiz, the program should still output the total number of correct answers and the total number of questions asked.

## CSV Format
CSV files may have questions with commas in them, for example:
```
"what 2+2, sir?",4
```

## Instructions
1. Clone the repository.
2. Navigate to the project directory.
3. Run the program using `go run main.go`.
4. Use flags `-csv` to specify the CSV file and `-time` to specify the time limit.

## Example Usage
```bash
go run main.go -csv problems.csv -time 45
```

## Dependencies
- This program uses the Go standard library's `encoding/csv` package for CSV parsing.

## Bonus
Bonus
As a bonus exercises you can also…

Add string trimming and cleanup to help ensure that correct answers with extra whitespace, capitalization, etc are not considered incorrect. Hint: Check out the strings package.
Add an option (a new flag) to shuffle the quiz order each time it is run.
